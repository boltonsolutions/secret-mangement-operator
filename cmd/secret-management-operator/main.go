package main

import (
	"context"
	"runtime"

	stub "github.com/boltonsolutions/secret-management-operator/pkg/stub"
	sdk "github.com/operator-framework/operator-sdk/pkg/sdk"
	sdkVersion "github.com/operator-framework/operator-sdk/version"

	"github.com/sirupsen/logrus"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

func printVersion() {
	logrus.Infof("Go Version: %s", runtime.Version())
	logrus.Infof("Go OS/Arch: %s/%s", runtime.GOOS, runtime.GOARCH)
	logrus.Infof("operator-sdk Version: %v", sdkVersion.Version)
}

func main() {
	printVersion()

	conf := stub.NewConfig()
	sdk.ExposeMetricsPort()

	logrus.Infof("Watching Secrets on all Namespaces")
	sdk.Watch("v1", "Secret", "", 1000000000)
	sdk.Handle(stub.NewHandler(conf))
	sdk.Run(context.TODO())
}
