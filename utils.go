package main

import (
	"runtime"
	"os/exec"
	"strings"
)

func startBrowser(url string) bool {
	// try to start the browser
	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open"}
	case "windows":
		args = []string{"cmd", "/c", "start"}
	default:
		args = []string{"xdg-open"}
	}
	cmd := exec.Command(args[0], append(args[1:], url)...)
	return cmd.Start() == nil
}

func checkSuffix(str string, suffix []string) (re string) {
	for _, v := range suffix {
		if strings.HasSuffix(str, v) {
			re = v
			return
		}
	}
	return
}

func findIntInSlice(sli []int, obj int) (re int) {
	re = -1
	for i, v := range sli {
		if v == obj {
			re = i
			break
		}
	}
	return
}

func getMin(x, y int) int {
	if x < y {
		return x
	}
	return y
}
