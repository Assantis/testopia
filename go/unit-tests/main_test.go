package main_test

import (
	"testing"
	main "testopia/m/v2"

	"github.com/stretchr/testify/require"
)

// Let's start simple with a basic unit test
// I like to explain tests with three simple terms:
// Given / When / Then
func TestGenerateRandomItem_HappyCase(t *testing.T) {
	// Given: your given environment or parameters for the test
	generator := main.NewDefaultGenerator()

	// When: the exuction of a function or method
	result := generator.GenerateRandomItem()

	// Then: the assertions you want to make about the result
	require.NotEmpty(t, result.Name)
	require.NotEmpty(t, result.Currency)
	require.NotZero(t, result.Value)
	require.NotEmpty(t, result.Tags)
}

// For stable code it is really worth covering the cases that indeed can go wrong
// and even more if the scenario is very likely or part of normal operation.
// This is called negative testing and I cannot stress enough what a powerful friend this is.
// I also recommend to not only test the error but also test the result.
func TestGenerateItem_ShouldFailOnObjectNotFound(t *testing.T) {
	// Given
	generator := main.NewDefaultGenerator()

	// When
	result, err := generator.GenerateItem("sword")

	// Then
	require.Equal(t, main.ErrNoSuchObject, err)
	require.Nil(t, result)
}
