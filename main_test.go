package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func RunProcessCommands(t *testing.T, inputfile string, expectedfile string) {
	input, err := os.Open(inputfile)
	if err != nil {
		assert.FailNow(t, "Failed to open test input fixture")
	}

	expected, err := ioutil.ReadFile(expectedfile)
	if err != nil {
		assert.FailNow(t, "Failed to read expected result fixture")
	}

	b := &strings.Builder{}
	ProcessCommands(input, b)
	assert.Equal(t, string(expected), b.String())
}

func TestExample1(t *testing.T) {
	RunProcessCommands(t, "testdata/example_1.txt", "testdata/example_1_expected.txt")
}

func TestExample2(t *testing.T) {
	RunProcessCommands(t, "testdata/example_2.txt", "testdata/example_2_expected.txt")
}
func TestExample3(t *testing.T) {
	RunProcessCommands(t, "testdata/example_3.txt", "testdata/example_3_expected.txt")
}
func TestExample4(t *testing.T) {
	RunProcessCommands(t, "testdata/example_4.txt", "testdata/example_4_expected.txt")
}
