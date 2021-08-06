package activities

import (
	"log"

	"go-temporal-example/pkg/common/models"
)

// ReturnSomeJSON simple activity returns arbitrary JSON.
func ReturnSomeJSON() (*models.SomeJSON, error) {
	log.Println("ACTIVITY:: ReturnSomeJSON")
	return &models.SomeJSON{
		SomeProp: "someVal",
	}, nil
}
