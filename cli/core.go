package cli

import (
	"fmt"
	"io"
	"os/exec"

	"github.com/pkg/errors"
)

func Run(cli CLI) error {
	for {
		cmd := exec.Command("peco")
		stdin, err := cmd.StdinPipe()
		if err != nil {
			return errors.Wrap(err, "Stdin Pipe Error")
		}
		io.WriteString(stdin, "exit\n"+cli.GetRowsAsString())
		stdin.Close()

		out, err := cmd.Output()
		if err != nil {
			return errors.Wrap(err, "Cmd Output Error")
		}
		row := string(out)
		if row == "" || row == "exit\n" {
			break
		}

		cli.SetID(row)
		cli.UpdateRows()
		go cli.Exec()

		if cli.isOnceCLI() {
			break
		}
	}

	if out := cli.Output(); out != "" {
		fmt.Println(out)
	}

	return nil
}
