package main

import (
	"os"
	"strconv"
	"sync"
	"time"
)

type statusMonitor struct {
	ready          bool
	chStop         chan int
	monitorRateSec int
	healthChecks   []func() bool
}

var monitorOnce sync.Once
var monitor *statusMonitor

func (monitor *statusMonitor) addCheck(f func() bool) {
	monitor.healthChecks = append(monitor.healthChecks, f)
}

func getStatusMonitor() *statusMonitor {
	monitorOnce.Do(func() {
		monitor = &statusMonitor{ready: true, chStop: make(chan int)}
		rate, v := os.LookupEnv("MONITOR_RATE_SEC")
		if !v {
			monitor.monitorRateSec = 10
		} else {
			i, err := strconv.Atoi(rate)
			if err == nil {
				monitor.monitorRateSec = i
			} else {
				monitor.monitorRateSec = 10
			}
		}
	})
	return monitor
}

func (monitor *statusMonitor) isReady() bool {
	return monitor.ready
}

func (monitor *statusMonitor) start() {
	go func() {
		for {
			select {
			case <-monitor.chStop:
				return
			default:
				monitor.ready = monitor.getStatus()
				time.Sleep(time.Duration(monitor.monitorRateSec) * time.Second)
			}
		}
	}()
}

func (monitor *statusMonitor) getStatus() bool {
	for _, f := range monitor.healthChecks {
		if !f() {
			return false
		}
	}
	return true
}

func (monitor *statusMonitor) stop() {
	close(monitor.chStop)
}
