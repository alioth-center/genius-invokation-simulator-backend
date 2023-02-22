package util

import (
	"net"
	"strconv"
	"strings"
	"time"
)

type SystemBits byte

const (
	BitsUnknown SystemBits = 0
	Bits32      SystemBits = 32
	Bits64      SystemBits = 64
)

var (
	zeroTimeStamp = new(time.Time)
)

func init() {
	// 初始化“零点”，为本项目的创建时间
	utc8timeZone := time.FixedZone("UTC+8", 8*60*60)
	*zeroTimeStamp = time.Date(2022, 12, 9, 23, 35, 0, 0, utc8timeZone)
}

// GetZeroTimeStamp 获取系统的零时
func GetZeroTimeStamp() time.Time {
	return *zeroTimeStamp
}

// GetSystemBits 获取操作系统的位数，结果为SystemBits
func GetSystemBits() SystemBits {
	bit := 32 << ((^uint(0)) >> 63)

	switch bit {
	case 32:
		return Bits32
	case 64:
		return Bits64
	default:
		return BitsUnknown
	}
}

// GetMacAddresses 使用net包获取本机mac地址
func GetMacAddresses() (macAddr []net.Interface, err error) {
	return net.Interfaces()
}

// GetUintMacAddress 从net.Interface中获取一个uint64类型的地址，长度为48位
func GetUintMacAddress(mac net.Interface) (addr uint64, err error) {
	macAddrArr := strings.Split(mac.HardwareAddr.String(), ":")
	macAddr := strings.Join(macAddrArr, "")

	// mac地址为空，考虑到权限问题，设置mac地址为随机变量
	if macAddr == "" {
		macPartBits := uint64(1<<48) - 1
		addr = GenerateHashWithOpts[net.Interface, uint64](mac)
		return addr & macPartBits, nil
	}

	// mac地址非空，则正常计算mac的48位数值
	if addr, err = strconv.ParseUint(macAddr, 16, 64); err != nil {
		return 0, err
	} else {
		macPartBits := uint64(1<<48) - 1
		return addr & macPartBits, nil
	}
}
