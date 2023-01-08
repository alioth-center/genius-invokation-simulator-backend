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
	entityPtr := &key
	hash = uint(0)
	start := uintptr(unsafe.Pointer(entityPtr))
	end := unsafe.Sizeof(key) + start
	offset := unsafe.Sizeof(byte(0))

	for i := start; i < end; i += offset {
		b := *(*byte)(unsafe.Pointer(i))
		hash = uint(b) + (hash << 6) + (hash << 16) - hash
	}

	return hash
}

// GeneratePrefixHash 将任意结构的前offset字节的内容进行哈希，越界部份不会被计算，使用SDBM算法作为实现
func GeneratePrefixHash[Key any](key Key, offset uintptr) (hash uint) {
	entityPtr := &key
	hash = uint(0)
	start := uintptr(unsafe.Pointer(entityPtr))
	end := start + offset
	if end > unsafe.Sizeof(key)+start {
		end = unsafe.Sizeof(key) + start
	}

	byteOffset := unsafe.Sizeof(byte(0))
	for i := start; i < end; i += byteOffset {
		b := *(*byte)(unsafe.Pointer(i))
		hash = uint(b) + (hash << 6) + (hash << 16) - hash
	}

	return hash
}

func GenerateMD5(source string) (md5CheckSum string) {
	hash := md5.New()
	hash.Write([]byte(source))
	return string(hash.Sum(nil))
}
