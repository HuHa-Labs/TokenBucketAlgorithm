package main

type TokenBucket interface {
	ConsumeMultipleToken(numToken uint32) bool
	ConsumeOneToken() bool
	tryConsumeOne() bool
	tryConsumeMultiple(numToken uint32)
	refill(numToken uint32)
	getUntilNextRefill() uint32
	getSize() uint32
	getCapacity() uint32
}

