package timer

import (
	"errors"
	"fmt"
	"time"

	"github.com/shawnHartsell/rabbit-probe/probe"
	"github.com/streadway/amqp"
)

//Start begins a timer that will invoke an operation at a rate of x times/sec over a duration of y secs.
//TODO: name is terrible, it implies an async operation
func Start(actions probe.Actions) (err error) {
	probe := actions.GetProbe()
	if err := actions.Validate(); err != nil {
		return err
	}

	//TODO: abstract, abstract, abstract
	fmt.Println("opening amqp connection")
	conn, err := amqp.Dial(probe.URI)
	if err != nil {
		return err
	}

	channel, err := conn.Channel()
	if err != nil {
		return err
	}

	tickerRate, err := getTickerRate(probe.Rate)
	if err != nil {
		return err
	}

	ticker := time.NewTicker(tickerRate)
	doneChan := make(chan bool)

	fmt.Printf("probe started at a rate of %d/s over %d seconds\n", probe.Rate, probe.Duration)
	go func() {
		time.Sleep(time.Second * time.Duration(probe.Duration))
		doneChan <- true
	}()

	for {
		select {
		case <-ticker.C:
			//TODO: handle channel errors (restablish connection, new channel, etc)
			actions.PublishMessage(channel)
		case <-doneChan:
			fmt.Println("probe has completed")
			ticker.Stop()
			actions.DisplayResults()
			return nil
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
