package usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"os/exec"
	"strings"
	"time"

	"github.com/VasySS/service-monitoring-vk-task/pinger/internal/config"
)

var httpClient = http.Client{Timeout: time.Second * 5}

func ScrapeContainersLoop(ctx context.Context) {
	ticker := time.NewTicker(time.Duration(config.C.ScrapeInterval) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			slog.Info("scraping containers...")

			containers, err := getContainers()
			if err != nil {
				slog.Error("failed to get container list", slog.Any("error", err))
				continue
			}

			results := StartWorkers(containers)
			if err := sendPingResults(results); err != nil {
				slog.Error("failed to send scrape results", slog.Any("error", err))
			}
		}
	}
}

func getContainers() ([]string, error) {
	cmd := exec.Command(
		"sh",
		"-c",
		"docker ps -q | xargs -n 1 docker inspect --format '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}'",
	)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return []string{}, err
	}

	var ips []string

	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		if trimmed := strings.TrimSpace(line); trimmed != "" {
			ips = append(ips, trimmed)
		}
	}

	return ips, nil
}

func sendPingResults(results []PingResult) error {
	resBytes, err := json.Marshal(results)
	if err != nil {
		return fmt.Errorf("failed to marshal ping results: %w", err)
	}

	backendURL, err := url.Parse(config.C.BackendURL)
	if err != nil {
		return fmt.Errorf("failed to parse backend url: %w", err)
	}

	backendURL.Path = "/v1/statuses"

	resp, err := httpClient.Do(&http.Request{
		Method: "POST",
		URL:    backendURL,
		Header: http.Header{
			"Content-Type": {"application/json"},
		},
		Body: io.NopCloser(bytes.NewReader(resBytes)),
	})
	if err != nil {
		return fmt.Errorf("failed to make request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send ping results, status: %s", resp.Status)
	}

	return nil
}
