package library

import "encoding/xml"

const (
	AlliPayXmlGbkHeader  = `<?xml version="1.0" encoding="GBK"?>`
	AlliPayXmlUtf8Header = `<?xml version="1.0" encoding="UTF-8"?>`
)

// 通联 请求头部信息
type AlliPayReqINFO struct {
	TRX_CODE    string `xml:"TRX_CODE,omitempty"`    // 交易代码
	VERSION     string `xml:"VERSION,omitempty"`     // 版本（04）
	DATA_TYPE   string `xml:"DATA_TYPE,omitempty"`   // 数据格式（2：xml格式）
	LEVEL       string `xml:"LEVEL,omitempty"`       // 处理级别（0-9  0优先级最低，默认为5）
	MERCHANT_ID string `xml:"MERCHANT_ID,omitempty"` // 商户代码
	USER_NAME   string `xml:"USER_NAME,omitempty"`   // 用户名
	USER_PASS   string `xml:"USER_PASS,omitempty"`   // 用户密码
	REQ_SN      string `xml:"REQ_SN,omitempty"`      // 交易流水号（必须全局唯一）
	SIGNED_MSG  string `xml:"SIGNED_MSG,omitempty"`  // 签名信息
}

// 通联 响应头部信息
type AlliPayResINFO struct {
	TRX_CODE   string `xml:"TRX_CODE,omitempty"`   // 交易代码
	VERSION    string `xml:"VERSION,omitempty"`    // 版本（04）
	DATA_TYPE  string `xml:"DATA_TYPE,omitempty"`  // 数据格式（2：xml格式）
	REQ_SN     string `xml:"REQ_SN,omitempty"`     // 交易流水号(原请求报文的流水号，原样返回)
	RET_CODE   string `xml:"RET_CODE,omitempty"`   // 返回代码
	ERR_MSG    string `xml:"ERR_MSG,omitempty"`    // 错误信息
	SIGNED_MSG string `xml:"SIGNED_MSG,omitempty"` // 签名信息
}

// 协议支付签约短信触发 请求
type AgreementSigningSmsReq struct {
	XMLName xml.Name                    `xml:"AIPG,omitempty"`
	INFO    AlliPayReqINFO              `xml:"INFO,omitempty"`
	FAGRA   AgreementSigningSmsReqFAGRA `xml:"FAGRA,omitempty"`
}

type AgreementSigningSmsReqFAGRA struct {
	MERCHANT_ID  string `xml:"MERCHANT_ID,omitempty"`  // 商户代码
	BANK_CODE    string `xml:"BANK_CODE,omitempty"`    // 银行代码
	ACCOUNT_TYPE string `xml:"ACCOUNT_TYPE,omitempty"` // 账号类型：00借记卡，02信用卡
	ACCOUNT_NO   string `xml:"ACCOUNT_NO,omitempty"`   // 账号（借记卡或信用卡）
	ACCOUNT_NAME string `xml:"ACCOUNT_NAME,omitempty"` // 账号名（借记卡或信用卡上的所有人姓名）
	ACCOUNT_PROP string `xml:"ACCOUNT_PROP,omitempty"` // 账号属性（0私人，1公司。不填时，默认为私人0）
	ID_TYPE      string `xml:"ID_TYPE,omitempty"`      // 开户证件类型（0身份证，1户口簿，2护照，3军官证，4士兵证...）
	ID           string `xml:"ID,omitempty"`           // 证件号
	TEL          string `xml:"TEL,omitempty"`          // 手机号
	CVV2         string `xml:"CVV2,omitempty"`         // CVV2（信用卡时必填）
	VAILDDATE    string `xml:"VAILDDATE,omitempty"`    // 有效期（信用卡时必填，格式MMYY（信用卡上的两位月两位年））
	MERREM       string `xml:"MERREM,omitempty"`       // 商户保留信息（商户保留信息）
	REMARK       string `xml:"REMARK,omitempty"`       // 备注（供商户填入参考信息）
}

