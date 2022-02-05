package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func createApp() *cli.App {
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:    "log-level",
			Usage:   "log-level \"debug\" (more on the supported levels here: https://github.com/sirupsen/logrus/blob/fdf1618bf7436ec3ee65753a6e2999c335e97221/logrus.go#L25)",
			Value:   "debug",
			EnvVars: []string{"LOG_LEVEL"},
		},
	}

	return &cli.App{
		Name:  "drain-my-spot",
		Usage: "drain-my-spot is aan app that sends the drain request to k8s API when worker node is scheduled for preemtion",
		Flags: flags,
		Action: func(c *cli.Context) error {
			return runAction(c, func(params *Parameters) error {
				return nil
			})
		},
	}
}

//Parameters struct that conveys the main parameters of the app
type Parameters struct {
}

func runAction(c *cli.Context, f func(params *Parameters) error) error {
	log.SetFormatter(&log.JSONFormatter{})
	lvl, err := log.ParseLevel(c.String("log-level"))
	if err != nil {
		log.Errorf("unable to parse log level (%s): %v",
			c.String("log-level"), err)
		return err
	}
	log.SetLevel(lvl)

	return f(&Parameters{})
}
