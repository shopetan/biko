package cli

import (
	"github.com/KeisukeYamashita/biko/browser"
	"github.com/KeisukeYamashita/biko/providers/youtube"
	"github.com/urfave/cli"
)

func newYoutubeCmd() cli.Command {
	return cli.Command{
		Name:    "youtube",
		Aliases: []string{"yt"},
		Usage:   "Open Youtube source",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "project",
				Usage: "Specify the project to open",
			},
		},
		Action: func(c *cli.Context) error {
			yt, err := youtube.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, yt)
		},
		Subcommands: []cli.Command{
			newYoutubeSearchCmd(),
		},
	}
}

func newYoutubeSearchCmd() cli.Command {
	return cli.Command{
		Name:    "search",
		Aliases: []string{"s"},
		Usage:   "Search a page",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "query, q",
				Usage: "Query a page",
			},
		},
		Action: func(c *cli.Context) error {
			yt, err := youtube.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, yt)
		},
	}

}
