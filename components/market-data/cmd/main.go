package main

import (
	nomics "github.com/giannoul/microtrader/internal/pkg/nomics"
	rsource "github.com/giannoul/microtrader/internal/pkg/rsource"
	"log"
	"time"
)

func main() {
	for {
		sources := []rsource.RatesSource{nomics.Nomics{}}
		for _, src := range sources {
			data, err := src.GetExchangeRates()
			if err != nil {
				log.Fatalln(err)
			}
			src.StoreExchangeRates(data)
		}

		time.Sleep(30 * time.Second)
	}

}
