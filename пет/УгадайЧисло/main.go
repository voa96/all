package main

import (
	"flag"
	"guessNumber/guessCycle"
	"guessNumber/random"
)

func main() {
	min := flag.Int("min", 1, "минимальное число")
	max := flag.Int("max", 100, "максимальное число")
	attempts := flag.Int("attempts", 10, "число попыток")
	flag.Parse()
	findNum := random.RandNum(*min, *max)
	settings := map[string]*int{
		"min":        min,
		"max":        max,
		"attempts":   attempts,
		"findNumber": &findNum,
	}
	guesscycle.Switcher(settings)
}
