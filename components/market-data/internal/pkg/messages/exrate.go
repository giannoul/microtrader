package exrate

import (
	"fmt"
	queue "github.com/giannoul/microtrader/internal/pkg/rabbitmq"
	"log"
)

// Rate is the common message structure for the exchange rates
type Rate struct {
	Source    string
	Currency  string
	Rate      float64
	Timestamp string
}

// Store will store the message to the queue
func (r *Rate) Store() error {
	log.Printf("%#v", r)
	msgBytes := []byte(fmt.Sprintf("%v", r))
	queue.Enqueue(msgBytes)
	return nil
}