// 协议支付签约短信触发 响应
type SigningAgreementSmsRes struct {
	XMLName  xml.Name                        `xml:"AIPG,omitempty"`
	INFO     AlliPayResINFO                  `xml:"INFO,omitempty"`
	FAGRARET SigningAgreementSmsResFAGRARET `xml:"FAGRARET,omitempty"`
}

type SigningAgreementSmsResFAGRARET struct {
	RET_CODE string `xml:"RET_CODE,omitempty"` // 返回代码
	ERR_MSG  string `xml:"ERR_MSG,omitempty"`  // 错误文本
}

// 协议支付签约 请求
type AgreementSigningConfirmReq struct {
	XMLName xml.Name                        `xml:"AIPG,omitempty"`
	INFO    AlliPayReqINFO                  `xml:"INFO,omitempty"`
	FAGRC   AgreementSigningConfirmReqFAGRC `xml:"FAGRC,omitempty"`
}

type AgreementSigningConfirmReqFAGRC struct {
	MERCHANT_ID string `xml:"MERCHANT_ID,omitempty"` // 商户代码
	SRCREQSN    string `xml:"SRCREQSN,omitempty"`    // 原请求流水（对应申请请求报文中的REQ_SN）
	VERCODE     string `xml:"VERCODE,omitempty"`     // 验证码（短信验证码：请兼容4或6位）
}

// 协议支付签约 响应
type SigningAgreementConfirmRes struct {
	XMLName  xml.Name                         `xml:"AIPG,omitempty"`
	INFO     AlliPayResINFO                   `xml:"INFO,omitempty"`
	FAGRCRET SigningAgreementConfirmFAGRCRET `xml:"FAGRCRET,omitempty"`
}

type SigningAgreementConfirmFAGRCRET struct {
	AGRMNO   string `xml:"AGRMNO,omitempty"`   // 协议号（成功时协议号不为空）
	RET_CODE string `xml:"RET_CODE,omitempty"` // 返回代码
	ERR_MSG  string `xml:"ERR_MSG,omitempty"`  // 错误文本
}

// 协议支付解约 请求
type AgreementSigningCancelReq struct {
	XMLName xml.Name                         `xml:"AIPG,omitempty"`
	INFO    AlliPayReqINFO                   `xml:"INFO,omitempty"`
	FAGRCNL AgreementSigningCancelReqFAGRCNL `xml:"FAGRCNL,omitempty"`
}

type AgreementSigningCancelReqFAGRCNL struct {
	MERCHANT_ID string `xml:"MERCHANT_ID,omitempty"` // 商户代码
	ACCOUNT_NO  string `xml:"ACCOUNT_NO,omitempty"`  // 账号（借记卡或信用卡）
	AGRMNO      string `xml:"AGRMNO,omitempty"`      // 协议号（签约时返回的协议号）
}

// 协议支付解约 响应
type SigningAgreementCancelRes struct {
	XMLName    xml.Name                          `xml:"AIPG,omitempty"`
	INFO       AlliPayResINFO                    `xml:"INFO,omitempty"`
	FAGRCNLRET SigningAgreementCancelFAGRCNLRET `xml:"FAGRCNLRET,omitempty"`
}

type SigningAgreementCancelFAGRCNLRET struct {
	RET_CODE string `xml:"RET_CODE,omitempty"` // 返回代码
	ERR_MSG  string `xml:"ERR_MSG,omitempty"`  // 错误文本
}

// 协议支付 请求
type AgreementSigningPayReq struct {
	XMLName xml.Name                      `xml:"AIPG,omitempty"`
	INFO    AlliPayReqINFO                `xml:"INFO,omitempty"`
	FASTTRX AgreementSigningPayReqFASTTRX `xml:"FASTTRX,omitempty"`
}

