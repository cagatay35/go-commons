package configuration

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

const (
	profileEnvironmentKey = "PROFILE"
)

var (
	profile = os.Getenv(profileEnvironmentKey)
)

func ReadConfigurationFiles(logger *logrus.Logger) {

	if profile == "" {
		viper.SetConfigName("application")
	} else {
		viper.SetConfigName(fmt.Sprintf("application-%s", profile))
		logger.Infof("Application Profile: %s", profile)
	}

	viper.SetConfigType("yml")
	viper.AddConfigPath("./config/")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		logger.Errorf("fatal error config file: default %s", err)
		os.Exit(1)
	}
}
