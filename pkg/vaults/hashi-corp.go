package vaults

import (
	"time"
)

type HashiCorpProvider struct {
}

func (p *HashiCorpProvider) Provision() (keypair KeyPair, certError error) {
	return KeyPair{
		Secret:   []byte{},
		Key:    []byte{},
		Expiry: time.Now(),
	}, nil
}

func (p *HashiCorpProvider) Deprovision(host string) error {
	return nil
}