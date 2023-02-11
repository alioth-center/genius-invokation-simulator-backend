package definition

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/util"
	"reflect"
	"runtime"
	"time"
)

type ModInfo interface {
	ID() uint64
	Name() string
	PackagePath() string
}

type modInfo struct {
	id   uint64
	name string
	pkg  string
}

func (m modInfo) ID() uint64 {
	return m.id
}

func (m modInfo) Name() string {
	return m.name
}

func (m modInfo) PackagePath() string {
	return m.pkg
}

func NewModInfo(name string) (info ModInfo, err error) {
	_info := &modInfo{}
	if pc, _, _, ok := runtime.Caller(1); !ok {
		return nil, GetPackagePathFunctionPointerFailed()
	} else {
		reflect.TypeOf(pc)
	}

	if macs, getMacErr := util.GetMacAddresses(); getMacErr != nil {
		return nil, GetMacAddressFailedWithError(getMacErr)
	} else if len(macs) == 0 {
		return nil, GetMacAddressFailedWithEmptyMac()
	} else if addr, convertMacErr := util.GetUintMacAddress(macs[0]); convertMacErr != nil {
		return nil, GetMacAddressFailedWithIncorrectMacAddress(macs[0], convertMacErr)
	} else {
		_info.id = util.GenerateUID(addr, time.Now())
	}

	return _info, nil
}
