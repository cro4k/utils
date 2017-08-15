/**
 * Created by 93201 on 2017/5/26.
 */
package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
)

func MD5(data []byte) []byte {
	md5Ctx := md5.New()
	md5Ctx.Write(data)
	return md5Ctx.Sum(nil)
}

func MD5String(data []byte) string {
	return hex.EncodeToString(MD5(data))
}

func Sha1(data []byte) []byte {
	hash := sha1.New()
	hash.Write(data)
	return hash.Sum(nil)
}

func Sha1String(data []byte) string {
	return hex.EncodeToString(Sha1(data))
}
