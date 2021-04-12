package main

import (
	"os"

	"github.com/sirupsen/logrus"

	"github.com/erda-project/erda-actions/actions/dice-deploy/1.0/internal/dice-deploy"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors:    true,
		DisableTimestamp: true,
	})
	logrus.SetOutput(os.Stdout)
	logrus.Info("Deploying...")
	if err := dice_deploy.Run(); err != nil {
		logrus.Warning("Unable to cancel dice deploy, err: %v", err)
		os.Exit(1)
	}
}
