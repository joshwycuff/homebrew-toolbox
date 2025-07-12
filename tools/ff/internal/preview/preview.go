package preview

import "os/exec"

func Preview(filename string) (string, error) {
	cmd := exec.Command("bat", "--style=plain,numbers", "--color=always", filename)
	out, err := cmd.Output()
	return string(out), err
}
