package vaults

import (
	"time"
)

type Provider interface {
	Provision() (KeyPair, error)
	Deprovision(host string) error
}

type KeyPair struct {
	Secret   []byte
	Key    []byte
	Expiry time.Time
}

type ProviderConfig struct {
	Kind string `json:"kind"`
	Token string `json:"token"`
	Address string `json:"address"`
	Engine string `json:"engine"`
}