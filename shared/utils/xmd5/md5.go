package xmd5

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

func MD5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return hex.EncodeToString(h.Sum(nil))
}
