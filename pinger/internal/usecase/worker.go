package usecase

import (
	"fmt"
	"log/slog"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/VasySS/service-monitoring-vk-task/pinger/internal/config"
)

type PingStatus struct {
	IP        string
	Healthy   bool
	CreatedAt time.Time
}

func StartWorkers(containerIPs []string) []PingStatus {
	inputCh := make(chan string)
	outputCh := make(chan PingStatus)

	wg := &sync.WaitGroup{}

	go func() {
		defer close(inputCh)

		for _, ip := range containerIPs {
			inputCh <- ip
		}
	}()

	go func() {
		defer close(outputCh)

		for i := 0; i < config.C.Workers; i++ {
			wg.Add(1)
			go worker(wg, inputCh, outputCh)
		}

		wg.Wait()
	}()

	results := make([]PingStatus, 0, len(containerIPs))
	for res := range outputCh {
		results = append(results, res)
	}

	return results
}

func worker(wg *sync.WaitGroup, inputCh <-chan string, outputCh chan<- PingStatus) {
	defer wg.Done()

	for ip := range inputCh {
		status, err := ping(ip)
		if err != nil {
			slog.Error("failed to ping container", slog.String("ip", ip), slog.Any("error", err))
		}

		outputCh <- status
	}
}

func ping(ip string) (PingStatus, error) {
	if ip == "" {
		return PingStatus{}, fmt.Errorf("ip is empty")
	}

	cmd := exec.Command("ping", "-c", "1", "-W", "2", ip)
	createdAt := time.Now().UTC()

	resp, err := cmd.CombinedOutput()
	if err != nil {
		return PingStatus{
			IP:        ip,
			Healthy:   false,
			CreatedAt: createdAt,
		}, err
	}

	respStr := string(resp)
	if !strings.Contains(respStr, "1 received") || !strings.Contains(respStr, "0% packet loss") {
		return PingStatus{
			IP:        ip,
			Healthy:   false,
			CreatedAt: createdAt,
		}, nil
	}

	return PingStatus{
		IP:        ip,
		Healthy:   true,
		CreatedAt: createdAt,
	}, nil
}
