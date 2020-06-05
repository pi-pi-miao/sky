package server

import (
	"sky/apis"
	"sky/pkg/sky"
)

func Sky(config,path string) error {
	apis.Apis()
	if err := sky.InitAll(config,path); err != nil {
		return err
	}
	return nil
}
