package main

import "testing"

func TestTokenBucketImplIsTokenBucket(t *testing.T) {
	var _ TokenBucket = &TokenBucketImpl{capacity: 10, size: 10}
}


