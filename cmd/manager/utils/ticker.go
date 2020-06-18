package utils

import "time"

type Ticker struct {
	f        func()
	stopChan chan struct{}
	hour     uint8
}

func NewTicker(f func(), hour uint8) *Ticker {
	t := &Ticker{
		f:    f,
		hour: hour,
	}
	return t
}

func (t *Ticker) Start() {
	ticker := time.NewTicker(time.Duration(t.hour) * time.Hour)

	for {
		select {
		case <-ticker.C:
			t.f()
		case <-t.stopChan:
			ticker.Stop()
			close(t.stopChan)
		}
	}

}

func (t *Ticker) Stop() {
	t.stopChan <- struct{}{}
}
