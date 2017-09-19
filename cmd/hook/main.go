package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/zdq0394/registryhook/hook"
	"github.com/zdq0394/registryhook/hook/version"
)

func main() {
	app := cli.NewApp()
	app.Name = "HookServer"
	app.Version = version.Version.String()
	app.Flags = []cli.Flag{}
	app.Action = hookAction
	if err := app.Run(os.Args); err != nil {
		logrus.Errorln(err)
	}
}

func hookAction(ctx *cli.Context) error {
	hook.Start()
	return nil
}
