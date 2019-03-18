
package stub

import (
	config "github.com/micro/go-config"
	"github.com/micro/go-config/source/env"
	"github.com/micro/go-config/source/file"
	"github.com/micro/go-config/source/flag"
	"github.com/micro/go-config/source/memory"
	"github.com/boltonsolutions/secret-management-operator/pkg/vaults"
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

	return conf
}

func getConfigFile() (configFile string) {
	return defaultConfigFile
}


