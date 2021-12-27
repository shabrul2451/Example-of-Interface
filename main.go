package main

import (
	"github.com/interfaceExample/ExampleOfInterface/core/logic"
	"log"
)

func main() {
	test := logic.NewTaxCalculator()
	t := test.CalculateTax(10)
	log.Println("tax Amount", t)
}
