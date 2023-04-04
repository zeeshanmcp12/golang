package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ExchangeRatesResponse struct {
	Base  string             `json:"base"`
	Rates map[string]float64 `json:"rates"`
	Date  string             `json:"date"`
}

func convertUSDToPKR(amount float64) (float64, error) {
	apiKey := "<apikey here>"
	// url := "https://api.exchangeratesapi.io/latest?base=USD&symbols=PKR"
	url := "https://api.apilayer.com/exchangerates_data/latest?symbols=PKR&base=USD"
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("apikey", apiKey)

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("unexpected status code %d", response.StatusCode)
	}

	var exchangeRatesResponse ExchangeRatesResponse
	err = json.NewDecoder(response.Body).Decode(&exchangeRatesResponse)
	if err != nil {
		return 0, err
	}

	rate, ok := exchangeRatesResponse.Rates["PKR"]
	if !ok {
		return 0, fmt.Errorf("exchange rate not found")
	}

	return amount * rate, nil
}

func main() {
	amount := 100.00
	pkrAmount, err := convertUSDToPKR(amount)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%.2f USD = %.2f PKR\n", amount, pkrAmount)
}
