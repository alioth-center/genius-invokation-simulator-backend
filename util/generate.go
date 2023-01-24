package util

import (
	"crypto/md5"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"reflect"
	"time"
	"unsafe"
)

type UintLike interface {
	uint | uint8 | uint16 | uint32 | uint64
}

// GenerateUUID 生成一个UUID，长度为36
func GenerateUUID() string {
	return uuid.Must(uuid.NewV4(), nil).String()
}

// GenerateTypeID 根据entity的包和结构名，生成类型ID
func GenerateTypeID[T any](entity T) (uid string) {
	typesOfT := reflect.TypeOf(entity)
	return fmt.Sprintf("%s@%s", typesOfT.PkgPath(), typesOfT.Name())
}

// GeneratePackageID 根据给定的包名，生成其Hash，不保证不会冲突
func GeneratePackageID[ID UintLike](packageName string) (id ID) {
	return GenerateHashWithOpts[string, ID](packageName)
}

// GenerateHash 将任意结构进行哈希，使用SDBM算法作为实现
func GenerateHash[Key any](key Key) (hash uint) {
	return GenerateHashWithOpts[Key, uint](key)
}

// GeneratePrefixHash 将任意结构的前offset字节的内容进行哈希，越界部份不会被计算，使用SDBM算法作为实现
func GeneratePrefixHash[Key any](key Key, offset uintptr) (hash uint) {
	return GeneratePrefixHashWithOpts[Key, uint](key, offset)
}

// GenerateHashWithOpts 将任意结构进行哈希，生成一个指定类型(无符号整形)的哈希值，使用SDBM算法作为实现
func GenerateHashWithOpts[Key any, Hash UintLike](key Key) (hash Hash) {
	entityPtr := &key
	sum := uint64(0)
	start := uintptr(unsafe.Pointer(entityPtr))
	end := unsafe.Sizeof(key) + start
	offset := unsafe.Sizeof(byte(0))

	for i := start; i < end; i += offset {
		byteData := *(*byte)(unsafe.Pointer(i))
		sum = uint64(byteData) + (sum << 6) + (sum << 16) - sum
	}

	return Hash(sum)
}

// GeneratePrefixHashWithOpts 将任意结构的前offset字节的内容进行哈希，生成一个指定类型(无符号整形)的哈希值，越界部份不会被计算，使用SDBM算法作为实现
func GeneratePrefixHashWithOpts[Key any, Hash UintLike](key Key, offset uintptr) (hash Hash) {
	entityPtr := &key
	sum := uint64(0)
	start := uintptr(unsafe.Pointer(entityPtr))
	end := start + offset
	if end > unsafe.Sizeof(key)+start {
		end = unsafe.Sizeof(key) + start
	}

	byteOffset := unsafe.Sizeof(byte(0))
	for i := start; i < end; i += byteOffset {
		b := *(*byte)(unsafe.Pointer(i))
		sum = uint64(b) + (sum << 6) + (sum << 16) - sum
	}

	return Hash(sum)
}

// GenerateMD5 将给定的字符串使用MD5摘要算法生成摘要
func GenerateMD5(source string) (md5CheckSum string) {
	return fmt.Sprintf("%x", md5.Sum([]byte(source)))
}

// GenerateUID 根据时间戳生成一个UID，毫秒级，48位，其中前6位为设备序列号，后42位为毫秒时间戳，可使用138年
func GenerateUID(mac uint64, timeStamp time.Time) (uid uint64) {
	timeStampID := uint64(timeStamp.Sub(GetZeroTimeStamp()).Milliseconds())
	timePartBits := uint64(1<<42) - 1
	deviceID := GenerateHashWithOpts[uint64, uint8](mac)
	devicePart := uint64(1<<6) - 1
	return ((uint64(deviceID) & devicePart) << 42) + (timeStampID & timePartBits)
}

// GenerateRealID 根据MOD的ID和MOD中的ID，拼接出真实ID
func GenerateRealID(uid uint64, sub uint16) (realID uint64) {
	uidPartBits := uint64(1<<48) - 1
	subPartBits := uint64(1<<16) - 1
	return ((uid & uidPartBits) << 16) + (uint64(sub) & subPartBits)
}
