package golang

import (
	"fmt"
	"testing"
	"time"
)

func Test_CountPrimes(t *testing.T) {
	start := time.Now()
	count1 := CountPrimes(1500000)
	fmt.Printf("CountPrimes用时: %dms，结果: %d\n", time.Since(start).Milliseconds(), count1)

	start = time.Now()
	count2 := CountPrimes2(1500000)
	fmt.Printf("CountPrimes2用时: %dms，结果: %d\n", time.Since(start).Milliseconds(), count2)
}