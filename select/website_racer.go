package website_racer

import (
	"net/http"
	"time"
)

const (
	tenSecondTimeout = 10 * time.Second
)

func Racer(a, b string) (string, error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", http.ErrHandlerTimeout
	}
}

func ping(url string) chan any {
	ch := make(chan any)
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