type AgreementSigningPayReqFASTTRX struct {
	BUSINESS_CODE string `xml:"BUSINESS_CODE,omitempty"` // 业务代码
	MERCHANT_ID   string `xml:"MERCHANT_ID,omitempty"`   // 商户代码
	SUBMIT_TIME   string `xml:"SUBMIT_TIME,omitempty"`   // 提交时间（YYYYMMDDHHMMSS）
	AGRMNO        string `xml:"AGRMNO,omitempty"`        // 协议号（签约时返回的协议号）
	ACCOUNT_NO    string `xml:"ACCOUNT_NO,omitempty"`    // 账号（借记卡或信用卡）
	ACCOUNT_NAME  string `xml:"ACCOUNT_NAME,omitempty"`  // 账号名（借记卡或信用卡上的所有人姓名）
	AMOUNT        string `xml:"AMOUNT,omitempty"`        // 金额(整数，单位分)
	CURRENCY      string `xml:"CURRENCY,omitempty"`      // 货币类型(人民币：CNY, 港元：HKD，美元：USD。不填时，默认为人民币)
	ID_TYPE       string `xml:"ID_TYPE,omitempty"`       // 开户证件类型（0身份证，1户口簿，2护照，3军官证，4士兵证...）
	ID            string `xml:"ID,omitempty"`            // 证件号
	TEL           string `xml:"TEL,omitempty"`           // 手机号
	CVV2          string `xml:"CVV2,omitempty"`          // CVV2（信用卡时必填）
	VAILDDATE     string `xml:"VAILDDATE,omitempty"`     // 有效期（信用卡时必填，格式MMYY（信用卡上的两位月两位年））
	CUST_USERID   string `xml:"CUST_USERID,omitempty"`   // 自定义用户号（商户自定义的用户号，开发人员可当作备注字段使用）
	SUMMARY       string `xml:"SUMMARY,omitempty"`       // 交易附言（填入网银的交易备注）
	REMARK        string `xml:"REMARK,omitempty"`        // 备注（供商户填入参考信息）
}

// 协议支付 响应
type SigningAgreementPayRes struct {
	XMLName    xml.Name                       `xml:"AIPG,omitempty"`
	INFO       AlliPayResINFO                 `xml:"INFO,omitempty"`
	FASTTRXRET SigningAgreementPayFASTTRXRET `xml:"FASTTRXRET,omitempty"`
}

type SigningAgreementPayFASTTRXRET struct {
	RET_CODE    string `xml:"RET_CODE,omitempty"`    // 返回代码
	SETTLE_DAY  string `xml:"SETTLE_DAY,omitempty"`  // 完成日期（YYYYMMDD）
	ERR_MSG     string `xml:"ERR_MSG,omitempty"`     // 错误文本
	ACCT_SUFFIX string `xml:"ACCT_SUFFIX,omitempty"` // 卡号后4位
}

// 直接支付短信触发 请求
type DirectPaySmsReq struct {
	XMLName xml.Name               `xml:"AIPG,omitempty"`
	INFO    AlliPayReqINFO         `xml:"INFO,omitempty"`
	FASTTRX DirectPaySmsReqFASTTRX `xml:"FASTTRX,omitempty"`
}

type DirectPaySmsReqFASTTRX struct {
	BUSINESS_CODE string `xml:"BUSINESS_CODE,omitempty"` // 业务代码
	MERCHANT_ID   string `xml:"MERCHANT_ID,omitempty"`   // 商户代码
	SUBMIT_TIME   string `xml:"SUBMIT_TIME,omitempty"`   // 提交时间（YYYYMMDDHHMMSS）
	ACCOUNT_TYPE  string `xml:"ACCOUNT_TYPE,omitempty"`  // 账号类型（00借记卡，02信用卡。不填默认为借记卡00）
	ACCOUNT_NO    string `xml:"ACCOUNT_NO,omitempty"`    // 账号（借记卡或信用卡）
	ACCOUNT_NAME  string `xml:"ACCOUNT_NAME,omitempty"`  // 账号名（借记卡或信用卡上的所有人姓名）
	BANK_CODE     string `xml:"BANK_CODE,omitempty"`     // 银行代码
	AMOUNT        string `xml:"AMOUNT,omitempty"`        // 金额(整数，单位分)
	CURRENCY      string `xml:"CURRENCY,omitempty"`      // 货币类型(人民币：CNY, 港元：HKD，美元：USD。不填时，默认为人民币)
	ID_TYPE       string `xml:"ID_TYPE,omitempty"`       // 开户证件类型（0身份证，1户口簿，2护照，3军官证，4士兵证...）
	ID            string `xml:"ID,omitempty"`            // 证件号
	TEL           string `xml:"TEL,omitempty"`           // 手机号
	CVV2          string `xml:"CVV2,omitempty"`          // CVV2（信用卡时必填）
	VAILDDATE     string `xml:"VAILDDATE,omitempty"`     // 有效期（信用卡时必填，格式MMYY（信用卡上的两位月两位年））
	REMARK        string `xml:"REMARK,omitempty"`        // 备注（供商户填入参考信息）
}

