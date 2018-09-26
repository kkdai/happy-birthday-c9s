package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// $ go run yifa.go

const (
	wishFile = "yifa.base64"
)

func main() {
	f, err := os.Open(wishFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	raw, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	plain, err := base64.StdEncoding.DecodeString(string(raw))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(plain))
}
