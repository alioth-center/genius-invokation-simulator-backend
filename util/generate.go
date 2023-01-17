package util

import (
	"crypto/md5"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"reflect"
	"unsafe"
)

// GenerateUUID 生成一个UUID，长度为36
func GenerateUUID() string {
	return uuid.Must(uuid.NewV4(), nil).String()
}

// GenerateTypeID 根据entity的包和结构名，生成类型ID
func GenerateTypeID[T any](entity T) (uid string) {
	typesOfT := reflect.TypeOf(entity)
	return fmt.Sprintf("%s@%s", typesOfT.PkgPath(), typesOfT.Name())
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
func GenerateHashWithOpts[Key any, Hash uint | uint16 | uint32 | uint64](key Key) (hash Hash) {
	entityPtr := &key
	hash = Hash(0)
	start := uintptr(unsafe.Pointer(entityPtr))
	end := unsafe.Sizeof(key) + start
	offset := unsafe.Sizeof(byte(0))

	for i := start; i < end; i += offset {
		byteData := *(*byte)(unsafe.Pointer(i))
		hash = Hash(byteData) + (hash << 6) + (hash << 16) - hash
	}

	return hash
}

// GeneratePrefixHashWithOpts 将任意结构的前offset字节的内容进行哈希，生成一个指定类型(无符号整形)的哈希值，越界部份不会被计算，使用SDBM算法作为实现
func GeneratePrefixHashWithOpts[Key any, Hash uint | uint16 | uint32 | uint64](key Key, offset uintptr) (hash Hash) {
	entityPtr := &key
	hash = Hash(0)
	start := uintptr(unsafe.Pointer(entityPtr))
	end := start + offset
	if end > unsafe.Sizeof(key)+start {
		end = unsafe.Sizeof(key) + start
	}

	byteOffset := unsafe.Sizeof(byte(0))
	for i := start; i < end; i += byteOffset {
		b := *(*byte)(unsafe.Pointer(i))
		hash = Hash(b) + (hash << 6) + (hash << 16) - hash
	}

	return hash
}

// GenerateMD5 将给定的字符串使用MD5摘要算法生成摘要
func GenerateMD5(source string) (md5CheckSum string) {
	return fmt.Sprintf("%x", md5.Sum([]byte(source)))
}
