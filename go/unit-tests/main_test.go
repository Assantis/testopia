package main_test

import (
	"testing"
	main "testopia/m/v2"

	"github.com/stretchr/testify/require"
)

// Let's start simple with a basic unit test
// I like to explain tests with three simple terms:
// Given / When / Then
func TestGenerateRandomItem_ShouldSucceed(t *testing.T) {
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

// Now here is one of the most valuable tips I got from one of my prior mentors.
// He noticed that I was frustrated, because the unit tests took me a lot of time.
// He sat down and asked me to demonstrate how I write tests, to identify why that is.
// And basically my issue was: I always started with the happy case
// This forced my brain to consider everything that could go wrong at once.
// As nearly every brain has an easier time processing small chunks of information,
// he told me to start with the error cases and work my way up to the final case: the happy one
// After just a week, I was able to write the tests in a fraction of the time.
// Because basically what this allows you to do is:
// - copy your prior setup
// - change the name of the test
// - change your input
// - assert for the next error / assert the invertion of your prior logic
func TestGenerateItem_ShouldSucceed(t *testing.T) {
	// Given
	generator := main.NewDefaultGenerator()

	// When
	result, err := generator.GenerateItem("helm")

	// Then
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotEmpty(t, result.Name)
	require.NotEmpty(t, result.Currency)
	require.NotZero(t, result.Value)
	require.NotEmpty(t, result.Tags)
}
