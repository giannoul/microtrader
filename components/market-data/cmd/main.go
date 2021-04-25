package main

import (
	nomics "github.com/giannoul/microtrader/internal/pkg/nomics"
	rsource "github.com/giannoul/microtrader/internal/pkg/rsource"
	"log"
)

func main() {
	n1 := nomics.Nomics{}
	n1.SetupQueue("hello1")
	n2 := nomics.Nomics{}
	n2.SetupQueue("byebye")

	sources := []rsource.RatesSource{&n1, &n2}
	for _, src := range sources {
		data, err := src.GetExchangeRates()
		if err != nil {
			log.Fatalln(err)
		}
		src.StoreExchangeRates(data)
		src.TeardownQueue()
	}

}
