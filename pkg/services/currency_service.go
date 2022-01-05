package services

import (
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
	// convert
	switch strings.ToUpper(source) {
	case nigeriaNaira:
		return "1.0", nil
	case ghanaCedis:
		return "2.0", nil
	case kenyanShilling:
		return "3.0", nil
	default:
		return "", nil
	}
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
