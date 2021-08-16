package adr

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/user"
	"strings"
	"text/template"

	"github.com/magefile/mage/sh"
)

const DOCKER_EXEC = ("docker run --rm -u {{.UID}}:{{.GID}} -v {{.AdrDir}}:/doc " +
	"-v {{.AdrConfigDir}}:/adr-config -e ADR_TEMPLATE=/adr-config/{{.AdrTemplate}} " +
	"brianskarda/adr-tools-docker adr {{.Cmd}}")

const (
	adrDirDefault       = "documentation/adr/"
	adrConfigDirDefault = ".adr-config/"
	adrTemplateDefault  = "template.rst"
)

type params struct {
	UID          string
	GID          string
	AdrDir       string
	AdrConfigDir string
	AdrTemplate  string
	Cmd          string
}

func buildParams(ctx context.Context, cmd, title string) (*params, error) {
	var adrDir, adrConfigDir, adrTemplate interface{}

	cmd = fmt.Sprintf("%s %s", cmd, title)
	usr, err := user.Current()
	if err != nil {
		return nil, err
	}
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	if adrDir = ctx.Value("adrDir"); adrDir == nil {
		adrDir = adrDirDefault
	}
	if adrConfigDir = ctx.Value("adrConfigDir"); adrConfigDir == nil {
		adrConfigDir = adrConfigDirDefault
	}
	if adrTemplate = ctx.Value("adrTemplate"); adrTemplate == nil {
		adrTemplate = adrTemplateDefault
	}

	return &params{
		Cmd:          cmd,
		AdrDir:       fmt.Sprintf("%s/%s", pwd, adrDir.(string)),
		AdrConfigDir: fmt.Sprintf("%s/%s", pwd, adrConfigDir.(string)),
		AdrTemplate:  adrTemplate.(string),
		UID:          usr.Uid,
		GID:          usr.Gid,
	}, nil
}

func buildCommand(ctx context.Context, cmd, path string) ([]string, error) {
	tmpl, err := template.New("RUN_ADR").Parse(DOCKER_EXEC)
	if err != nil {
		return []string{}, err
	}

	params, err := buildParams(ctx, "init", path)
	if err != nil {
		return []string{}, err
	}

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, params)
	if err != nil {
		return []string{}, err
	}

	command := strings.Split(buf.String(), " ")
	return command, nil

}

// New creates a new ADR entry. Specify title of an ADR entry.
// By default adrDir, adrConfigDir, and adrTemplate are set.
// If you wish to override them use ctx.WithValue() to
// set corresponding values and pass the context to this function.
//   * adrDir - directory which docker mounts as /docs
//   * adrConfigDir - directory with configuration which docker mounts as /adr-config
//   * adrTemplate - template name which will be used to produce ADRs.
func New(ctx context.Context, title string) error {
	var command []string
	var err error
	if command, err = buildCommand(ctx, "new", title); err != nil {
		return err
	}
	return sh.Run(command[0], command[1:]...)
}
