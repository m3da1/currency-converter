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

func ProcessConversion(source, target string) (string, error) {
	// santizing currencies
	source = strings.TrimSpace(source)
	source = strings.ToUpper(source)
	target = strings.TrimSpace(target)
	target = strings.ToUpper(target)
	// validate request
	if validationErr := validateRequest(source, target); validationErr != nil {
		return "", validationErr
	}
	// get exchange rate
	return findExchangeRate(source, target)
}

func findExchangeRate(source, target string) (string, error) {
	var exchange model.ExchangeRate
	if err := model.DB.Where("source_ccy = ? and target_ccy = ?", source, target).Find(&exchange); err != nil {
		return "", fmt.Errorf("couldn't find exchange rate; reason: %+v", err)
	}
	return fmt.Sprintf("%v", exchange.Rate), nil
}

func supportedCurrencies() []string {
	return []string{nigeriaNaira, ghanaCedis, kenyanShilling}
}

func validateRequest(source, target string) error {
	// same source and target currencies
	if source == target {
		return fmt.Errorf("source and target currencies must not be identical")
	}
	// length check
	if sourceErr := lengthCheck(source); sourceErr != nil {
		return sourceErr
	}
	if targetErr := lengthCheck(target); targetErr != nil {
		return targetErr
	}
	// supported currency check
	if verifySourceErr := verifySupportedCurrency(source); verifySourceErr != nil {
		return verifySourceErr
	}
	if verifyTargetErr := verifySupportedCurrency(target); verifyTargetErr != nil {
		return verifyTargetErr
	}
	return nil
}

func lengthCheck(ccy string) error {
	if len(ccy) != 3 {
		return fmt.Errorf("invalid currency length: %+v", ccy)
	}
	return nil
}

func verifySupportedCurrency(ccy string) error {
	for _, v := range supportedCurrencies() {
		if v == ccy {
			return nil
		}
	}
	return fmt.Errorf("invalid currency: %+v", ccy)
}
