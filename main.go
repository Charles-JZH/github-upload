package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for {
			fmt.Println("1")
			time.Sleep(time.Duration(2) * time.Second)
		}
	}()
	fmt.Println("JZH")
}

