package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"strings"
)

type Config struct {
	Name string
}

func (c *Config) init() error {
	if c.Name != "" {
		viper.SetConfigName(c.Name)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath("conf")
	}
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("GOTE")
	replacer := strings.NewReplacer(".", "-")
	viper.SetEnvKeyReplacer(replacer)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Printf("Config file changed: %s", in.Name)
	})
}

func InitConfig(name string) error {
	c := Config{Name: name}

	if err := c.init(); err != nil {
		return err
	}
	c.watchConfig()

	return nil
}
