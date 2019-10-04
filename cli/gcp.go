package cli

import (
	"github.com/KeisukeYamashita/biko/browser"
	"github.com/KeisukeYamashita/biko/providers/gcp"
	"github.com/urfave/cli"
)

func newGCPCmd() cli.Command {
	return cli.Command{
		Name:  "gcp",
		Usage: "Open GCP resource",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "project",
				Usage: "Specify the project to open",
			},
		},
		Action: func(c *cli.Context) error {
			gcp := &gcp.Provider{}
			return browser.Open(c, gcp)
		},
		Subcommands: []cli.Command{
			newGCPGAECmd(),
			newGCPBQCmd(),
			newGCPGKECmd(),
			newGCPSpannerCmd(),
			newGCPGCRCmd(),
			newGCPFunctionsCmd(),
			newGCPCloudRunCmd(),
			newGCPGCECmd(),
		},
	}
}

func newGCPGAECmd() cli.Command {
	return cli.Command{
		Name:    "appengine",
		Aliases: []string{"gae"},
		Usage:   "Open GAE page",
		Flags:   []cli.Flag{},
		Action: func(c *cli.Context) error {
			gcp := &gcp.Provider{}
			return browser.Open(c, gcp)
		},
	}
}

func newGCPBQCmd() cli.Command {
	return cli.Command{
		Name:    "biqquery",
		Aliases: []string{"bq"},
		Usage:   "Open Bigquery page",
		Flags:   []cli.Flag{},
		Action: func(c *cli.Context) error {
			gcp := &gcp.Provider{}
			return browser.Open(c, gcp)
		},
	}
}

func newGCPGKECmd() cli.Command {
	return cli.Command{
		Name:    "kubernetes",
		Aliases: []string{"gke"},
		Usage:   "Open GKE page",
		Flags:   []cli.Flag{},
		Action: func(c *cli.Context) error {
			gcp := &gcp.Provider{}
			return browser.Open(c, gcp)
		},
	}
}

func newGCPSpannerCmd() cli.Command {
	return cli.Command{
		Name:  "spanner",
		Usage: "Open Cloud Spanner page",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "project",
				Usage: "Specify the project to open",
			},
			cli.StringFlag{
				Name:  "instance",
				Usage: "Instance name",
			},
			cli.StringFlag{
				Name:  "database",
				Usage: "Database name",
			},
		},
		Action: func(c *cli.Context) error {
			gcp := &gcp.Provider{}
			return browser.Open(c, gcp)
		},
	}
}

func newGCPGCRCmd() cli.Command {
	return cli.Command{
		Name:  "gcr",
		Usage: "Open Cloud Registry page",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			gcp := &gcp.Provider{}
			return browser.Open(c, gcp)
		},
	}
}

func newGCPFunctionsCmd() cli.Command {
	return cli.Command{
		Name:  "functions",
		Usage: "Open Cloud Functions page",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			gcp := &gcp.Provider{}
			return browser.Open(c, gcp)
		},
	}
}

func newGCPCloudRunCmd() cli.Command {
	return cli.Command{
		Name:  "run",
		Usage: "Open Cloud Run page",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			gcp := &gcp.Provider{}
			return browser.Open(c, gcp)
		},
	}
}

func newGCPGCECmd() cli.Command {
	return cli.Command{
		Name:    "compute",
		Aliases: []string{"gce"},
		Usage:   "Open Compute Engine page",
		Flags:   []cli.Flag{},
		Action: func(c *cli.Context) error {
			gcp := &gcp.Provider{}
			return browser.Open(c, gcp)
		},
	}
}
