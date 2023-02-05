package belajar_golang_json

import (
	"encoding/json"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("customerOut.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "Paulus",
		MiddleName: "Lestyo",
		LastName:   "Adhiatma",
	}

	encoder.Encode(customer)

}
