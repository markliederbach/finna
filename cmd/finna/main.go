package main

import (
	"os"
	"time"

	"github.com/markliederbach/finna/pkg/commands"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var (
	Version  = "latest"
	Commands = []*cli.Command{
		commands.NewFormatCommand().ToCliCommand(),
	}
)

const (
	// LogLevelFlag wraps the name of the command flag
	LogLevelFlag = "log-level"
)

func main() {
	app := &cli.App{
		Name:     "finna",
		Version:  Version,
		Usage:    "Financial analysis tool",
		Compiled: time.Now(),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    LogLevelFlag,
				Aliases: []string{"l"},
				Usage:   "Set the log level",
				Value:   "INFO",
				EnvVars: []string{"LOG_LEVEL"},
			},
		},
		Before: func(ctx *cli.Context) error {
			logrus.SetFormatter(&logrus.JSONFormatter{})
			level, err := logrus.ParseLevel(ctx.String(LogLevelFlag))
			if err != nil {
				return err
			}
			logrus.SetLevel(level)
			logrus.WithField("version", Version).Info("Running finna")
			return nil
		},
		Commands: Commands,
	}
	err := app.Run(os.Args)
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
	os.Exit(0)
}
