package _select

import (
	"fmt"
	"net/http"
	"time"
)

// Now why are we doing this synchronously when we are only concerned of which
// request comes back first? Can we call Get() on both websites at the same time?

// func Racer(urlA, urlB string) (url string){

// 	if measureResponseTime(urlA) < measureResponseTime(urlB) {
// 		return urlA
// 	}

// 	return urlB

// }

// func measureResponseTime(url string) time.Duration {
// 	start := time.Now()
// 	http.Get(url)
// 	return time.Since(start)
// }

// Learning
// 1) Select to wait on multiple chann. waiting for a channel is blocking though
// 1.5) time.After to set a fail safe timing to stop blocking
// 2) httptest -> good for simuating server locally

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

// chan struct{} usually used for notification because 
// struct{} is 0 bytes! 
func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}