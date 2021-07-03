package main

type Card struct {
	Next, Previous *Card
	Value int
}