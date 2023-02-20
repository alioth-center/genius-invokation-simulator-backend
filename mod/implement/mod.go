package implement

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/util"
	"time"
)

var (
	modID     uint64 = 0
	debugFlag        = false
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
