package varnish

import (
	"log"
	"os"
	"testing"
)

func setup() {
	log.Println("setup tests")
}

func teardown() {
	log.Println("teardown tests")
}

func TestMain(m *testing.M) {
	setup()
	results := m.Run()
	teardown()
	os.Exit(results)
}
