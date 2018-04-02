package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"
)

var (
	Version string
)

func move(src, dest string) error {
	cpCmd := exec.Command("mv", src, dest)
	err := cpCmd.Run()
	return err
}

func autoUpdate() {
	binPath, _ := os.Executable()
	updatedBinPath := "./updates/go-self-update"

	if _, err := os.Stat(updatedBinPath); err == nil {
		out, _ := exec.Command(updatedBinPath, "-version").Output()
		remoteVersionStr := strings.TrimSpace(string(out))
		remoteVersion, err := strconv.ParseInt(remoteVersionStr, 10, 32)
		if err != nil {
			log.Fatal("Invalid remote version formatting")
		}

		localVersion, err := strconv.ParseInt(Version, 10, 32)
		if err != nil {
			log.Fatal("Invalid local version formatting")
		}

		if remoteVersion > localVersion {
			move(updatedBinPath, binPath)
			log.Printf("Updated...\nRestarting application...\n")
			syscall.Exec(os.Args[0], os.Args, os.Environ())
		}
	}
}

func main() {
	version := flag.Bool("version", false, "Prints current version")
	flag.Parse()
	if *version {
		fmt.Println(Version)
		os.Exit(0)
	}

	counter := 0
	for true {
		log.Printf("Current Version: %s", Version)
		counter = counter + 1
		time.Sleep(time.Second * 5)

		if counter%5 == 0 {
			autoUpdate()
		}
	}
}
