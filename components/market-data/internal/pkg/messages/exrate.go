package exrate

import (
	"fmt"
	queue "github.com/giannoul/microtrader/internal/pkg/rabbitmq"
	"log"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// Rate is the common message structure for the exchange rates
type Rate struct {
	Source    string
	Currency  string
	Rate      float64
	Timestamp string
}

// Store will store the message to the queue
func (r *Rate) Store(queue *queue.Queue) error {
	msgBytes := []byte(fmt.Sprintf("%v", r))
	queue.Enqueue(msgBytes)
	return nil
}
