package vaults

import (
	"time"
	api "github.com/hashicorp/vault/api"
	"fmt"
)

type HashiCorpProvider struct {
}

func (p *HashiCorpProvider) Provision() (keypair KeyPair, certError error) {

	fmt.Printf("Hashi Corp Start")
	client, err := api.NewClient(&api.Config{
	    Address: "http://127.0.0.1:8200",
	})

	if err != nil {
		fmt.Printf("error...")
	}

	client.SetToken("s.gZFIyinkDU8z3MB8vHwEjZoW")

	secretValues, err := client.Logical().Read("secret/data/com/bolton")
	if err != nil {
		fmt.Printf("error...")
	}

	fmt.Printf("secret %s -> %v", "username", secretValues)
	
	return KeyPair{
		Secret:   []byte{},
		Key:    []byte{},
		Expiry: time.Now(),
	}, nil
}

func (p *HashiCorpProvider) Deprovision(host string) error {
	return nil
}