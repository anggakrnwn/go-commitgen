package git

import (
	"bytes"
	"fmt"
	"os/exec"
)

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) DiffCached() (string, error) {
	cmd := exec.Command("git", "diff", "--cached")

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("failed to execute git diff: %v, stderr: %s", err, stderr.String())
	}

	return out.String(), nil

}
