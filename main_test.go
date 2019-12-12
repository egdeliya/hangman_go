package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_HangmanPositive(t *testing.T) {
	var stdin bytes.Buffer

	stdin.Write([]byte("w\no\nr\nl\nd\n"))

	result := hangman(&stdin, []string{"world"})
	require.True(t, result)
}

func Test_HangmanNegative(t *testing.T) {
	var stdin bytes.Buffer

	stdin.Write([]byte("w\no\nr\nl\nd\nb\nf\n"))

	result := hangman(&stdin, []string{"hello"})
	require.False(t, result)
}