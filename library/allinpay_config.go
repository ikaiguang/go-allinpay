package library

import (
	"crypto/rsa"
	"io/ioutil"
	"encoding/pem"
	"crypto/x509"
	"errors"
	"fmt"
	"os"
)

// 配置
type AllinPayConfig struct {
	TestWeb        string // 账户系统web前端(测试)
	TestAddress    string // 账户对接测试地址
	ProductWeb     string // 账户系统web前端(生产)
	ProductAddress string // 账户对接生产地址
	WebUsername    string // web端用户名
	WebPassword    string // web端密码
	MerchantCode   string // 商户号
	Username       string // 用户
	Password       string // 密码
	TestAccount    string // 虚拟账号(用于查询虚拟户余额)
	CertFile       string // 通联公钥
	PrivateKey     string // 商户私钥
}

var (
	AllinPayAddress    string          // request url
	AllinPayCfg        *AllinPayConfig // config
	AllinPayPrivateKey *rsa.PrivateKey // private key
	AllinPayPublicKey  *rsa.PublicKey  // public key
)

func init() {
	var err error
	// init config
	AllinPayCfg = InitAllinPayConfig()
	// init private key
	AllinPayPrivateKey, err = InitAllinPayPrivateKey()
	if err != nil {
		fmt.Println("InitAllinPayPrivateKey error : ", err)
	}
	// init public key
	AllinPayPublicKey, err = InitAllinPayPublicKey()
	if err != nil {
		fmt.Println("InitAllinPayPublicKey error : ", err)
	}
}

// 初始化配置
func InitAllinPayConfig() *AllinPayConfig {
	// config
	cfg := &AllinPayConfig{
		TestWeb:        os.Getenv("AllinPayTestWeb"),        // 账户系统web前端(测试)
		TestAddress:    os.Getenv("AllinPayTestAddress"),    // 账户对接测试地址
		ProductWeb:     os.Getenv("AllinPayProductWeb"),     // 账户系统web前端(生产)
		ProductAddress: os.Getenv("AllinPayProductAddress"), // 账户对接生产地址
		WebUsername:    os.Getenv("AllinPayWebUsername"),    // web端用户名
		WebPassword:    os.Getenv("AllinPayWebPassword"),    // web端密码
		MerchantCode:   os.Getenv("AllinPayMerchantCode"),   // 商户号
		Username:       os.Getenv("AllinPayUsername"),       // 用户
		Password:       os.Getenv("AllinPayPassword"),       // 密码
		TestAccount:    os.Getenv("AllinPayTestAccount"),    // 虚拟账号(用于查询虚拟户余额)
		CertFile:       os.Getenv("AllinPayCertFile"),       // 通联公钥
		PrivateKey:     os.Getenv("AllinPayPrivateKey"),     // 商户私钥
	}
	// request url
	if os.Getenv("AllinPayEnv") == "dev" {
		AllinPayAddress = cfg.TestAddress
	} else {
		AllinPayAddress = cfg.ProductAddress
	}
	return cfg
}

// 初始化证书 私钥
func InitAllinPayPrivateKey() (privateKey *rsa.PrivateKey, err error) {
	// pem
	privatePemByte, err := ioutil.ReadFile(AllinPayCfg.PrivateKey)
	if err != nil {
		err = errors.New("ioutil.ReadFile error : " + err.Error())
		return privateKey, err
	}
	// pemBlock
	pemBlock, _ := pem.Decode(privatePemByte)
	if pemBlock == nil {
		err = errors.New("private key file error : Invalid Key")
		return privateKey, err
	}
	// pkcs8
	parsedKey, err := x509.ParsePKCS8PrivateKey(pemBlock.Bytes)
	if err != nil {
		err = errors.New("x509.ParsePKCS8PrivateKey error : " + err.Error())
		return privateKey, err
	}
	// private key
	privateKey, privateKeyOk := parsedKey.(*rsa.PrivateKey)
	if !privateKeyOk {
		err = errors.New("key is not a valid RSA private key")
		return privateKey, err
	}
	return privateKey, err
}

// 初始化证书 公钥
func InitAllinPayPublicKey() (publicKey *rsa.PublicKey, err error) {
	// pem
	publicPemByte, err := ioutil.ReadFile(AllinPayCfg.CertFile)
	if err != nil {
		err = errors.New("ioutil.ReadFile error : " + err.Error())
		return publicKey, err
	}
	// pemBlock
	pemBlock, _ := pem.Decode(publicPemByte)
	if pemBlock == nil {
		err = errors.New("public key file error : Invalid Key")
		return publicKey, err
	}
	// x509.ParsePKIXPublicKey() // not a public pem
	// cert
	cert, err := x509.ParseCertificate(pemBlock.Bytes)
	if err != nil {
		err = errors.New("x509.ParseCertificate error : " + err.Error())
		return publicKey, err
	}
	// public key
	publicKey, publicKeyOk := cert.PublicKey.(*rsa.PublicKey)
	if !publicKeyOk {
		err = errors.New("key is not a valid RSA public key")
		return publicKey, err
	}
	return publicKey, err
}

// 通联请求头
type AllinPayReqHeaderReq struct {
	TRX_CODE   string // 交易代码
	LEVEL      string // 处理级别（0-9  0优先级最低，默认为5）
	REQ_SN     string // 交易流水号（必须全局唯一）
	SIGNED_MSG string // 签名信息
}

// 通联请求头
func InitAllinPayReqHeader(req *AllinPayReqHeaderReq) *AllinPayReqINFO {
	return &AllinPayReqINFO{
		TRX_CODE:    req.TRX_CODE,             // 交易代码
		VERSION:     "04",                     // 版本（04）
		DATA_TYPE:   "2",                      // 数据格式（2：xml格式）
		LEVEL:       req.LEVEL,                // 处理级别（0-9  0优先级最低，默认为5）
		MERCHANT_ID: AllinPayCfg.MerchantCode, // 商户代码
		USER_NAME:   AllinPayCfg.Username,     // 用户名
		USER_PASS:   AllinPayCfg.Password,     // 用户密码
		REQ_SN:      req.REQ_SN,               // 交易流水号（必须全局唯一）
		SIGNED_MSG:  req.SIGNED_MSG,           // 签名信息
	}
}
