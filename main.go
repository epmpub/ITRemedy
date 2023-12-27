package main

import (
	"log"
	"runtime"
)

func main() {

	switch runtime.GOOS {
	case "windows":
		log.Println("Windows")
		err := check_service()
		if err != nil {
			log.Println("Run Check Script Fail")
		}
	case "Linux":
		log.Println("Linux")
	default:
		log.Println("Error Type OS")
	}
}
