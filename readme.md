# 通联支付

## install

go get -u -v github.com/ikaiguang/go-allinpay

## go version go1.10.1 darwin/amd64

go run main.go

## 生成私钥

openssl genrsa -out rsa_private_key.pem 1024

## 生成公钥

openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem

## 转换私钥

openssl pkcs12 -nocerts -nodes -in 20060400000044502.p12 -out 20060400000044502.pem

> 转换的时候，密码是：111111

## 转换公钥(或证书)

openssl x509 -inform DER -in allinpay-pds.cer -out allinpay-pds.pem

