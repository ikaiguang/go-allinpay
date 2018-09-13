package library

import (
	"encoding/xml"
	"reflect"
	"errors"
	"regexp"
	"strings"
)

// utf8 conversion gbk
func ToAllinPayGbkXmlByte(data interface{}) ([]byte, error) {
	// req to xml
	b, err := xml.Marshal(data)
	if err != nil {
		err = errors.New("xml.Marshal error : " + err.Error())
		return b, err
	}
	// xml header
	xmlString := AllinPayXmlGbkHeader + string(b)
	// utf8 conversion gbk
	return Utf8ToGbk([]byte(xmlString))
}

// set req.INFO.SIGNED_MSG value
// req : request param pointer
func SetAllinPayRequestSignValue(req interface{}, sign string) (err error) {
	// reflect value
	reflectValue := reflect.ValueOf(req)
	// pointer
	if reflectValue.Kind() != reflect.Ptr {
		return errors.New("SetAllinPayRequestSignValue request interface not a pointer type : reflect.Ptr")
	}
	// info reflect value
	infoValue := reflectValue.Elem().FieldByName("INFO")
	// Info
	if ! infoValue.IsValid() {
		return errors.New("SetAllinPayRequestSignValue error : INFO flied not exist")
	}
	// pointer
	//if infoValue.Kind() != reflect.Ptr {
	//	return errors.New("SetAllinPayRequestSignValue infoValue not a pointer type : reflect.Ptr")
	//}
	// signValue
	//signValue := infoValue.Elem().FieldByName("SIGNED_MSG")
	// not pointer
	signValue := infoValue.FieldByName("SIGNED_MSG")
	// SIGNED_MSG
	if ! signValue.IsValid() {
		return errors.New("SetAllinPayRequestSignValue error : SIGNED_MSG flied not exist")
	}
	// string
	if signValue.Kind() != reflect.String {
		return errors.New("SetAllinPayRequestSignValue signValue not a string type : reflect.String")
	}
	// assignment value
	signValue.SetString(sign)
	// return
	return err
}

// request to AllinPay xml
func ToAllinPayRequestXmlByte(req interface{}) (r []byte, err error) {
	// xml
	xmlByte, err := ToAllinPayGbkXmlByte(req)
	if err != nil {
		err = errors.New("req ToAllinPayGbkXmlByte error : " + err.Error())
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
	if err = SetAllinPayRequestSignValue(req, sign); err != nil {
		err = errors.New("req AssignmentSignValue error : " + err.Error())
		return r, err
	}
	// return
	return ToAllinPayGbkXmlByte(req)
}

// get sign string then set sign empty value
func GetSignStrAndThenSetSignStrEmpty(res interface{}) (sign string, err error) {
	// reflect value
	reflectValue := reflect.ValueOf(res)
	// pointer
	if reflectValue.Kind() != reflect.Ptr {
		err = errors.New("GetSignStrAndThenSetSignStrEmpty result interface not a pointer type : reflect.Ptr")
		return sign, err
	}
	// info reflect value
	infoValue := reflectValue.Elem().FieldByName("INFO")
	// Info
	if ! infoValue.IsValid() {
		err = errors.New("GetSignStrAndThenSetSignStrEmpty error : INFO flied not exist")
		return sign, err
	}
	// pointer
	//if infoValue.Kind() != reflect.Ptr {
	//	return errors.New("GetSignStrAndThenSetSignStrEmpty infoValue not a pointer type : reflect.Ptr")
	//}
	// signValue
	//signValue := infoValue.Elem().FieldByName("SIGNED_MSG")
	// not pointer
	signValue := infoValue.FieldByName("SIGNED_MSG")
	// SIGNED_MSG
	if ! signValue.IsValid() {
		err = errors.New("GetSignStrAndThenSetSignStrEmpty error : SIGNED_MSG flied not exist")
		return sign, err
	}
	// string
	if signValue.Kind() != reflect.String {
		err = errors.New("GetSignStrAndThenSetSignStrEmpty signValue not a string type : reflect.String")
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
func VerifyAndSetAllinPayResponse(bodyByte []byte, res interface{}) (err error) {
	// gbk to utf8
	utf8Byte, err := GbkToUtf8(bodyByte)
	if err != nil {
		return errors.New("GbkToUtf8 error : " + err.Error())
	}
	// replace gbk header to utf8 header
	utf8Byte = []byte(strings.Replace(string(utf8Byte), AllinPayXmlGbkHeader, AllinPayXmlUtf8Header, 1))
	// decode
	if err = xml.Unmarshal(utf8Byte, res); err != nil {
		return errors.New("xml.Unmarshal error : " + err.Error())
	}
	// get sign string and set empty value
	signString, err := GetSignStrAndThenSetSignStrEmpty(res)
	if err != nil {
		return errors.New("GetSignStrAndThenSetSignStrEmpty error : " + err.Error())
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
		return errors.New("ToAllinPayGbkXmlByte error : " + err.Error())
	}
	return err
}
