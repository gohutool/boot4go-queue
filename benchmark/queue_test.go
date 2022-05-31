package benchmarks

import (
	queue "github.com/gohutool/boot4go-queue"
	"testing"
)

func BenchmarkQueueNonBlockingOneGoroutine(b *testing.B) {
	benchmarkQueueNonBlocking(b, 1)
}
func BenchmarkQueueNonBlockingTwoGoroutines(b *testing.B) {
	benchmarkQueueNonBlocking(b, 1)
}
func BenchmarkQueueNonBlockingThreeGoroutinesWithContendedWrite(b *testing.B) {
	benchmarkQueueNonBlocking(b, 2)
}

func TestQueuePutDoGet(t *testing.T) {
	testOne(10, 1)
}

func testOne(iterations, writers int) {
	maxReads := iterations * writers

	var q *queue.Queue[int] = queue.NewQueue(uint32(maxReads), 1)

	for x := 0; x < writers; x++ {
		for i := 0; i < iterations; i++ {
			q.Put(&i)
		}
	}

	for i := 0; i < maxReads; i++ {
		for {
			_, ok, _ := q.Get()
			if ok {
				break
			}
		}
	}
}

func benchmarkQueueNonBlocking(b *testing.B, writers int) {
	iterations := b.N

	b.ReportAllocs()
	b.ResetTimer()

	testOne(iterations, writers)

}
