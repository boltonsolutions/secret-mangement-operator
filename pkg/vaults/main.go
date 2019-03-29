package vaults

import (
)

type Provider interface {
	Provision() (UsernamePassword, error)
	Deprovision() error
}

type UsernamePassword struct {
	Username []byte
	Password []byte
}

type ProviderConfig struct {
	Kind string `json:"kind"`
	Token string `json:"token"`
	Address string `json:"address"`
	Engine string `json:"engine"`
}