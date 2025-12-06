package main

import (
	"errors"
	"fmt"
	"math/rand"
	"slices"
	"strings"
	"time"
)

// As an example for code we can test here
// I came up with a fun little loot generator
// This will get more complex on the way

var (
	ErrNoSuchObject     = errors.New("object with this name not found in pool")
	ErrEmptyObjectInput = errors.New("provided object name is empty")
)

type Item struct {
	Name     string
	Value    int
	Currency string
	Tags     map[string]string
}

type Generator struct {
	Adjectives  []string
	ObjectNames []string
	Currencies  []string
	Rand        *rand.Rand
}

// NewDefaultGenerator builds a generator with common word lists.
func NewDefaultGenerator() *Generator {
	return &Generator{
		Adjectives:  []string{"ancient", "mystic", "rusty", "shiny", "dull"},
		ObjectNames: []string{"greatsword", "amulet", "shield", "dagger", "helm"},
		Currencies:  []string{"gold", "silver", "copper"},
		Rand:        rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// GenerateRandomItem returns a random item from the pool
func (g *Generator) GenerateRandomItem() Item {
	adj := g.Adjectives[g.Rand.Intn(len(g.Adjectives))]
	objectName := g.ObjectNames[g.Rand.Intn(len(g.ObjectNames))]
	itemName := adj + " " + objectName

	return g.generate(objectName, itemName)
}

// GenerateItem returns a random item of this specific objectName if part of the pool
func (g *Generator) GenerateItem(objectName string) (*Item, error) {

	if objectName == "" {
		return nil, ErrEmptyObjectInput
	}

	if ok := slices.Contains(g.ObjectNames, strings.ToLower(objectName)); !ok {
		return nil, ErrNoSuchObject
	}

	adj := g.Adjectives[g.Rand.Intn(len(g.Adjectives))]
	itemName := adj + " " + objectName

	result := g.generate(objectName, itemName)
	return &result, nil
}

func (g *Generator) generate(objectName, itemName string) Item {
	currency := g.Currencies[g.Rand.Intn(len(g.Currencies))]
	value := g.randomValueForCurrency(currency)

	tags := map[string]string{
		"type":   objectName,
		"rarity": g.randomRarity(),
	}

	return Item{
		Name:     itemName,
		Value:    value,
		Currency: currency,
		Tags:     tags,
	}
}

func (g *Generator) randomValueForCurrency(currency string) int {
	switch currency {
	case "gold":
		return g.Rand.Intn(5) + 1 // 1–5
	case "silver":
		return g.Rand.Intn(20) + 1 // 1–20
	default:
		return g.Rand.Intn(50) + 1 // copper: 1–50
	}
}

func (g *Generator) randomRarity() string {
	rarities := []string{"common", "uncommon", "rare", "epic"}
	return rarities[g.Rand.Intn(len(rarities))]
}

func main() {
	generator := NewDefaultGenerator()
	fmt.Printf("%v", generator.GenerateRandomItem())
}
