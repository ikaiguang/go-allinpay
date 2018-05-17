package library

import (
	"golang.org/x/text/encoding/simplifiedchinese"
	"errors"
)

// Utf8ToGbk
func GbkToUtf8(gbkByte []byte) ([]byte, error) {
	b, err := simplifiedchinese.GBK.NewDecoder().Bytes(gbkByte)
	if err != nil {
		err = errors.New("simplifiedchinese.GBK.NewDecoder().Bytes() error : " + err.Error())
	}
	return b, err
}

// Utf8ToGbk
func Utf8ToGbk(utf8Byte []byte) ([]byte, error) {
	b, err := simplifiedchinese.GBK.NewEncoder().Bytes(utf8Byte)
	if err != nil {
		err = errors.New("simplifiedchinese.GBK.NewEncoder().Bytes() error : " + err.Error())
	}
	return b, err
}
