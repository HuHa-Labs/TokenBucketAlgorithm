package main

type TokenBucket interface {
	ConsumeMultipleToken(numToken uint32) bool
	ConsumeOneToken() bool
	tryConsumeOneToken() bool
	tryConsumeMultipleToken(numToken uint32)
	refill(numToken uint32)
	getUntilNextRefill() uint32
	getSize() uint32
	getCapacity() uint32
}

type TokenBucketImpl struct {
	capacity uint32
	size uint32
	refillStrategy RefillStrategy
	sleepStrategy SleepStrategy
}