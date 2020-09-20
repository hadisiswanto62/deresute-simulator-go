package parser

import (
	"fmt"

	"github.com/hadisiswanto62/deresute-simulator-go/jsonmodels"
)

type dataInitializer interface {
	InitData() error
}

func InitData(dp dataInitializer) error {
	if err := dp.InitData(); err != nil {
		return fmt.Errorf("could not initialize data: %v", err)
	}
	return nil
}

var _ dataInitializer = jsonmodels.JSONDataParser{}
