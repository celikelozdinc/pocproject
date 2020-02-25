package main

import (
	"fmt"
	"pocproject/entity"
)

func main() {

	fmt.Println("..Main package...")

	// Do not create a value of struct
	// Use address of our value
	PtrToPhoneRecord := &entity.PhoneRecord{
		Phone:   "05332793271",
		Name:    "Simge",
		Surname: "Alkin",
		Country: "France",
		ID:      41,
	}

	PtrToPhoneRecord.Printer()

}
