package fetcher

import (
	"path/filepath"
	"regexp"

	"github.com/bashidogames/gdvm/config"
	"github.com/bashidogames/gdvm/internal/utils"
	"github.com/bashidogames/gdvm/semver"
)

const EXECUTABLE_REGEX_PATTERN = "Godot((?!console).)*?[.]exe"
const LINK_FILENAME = "godot"

var ExecutableRegex = regexp.MustCompile(EXECUTABLE_REGEX_PATTERN)

type Fetcher struct {
	Config *config.Config
}

func (f *Fetcher) TargetPath(semver semver.Semver) (string, error) {
	return f.locateExecutable(filepath.Join(f.Config.GodotRootDirectory, semver.GodotString()))
}

func (f *Fetcher) LinkPath(semver semver.Semver) string {
	return filepath.Join(f.Config.BinDirectory, LINK_FILENAME)
}

func (f *Fetcher) locateExecutable(root string) (string, error) {
	return utils.LocateExecutable(ExecutableRegex, root, true)
}

func New(config *config.Config) *Fetcher {
	return &Fetcher{
		Config: config,
	}
}
