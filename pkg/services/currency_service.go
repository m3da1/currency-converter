package services

import (
	"currency-converter/pkg/model"
	"fmt"
	"strings"
)

const (
	nigeriaNaira   = "NGN"
	ghanaCedis     = "GHS"
	kenyanShilling = "KSH"
)

// List of supported currency codes
func supportedCurrencies() []string {
	return []string{nigeriaNaira, ghanaCedis, kenyanShilling}
}

// Process exchange rate conversion
func ProcessConversion(source, target string) (string, error) {
	// sanitizing input currency
	source = strings.TrimSpace(source)
	source = strings.ToUpper(source)
	target = strings.TrimSpace(target)
	target = strings.ToUpper(target)
	// validate request
	if validationErr := validateRequest(&source, &target); validationErr != nil {
		return "", validationErr
	}
	// get exchange rate
	return findExchangeRate(&source, &target)
}

// Performs validation of currency values
func validateRequest(source, target *string) error {
	// supported currency check
	if verifySourceErr := validateCurrency(source); verifySourceErr != nil {
		return verifySourceErr
	}
	if verifyTargetErr := validateCurrency(target); verifyTargetErr != nil {
		return verifyTargetErr
	}
	return nil
}

// validate currency code length and check whether it is supported
func validateCurrency(ccy *string) error {
	// perform validation of the currency code length
	if len(*ccy) != 3 {
		return fmt.Errorf("invalid currency length: %+v", *ccy)
	}
	// check whether currency is supported
	for _, v := range supportedCurrencies() {
		if v == *ccy {
			return nil
		}
	}
	return fmt.Errorf("invalid currency: %+v", *ccy)
}

// Get exchange rate from local conversion table
func findExchangeRate(source, target *string) (string, error) {
	var exchange model.ExchangeRate
	if err := model.DB.Where("source_ccy = ? and target_ccy = ?", source, target).Find(&exchange).Error; err != nil {
		return "", fmt.Errorf("couldn't find exchange rate; reason: %+v", err)
	}
	return fmt.Sprintf("%.3f", exchange.Rate), nil
}
