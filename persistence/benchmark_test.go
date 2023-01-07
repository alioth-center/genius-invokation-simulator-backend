package persistence

import (
	"math/rand"
	"strconv"
	"testing"
)

func BenchmarkTestPerformanceMapQueryByID(b *testing.B) {
	performanceMap := NewPerformanceMap[emptyProductInterface]()
	var loadList []PerformanceMapRecord
	for i := 0; i < 114514; i++ {
		loadList = append(loadList, PerformanceMapRecord{ID: uint(i), UID: strconv.Itoa(i)})
	}
	performanceMap.Load(loadList)
	randomX := uint(rand.Intn(114514))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		performanceMap.QueryByID(randomX)
	}
}

func BenchmarkTestPerformanceMapQueryByUID(b *testing.B) {
	performanceMap := NewPerformanceMap[emptyProductInterface]()
	var loadList []PerformanceMapRecord
	for i := 0; i < 114514; i++ {
		loadList = append(loadList, PerformanceMapRecord{ID: uint(i), UID: strconv.Itoa(i)})
	}
	performanceMap.Load(loadList)
	randomX := strconv.Itoa(rand.Intn(114514))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		performanceMap.QueryByUID(randomX)
	}
}

func BenchmarkTestPerformanceMapFlush(b *testing.B) {
	performanceMap := NewPerformanceMap[emptyProductInterface]()
	var loadList []PerformanceMapRecord
	for i := 0; i < 114514; i++ {
		loadList = append(loadList, PerformanceMapRecord{ID: uint(i), UID: strconv.Itoa(i)})
	}
	performanceMap.Load(loadList)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		performanceMap.Flush()
	}
}

func BenchmarkTestPerformanceMapLoad(b *testing.B) {
	performanceMap := NewPerformanceMap[emptyProductInterface]()
	var loadList []PerformanceMapRecord
	for i := 0; i < 114514; i++ {
		loadList = append(loadList, PerformanceMapRecord{ID: uint(i), UID: strconv.Itoa(i)})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		performanceMap.Load(loadList)
	}
}

func BenchmarkTestPerformanceMapRegister(b *testing.B) {
	performanceMap := NewPerformanceMap[testProductInterface]()
	var loadList []PerformanceMapRecord
	for i := 0; i < 114514; i++ {
		loadList = append(loadList, PerformanceMapRecord{ID: uint(i), UID: strconv.Itoa(i)})
	}
	performanceMap.Load(loadList)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		performanceMap.Register(testIntTypeFactory)
	}
}
