
package stub

import (
	"github.com/boltonsolutions/secret-management-operator/pkg/vaults"
	"encoding/json"
	"os"
	"fmt"
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
	var config Config
    configFile, err := os.Open("/etc/secret-management-operator/config.yaml")
    defer configFile.Close()
    if err != nil {
        fmt.Println(err.Error())
    }
    jsonParser := json.NewDecoder(configFile)
    jsonParser.Decode(&config)
    return config
}

func (c *Config) String() string {
	out, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}
	return string(out)
}


