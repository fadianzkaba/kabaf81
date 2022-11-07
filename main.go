package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {

	out, err := exec.Command("kubectl", "get", "svc", "-n", "bah-dev").Output()

	if err != nil {
		log.Fatal(err)
		out, err := exec.Command("gcloud", "auth", "login").Output()
	}
	fmt.Printf(string(out))
}
