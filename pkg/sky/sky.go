package sky

import (
	"go.uber.org/zap"
	corev1informers "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"
)

var (
	Sky *Config
)

type Config struct {
	SkyConfig *SkyConfig
}

type SkyConfig struct {
	SkyAddr      string   `toml:"sky_addr" json:"sky_addr"`
	SkyPProfAddr string   `toml:"sky_pprof_addr" json:"sky_pprof_addr "`
	AlarmUrl     string   `toml:"alarm_url" json:"alarm_url"`
	IsDebug      bool     `toml:"is_debug" json:"is_debug"`
	LogLevel     string   `toml:"log_level" json:"log_level"`
	LogPath      string   `toml:"log_pah" json:"log_pah"`
	NameSpace    string   `toml:"namespace" json:"namespace"`
	Client *kubernetes.Clientset
	SugaredLogger *zap.SugaredLogger
	Stop         chan struct{}
	Informer     corev1informers.ConfigMapInformer
}

func (c *Config)Env()bool{
	return c.SkyConfig.IsDebug
}

func (c *Config)GetLogLevel()string{
	return c.SkyConfig.LogLevel
}

func (c *Config)GetLogPath()string {
	return c.SkyConfig.LogPath
}

func (c *Config)SetSugaredLogger(sugaredLogger *zap.SugaredLogger) {
	c.SkyConfig.SugaredLogger = sugaredLogger
	return
}

func (c *Config)GetSugaredLogger()*zap.SugaredLogger{
	return c.SkyConfig.SugaredLogger
}

func (c *Config)GetAlarmUrl()string{
	return c.SkyConfig.AlarmUrl
}

func (s *SkyConfig)OnAdd(object interface{}){

}


func (s *SkyConfig)OnUpdate(oldObject, newObject interface{}){

}


func (s *SkyConfig)OnDelete(object interface{}){

}