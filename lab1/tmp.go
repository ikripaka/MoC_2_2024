package main

import (
	"os"
	"strings"
)

func main() {
	f, err := os.ReadFile("./gamma_v1.txt")
	if err != nil {
		panic(err)
	}

	N := 2000

	err = os.WriteFile(
		"./gamma_prepared.txt",
		[]byte(strings.Join(strings.Split(string(f), ""), ",")[:N]),
		0666)

	if err != nil {
		panic(err)
	}

}
