package adr

import (
	"fmt"
	"os"
	"os/user"
)

const DOCKER_EXEC = ("docker run --rm -u {{.UID}}:{{.GID}} -v {{.AdrDir}}:/doc " +
	"-v {{.AdrConfigDir}}:/adr-config -e ADR_TEMPLATE=/adr-config/{{.AdrTemplate}} " +
	"brianskarda/adr-tools-docker adr {{.Cmd}}")

const (
	adrDirDefault       = "documentation/"
	adrConfigDirDefault = ".adr-config/"
	adrTemplateDefault  = "template.md"
)

type params struct {
	UID          string
	GID          string
	AdrDir       string
	AdrConfigDir string
	AdrTemplate  string
	Cmd          string
}

func buildParams(cmd, title, adrDir, adrConfigDir, adrTemplate string) (*params, error) {

	cmd = fmt.Sprintf("%s %s", cmd, title)
	usr, err := user.Current()
	if err != nil {
		return nil, err
	}
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	if adrDir == "" {
		adrDir = adrDirDefault
	}
	if adrConfigDir == "" {
		adrConfigDir = adrConfigDirDefault
	}
	if adrTemplate == "" {
		adrTemplate = adrTemplateDefault
	}

	return &params{
		Cmd:          cmd,
		AdrDir:       fmt.Sprintf("%s/%s", pwd, adrDir),
		AdrConfigDir: fmt.Sprintf("%s/%s", pwd, adrConfigDir),
		AdrTemplate:  adrTemplate,
		UID:          usr.Uid,
		GID:          usr.Gid,
	}, nil
}
