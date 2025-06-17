package scheduler

import (
	"log"
	"time"
)

type TaskFunc func() error

type Scheduler struct {
	interval time.Duration
	task     TaskFunc
	stopChan chan struct{}
}

func NewScheduler(intervalSeconds int, task TaskFunc) *Scheduler {
	return &Scheduler{
		interval: time.Duration(intervalSeconds) * time.Second,
		task:     task,
		stopChan: make(chan struct{}),
	}
}

func (s *Scheduler) Start() {
	ticker := time.NewTicker(s.interval)
	defer ticker.Stop()

	go func() {
		log.Println("[Scheduler] Starting scheduled task immediately")
		if err := s.task(); err != nil {
			log.Println("[Scheduler] Task error:", err)
		}
	}()

	for {
		select {
		case <-ticker.C:
			log.Println("[Scheduler] Running scheduled task")
			if err := s.task(); err != nil {
				log.Println("[Scheduler] Task error:", err)
			}
		case <-s.stopChan:
			log.Println("[Scheduler] Stopping scheduler")
			return
		}
	}
}

func (s *Scheduler) Stop() {
	close(s.stopChan)
}
