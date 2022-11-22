package xmd5_test

import (
	"due-mahjong-server/shared/utils/xmd5"
	"testing"
)

func TestMD5(t *testing.T) {
	t.Log(xmd5.MD5("123456"))
}
