package adr

import (
	"bytes"
	"github.com/magefile/mage/sh"
	"strings"
	"text/template"
)

// RunADRCommand triggers given ADR command
// This function allows for these parameters to be configured:
//   * adrDir - directory which docker mounts as /docs
//   * adrConfigDir - directory with configuration which docker mounts as /adr-config
//   * adrTemplate - template name which will be used to produce ADRs.
func RunADRCommand(cmd, title, adrDir, adrConfigDir, adrTemplate string) error {
	tmpl, err := template.New("RUN_ADR").Parse(DOCKER_EXEC)
	if err != nil {
		return err
	}

	params, err := buildParams(cmd, title, adrDir, adrConfigDir, adrTemplate)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, params)
	if err != nil {
		return err
	}

	command := strings.Split(buf.String(), " ")
	return sh.Run(command[0], command[1:]...)

}
