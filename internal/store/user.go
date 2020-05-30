package store

import (
	"errors"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/util/retry"
	"sky/pkg/sky"
)

func CreateUser(email string, data []byte) (err error) {
	user := &corev1.ConfigMap{}
	if user, err = sky.Sky.SkyConfig.Informer.Lister().ConfigMaps(sky.Sky.SkyConfig.NameSpace).Get(sky.Sky.SkyConfig.User.UserStoreName); err != nil {
		return err
	}

	if user.BinaryData == nil {
		m := make(map[string][]byte)
		m[email] = data
	} else {
		user.BinaryData[email] = data
	}
	err = retry.RetryOnConflict(retry.DefaultRetry, func() error {
		_, err := sky.Sky.SkyConfig.Client.CoreV1().ConfigMaps(sky.Sky.SkyConfig.NameSpace).Update(user)
		return err
	})
	return
}

func GetUser(email string) (data []byte, err error) {
	user := &corev1.ConfigMap{}
	if user, err = sky.Sky.SkyConfig.Informer.Lister().ConfigMaps(sky.Sky.SkyConfig.NameSpace).Get(sky.Sky.SkyConfig.User.UserStoreName); err != nil {
		return
	}
	if user.BinaryData == nil {
		return nil, errors.New("Illegal user")
	}
	err = retry.RetryOnConflict(retry.DefaultRetry, func() error {
		user, err = sky.Sky.SkyConfig.Informer.Lister().ConfigMaps(sky.Sky.SkyConfig.NameSpace).Get(sky.Sky.SkyConfig.User.UserStoreName)
		return err
	})
	data = user.BinaryData[email]
	return
}
