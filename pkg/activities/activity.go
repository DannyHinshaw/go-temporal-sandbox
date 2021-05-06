package activities

import (
	"fmt"
	"go-temporal-example/app/pkg/common"
)

// ReturnNonSerializableJSON never returns an error, but returns BadJSON which will cause an error in Temporal.
func ReturnNonSerializableJSON() (*common.BadJSON, error) {
	return &common.BadJSON{
		SomeProp: "",
		Error:    fmt.Errorf("test error"),
	}, nil
}
