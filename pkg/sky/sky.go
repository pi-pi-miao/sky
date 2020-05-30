package sky

import (
	"go.uber.org/zap"
	corev1informers "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

var (
	Sky *Config
)

/**
* todo ip current-limiting
* todo jwt
*
*/

type Config struct {
	SkyConfig  *SkyConfig            `toml:"sky_config" json:"sky_config"`
	UserLabels map[string]UserLabels `toml:"user_labels"`
	User       *User                 `toml:"user"`
	Code       *Code                 `toml:"code_error"`
}

type SkyConfig struct {
	SkyAddr      string `toml:"sky_addr" json:"sky_addr"`
	SkyPProfAddr string `toml:"sky_pprof_addr" json:"sky_pprof_addr "`
	AlarmUrl     string `toml:"alarm_url" json:"alarm_url"`
	IsDebug      bool   `toml:"is_debug" json:"is_debug"`
	LogLevel     string `toml:"log_level" json:"log_level"`
	LogPath      string `toml:"log_pah" json:"log_pah"`
	NameSpace    string `toml:"namespace" json:"namespace"`
	*Config

	Client            *kubernetes.Clientset
	SugaredLogger     *zap.SugaredLogger
	Stop              chan struct{}
	Informer          corev1informers.ConfigMapInformer
	NamespaceInformer corev1informers.NamespaceInformer
	Informers         []cache.InformerSynced
}

type Code struct {
	RequestError   string `toml:"request_error"`
	RequestSuccess string `toml:"request_success"`
	IllegalRequest string `toml:"illegal_request"`
}

type User struct {
	UserStoreName string `toml:"user_store_name"`
	UserVersion   string `toml:"user_version"`
	UserKind      string `toml:"user_kind"`
}

type UserLabels struct {
	Value string `toml:"value"`
}

func (c *Config) Env() bool {
	return c.SkyConfig.IsDebug
}

func (c *Config) GetLogLevel() string {
	return c.SkyConfig.LogLevel
}

func (c *Config) GetLogPath() string {
	return c.SkyConfig.LogPath
}

func (c *Config) SetSugaredLogger(sugaredLogger *zap.SugaredLogger) {
	c.SkyConfig.SugaredLogger = sugaredLogger
	return
}

func (c *Config) GetSugaredLogger() *zap.SugaredLogger {
	return c.SkyConfig.SugaredLogger
}

func (c *Config) GetAlarmUrl() string {
	return c.SkyConfig.AlarmUrl
}

func (s *SkyConfig) OnAdd(object interface{}) {

}

func (s *SkyConfig) OnUpdate(oldObject, newObject interface{}) {

}

func (s *SkyConfig) OnDelete(object interface{}) {

}
