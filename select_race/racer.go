package select_race

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

// PROBLEMS with real HTTP tests: Slow, Flaky, Can't test edge cases
func RacerOld(urlA, urlB string) (winner string) {
	aDuration := measureResponseTime(urlA)
	bDuration := measureResponseTime(urlB)

	if aDuration < bDuration {
		return urlA
	}
	return urlB
}

func Racer(a, b string) (winner string, timeTaken time.Duration, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, timeTaken time.Duration, error error) {
    start := time.Now()
    select {
    case timeA := <-ping(a):
        return a, time.Since(start) - timeA, nil
    case timeB := <-ping(b):
        return b, time.Since(start) - timeB, nil
    case <-time.After(timeout):
        return "", 0, fmt.Errorf("timed out waiting for %s and %s", a, b)
    }
}

// MICROLESSON: adding enhanced timing to ping function
// using a type chan struct{} is the smallest data type in Go from a Mem. perspective
// chan struct{} is a signal only channel, it is used to signal that an event has occurred
// as an extension I used channel type of time.Duration here to measure the time taken to get a response from a URL
func ping(url string) chan time.Duration {
    ch := make(chan time.Duration)
    go func() {
        start := time.Now()
        _, err := http.Get(url)
        if err != nil {
            ch <- 0
            return
        }
        ch <- time.Since(start)
    }()
    return ch
}

// measureResponseTime returns the time taken to get a response from a URL
func measureResponseTime(url string) time.Duration {
	startA := time.Now()
	http.Get(url)
	return time.Since(startA)
}
