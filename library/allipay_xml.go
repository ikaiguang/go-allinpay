package library

import (
	"encoding/xml"
	"reflect"
	"errors"
	"regexp"
	"strings"
)

// utf8 conversion gbk
func ToAlliPayGbkXmlByte(data interface{}) ([]byte, error) {
	// req to xml
	b, err := xml.Marshal(data)
	if err != nil {
		err = errors.New("xml.Marshal error : " + err.Error())
		return b, err
	}
	// xml header
	xmlString := AlliPayXmlGbkHeader + string(b)
	// utf8 conversion gbk
	return Utf8ToGbk([]byte(xmlString))
}

// set req.INFO.SIGNED_MSG value
// req : request param pointer
func SetAlliPayRequestSignValue(req interface{}, sign string) (err error) {
	// reflect value
	reflectValue := reflect.ValueOf(req)
	// pointer
	if reflectValue.Kind() != reflect.Ptr {
		return errors.New("SetAlliPayRequestSignValue request interface not a pointer type : reflect.Ptr")
	}
	// info reflect value
	infoValue := reflectValue.Elem().FieldByName("INFO")
	// Info
	if ! infoValue.IsValid() {
		return errors.New("SetAlliPayRequestSignValue error : INFO flied not exist")
	}
	// pointer
	//if infoValue.Kind() != reflect.Ptr {
	//	return errors.New("SetAlliPayRequestSignValue infoValue not a pointer type : reflect.Ptr")
	//}
	// signValue
	//signValue := infoValue.Elem().FieldByName("SIGNED_MSG")
	// not pointer
	signValue := infoValue.FieldByName("SIGNED_MSG")
	// SIGNED_MSG
	if ! signValue.IsValid() {
		return errors.New("SetAlliPayRequestSignValue error : SIGNED_MSG flied not exist")
	}
	// string
	if signValue.Kind() != reflect.String {
		return errors.New("SetAlliPayRequestSignValue signValue not a string type : reflect.String")
	}
	// assignment value
	signValue.SetString(sign)
	// return
	return err
}

// request to alliPay xml
func ToAlliPayRequestXmlByte(req interface{}) (r []byte, err error) {
	// xml
	xmlByte, err := ToAlliPayGbkXmlByte(req)
	if err != nil {
		err = errors.New("req ToAlliPayGbkXmlByte error : " + err.Error())
		return r, err
	}
	//fmt.Println(string(xmlByte))
	// sign
	sign, err := SignXmlByteFromPrivateKey(xmlByte)
	if err != nil {
		err = errors.New("req SignXmlByteFromPrivateKey error : " + err.Error())
		return r, err
	}
	// AssignmentSignValue
	if err = SetAlliPayRequestSignValue(req, sign); err != nil {
		err = errors.New("req AssignmentSignValue error : " + err.Error())
		return r, err
	}
	// return
	return ToAlliPayGbkXmlByte(req)
}

// get sign string then set sign empty value
func GetThenEmptyAlliPayResultSignValue(res interface{}) (sign string, err error) {
	// reflect value
	reflectValue := reflect.ValueOf(res)
	// pointer
	if reflectValue.Kind() != reflect.Ptr {
		err = errors.New("GetThenEmptyAlliPayResultSignValue result interface not a pointer type : reflect.Ptr")
		return sign, err
	}
	// info reflect value
	infoValue := reflectValue.Elem().FieldByName("INFO")
	// Info
	if ! infoValue.IsValid() {
		err = errors.New("GetThenEmptyAlliPayResultSignValue error : INFO flied not exist")
		return sign, err
	}
	// pointer
	//if infoValue.Kind() != reflect.Ptr {
	//	return errors.New("GetThenEmptyAlliPayResultSignValue infoValue not a pointer type : reflect.Ptr")
	//}
	// signValue
	//signValue := infoValue.Elem().FieldByName("SIGNED_MSG")
	// not pointer
	signValue := infoValue.FieldByName("SIGNED_MSG")
	// SIGNED_MSG
	if ! signValue.IsValid() {
		err = errors.New("GetThenEmptyAlliPayResultSignValue error : SIGNED_MSG flied not exist")
		return sign, err
	}
	// string
	if signValue.Kind() != reflect.String {
		err = errors.New("GetThenEmptyAlliPayResultSignValue signValue not a string type : reflect.String")
		return sign, err
	}
	// get sign
	sign = signValue.String()
	// assignment value
	signValue.SetString("")
	// return
	return sign, err
}

// verify and set response
func VerifyAndSetAlliPayResponse(bodyByte []byte, res interface{}) (err error) {
	// gbk to utf8
	utf8Byte, err := GbkToUtf8(bodyByte)
	if err != nil {
		return errors.New("GbkToUtf8 error : " + err.Error())
	}
	// replace gbk header to utf8 header
	utf8Byte = []byte(strings.Replace(string(utf8Byte), AlliPayXmlGbkHeader, AlliPayXmlUtf8Header, 1))
	// decode
	if err = xml.Unmarshal(utf8Byte, res); err != nil {
		return errors.New("xml.Unmarshal error : " + err.Error())
	}
	// get sign string and set empty value
	signString, err := GetThenEmptyAlliPayResultSignValue(res)
	if err != nil {
		return errors.New("GetThenEmptyAlliPayResultSignValue error : " + err.Error())
	}
	// 提换加密字符串
	replacePattern := `<SIGNED_MSG>.*<\/SIGNED_MSG>`
	replaceRegexp, err := regexp.Compile(replacePattern)
	if err != nil {
		return errors.New("regexp.Compile error : " + err.Error())
	}
	verifyByte := []byte(replaceRegexp.ReplaceAllString(string(bodyByte), ""))
	// verify
	if err = VerifyXmlByteFromPublicKey(verifyByte, signString); err != nil {
		return errors.New("ToAlliPayGbkXmlByte error : " + err.Error())
	}
	return err
}
