//+build !production

package main

import (
	"log"
	"os"
	"os/signal"
	"runtime"

	"github.com/zserge/lorca"
	"github.com/mitchellh/go-homedir"
)

func main() {
	args := []string{}
	if runtime.GOOS == "linux" {
		args = append(args, "--class=Lorca")
	}
	ui, err := lorca.New("", "", 800, 600, args...)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	ui.Load("http://localhost:3000")

	home, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}
	ui.Bind("profiles", func() []string {
		cr := os.Getenv("AWS_SHARED_CREDENTIALS_FILE")
		if cr != "" {
			log.Fatalf("Error: %s", home)
		}
		return []string{}
	})

	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)
	select {
	case <-sigc:
	case <-ui.Done():
	}

	log.Println("exiting...")
}
