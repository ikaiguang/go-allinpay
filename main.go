package main

import (
	_ "github.com/ikaiguang/go-allinpay/config"
	"github.com/ikaiguang/go-allinpay/library"
	"fmt"
)

func main() {
	// 协议签约
	TestAgreementSigningSms()
	// 协议支付
	TestAgreementSigningPay()
}

func TestAgreementSigningPay() {
	// 请求头
	headerReq := &library.AllinPayReqHeaderReq{
		TRX_CODE:   "310011",                            // 交易代码
		LEVEL:      "6",                                 // 处理级别（0-9  0优先级最低，默认为5）
		REQ_SN:     "200604000000445-20180518-99999999", // 交易流水号（必须全局唯一）
		SIGNED_MSG: "",                                  // 签名信息
	}
	// 请求参数
	payReq := &library.AgreementSigningPayReq{
		INFO: *library.InitAllinPayReqHeader(headerReq),
		FASTTRX: library.AgreementSigningPayReqFASTTRX{
			BUSINESS_CODE: "01",                            // 业务代码
			MERCHANT_ID:   library.AllinPayCfg.MerchantCode, // 商户代码
			SUBMIT_TIME:   "20180518121212",                // 提交时间（YYYYMMDDHHMMSS）
			AGRMNO:        "abc",                           // 协议号（签约时返回的协议号）
			ACCOUNT_NO:    "6217000010064449999",           // 账号（借记卡或信用卡）
			ACCOUNT_NAME:  "小明",                            // 账号名（借记卡或信用卡上的所有人姓名）
			AMOUNT:        "12300",                         // 金额(整数，单位分)
			CURRENCY:      "CNY",                           // 货币类型(人民币：CNY, 港元：HKD，美元：USD。不填时，默认为人民币)
			ID_TYPE:       "0",                             // 开户证件类型（0身份证，1户口簿，2护照，3军官证，4士兵证...）
			ID:            "440882199909099999",            // 证件号
			TEL:           "15992122300",                   // 手机号
			CUST_USERID:   "CUST_USERID-15992122300",       // 自定义用户号（商户自定义的用户号，开发人员可当作备注字段使用）
			SUMMARY:       "SUMMARY-15992122300",           // 交易附言（填入网银的交易备注）
			REMARK:        "REMARK-15992122300",            // 备注（供商户填入参考信息）
		},
	}
	// xml
	reqXmlByte, err := library.ToAllinPayRequestXmlByte(payReq)
	if err != nil {
		fmt.Println("library.ToAllinPayRequestXmlByte error : ", err)
		return
	}
	//fmt.Println(string(reqXmlByte))
	// post
	bodyByte, err := library.PostAllinPayXmlByte(library.AllinPayAddress, reqXmlByte)
	if err != nil {
		fmt.Println("library.PostAllinPayXmlByte error : ", err)
		return
	}
	//fmt.Println(string(bodyByte))
	// verify and assignment result
	payRes := &library.SigningAgreementPayRes{}
	if err = library.VerifyAndSetAllinPayResponse(bodyByte, payRes); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", payRes)
}

func TestAgreementSigningSms() {
	// 请求头
	headerReq := &library.AllinPayReqHeaderReq{
		TRX_CODE:   "310001",                           // 交易代码
		LEVEL:      "6",                                // 处理级别（0-9  0优先级最低，默认为5）
		REQ_SN:     "200604000000445-2018-05-21TestSn", // 交易流水号（必须全局唯一）
		SIGNED_MSG: "",                                 // 签名信息
	}
	// 请求参数
	smsReq := &library.AgreementSigningSmsReq{
		INFO: *library.InitAllinPayReqHeader(headerReq),
		FAGRA: library.AgreementSigningSmsReqFAGRA{
			MERCHANT_ID:  library.AllinPayCfg.MerchantCode,      // 商户代码
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
	reqXmlByte, err := library.ToAllinPayRequestXmlByte(smsReq)
	if err != nil {
		fmt.Println("library.ToAllinPayRequestXmlByte error : ", err)
		return
	}
	//fmt.Println(string(reqXmlByte))
	// post
	bodyByte, err := library.PostAllinPayXmlByte(library.AllinPayAddress, reqXmlByte)
	if err != nil {
		fmt.Println("library.PostAllinPayXmlByte error : ", err)
		return
	}
	//fmt.Println(string(bodyByte))
	// verify and assignment result
	smsRes := &library.SigningAgreementSmsRes{}
	if err = library.VerifyAndSetAllinPayResponse(bodyByte, smsRes); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", smsRes)
}