// 直接支付短信触发 响应
type DirectPaySmsRes struct {
	XMLName  xml.Name                 `xml:"AIPG,omitempty"`
	INFO     AlliPayResINFO           `xml:"INFO,omitempty"`
	TRANSRET DirectPaySmsResTRANSRET `xml:"TRANSRET,omitempty"`
}

type DirectPaySmsResTRANSRET struct {
	RET_CODE string `xml:"RET_CODE,omitempty"` // 返回代码
	ERR_MSG  string `xml:"ERR_MSG,omitempty"`  // 错误文本
}

// 直接支付确认 请求
type DirectPayConfirmReq struct {
	XMLName xml.Name                   `xml:"AIPG,omitempty"`
	INFO    AlliPayReqINFO             `xml:"INFO,omitempty"`
	FASTTRX DirectPayConfirmReqFASTTRX `xml:"FASTTRX,omitempty"`
}

type DirectPayConfirmReqFASTTRX struct {
	SRC_REQ_SN    string `xml:"SRC_REQ_SN,omitempty"`    // 原交易流水号（发起直接支付短信触发的交易流水）
	VER_CODE      string `xml:"VER_CODE,omitempty"`      // 银行验证码
	BUSINESS_CODE string `xml:"BUSINESS_CODE,omitempty"` // 业务代码
	MERCHANT_ID   string `xml:"MERCHANT_ID,omitempty"`   // 商户代码
	SUBMIT_TIME   string `xml:"SUBMIT_TIME,omitempty"`   // 提交时间（YYYYMMDDHHMMSS）
	ACCOUNT_NO    string `xml:"ACCOUNT_NO,omitempty"`    // 账号（借记卡或信用卡，须与触发短信填值一致）
	ACCOUNT_NAME  string `xml:"ACCOUNT_NAME,omitempty"`  // 账号名（借记卡或信用卡上的所有人姓名，须与触发短信填值一致）
	AMOUNT        string `xml:"AMOUNT,omitempty"`        // 金额(整数，单位分，须与触发短信填值一致)
}

// 直接支付确认 响应
type DirectPayConfirmRes struct {
	XMLName  xml.Name                     `xml:"AIPG,omitempty"`
	INFO     AlliPayResINFO               `xml:"INFO,omitempty"`
	TRANSRET DirectPayConfirmResTRANSRET `xml:"TRANSRET,omitempty"`
}

type DirectPayConfirmResTRANSRET struct {
	RET_CODE   string `xml:"RET_CODE,omitempty"`   // 返回代码
	SETTLE_DAY string `xml:"SETTLE_DAY,omitempty"` // 完成日期（YYYYMMDD）
	ERR_MSG    string `xml:"ERR_MSG,omitempty"`    // 错误文本
}

// 退款 请求
type RefundReq struct {
	XMLName xml.Name        `xml:"AIPG,omitempty"`
	INFO    AlliPayReqINFO  `xml:"INFO,omitempty"`
	REFUND  RefundReqREFUND `xml:"REFUND,omitempty"`
}

