package currencyapi_test

import (
	"testing"

	"frontendmasters.com/cryptomasters/currencyapi"
)

func TestGetRate(t *testing.T) {
	_, err := currencyapi.GetRate("")
	if err == nil {
		t.Errorf("Empty currencies should return error")
	}
}

// func TestAPI(t *testing.T) {
// 	_, err := currencyapi.GetRate("BTC")
// 	if err == nil {
// 		t.Errorf("Empty currencies should return an error")
// 	}
// }
