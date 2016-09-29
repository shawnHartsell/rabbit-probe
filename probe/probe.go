package probe

import "github.com/streadway/amqp"

//TODO: improve the name
type Actions interface {
	GetProbe() *Probe
	Validate() error
	PublishMessage(channel *amqp.Channel) error
	DisplayResults()
}

type Probe struct {
	Duration   int
	Rate       int
	URI        string
	Exchange   string
	RoutingKey string
}

func (p *Probe) GetProbe() *Probe {
	return p
}
