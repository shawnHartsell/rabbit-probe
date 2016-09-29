package probe

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

type Quick struct {
	Payload string
	Probe
}

func (p *Quick) Validate() error {
	var js map[string]interface{}

	if err := json.Unmarshal([]byte(p.Payload), &js); err != nil {
		errMsg := fmt.Sprintf("payload must be valid JSON \n %s \n", err.Error())
		return errors.New(errMsg)
	}

	if _, err := amqp.ParseURI(p.URI); err != nil {
		errMsg := fmt.Sprintf("uri is not a valid amqp uri \n %s \n", err.Error())
		return errors.New(errMsg)
	}

	return nil
}

func (p *Quick) PublishMessage(channel *amqp.Channel) error {

	msg := amqp.Publishing{
		DeliveryMode: amqp.Transient,
		ContentType:  "application/json",
		Timestamp:    time.Now(),
		Body:         []byte(p.Payload),
		MessageId:    "muh id",
	}

	return channel.Publish(p.Exchange, p.RoutingKey, false, false, msg)
}

func (p *Quick) DisplayResults() {
	fmt.Println("we done!")
}
