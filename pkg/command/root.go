package command

import (
	"bufio"
	"context"
	"embed"
	"fmt"
	"github.com/dragonchaser/yt-kiosk/pkg/runner"
	"github.com/oklog/run"
	"github.com/urfave/cli/v2"
	"os"
)

var (
	TimeInterval = 60
	//go:embed assets/*
	res embed.FS
)

func Execute() error {
	app := &cli.App{
		Name:  "yt-kiosk",
		Usage: "Run the yt-kiosk",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "playlist",
				Aliases:     []string{"p"},
				Usage:       "The playlist to play",
				DefaultText: "Embedded Playlist",
			},
			&cli.IntFlag{
				Name:        "interval",
				Aliases:     []string{"i"},
				Usage:       "The time interval to switch videos",
				DefaultText: "30 seconds",
			},
		},
		Action: func(c *cli.Context) error {
			r := &runner.Runner{
				TimeInterval: TimeInterval,
			}
			if c.IsSet("playlist") {
				f, err := os.Open(c.String("playlist"))
				if err != nil {
					return err
				}
				defer f.Close()
				scanner := bufio.NewScanner(f)
				for scanner.Scan() {
					r.PlayList = append(r.PlayList, scanner.Text())
				}
			} else {
				f, err := res.Open("assets/playlist.txt")
				if err != nil {
					return err
				}
				defer f.Close()
				scanner := bufio.NewScanner(f)
				for scanner.Scan() {
					r.PlayList = append(r.PlayList, scanner.Text())
				}
			}
			if c.IsSet("interval") {
				r.TimeInterval = c.Int("interval")
			}
			fmt.Printf("Running yt-kiosk with %v items and %vs interval\n", len(r.PlayList), r.TimeInterval)
			return Run(r)
		},
	}
	return app.Run(os.Args)
}

func Run(r *runner.Runner) error {
	ctx, cancel := func() (context.Context, context.CancelFunc) {
		return context.WithCancel(context.Background())
	}()

	defer cancel()
	err := make(chan error)
	gr := run.Group{}
	gr.Add(func() error {
		select {
		case <-ctx.Done():
			return nil

		case err <- r.RunSwitcher(ctx):
			return <-err
		}
	}, func(_ error) {
		fmt.Println("Shutting down switcher")
		cancel()
	})

	gr.Add(r.RunVlc, func(_ error) {
		fmt.Println("Shutting down cvlc")
		cancel()
	})

	return gr.Run()
}
