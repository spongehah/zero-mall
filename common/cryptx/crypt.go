package cryptx

import (
	"fmt"

	"golang.org/x/crypto/scrypt"
)

// PasswordEncrypt 密码加密：给需要存储的密码加盐，转换为不可逆的散列值
func PasswordEncrypt(salt, password string) string {
	dk, _ := scrypt.Key([]byte(password), []byte(salt), 32768, 8, 1, 32)
	return fmt.Sprintf("%x", dk)
}
