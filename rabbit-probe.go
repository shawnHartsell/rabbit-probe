package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"

	"github.com/streadway/amqp"
)

var host = flag.String("h", "localhost", "server")
var port = flag.Int("p", 5672, "port")
var exchange = flag.String("e", "/", "exchange to publish to")
var vhost = flag.String("vh", "/", "virtual host")
var user = flag.String("u", "guest:guest", "user:pass combo")
var data = flag.String("d", "", "message data")

func main() {
	flag.Parse()

	uri := fmt.Sprintf("amqp://%s@%s:%d/%s", *user, *host, *port, url.QueryEscape(*vhost))
	log.Printf("Connecting to %s", uri)
	conn, err := amqp.Dial(uri)

	if err != nil {
		log.Fatalf("connection.open: %s", err)
	}

	defer conn.Close()
}
