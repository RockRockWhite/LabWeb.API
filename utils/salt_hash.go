package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"time"
)

// GenerateSalt 生成随机盐值
func GenerateSalt() string {
	// 发生随机数
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(999_999_999)

	// 获得时间戳
	timestamp := time.Now().Unix()

	// 计算哈希
	saltHash := md5.New()
	if _, err := io.WriteString(saltHash, strconv.FormatInt(int64(randNum), 10)); err != nil {
		panic(fmt.Sprintf("Failed to write string to %v", saltHash))
	}
	if _, err := io.WriteString(saltHash, strconv.FormatInt(timestamp, 10)); err != nil {
		panic(fmt.Sprintf("Failed to write string to %v", saltHash))
	}
	salt := hex.EncodeToString(saltHash.Sum(nil))

	return salt
}

// EncryptPasswordHash 为密码生成加盐hash
func EncryptPasswordHash(password string, salt string) string {
	// 第一层: MD5加密
	passwordMd5 := fmt.Sprintf("%x", md5.Sum([]byte(password)))
	saltMd5 := fmt.Sprintf("%x", md5.Sum([]byte(salt)))

	// 第二层: SHA256加密
	shaHash := sha256.New()
	shaHash.Write([]byte(passwordMd5))
	shaHash.Write([]byte(saltMd5))

	return fmt.Sprintf("%x", shaHash.Sum(nil))
}

// ValifyPasswordHash 验证密码是否正确
func ValifyPasswordHash(password string, salt string, passwordHash string) bool {
	return EncryptPasswordHash(password, salt) == passwordHash
}
