package usecase

import (
	"context"
	"fmt"
	"log/slog"
	"os/exec"
	"strings"
	"time"
)

type PingStatus struct {
	IP        string
	Healthy   bool
	CreatedAt time.Time
}

func ScrapeContainersLoop(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			containers, err := getContainers()
			if err != nil {
				slog.Error("failed to get containers", slog.Any("error", err))
				time.Sleep(time.Second * 5)
				continue
			}

			for _, container := range containers {
				pingContainer(container)
			}

			time.Sleep(time.Second * 5)
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

	return strings.Split(string(out), "\n"), nil
}

func pingContainer(ip string) (PingStatus, error) {
	if ip == "" {
		return PingStatus{}, fmt.Errorf("ip is empty")
	}

	cmd := exec.Command("ping", "-c", "1", "-W", "2", ip)

	resp, err := cmd.CombinedOutput()
	if err != nil {
		return PingStatus{}, err
	}

	respStr := string(resp)
	if !strings.Contains(respStr, "1 received") || !strings.Contains(respStr, "0% packet loss") {
		return PingStatus{
			IP:        ip,
			Healthy:   false,
			CreatedAt: time.Now().UTC(),
		}, nil
	}

	return PingStatus{
		IP:        ip,
		Healthy:   true,
		CreatedAt: time.Now().UTC(),
	}, nil
}
