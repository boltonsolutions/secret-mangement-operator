
package stub

import (
	config "github.com/micro/go-config"
	"github.com/micro/go-config/source/env"
	"github.com/micro/go-config/source/file"
	"github.com/micro/go-config/source/flag"
	"github.com/micro/go-config/source/memory"
	"github.com/boltonsolutions/secret-management-operator/pkg/vaults"
	"github.com/sirupsen/logrus"
	"encoding/json"
)

type Config struct {
	Provider vaults.ProviderConfig `json:"provider"`
	General  GeneralConfig        `json:"general"`
}

type GeneralConfig struct {
	Annotations AnnotationConfig `json:"annotations"`
}

type AnnotationConfig struct {
	Status        string `json:"status"`
}

const (
	defaultConfigFile = "/etc/secret-management-operator/config.yaml"
	defaultProvider   = "hashicorp"
	defaultConfig = `
  {
    "general": {
      "annotations": {
        "status": "openshift.io/secret-status"
      }
    },
    "provider": {
      "kind": "hashicorp"
    }
  }`
)

func NewConfig() Config {

	tmpConfig := config.NewConfig()

	data := []byte(defaultConfig)

	memorySource := memory.NewSource(
		memory.WithData(data),
	)
	// Load json config file
	tmpConfig.Load(
		memorySource,
		file.NewSource(
			file.WithPath(getConfigFile()),
		),
		env.NewSource(),
		flag.NewSource(),
	)
	var conf Config

	tmpConfig.Scan(&conf)
	logrus.Infof(conf.String())
	return conf
}

func getConfigFile() (configFile string) {
	logrus.Infof("Loading default config file from %v", defaultConfigFile)
	return defaultConfigFile
}

func (c *Config) String() string {
	out, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}
	return string(out)
}


