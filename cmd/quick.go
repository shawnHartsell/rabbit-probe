package cmd

import (
	"encoding/json"
	"errors"

	"github.com/shawnHartsell/rabbit-probe/timer"
	"github.com/spf13/cobra"
)

var (
	duration int
	rate     int
	uri      string
	exchange string
	payload  string
)

// quickCmd represents the quick command
var quickCmd = &cobra.Command{
	Use:   "quick",
	Short: "start a simple job to publish messages to an excahnge",
	Long: `
The quick command will publish the provided message to an exchange an an iterval determined by rate and 
duration. The command only supports JSON message bodies and does not modify the body
between publishes (to simulate malformed messages, etc).
Use quick mode as an easy way to smoke/stress test consumers
  
Example usage:

1) rabbit-probe quick

  start a probe using defaults
    

2) rabbit-probe quick -r 2 -d 10 -u amqp://guest:guest@localhost:5672/&2f -e default -p "{\"foo\":\"bar\"}"
   
   publish the provided message at a rate of 2 messages/second over a duration of 10 secs
`,
}

func init() {
	RootCmd.AddCommand(quickCmd)
	quickCmd.Flags().IntVarP(&duration, "duration", "d", 60, "duration (in seconds) that the probe should run")
	quickCmd.Flags().IntVarP(&rate, "rate", "r", 20, "message publish rate (per second) ")
	quickCmd.Flags().StringVarP(&uri, "uri", "u", "amqp://guest:guest@localhost:5672/%2f", "uri of the rabbitMQ broker to probe")
	quickCmd.Flags().StringVarP(&exchange, "exchange", "e", "default", "exchange to publish the payload to")
	quickCmd.Flags().StringVarP(&payload, "payload", "p", "{}", "message to publish (json)")
	quickCmd.RunE = quick
}

func quick(cmd *cobra.Command, args []string) error {
	if !isJSON(payload) {
		return errors.New("payload is not valid JSON")
	}

	timer.Start(duration, rate)
	return nil
}

func isJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}
