package stub

import (
	"context"

	"github.com/operator-framework/operator-sdk/pkg/sdk"
	"github.com/sirupsen/logrus"
	"github.com/boltonsolutions/secret-management-operator/pkg/vaults"
	corev1 "k8s.io/api/core/v1"
)

func NewHandler(config Config) sdk.Handler {
	var provider vault.Provider

	switch config.Provider.Kind {
		case "hashicorp":
			logrus.Infof("Hashi Corp Provider.")
			provider = new(vault.HashiCorpProvider)
		default:
			panic("Well that didn't work.")
	}

	return &Handler{
		config:   config,
		provider: provider,
	}
}

type Handler struct {
	config   Config
	provider vaults.Provider
}

func (h *Handler) Handle(ctx context.Context, event sdk.Event) error {
	switch o := event.Object.(type) {
	case *corev1.Secret:
		h.handleSecret(o)
	}
	return nil
}

func (h *Handler) handleSecret(secret *corev1.Secret) error {

	if route.ObjectMeta.Annotations == nil || route.ObjectMeta.Annotations[h.config.General.Annotations.Status] == "" {
		return nil
	}

	if route.ObjectMeta.Annotations[h.config.General.Annotations.Status] == "need" {
		logrus.Infof("We need a secret from the vault.")
	}
}


