package library

import (
	"crypto"
	"crypto/rsa"
	"crypto/rand"
	"encoding/hex"
	"errors"
)


func VerifyXmlByteFromPublicKey(xmlByte []byte, sign string) error {
	// hex2bin
	signByte, err := hex.DecodeString(sign)
	if err != nil {
		return errors.New("hex.DecodeString error : "+err.Error())
	}
	// sha1
	hash := crypto.SHA1
	h := hash.New()
	if _, err = h.Write(xmlByte); err != nil{
		return errors.New("hash.Write error : "+err.Error())
	}
	hashed := h.Sum(nil)
	// verify
	if err = rsa.VerifyPKCS1v15(AlliPayPublicKey, hash, hashed, signByte); err != nil {
		return errors.New("rsa.VerifyPKCS1v15 error : "+err.Error())
	}
	return err
}

// sign
func SignXmlByteFromPrivateKey(xmlByte []byte) (sign string, err error) {
	// sha1
	hash := crypto.SHA1
	h := hash.New()
	if _, err = h.Write(xmlByte); err != nil{
		err = errors.New("hash.Write error : "+err.Error())
		return sign, err
	}
	hashed := h.Sum(nil)
	// sign
	singByte, err := rsa.SignPKCS1v15(rand.Reader, AlliPayPrivateKey, hash, hashed)
	if err != nil {
		err = errors.New("rsa.SignPKCS1v15 error : "+err.Error())
		return sign, err
	}
	// bin2hex
	sign = hex.EncodeToString(singByte)
	// return
	return sign, err
}
