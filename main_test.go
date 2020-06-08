package main

import (
	"os"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFoo(t *testing.T) {
	assert := assert.New(t)

	var envFoo string = os.Getenv("FOO")

	assert.Equal("foo", envFoo, "The two words should be the same.")
}

func TestBar(t *testing.T) {
	assert := assert.New(t)

	var envFoo string = os.Getenv("BAR")

	assert.Equal("bar", envFoo, "The two words should be the same.")
}