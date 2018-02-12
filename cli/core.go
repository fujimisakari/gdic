package cli

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func Run(cli CLI) {
	for {
		cmd := exec.Command("peco")
		stdin, err := cmd.StdinPipe()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		io.WriteString(stdin, cli.GetRowsAsString())
		stdin.Close()
		out, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		row := string(out)
		if row == "" {
			break
		}
		if cli.IsExitCLI(row) {
			break
		}

		id := cli.PickUpID(row)
		cli.UpdateRows(id)
		go cli.Exec(id)
	}
	fmt.Println("done")
}
