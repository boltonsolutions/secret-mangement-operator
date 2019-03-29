package vaults

import (
	api "github.com/hashicorp/vault/api"
	"fmt"
)

type HashiCorpProvider struct {
}

func (p *HashiCorpProvider) Provision() (usernamePassword UsernamePassword, certError error) {

	fmt.Printf("Hashi Corp Start\n")
	client, err := api.NewClient(&api.Config{
	    Address: "http://127.0.0.1:8200",
	})

	if err != nil {
		fmt.Printf("error...")
	}

	client.SetToken("s.P7WHoiC2rYz4Hvu5WdXHRnhB")

	secretValues, err := client.Logical().Read("secret/data/com/bolton")
	if err != nil {
		fmt.Printf("error...")
	}

	var username []byte
	var password []byte

	for _, record := range secretValues.Data {

	    if rec, ok := record.(map[string]interface{}); ok {
	        for key, value := range rec {
	        	if key == "username" {
	        		username = []byte(value.(string))
	        	}
	        	if key == "password" {
	        		password = []byte(value.(string))
	        	}
	        }
	    }
	}
	
	return UsernamePassword{
		username,
		password }, nil
}

func (p *HashiCorpProvider) Deprovision() error {
	return nil
}