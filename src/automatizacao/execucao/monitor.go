package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {

	binary, lookErr := exec.LookPath("date")
	if lookErr != nil {
		panic(lookErr)
	}

	binary, lookErr := exec.LookPath("vmstat")
	if lookErr != nil {
		panic(lookErr)
	}

	date := []string{"date"}

	env := os.Environ()

	vms := []string{"vmstat"}

	execErr := syscall.Exec(binary, date, env, vmstat)
	if execErr != nil {
		panic(execErr)
	}

}
