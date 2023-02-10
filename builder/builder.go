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
}

type Builder struct {
	config BuildConfig
}

// Implement Configurable
func (b *Builder) Config() (interface{}, error) {
	return &b.config, nil
}

// Implement ConfigurableNotify
func (b *Builder) ConfigSet(config interface{}) error {
	// TODO(kevin): validate config
	return nil
}

// Implement Builder
func (b *Builder) BuildFunc() interface{} {
	// return a function which will be called by Waypoint
	return b.Build
}

func (b *Builder) Build(
	ctx context.Context,
	ui terminal.UI,
	src *component.Source,
	log hclog.Logger,
) (*Image, error) {
	u := ui.Status()
	defer u.Close()
	u.Update("Building application")

	if b.config.Source == "" {
		b.config.Source = "./"
	}

	// check for nixpack
	nixpacksPath, err := exec.LookPath("nixpacks")
	if err != nil {
		u.Step(terminal.StatusError, fmt.Sprintf("Nixpacks not found: %s", err))
		return nil, err
	}

	u.Step(terminal.StatusOK, fmt.Sprintf("Nixpacks installed at: %q", nixpacksPath))

	cmd := exec.Command(nixpacksPath, "build", ".", "--name", fmt.Sprintf("waypoint.local/%s", src.App))

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
