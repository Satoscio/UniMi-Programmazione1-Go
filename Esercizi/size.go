// thanks to lnmx @ stackoverflow

package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	var h, w int
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin

	out, _ := cmd.Output()

	h, _ = strconv.Atoi(string(out[0]) + string(out[1]))

	if out[5] == 10 {
		w, _ = strconv.Atoi(string(out[3]) + string(out[4]))
	} else {
		w, _ = strconv.Atoi(string(out[3]) + string(out[4]) + string(out[5]))
	}
	fmt.Printf("Terminal Width: %d\n", w)
	fmt.Printf("Terminal Height: %d\n", h)
}
