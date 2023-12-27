package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func main() {

	fmt.Println(os.Getenv("LOCALAPPDATA"))

	// msg := "helloworld" + os.Getenv("LOCALAPPDATA") + "test"

	T5()

}
func T5() {
	ticker := time.NewTicker(5 * time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(12 * time.Hour)
	done <- true
	fmt.Println("Ticker stopped")

}

func T4() {
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		// <-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	// time.Sleep(2 * time.Second)
}

func T3(src, dst string) {
	bytes, err := os.ReadFile(src)
	if err != nil {
		log.Fatal(err)
	}

	os.WriteFile(dst, bytes, 0755)
	if err != nil {
		log.Fatal(err)
	}

}

func T2(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}
	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()

	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func T1() {
	src := strings.NewReader("hello world\n")

	dst := os.Stdout

	bytes, err := io.Copy(dst, src)

	if err != nil {
		panic(err)
	}

	fmt.Printf("The number ofo bytes are : %d\n", bytes)
}
