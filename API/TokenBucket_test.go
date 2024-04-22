package main

import "testing"

func TestTokenBucketImplIsTokenBucket(t *testing.T) {
	var _ TokenBucket = &TokenBucketImpl{capacity: 10, size: 10}
}

func TestTokenBucketImplGetSize(t *testing.T) {
	tb := TokenBucketImpl{capacity: 10, size: 10}
	if tb.getSize() != 10 {
		t.Errorf("Expected 10, got %d", tb.getSize())
	}
}

func TestTokenBucketImplGetCapacity(t *testing.T) {
	tb := TokenBucketImpl{capacity: 10, size: 10}
	if tb.getCapacity() != 10 {
		t.Errorf("Expected 10, got %d", tb.getCapacity())
	}
}