package tasks

import (
	"context"
	"fmt"
	"time"

	"sync/atomic"
)

type ScanStatus struct {
	ScanId      string `json:"scan_id,omitempty"`
	ScanStatus  string `json:"scan_status,omitempty"`
	ScanMessage string `json:"scan_message,omitempty"`
}

type StatusValues struct {
	IN_PROGRESS string
	CANCELLED   string
	FAILED      string
	SUCCESS     string
}

type ScanContext struct {
	Res           chan error
	StopTriggered atomic.Bool
	IsAlive       atomic.Bool
	Context       context.Context
	Cancel        context.CancelFunc
}

func (sc *ScanContext) Checkpoint(stage string) error {
	if sc != nil {
		err := sc.Context.Err()
		if err != nil {
			return fmt.Errorf("%v: %v", stage, err)
		}
		sc.IsAlive.Store(true)
	}
	return nil
}

func newScanContext(res chan error) *ScanContext {

	ctx, cancel := context.WithCancel(context.Background())

	obj := ScanContext{
		Res:     res,
		IsAlive: atomic.Bool{},
		Context: ctx,
		Cancel:  cancel,
	}
	return &obj
}

func StartStatusReporter(
	scanId string,
	sendScanStatus func(ScanStatus) error,
	sv StatusValues,
	threshold time.Duration) (chan error, *ScanContext) {

	res := make(chan error)
	scanCtx := newScanContext(res)

	go func() {
		ticker := time.NewTicker(30 * time.Second)
		var err error
		ttl := time.Now().Add(threshold)
	loop:
		for {
			select {
			case err = <-res:
				break loop
			case <-ticker.C:
				if scanCtx.IsAlive.Swap(false) {
					ttl = time.Now().Add(threshold)
				}

				if time.Now().After(ttl) {
					scanCtx.Cancel()
					ticker.Stop()
					// Wait for the job to finish
					continue
				}

				sendScanStatus(
					ScanStatus{
						scanId,
						sv.IN_PROGRESS,
						""})
			}
		}

		status := sv.SUCCESS
		status_message := ""
		if err != nil {
			if scanCtx.StopTriggered.Load() {
				status = sv.CANCELLED
			} else {
				status = sv.FAILED
			}
			status_message = err.Error()
		}

		sendScanStatus(
			ScanStatus{
				scanId,
				status,
				status_message})
	}()
	return res, scanCtx
}
