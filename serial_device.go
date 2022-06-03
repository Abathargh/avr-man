package main

import (
	"fmt"
	"go.bug.st/serial/enumerator"
	"log"
)

var deviceNames map[string][]string = map[string][]string{
	"Arduino Mega 2560": []string{"-D", "-c", "stk500v2"},
}

func findDevice() (string, error) {
	ports, err := enumerator.GetDetailedPortsList()
	if err != nil {
		log.Fatal(err)
	}

	for _, port := range ports {
		if port.IsUSB {
			fmt.Printf("%+v\n", port.Product)
		}
	}
	return "", nil
}

func main() {
	findDevice()
}
