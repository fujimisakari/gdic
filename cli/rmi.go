package cli

import (
	"bufio"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

type RmiCLI struct {
	id   string
	rows []string
}

func (c *RmiCLI) SetID(row string) {
	c.id = c.pickUpID(row)
}

func (c *RmiCLI) GetRowsAsString() string {
	return strings.Join(c.rows, "\n")
}

func (c *RmiCLI) UpdateRows() {
	newRows := []string{}
	for i := range c.rows {
		id := c.pickUpID(c.rows[i])
		if c.id != id {
			newRows = append(newRows, c.rows[i])
		}
	}
	c.rows = newRows
}

func (c *RmiCLI) pickUpID(row string) string {
	r := strings.Fields(row)
	return r[0]
}

func (c *RmiCLI) Exec() error {
	return exec.Command("docker", "rmi", c.id).Run()
}

func (c *RmiCLI) Output() string {
	return ""
}

func (c *RmiCLI) isOnceCLI() bool {
	return false
}

func NewRmiCLI() (*RmiCLI, error) {
	cmd := exec.Command("docker", "images")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, errors.Wrap(err, "NewRmiCLI Error")
	}

	var i uint64
	rows := []string{}
	cmd.Start()
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		if i == 0 {
			i++
			continue
		}
		rows = append(rows, scanner.Text())
	}
	cmd.Wait()
	return &RmiCLI{rows: rows}, nil
}
