package implement

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/util"
	"time"
)

var (
	modID     uint64 = 0
	debugFlag        = false
	maxID     uint16 = 1<<16 - 1
	lastID    uint16 = 0
	usedID           = map[uint16]bool{}
)

func initModID() {
	if !debugFlag {
		if macAddresses, err := util.GetMacAddresses(); err != nil {
			panic("couldn't get mac addresses")
		} else if len(macAddresses) == 0 {
			panic("couldn't get mac addresses")
		} else if uintMacAddress, err := util.GetUintMacAddress(macAddresses[0]); err != nil {
			panic("couldn't get mac address'")
		} else {
			modID = util.GenerateUID(uintMacAddress, time.Now())
		}
	}
}

func SetDebugFlag(flag bool) {
	debugFlag = flag
}

func ModID() uint64 {
	if debugFlag {
		return 109951162778600
	} else if modID != 0 {
		return modID
	} else {
		initModID()
		return modID
	}
}

// NextID 使用内置的ID分配工具获取一个可用的不与其余被托管ID冲突的自增ID
func NextID() uint16 {
	if !usedID[lastID+1] {
		lastID += 1
		usedID[lastID] = true
		return lastID
	} else {
		for i := uint16(1); i <= maxID; i++ {
			if !usedID[i] {
				lastID = i
				usedID[i] = true
				return i
			}
		}
	}

	panic("entity id overflow")
}

// UseID 使用内置的ID分配工具分配一个指定的ID，若不可用，则分配一个自增的ID
func UseID(want uint16) (success bool, result uint16) {
	if !usedID[want] {
		usedID[want] = true
		return true, want
	} else {
		return false, NextID()
	}
}
