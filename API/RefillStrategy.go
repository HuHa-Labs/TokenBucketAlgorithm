package main

type RefillStrategy interface {
	Refill() uint32
	GetUntilNextRefill(timeUnit string)
}
