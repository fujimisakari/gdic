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
		io.WriteString(stdin, cli.GetRowsAsString())
		stdin.Close()

		out, err := cmd.Output()
		if err != nil {
			return errors.Wrap(err, "Cmd Output Error")
		}
		row := string(out)
		if row == "" || cli.IsExitCLI(row) {
			break
		}

		id := cli.PickUpID(row)
		cli.UpdateRows(id)
		go cli.Exec(id)
	}
	fmt.Println("done")
	return nil
}
