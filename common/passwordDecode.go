package common

import "encoding/base64"

func DecodePass(decode string)string{
	decoded, err := base64.StdEncoding.DecodeString(decode)
	ErrHandler("密码解密失败",err)
	return string(decoded)
}
