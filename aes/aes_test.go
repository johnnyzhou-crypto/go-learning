package aes

import (
	"encoding/base64"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAes(t *testing.T) {
	var aeskey = []byte("Qjn*ZJdX*dZLd.u7*HL8G2H9")
	pass := []byte("rails|1634025496933|131211")
	xpass, err := AesEncrypt(pass, aeskey)
	if err != nil {
		fmt.Println(err)
		return
	}
	pass64 := base64.StdEncoding.EncodeToString(xpass)
	assert.Equal(t, "MbW9ncEtj6y7mNATb6cBdMovgvTbReriC0oBhckHl8s=", pass64)
	bytesPass, _ := base64.StdEncoding.DecodeString(pass64)
	tpass, _ := AesDecrypt(bytesPass, aeskey)
	assert.Equal(t, pass, tpass)
}
