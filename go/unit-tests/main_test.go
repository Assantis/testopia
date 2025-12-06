package main_test

import (
	"testing"
	main "testopia/m/v2"

	"github.com/stretchr/testify/require"
)

// Let's start simple with a basic unit test
// I like to explain tests with three simple terms:
// Given / When / Then
func TestGenerator(t *testing.T) {
	// Given: your given environment or parameters for the test
	generator := main.NewDefaultGenerator()

	// When: the exuction of a function or method
	result := generator.GenerateItem()

	// Then: the assertions you want to make about the result
	require.NotEmpty(t, result.Name)
	require.NotEmpty(t, result.Currency)
	require.NotZero(t, result.Value)
	require.NotEmpty(t, result.Tags)
}
