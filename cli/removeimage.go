package cli

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type RemoveImageCLI struct {
	rows []string
}

func (c *RemoveImageCLI) GetRowsAsString() string {
	return "exit\n" + strings.Join(c.rows, "\n")
}

func (c *RemoveImageCLI) Exec(id string) error {
	return exec.Command("docker", "rmi", id).Run()
}

func (c *RemoveImageCLI) PickUpID(row string) string {
	return c.pickUpID(row)
}

func (c *RemoveImageCLI) pickUpID(row string) string {
	r := strings.Fields(row)
	return r[2]
}

func (c *RemoveImageCLI) UpdateRows(id string) {
	newRows := []string{}
	for i := range c.rows {
		_id := c.pickUpID(c.rows[i])
		if id != _id {
			newRows = append(newRows, c.rows[i])
		}
	}
	c.rows = newRows
}

func (c *RemoveImageCLI) IsExitCLI(row string) bool {
	if row == "exit\n" {
		return true
	}
	return false
}

func NewRemoveImageCLI() *RemoveImageCLI {
	cmd := exec.Command("docker", "images")
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
	return &RemoveImageCLI{rows}
}
