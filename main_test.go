package main

import (
	"log"
	"os"
	"strings"
	"testing"
)

func TestGetOsArch(t *testing.T) {
	oa := getOsArch()
	log.Println("OS/ARCH: ", oa)

	if !strings.HasSuffix(oa, os.Getenv("GOARCH")) {
		t.Error("Invalid OS/ARCH string - does not contain ${GOARCH}")
	}
}
