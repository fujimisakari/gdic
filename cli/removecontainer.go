package cli

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type RemoveContainerCLI struct {
	rows []string
}

func (c *RemoveContainerCLI) GetRowsAsString() string {
	return "exit\n" + strings.Join(c.rows, "\n")
}

func (c *RemoveContainerCLI) Exec(id string) error {
	return exec.Command("docker", "rm", id).Run()
}

func (c *RemoveContainerCLI) PickUpID(row string) string {
	return c.pickUpID(row)
}

func (c *RemoveContainerCLI) pickUpID(row string) string {
	r := strings.Fields(row)
	return r[0]
}

func (c *RemoveContainerCLI) UpdateRows(id string) {
	newRows := []string{}
	for i := range c.rows {
		_id := c.pickUpID(c.rows[i])
		if id != _id {
			newRows = append(newRows, c.rows[i])
		}
	}
	c.rows = newRows
}

func (c *RemoveContainerCLI) IsExitCLI(row string) bool {
	if row == "exit\n" {
		return true
	}
	return false
}

func NewRemoveContainerCLI() *RemoveContainerCLI {
	cmd := exec.Command("docker", "ps", "-a")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	cmd.Start()

	var i uint64
	rows := []string{}
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		if i == 0 {
			i++
			continue
		}
		rows = append(rows, scanner.Text())
	}
	cmd.Wait()
	return &RemoveContainerCLI{rows}
}
