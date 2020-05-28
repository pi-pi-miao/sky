package server

import (
	"sky/pkg/sky"
)

func Sky(path string)error{
	if err := sky.InitAll(path);err != nil {
		return err
	}
	return nil
}


