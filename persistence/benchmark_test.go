package persistence

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func BenchmarkTestPerformanceMapQueryByID(b *testing.B) {
	performanceMap := newPerformanceMap[emptyProductInterface]()
	var loadList []FactoryPersistenceRecord
	for i := 0; i < 114514; i++ {
		loadList = append(loadList, FactoryPersistenceRecord{ID: uint(i), UID: strconv.Itoa(i)})
	}
	performanceMap.Load(loadList)
	randomX := uint(rand.Intn(114514))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		performanceMap.QueryByID(randomX)
	}
}

func BenchmarkTestPerformanceMapQueryByUID(b *testing.B) {
	performanceMap := newPerformanceMap[emptyProductInterface]()
	var loadList []FactoryPersistenceRecord
	for i := 0; i < 114514; i++ {
		loadList = append(loadList, FactoryPersistenceRecord{ID: uint(i), UID: strconv.Itoa(i)})
	}
	performanceMap.Load(loadList)
	randomX := strconv.Itoa(rand.Intn(114514))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		performanceMap.QueryByUID(randomX)
	}
}

func BenchmarkTestPerformanceMapFlush(b *testing.B) {
	performanceMap := newPerformanceMap[emptyProductInterface]()
	var loadList []FactoryPersistenceRecord
	for i := 0; i < 114514; i++ {
		loadList = append(loadList, FactoryPersistenceRecord{ID: uint(i), UID: strconv.Itoa(i)})
	}
	performanceMap.Load(loadList)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		performanceMap.Flush()
	}
}

func BenchmarkTestPerformanceMapLoad(b *testing.B) {
	performanceMap := newPerformanceMap[emptyProductInterface]()
	var loadList []FactoryPersistenceRecord
	for i := 0; i < 114514; i++ {
		loadList = append(loadList, FactoryPersistenceRecord{ID: uint(i), UID: strconv.Itoa(i)})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		performanceMap.Load(loadList)
	}
}

func BenchmarkTestPerformanceMapRegister(b *testing.B) {
	performanceMap := newPerformanceMap[testProductInterface]()
	var loadList []FactoryPersistenceRecord
	for i := 0; i < 114514; i++ {
		loadList = append(loadList, FactoryPersistenceRecord{ID: uint(i), UID: strconv.Itoa(i)})
	}
	performanceMap.Load(loadList)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		performanceMap.Register(testIntTypeFactory)
	}
}

func BenchmarkTestMemoryCacheQueryByID(b *testing.B) {
	memoryCache := newMemoryCache[uint, struct{}]()
	for i := 0; i < 114514; i++ {
		memoryCache.InsertOne(uint(i), struct{}{})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		memoryCache.QueryByID(uint(i))
	}
}

func BenchmarkTestMemoryCacheInsertOne(b *testing.B) {
	memoryCache := newMemoryCache[uint, struct{}]()
	for i := 0; i < 114514; i++ {
		memoryCache.InsertOne(uint(i), struct{}{})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		memoryCache.InsertOne(uint(i), struct{}{})
	}
}

func BenchmarkTestMemoryCacheUpdateOne(b *testing.B) {
	memoryCache := newMemoryCache[uint, struct{}]()
	for i := 0; i < 114514; i++ {
		memoryCache.InsertOne(uint(i), struct{}{})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		memoryCache.UpdateByID(uint(i), struct{}{})
	}
}

func BenchmarkTestMemoryCacheDeleteOne(b *testing.B) {
	memoryCache := newMemoryCache[uint, struct{}]()
	for i := 0; i < 114514; i++ {
		memoryCache.InsertOne(uint(i), struct{}{})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		memoryCache.DeleteOne(uint(i))
	}
}

func BenchmarkTestTimingMemoryCacheQueryByID(b *testing.B) {
	memoryCache := newTimingMemoryCache[int, struct{}]()
	for i := 0; i < 114514; i++ {
		memoryCache.InsertOne(i, struct{}{}, time.Millisecond*500)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		memoryCache.QueryByID(i)
	}
}

func BenchmarkTestTimingMemoryCacheInsertOne(b *testing.B) {
	memoryCache := newTimingMemoryCache[int, struct{}]()
	for i := 0; i < 114514; i++ {
		memoryCache.InsertOne(i, struct{}{}, time.Millisecond*500)
	}

	b.ResetTimer()
	for i := b.N; i > 0; i-- {
		memoryCache.InsertOne(i, struct{}{}, time.Millisecond*500)
	}
}

func BenchmarkTestTimingMemoryCacheRefreshByID(b *testing.B) {
	memoryCache := newTimingMemoryCache[int, struct{}]()
	for i := 0; i < 114514; i++ {
		memoryCache.InsertOne(i, struct{}{}, time.Millisecond*500)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		memoryCache.RefreshByID(i, time.Millisecond*500)
	}
}

func BenchmarkTestTimingMemoryCacheDeleteByID(b *testing.B) {
	memoryCache := newTimingMemoryCache[int, struct{}]()
	for i := 0; i < 114514; i++ {
		memoryCache.InsertOne(i, struct{}{}, time.Millisecond*500)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		memoryCache.DeleteByID(i)
	}
}

func BenchmarkTestTimingMemoryCacheUpdateByID(b *testing.B) {
	memoryCache := newTimingMemoryCache[int, struct{}]()
	for i := 0; i < 114514; i++ {
		memoryCache.InsertOne(i, struct{}{}, time.Millisecond*500)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		memoryCache.UpdateByID(i, struct{}{})
	}
}

func BenchmarkTestTimingMemoryCacheGenerationUsage(b *testing.B) {
	memoryCache := newTimingMemoryCache[int, struct{}]()
	for i := 0; i < 114514; i++ {
		memoryCache.InsertOne(i, struct{}{}, time.Millisecond*500)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		switch i % 4 {
		case 0:
			memoryCache.UpdateByID(i, struct{}{})
		case 1:
			memoryCache.InsertOne(i, struct{}{}, time.Millisecond*500)
		case 2:
			memoryCache.DeleteByID(i)
		case 3:
			memoryCache.RefreshByID(i, time.Millisecond*500)
		case 4:
			memoryCache.QueryByID(i)
		}

	}
}
