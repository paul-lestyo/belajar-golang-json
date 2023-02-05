package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {
	customer := Customer{
		FirstName:  "Paulus",
		MiddleName: "Lestyo",
		LastName:   "Adhiatma",
		Age:        20,
		Married:    false,
		Hobbies:    []string{"Gaming", "Reading", "Coding"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Paulus","MiddleName":"Lestyo","LastName":"Adhiatma","Age":20,"Married":false,"Hobbies":["Gaming","Reading","Coding"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)

}

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName:  "Paulus",
		MiddleName: "Lestyo",
		LastName:   "Adhiatma",
		Age:        20,
		Married:    false,
		Hobbies:    []string{"Gaming", "Reading"},
		Addresses: []Address{
			{
				Street:     "Jalan Belum Ada",
				Country:    "Indonesia",
				PostalCode: "57155",
			},
			{
				Street:     "Jalan Tidak tahu juga",
				Country:    "Indonesia",
				PostalCode: "888888",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))

}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Paulus","MiddleName":"Lestyo","LastName":"Adhiatma","Age":20,"Married":false,"Hobbies":["Gaming","Reading"],"Addresses":[{"Street":"Jalan Belum Ada","Country":"Indonesia","PostalCode":"57155"},{"Street":"Jalan Tidak tahu juga","Country":"Indonesia","PostalCode":"888888"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Addresses)

}

func TestOnlyJSONArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"Jalan Belum Ada","Country":"Indonesia","PostalCode":"57155"},{"Street":"Jalan Tidak tahu juga","Country":"Indonesia","PostalCode":"888888"}]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}
	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(addresses)

}

func TestOnlyJSONArrayComplex(t *testing.T) {
	customer := []Address{
		{
			Street:     "Jalan Belum Ada",
			Country:    "Indonesia",
			PostalCode: "57155",
		},
		{
			Street:     "Jalan Tidak tahu juga",
			Country:    "Indonesia",
			PostalCode: "888888",
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))

}
