package util

import (
	"math/rand"
	"sync/atomic"
	"testing"
	"time"
)

func BenchmarkTestGenerateID(b *testing.B) {
	mac := rand.Uint64() % (1 << 48)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GenerateUID(mac, time.Now())
	}

}

func BenchmarkGenerateRealID(b *testing.B) {
	mac := rand.Uint64() % (1 << 48)
	uids := make([]uint64, 114514)
	for i := 0; i < 114514; i++ {
		uids[i] = GenerateUID(mac, time.Now())
	}

	var index int32 = 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if index >= 114514 {
			atomic.StoreInt32(&index, 0)
		}
		GenerateRealID(uids[index], uint16(index))
		atomic.AddInt32(&index, 1)
	}
}
