package runner

import (
	"container/list"
	"context"
	"fmt"
	"os/exec"
	"time"
)

type Runner struct {
	TimeInterval int
	PlayList     []string
	pl           *list.List
}

func (r *Runner) RunVlc() error {
	cmd := exec.Command("xset", "s", "off")
	err := cmd.Run()
	if err != nil {
		return err
	}
	cmd = exec.Command("xset", "-dpms")
	err = cmd.Run()
	if err != nil {
		return err
	}
	cmd = exec.Command("xset", "s", "noblank")
	err = cmd.Run()
	if err != nil {
		return err
	}
	cmd = exec.Command("cvlc", "-f")
	err = cmd.Run()
	return err
}

func (r *Runner) RunSwitcher(ctx context.Context) error {
	r.pl = list.New()
	for _, v := range r.PlayList {
		r.pl.PushBack(v)
	}
	return r.runSwitcher(ctx)
}

func (r *Runner) runSwitcher(ctx context.Context) error {
	time.Sleep(1 * time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	fmt.Printf("Playing %s\n", r.pl.Front().Value.(string))
	err := r.PlayVideo(r.pl.Front().Value.(string))
	if err != nil {
		return err
	}
	sleepWithContext(ctx, time.Duration(r.TimeInterval)*time.Second)
	r.pl.PushBack(r.pl.Remove(r.pl.Front()))
	return r.runSwitcher(ctx)
}

func (r *Runner) PlayVideo(url string) error {
	cmd := exec.Command("playerctl", "-p", "vlc", "open", url)
	err := cmd.Run()
	return err
}

func sleepWithContext(ctx context.Context, duration time.Duration) {
	timer := time.NewTimer(duration)
	select {
	case <-ctx.Done():
		if !timer.Stop() {
			<-timer.C
		}
	case <-timer.C:
	}
}
