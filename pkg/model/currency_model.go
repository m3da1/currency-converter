package model

// Exchange Rate database object
type ExchangeRate struct {
	ID        uint `gorm:"primary_key"`
	SourceCcy string
	TargetCcy string
	Rate      float64
}
