package util

import (
	"errors"

	"github.com/subosito/gotenv"
)

func LoadVars() error {
	if gotenv.Load("./config/env") != nil && gotenv.Load("../../config/env") != nil && gotenv.Load("../../../config/env") != nil {
		return errors.New("env not found")
	}
	return nil
}
