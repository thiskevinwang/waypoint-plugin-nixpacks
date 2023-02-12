package builder

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/waypoint-plugin-sdk/component"
	"github.com/hashicorp/waypoint-plugin-sdk/terminal"
)

type BuildConfig struct {
	Source string `hcl:"source,optional"`
	// linux/amd64 (default) | linux/arm64
	Platform string `hc:"platform,optional"`
}

type Builder struct {
	config BuildConfig
}

// Implement Configurable
func (b *Builder) Config() (interface{}, error) {
	return &b.config, nil
}

// ConfigSet is called after a configuration has been decoded
// we can use this to validate the config
func (b *Builder) ConfigSet(config interface{}) error {
	// TODO(kevin): validate config
	return nil
}

// Implement Builder
func (b *Builder) BuildFunc() interface{} {
	// return a function which will be called by Waypoint
	return b.Build
}

const (
	DEFAULT_PLATFORM = "linux/amd64"
)

func (b *Builder) Build(
	ctx context.Context,
	ui terminal.UI,
	src *component.Source,
	log hclog.Logger,
) (*Image, error) {
	u := ui.Status()
	defer u.Close()
	u.Update("Building application")

	// set config defaults
	if b.config.Source == "" {
		// this should be the directory where the waypoint.hcl file is located
		b.config.Source = src.Path
	}

	if b.config.Platform == "" {
		b.config.Platform = DEFAULT_PLATFORM
	}

	log.Info("config values", "source", b.config.Source, "platform", b.config.Platform)

	// check for nixpack
	nixpacksPath, err := exec.LookPath("nixpacks")
	if err != nil {
		u.Step(terminal.StatusError, fmt.Sprintf("Nixpacks not found: %s", err))
		return nil, err
	}

	u.Step(terminal.StatusOK, fmt.Sprintf("Nixpacks installed at: %q", nixpacksPath))

	cmd := exec.Command(nixpacksPath,
		"build", b.config.Source,
		"--platform", b.config.Platform,
		"--name", fmt.Sprintf("waypoint.local/%s", src.App))

	sg := ui.StepGroup()
	step := sg.Add("Building app with nixpacks...")
	defer func() {
		step.Abort()
	}()

	// pipe nixpacks output to waypoint output???
	cmd.Stdout = step.TermOutput()

	if err := cmd.Run(); err != nil {
		u.Step(terminal.StatusError, fmt.Sprintf("Nixpacks failed to build: %s", err))
		return nil, err
	}

	step.Update("Nixpacks build succeeded")
	step.Done()

	return &Image{
		Image: fmt.Sprintf("waypoint.local/%s", src.App),
		Tag:   "latest",
	}, nil
}
