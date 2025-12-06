package main

import (
	"fmt"
	"math/rand"
	"time"
)

// As an example for code we can test here
// I came up with a fun little loot generator
// This will get more complex on the way

type Item struct {
	Name     string
	Value    int
	Currency string
	Tags     map[string]string
}

type Generator struct {
	Adjectives []string
	Nouns      []string
	Currencies []string
	Rand       *rand.Rand
}

// NewDefaultGenerator builds a generator with common word lists.
func NewDefaultGenerator() *Generator {
	return &Generator{
		Adjectives: []string{"ancient", "mystic", "rusty", "shiny", "dull"},
		Nouns:      []string{"greatsword", "amulet", "shield", "dagger", "helm"},
		Currencies: []string{"gold", "silver", "copper"},
		Rand:       rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (g *Generator) GenerateItem() Item {
	adj := g.Adjectives[g.Rand.Intn(len(g.Adjectives))]
	noun := g.Nouns[g.Rand.Intn(len(g.Nouns))]
	name := adj + " " + noun

	currency := g.Currencies[g.Rand.Intn(len(g.Currencies))]
	value := g.randomValueForCurrency(currency)

	tags := map[string]string{
		"type":   noun,
		"rarity": g.randomRarity(),
	}

	return Item{
		Name:     name,
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
	fmt.Printf("%v", generator.GenerateItem())
}
