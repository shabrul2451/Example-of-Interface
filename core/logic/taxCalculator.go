package logic

import (
	"github.com/interfaceExample/ExampleOfInterface/core/service"
	"log"
)

type taxCalculator struct {
}

func (t taxCalculator) CalculateTax(amount float64) float64 {
	log.Println("Calculating tax for amount: ", amount)
	return amount * 0.2
}

func NewTaxCalculator() service.TaxCalculator {
	return &taxCalculator{}
}
