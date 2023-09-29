package tasks

import (
	"context"
	"errors"
	"fmt"
	"time"

	"sync/atomic"
)

type ScanContext struct {
	ScanID  string
	Res     chan error
	IsAlive atomic.Bool
	Context context.Context
	Cancel  context.CancelFunc
}

func (sc *ScanContext) IamAlive() {
	sc.IsAlive.Store(true)
}

func newScanContext(scanID string, res chan error) *ScanContext {

	ctx, cancel := context.WithCancel(context.Background())

	obj := ScanContext{
		ScanID:  scanID,
		Res:     res,
		IsAlive: atomic.Bool{},
		Context: ctx,
		Cancel:  cancel,
	}
	return &obj
}

var (
	StoppedError = errors.New("Job stopped")
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

func StartStatusReporter(
	scanId string,
	sendScanStatus func(ScanStatus) error,
	sv StatusValues) (chan error, *ScanContext) {

	res := make(chan error)
	scanCtx := newScanContext(scanId, res)

	//If we don't get any active status back within threshold,
	//we consider the scan job as dead
	threshold := 20 * time.Minute
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
					err = fmt.Errorf("Scan aborted as inactive since %v", ttl.Add(-threshold))
					break
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
			if errors.Is(err, StoppedError) {
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
