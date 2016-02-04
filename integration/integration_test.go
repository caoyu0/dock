package integration

import (
	"fmt"
	"os"
	"testing"
)

var testImage string //must match with what is in the Makefile

func init() {
	testImage = os.Getenv("TEST_IMAGE")
}

func TestSimpleCommand(t *testing.T) {
	d := newDocker()
	if err := d.start(true, "run", testImage, "dock", "--debug", "ls"); err != nil {
		fmt.Println(d.debugInfo())
		t.Fatal(err)
	}
}

func TestSimpleInteractiveCommand(t *testing.T) {
	d := newDocker()
	if err := d.start(true, "run", "-t", testImage, "dock", "-i", "--debug", "ls"); err != nil {
		fmt.Println(d.debugInfo())
		t.Fatal(err)
	}
}

func TestInteractiveNoTerminal(t *testing.T) {
	d := newDocker()
	if err := d.start(true, "run", testImage, "dock", "-i", "--debug", "ls"); err == nil { //fail because no terminal in docker container
		t.Fatal(err)
	}
}

func TestOrphanProcessReaping(t *testing.T) {
	d := newDocker()
	// using --debug, dock will have a 999 exit code if more than one process exists when exiting
	if err := d.start(true, "run", testImage, "dock", "--debug", "bash", "/assets/spawn_orphaned.sh"); err != nil {
		fmt.Println(d.debugInfo())
		t.Fatal(err)
	}
}
