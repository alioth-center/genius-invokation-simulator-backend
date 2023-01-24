package util

import (
	"sync/atomic"
	"testing"
	"time"
)

func BenchmarkTestGenerateID(b *testing.B) {
	if macs, err := GetMacAddresses(); err != nil || len(macs) == 0 {
		b.Fatalf("get mac addresses failed: %v", err)
	} else if mac, parsed := GetUintMacAddress(macs[0]); parsed != nil {
		b.Fatalf("parse mac address failed: %v", parsed)
	} else {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			GenerateUID(mac, time.Now())
		}
	}
}

func BenchmarkGenerateRealID(b *testing.B) {
	var macAddr uint64
	if macs, err := GetMacAddresses(); err != nil || len(macs) == 0 {
		b.Fatalf("get mac addresses failed: %v", err)
	} else if mac, parsed := GetUintMacAddress(macs[0]); parsed != nil {
		b.Fatalf("parse mac address failed: %v", parsed)
	} else {
		macAddr = mac
	}

	uids := make([]uint64, 114514)
	for i := 0; i < 114514; i++ {
		uids[i] = GenerateUID(macAddr, time.Now())
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
