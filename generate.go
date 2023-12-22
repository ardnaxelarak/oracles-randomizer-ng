package main

// the only point of this package is so that the top-level project directory
// doesn't get cluttered with a ton of .go files.

//go:generate esc -o randomizer/embed.go -pkg randomizer hints/ logic/ romdata/
