package nomics

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	exrate "github.com/giannoul/microtrader/internal/pkg/messages"
)

// Nomics the struct for this source
type Nomics struct{}

type nomicsRate struct {
	Currency  string `json:"currency"`
	Rate      string `json:"rate"`
	Timestamp string `json:"timestamp"`
}

// GetExchangeRates implements the GetExchangeRates interface from source
func (n Nomics) GetExchangeRates() ([]byte, error) {
	url := fmt.Sprintf("https://api.nomics.com/v1/exchange-rates?key=%s", os.Getenv("NOMICS_API_KEY"))
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return body, nil
}

// StoreExchangeRates implements the GetExchangeRates interface from source
func (n Nomics) StoreExchangeRates(data []byte) error {
	rates, err := convertRateToMessage(data)
	if err != nil {
		log.Println(err)
		return err
	}
	for _, r := range rates {
		ss := nomicsRateMapper(r)
		ss.Store()
	}
	return nil
}

func nomicsRateMapper(r nomicsRate) exrate.Rate {
	parsedRate, err := strconv.ParseFloat(r.Rate, 64)
	if err != nil {
		log.Println(err)
	}
	return exrate.Rate{
		Source:    "NOMICS",
		Currency:  r.Currency,
		Rate:      parsedRate,
		Timestamp: r.Timestamp,
	}
}

func convertRateToMessage(jsonData []byte) ([]nomicsRate, error) {
	var rates []nomicsRate
	err := json.Unmarshal(jsonData, &rates)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return rates, nil
}
