package service

type TaxCalculator interface {
	CalculateTax(amount float64) float64
}
