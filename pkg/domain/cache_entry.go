package domain

import (
	"context"
	"time"
)

type Entry struct {
	Key     interface{}
	Value   interface{}
	Element interface{}
	Exp     time.Time
	timer   *time.Timer
	cancel  context.CancelFunc
}

func (e *Entry) startTimer(d time.Duration, f func(key, value interface{})) {
	ctx, cancel := context.WithCancel(context.TODO())
	e.cancel = cancel
	e.timer = time.AfterFunc(d, func() {
		if ctx.Err() != nil {
			return
		}

		f(e.Key, e.Value)
	})
}

func (e *Entry) stopTimer() {
	if e.timer == nil {
		return
	}
	e.timer.Stop()
	e.cancel()
}