type RefundReqREFUND struct {
	MERCHANT_ID  string `xml:"MERCHANT_ID,omitempty"`  // 商户代码
	ORGBATCHID   string `xml:"ORGBATCHID,omitempty"`   // 原批次（原交易的REQ_SN）
	ORGFILEID    string `xml:"ORGFILEID,omitempty"`    // 原文件号(原文件号)
	ORGBATCHSN   string `xml:"ORGBATCHSN,omitempty"`   // 原批次序号（原交易的记录序号，原交易为单笔实时交易时填0）
	ACCOUNT_NO   string `xml:"ACCOUNT_NO,omitempty"`   // 账号（借记卡或信用卡）
	ACCOUNT_NAME string `xml:"ACCOUNT_NAME,omitempty"` // 账号名（借记卡或信用卡上的所有人姓名）
	AMOUNT       string `xml:"AMOUNT,omitempty"`       // 金额(整数，单位分)
	REMARK       string `xml:"REMARK,omitempty"`       // 备注（预留）
}

// 退款 响应
type RefundRes struct {
	XMLName  xml.Name           `xml:"AIPG,omitempty"`
	INFO     AlliPayResINFO     `xml:"INFO,omitempty"`
	TRANSRET RefundResTRANSRET `xml:"TRANSRET,omitempty"`
}

type RefundResTRANSRET struct {
	RET_CODE   string `xml:"RET_CODE,omitempty"`   // 返回代码
	SETTLE_DAY string `xml:"SETTLE_DAY,omitempty"` // 清算日期（YYYYMMDD）
	ERR_MSG    string `xml:"ERR_MSG,omitempty"`    // 错误文本
}

// 交易结果查询 请求
type QueryTransReq struct {
	XMLName   xml.Name               `xml:"AIPG,omitempty"`
	INFO      AlliPayReqINFO         `xml:"INFO,omitempty"`
	QTRANSREQ QueryTransReqQTRANSREQ `xml:"QTRANSREQ,omitempty"`
}

type QueryTransReqQTRANSREQ struct {
	MERCHANT_ID string `xml:"MERCHANT_ID,omitempty"` // 商户代码
	QUERY_SN    string `xml:"QUERY_SN,omitempty"`    // 要查询的交易流水（原交易的REQ_SN）
}

// 交易结果查询 响应
type QueryTransRes struct {
	XMLName   xml.Name                `xml:"AIPG,omitempty"`
	INFO      AlliPayResINFO          `xml:"INFO,omitempty"`
	QTRANSRSP QueryTransResQTRANSRSP `xml:"QTRANSRSP,omitempty"`
}

type QueryTransResQTRANSRSP struct {
	BATCHID      string `xml:"BATCHID,omitempty"`      // 交易批次号（原交易的REQ_SN）
	SN           string `xml:"SN,omitempty"`           // 记录序号（也就是原请求交易中的SN的值）
	TRXDIR       string `xml:"TRXDIR,omitempty"`       // 交易方向（0 付 1收）
	SETTLE_DAY   string `xml:"SETTLE_DAY,omitempty"`   // 清算日期（YYYYMMDD）
	FINTIME      string `xml:"FINTIME,omitempty"`      // 完成时间（yyyyMMddHHmmss）
	SUBMITTIME   string `xml:"SUBMITTIME,omitempty"`   // 提交时间（yyyyMMddHHmmss）
	ACCOUNT_NO   string `xml:"ACCOUNT_NO,omitempty"`   // 账号（只返回卡号后4位）
	ACCOUNT_NAME string `xml:"ACCOUNT_NAME,omitempty"` // 账号名
	AMOUNT       string `xml:"AMOUNT,omitempty"`       // 金额（整数，单位分）
	CUST_USERID  string `xml:"CUST_USERID,omitempty"`  // 自定义用户号（原代收付请求报文中的CUST_USERID字段）
	REMARK       string `xml:"REMARK,omitempty"`       // 备注（交易请求中的原样返回）
	SUMMARY      string `xml:"SUMMARY,omitempty"`      // 交易附言
	RET_CODE     string `xml:"RET_CODE,omitempty"`     // 返回代码（0000处理成功）
	ERR_MSG      string `xml:"ERR_MSG,omitempty"`      // 错误文本
}

