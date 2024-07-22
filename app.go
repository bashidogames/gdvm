package gdvm

import (
	"fmt"

	"github.com/bashidogames/gdvm/config"
	"github.com/bashidogames/gdvm/internal/environment"
	"github.com/bashidogames/gdvm/internal/github"
	"github.com/bashidogames/gdvm/internal/services/godot"
	"github.com/bashidogames/gdvm/internal/services/versions"
)

type App struct {
	Versions *versions.Service
	Godot    *godot.Service
}

func New(config *config.Config) (*App, error) {
	environment, err := environment.New(github.New(config))
	if err != nil {
		return nil, fmt.Errorf("failed to create environment: %w", err)
	}

	return &App{
		Versions: versions.New(environment, config),
		Godot:    godot.New(environment, config),
	}, nil
}
