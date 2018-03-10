package cli

import (
	"fmt"
	"io"
	"os/exec"

	"github.com/pkg/errors"
)

func Run(cli CLI) error {
	var stdErr string

	for {
		cmd := exec.Command("peco")
		stdin, err := cmd.StdinPipe()
		if err != nil {
			return errors.Wrap(err, "Stdin Pipe Error")
		}
		if stdErr != "" {
			io.WriteString(stdin, "back\n"+stdErr)
			stdErr = ""
		} else {
			io.WriteString(stdin, "exit\n"+cli.GetRowsAsString())
		}
		stdin.Close()

		out, err := cmd.Output()
		if err != nil {
			return errors.Wrap(err, "Cmd Output Error")
		}
		row := string(out)
		if row == "" || row == "exit\n" {
			break
		} else if row == "back\n" {
			continue
		}

		cli.SetID(row)
		if err := cli.Exec(); err != "" {
			stdErr = err
		} else {
			cli.UpdateRows()
		}

		if cli.isOnceCLI() {
			break
		}
	}

	if out := cli.Output(); out != "" {
		fmt.Println(out)
	}

	return nil
}
