package util

import (
	"github.com/subosito/gotenv"
)

func LoadVars() error {
	err1 := gotenv.Load("./config/env")
	err2 := gotenv.Load("../../config/env")
	if err1 != nil && err2 != nil {
		return err1
	}
	return nil
}
