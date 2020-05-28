package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"

	"nes/nes"
)

func testRom(path string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	console, err := nes.NewConsole(path)
	if err != nil {
		return err
	}
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	console.Load(decoder)

	address, kind, err := console.Mapper.ShowROMAddress(0x906A)
	if err != nil {

	}
	fmt.Fprintf(os.Stdout, "ADDRESS: 0x%04X KIND: %s\n", address, kind)
	console.StepSeconds(3)
	return nil
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatalln("Usage: go run util/rom_address.go rom_file")
	}
	file := args[0]
	err := testRom(file)
	if err == nil {
		fmt.Println("OK  ", file)
	} else {
		fmt.Println("FAIL", file)
		fmt.Println(err)
	}
}