// 此交易由通联向商户指定的url发起，使用HTTP GET方式提交到商户系统，仅适用单笔交易的情况
// 单笔交易结果通知 请求
type AlliPayNotifyReq struct {
	RETCODE    string // 返回码
	RETMSG     string // 错误信息
	ACCOUNT_NO string // 账号后4位（借记卡或信用卡后4位）
	MOBILE     string // 手机号/小灵通（小灵通带区号，不带括号，减号）
	AMOUNT     int64  // 金额（整数，单位分）
	SETTLE_DAY string // 清算日期（YYYYMMDD）
	FINTIME    string // 完成时间（yyyyMMddHHmmss）
	SUBMITTIME string // 提交时间（yyyyMMddHHmmss）
	BATCHID    string // 交易批次号（原交易的REQ_SN）
	SN         string // 序号（也就是原请求交易中的SN的值）
	USERCODE   string // 用户代码（商户客户ID）
	SIGN       string // 签名（使用SHA1withRSA签名。签名原始内容为返回码|账号|手机号|金额|交易批次号|序号）
}

// 单笔交易结果通知 响应
type AlliPayNotifyRes struct {
	Message string // 商户系统返回一行内容（商户系统返回一行内容 ）
}

// 本接口通过HTTPS GET下载即可
// https://服.务.器.地址/aipg/GetConFile.do?SETTDAY=xxx&REQTIME=yyy&MERID=zzz&SIGN=sss
// 简单对账文件下载 请求
type AlliPayAccountCheckingReq struct {
	Url        string // 请求url
	SETTLE_DAY string // 清算日期（YYYYMMDD）
	REQTIME    string // 请求时间（yyyyMMddHHmmss）
	CONTFEE    string // 是否包含手续费（0.不需手续费，1.包含手续费，空则默认为0）
	MERID      string // 商户号
	SIGN       string // 签名（使用SHA1withRSA签名。签名原始内容为 清算日期|请求时间|商户号）
}

// 文件名规范：PDS+商户号+日期(yyyymmdd)+.txt
// 对账文件分成不同的字段，字段之间用空格分开
// 简单对账文件下载 响应
type AlliPayAccountCheckingRes struct {
	Header          AlliPayAccountCheckingResHeader  // 对账文件的第一行是总摘要信息
	AccountChecking []*AlliPayAccountCheckingResBody // 对账文件的第二行起是对账的明细内容
}

// 对账文件的第一行是总摘要信息
type AlliPayAccountCheckingResHeader struct {
	PDSMK                 string // PDS对账文件标记，固定为PDSMK
	V200                  string // 版本号,本说明的版本固定为V200
	AgentReceivableNumber string // 代收总笔数（对账文件的代收总笔数）
	AgentReceivablePrice  string // 代收总金额（对账文件的代收总金额（分））
	AgentPayableNumber    string // 代付总笔数（对账文件的代付总笔数）
	AgentPayablePrice     string // 代付总金额（对账文件的代付总金额（分））
}

// 对账文件的第二行起是对账的明细内容
type AlliPayAccountCheckingResBody struct {
	BATCHID           string // 交易批次号（原交易的REQ_SN）
	SN                string // 记录序号（也就是原请求交易中的SN的值）
	TRXDIR            string // 交易类型（0 付 1收）
	RET_CODE          string // 交易状态
	AMOUNT            string // 交易金额(整数，单位分)
	ACCOUNT_NO        string // 对方账号（被付或被扣帐户）
	SUBMITTIME        string // 交易时间（yyyyMMddHHmmss）
	SETTLE_DAY        string // 清算日期（YYYYMMDD）
	CUST_USERID       string // 自定义用户号（Xml中的CUST_USERID）
	HandlingFree      string // 手续费（下载对账文件时手续费为空时以0表示, CONTFEE为1时才包含该字段）
	TRX_CODE          string // 交易代码（版本号为05才包含该字段）
	SettlementAccount string // 结算账号（版本号为05才包含该字段）
}
