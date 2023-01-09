package util

import "golang.org/x/crypto/scrypt"

// EncodeRandomSalt 对任意类型的key进行编码生成一个长度为8的[]byte结果
func EncodeRandomSalt[PK any](key PK) []byte {
	hashed := GenerateHash(key)
	salt := make([]byte, 8)
	for i := 0; i < 8; i++ {
		salt[i] = byte(hashed % 256)
		hashed = hashed >> 8
	}

	return salt
}

// EncodePassword 使用crypto/scrypt密钥库对密码进行单向加密
func EncodePassword[PK any](original []byte, key PK) (success bool, result []byte) {
	if encrypted, err := scrypt.Key(original, EncodeRandomSalt(key), 11451, 8, 1, 64); err != nil {
		return false, []byte{}
	} else {
		return true, encrypted
	}
}
