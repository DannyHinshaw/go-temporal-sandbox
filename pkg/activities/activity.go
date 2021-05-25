package activities

import (
	"go-temporal-example/app/pkg/common"
	"log"
)

// ReturnSomeJSON simple activity returns arbitrary JSON.
func ReturnSomeJSON() (*common.SomeJSON, error) {
	log.Println("ACTIVITY:: ReturnSomeJSON")
	return &common.SomeJSON{
		SomeProp: "someVal",
	}, nil
}
