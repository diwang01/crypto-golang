package tools

import (
	"crypto/md5"
	"encoding/hex"
)

const (
	base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
)

/** md5 encrypt */
func Md5Encrypt(data string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(data))
	cipherStr := md5Ctx.Sum(nil)
	encryptStr := hex.EncodeToString(cipherStr)
	return encryptStr
}
