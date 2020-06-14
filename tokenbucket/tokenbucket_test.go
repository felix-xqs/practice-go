package tokenbucket

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestBucket(t *testing.T) {
	bucket := NewBucket(5, 5, time.Second)
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		time.Sleep(time.Millisecond * 150)
		go getBucket(bucket, i, wg)
	}
	wg.Wait()
}

func getBucket(bucket *Bucket, i int, wg sync.WaitGroup) {
	if bucket.Get() {
		fmt.Printf("success: %d \n", i)
	} else {
		fmt.Printf("false: %d \n", i)
	}
	wg.Done()
}
