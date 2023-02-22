package util

import (
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestUIDGenerator(t *testing.T) {
	t.Run("TestUIDGenerator-1", func(t *testing.T) {
		timeStamp := GetZeroTimeStamp().Add(time.Second)
		macString := "b4:0e:de:fc:1b:ca"
		macAddrArr := strings.Split(macString, ":")
		macAddr := strings.Join(macAddrArr, "")
		mac, _ := strconv.ParseUint(macAddr, 16, 64)
		uid := GenerateUID(mac, timeStamp)
		if uid != 109951162778600 {
			t.Errorf("unexpected uid: %d, expected 109951162778600", uid)
		}
	})
	t.Run("TestUIDGenerator-2", func(t *testing.T) {
		timeStamp := GetZeroTimeStamp().Add(time.Second)
		macString := "b40edefc1bca"
		macAddrArr := strings.Split(macString, ":")
		macAddr := strings.Join(macAddrArr, "")
		mac, _ := strconv.ParseUint(macAddr, 16, 64)
		uid := GenerateUID(mac, timeStamp)
		if uid != 109951162778600 {
			t.Errorf("unexpected uid: %d, expected 109951162778600", uid)
		}
	})
	t.Run("TestUIDGenerator-3", func(t *testing.T) {
		timeStamp := GetZeroTimeStamp().Add(time.Second)
		macString := "b40edefc1b:ca"
		macAddrArr := strings.Split(macString, ":")
		macAddr := strings.Join(macAddrArr, "")
		mac, _ := strconv.ParseUint(macAddr, 16, 64)
		uid := GenerateUID(mac, timeStamp)
		if uid != 109951162778600 {
			t.Errorf("unexpected uid: %d, expected 109951162778600", uid)
		}
	})
}

func TestRealIDGenerator(t *testing.T) {
	t.Run("TestRealIDGenerator", func(t *testing.T) {
		realID := GenerateRealID(109951162778600, 1)
		if realID != 7205759403858329601 {
			t.Errorf("unexpected real-id: %v, expected 7205759403858329601", realID)
		}
	})
}
