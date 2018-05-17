package main

import (
	_ "./config"
	"./library"
	"fmt"
)

func main() {
	AgreementSigningSms()
}

func AgreementSigningSms() {
	// 请求头
	headerReq := &library.AlliPayReqHeaderReq{
		TRX_CODE:   "310001",                             // 交易代码
		LEVEL:      "6",                                  // 处理级别（0-9  0优先级最低，默认为5）
		REQ_SN:     "200604000000445-rrrr1356732135xxxx1", // 交易流水号（必须全局唯一）
		SIGNED_MSG: "",                                   // 签名信息
	}
	// 请求参数
	smsReq := &library.AgreementSigningSmsReq{
		INFO: *library.InitAlliPayReqHeader(headerReq),
		FAGRA: library.AgreementSigningSmsReqFAGRA{
			MERCHANT_ID:  library.AlliPayCfg.MerchantCode,      // 商户代码
			BANK_CODE:    "105",                                // 银行代码
			ACCOUNT_TYPE: "00",                                 // 账号类型：00借记卡，02信用卡
			ACCOUNT_NO:   "6217000010064449999",                // 账号（借记卡或信用卡）
			ACCOUNT_NAME: "小明",                                 // 账号名（借记卡或信用卡上的所有人姓名）
			ACCOUNT_PROP: "0",                                  // 账号属性（0私人，1公司。不填时，默认为私人0）
			ID_TYPE:      "0",                                  // 开户证件类型（0身份证，1户口簿，2护照，3军官证，4士兵证...）
			ID:           "440882199909099999",                 // 证件号
			TEL:          "15992122300",                        // 手机号
			CVV2:         "",                                   // CVV2（信用卡时必填）
			VAILDDATE:    "",                                   // 有效期（信用卡时必填，格式MMYY（信用卡上的两位月两位年））
			MERREM:       "商户保留信息，建行：6217000010064449999（105）", // 商户保留信息（商户保留信息）
			REMARK:       "备注，身份证：440882199909099999",          // 备注（供商户填入参考信息）
		},
	}
	// xml
	reqXmlByte, err := library.ToAlliPayRequestXmlByte(smsReq)
	if err != nil {
		fmt.Println("library.ToAlliPayRequestXmlByte error : ", err)
		return
	}
	// post
	bodyByte, err := library.PostAlliPayXmlByte(library.AlliPayCfg.TestAddress, reqXmlByte)
	if err != nil {
		fmt.Println("library.PostAlliPayXmlByte error : ", err)
		return
	}
	// verify and assignment result
	smsRes := &library.SigningAgreementSmsRes{}
	if err = library.VerifyAndSetAlliPayResponse(bodyByte, smsRes); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", smsRes)
}
