package library

import (
	"net/http"
	"crypto/tls"
	"bytes"
	"errors"
	"io/ioutil"
)

func PostAllinPayXmlByte(url string, xml []byte) (r []byte, err error) {
	// tls
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	// post
	resp, err := client.Post(url, "application/xml", bytes.NewReader(xml))
	if err != nil {
		err = errors.New("http post error : " + err.Error())
		return r, err
	}
	defer resp.Body.Close()
	// status ok
	if resp.StatusCode != http.StatusOK {
		err = errors.New("http post fail, error code : " + resp.Status)
		return r, err
	}
	// return
	if r, err = ioutil.ReadAll(resp.Body); err != nil {
		err = errors.New("ioutil.ReadAll error : " + resp.Status)
		return r, err
	}
	return r, err
}
