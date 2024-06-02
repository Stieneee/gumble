//go:build ignore
// +build ignore

package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	// Get Mumble.proto
	resp, err := http.Get("https://raw.githubusercontent.com/mumble-voip/mumble/master/src/Mumble.proto")
	if err != nil {
		log.Fatalf("could not download Mumble.proto: %s\n", err)
	}
	defer resp.Body.Close()

	// Write Mumble.proto
	f, err := os.Create("Mumble.proto")
	if err != nil {
		log.Fatalf("could not create Mumble.proto: %s\n", err)
	}
	if _, err := io.Copy(f, resp.Body); err != nil {
		log.Fatalf("could not write Mumble.proto: %s\n", err)
	}
	if err := f.Close(); err != nil {
		log.Fatalf("could not close Mumble.proto: %s\n", err)
	}

	// Get MumbleUDP.proto
	resp, err = http.Get("https://raw.githubusercontent.com/mumble-voip/mumble/master/src/MumbleUDP.proto")
	if err != nil {
		log.Fatalf("could not download MumbleUDP.proto: %s\n", err)
	}
	defer resp.Body.Close()

	// Write MumbleUDP.proto
	f, err = os.Create("MumbleUDP.proto")
	if err != nil {
		log.Fatalf("could not create MumbleUDP.proto: %s\n", err)
	}
	if _, err := io.Copy(f, resp.Body); err != nil {
		log.Fatalf("could not write MumbleUDP.proto: %s\n", err)
	}
	if err := f.Close(); err != nil {
		log.Fatalf("could not close MumbleUDP.proto: %s\n", err)
	}

	// download the latest license file from the mumble repo
	resp, err = http.Get("https://raw.githubusercontent.com/mumble-voip/mumble/master/LICENSE")
	if err != nil {
		log.Fatalf("could not download LICENSE: %s\n", err)
	}
	defer resp.Body.Close()

	// Write LICENSE
	f, err = os.Create("LICENSE")
	if err != nil {
		log.Fatalf("could not create LICENSE: %s\n", err)
	}
	if _, err := io.Copy(f, resp.Body); err != nil {
		log.Fatalf("could not write LICENSE: %s\n", err)
	}
	if err := f.Close(); err != nil {
		log.Fatalf("could not close LICENSE: %s\n", err)
	}

	// Build proto-gen-go
	if err := exec.Command("go", "install", "google.golang.org/protobuf/cmd/protoc-gen-go@latest").Run(); err != nil {
		log.Fatalf("could not build protoc-gen-go: %s\n", err)
	}

	// Generate code
	if err := exec.Command("protoc", "--go_opt=MMumble.proto=./MumbleProto", "--go_out=.", "Mumble.proto").Run(); err != nil {
		log.Fatalf("could not run protoc: %s\n", err)
	}
	if err := exec.Command("protoc", "--go_opt=MMumbleUDP.proto=./MumbleUDPProto", "--go_out=.", "MumbleUDP.proto").Run(); err != nil {
		log.Fatalf("could not run protoc: %s\n", err)
	}
}
