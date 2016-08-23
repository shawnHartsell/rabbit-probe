package timer

import (
	"errors"
	"fmt"
	"time"
)

//Start begins a timer that will invoke an operation at a rate of x times/sec over a duration of y secs.
func Start(duration int, rate int) (err error) {

	tickerRate, err := getTickerRate(rate)
	if err != nil {
		return err
	}

	ticker := time.NewTicker(tickerRate)
	doneChan := make(chan bool)

	go func() {
		time.Sleep(time.Second * time.Duration(duration))
		doneChan <- true
	}()

	for {
		select {
		case <-ticker.C:
			fmt.Println("got a tick!")
		case <-doneChan:
			fmt.Println("we done")
			ticker.Stop()
			return
		}
	}
}

func getTickerRate(rate int) (time.Duration, error) {
	numPlaces := (rate / 10) % 10

	if numPlaces > 3 {
		return time.Duration(0), errors.New("rates over 1000/s are not supported")
	}

	ticksPerMillisecond := (1.0 / float32(rate)) * 1000.0
	duration := time.Millisecond * time.Duration(ticksPerMillisecond)
	return duration, nil
}
