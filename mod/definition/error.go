package definition

import (
	"fmt"
	"net"
)

var (
	GetMacAddressFailedError error = GetMacAddressFailed{}
)

type baseError struct {
	errInfo string
}

func (b baseError) Error() string {
	return b.errInfo
}

type GetMacAddressFailed struct {
	baseError
}

func GetMacAddressFailedWithError(err error) GetMacAddressFailed {
	return GetMacAddressFailed{baseError{errInfo: fmt.Sprintf("get mac address failed with error: %v", err)}}
}

func GetMacAddressFailedWithEmptyMac() GetMacAddressFailed {
	return GetMacAddressFailed{baseError{errInfo: fmt.Sprintf("get mac address failed with empty mac address")}}
}

func GetMacAddressFailedWithIncorrectMacAddress(mac net.Interface, err error) GetMacAddressFailed {
	return GetMacAddressFailed{baseError{errInfo: fmt.Sprintf("get mac address failed with incorrect mac %v, error: %v", mac.HardwareAddr.String(), err)}}
}

type GetPackagePathFailed struct {
	baseError
}

func GetPackagePathFunctionPointerFailed() GetPackagePathFailed {
	return GetPackagePathFailed{baseError{errInfo: fmt.Sprintf("cannot get caller's function pointer")}}
}
