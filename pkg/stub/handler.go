package stub

import (
	"context"

	"github.com/operator-framework/operator-sdk/pkg/sdk"
	"github.com/sirupsen/logrus"
	"github.com/boltonsolutions/secret-management-operator/pkg/vaults"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NewHandler(config Config) sdk.Handler {
	var provider vaults.Provider
	
	switch config.Provider.Kind {
		case "hashicorp":
			logrus.Infof("Hashi Corp Provider.")
			provider = new(vaults.HashiCorpProvider)
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

func (h *Handler) handleSecret(secret *corev1.Secret) (vaults.UsernamePassword, error) {

	if secret.ObjectMeta.Annotations == nil || secret.ObjectMeta.Annotations[h.config.General.Annotations.Status] == "" {
		return vaults.UsernamePassword{}, nil
	}

	if secret.ObjectMeta.Annotations[h.config.General.Annotations.Status] == "need" {
		logrus.Infof("We need a secret from the vault.")
		usernamePassword, err := h.provider.Provision()

		logrus.Infof("username = %s", usernamePassword.Username)
		if err != nil {
			return vaults.UsernamePassword{}, err
		}

		dm := make(map[string][]byte)
		dm["username"] = usernamePassword.Username
		dm["password"] = usernamePassword.Password
		secret := &corev1.Secret{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Secret",
				APIVersion: "v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      secret.ObjectMeta.Name,
				Namespace: secret.ObjectMeta.Namespace,
			},
			Data: dm,
		}

		//We must delete and recreate...
		sdk.Delete(secret)
		err = sdk.Create(secret)
		if err != nil {
			logrus.Errorf("Failed to create secret: " + err.Error())
			// return vaults.UsernamePassword{}, err
		}

		return usernamePassword, nil
	}
	return vaults.UsernamePassword{}, nil
}


