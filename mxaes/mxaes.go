// Copyright (c) 2020.
// ALL Rights reserved.
// @Description mxaes.go
// @Author moxiao
// @Date 2020/11/21 18:19

package mxaes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

var (
	ErrInputNotFullBlocks = errors.New("crypto/cipher: input not full blocks")
)

//加密
func Encrypt(srcData, key, iv []byte) (encryptData []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}
	//填充原文
	blockSize := block.BlockSize()
	srcData = PKCS5Padding(srcData, blockSize)

	// 验证输入参数
	// 必须为aes.Blocksize的倍数
	if len(srcData)%aes.BlockSize != 0 {
		return nil, ErrInputNotFullBlocks
	}

	//初始向量IV
	cipherText := make([]byte, blockSize+len(srcData))
	if iv == nil {
		iv = cipherText[:blockSize]
		if _, err := io.ReadFull(rand.Reader, iv); err != nil {
			panic(err)
		}
	} else {
		iv = iv[:blockSize]
	}

	blockMode := cipher.NewCBCEncrypter(block, iv)
	encryptData = make([]byte, len(srcData))
	blockMode.CryptBlocks(encryptData, srcData)
	return
}

//解密
func Decrypt(decryptData, key, iv []byte) (srcData []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}
	blockSize := block.BlockSize()

	// 长度不能小于aes.Blocksize
	if len(decryptData) < blockSize {
		return nil, errors.New("crypto/cipher: ciphertext too short")
	}

	if iv == nil {
		iv = decryptData[:blockSize]
		decryptData = decryptData[blockSize:]
	} else {
		iv = iv[:blockSize]
		decryptData = decryptData[:]
	}

	// 验证输入参数,必须为aes.Blocksize的倍数
	if len(decryptData)%blockSize != 0 {
		return nil, errors.New("crypto/cipher: ciphertext is not a multiple of the block size")
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(decryptData))
	blockMode.CryptBlocks(origData, decryptData)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

func Decrypt2(decryptData, key, iv []byte) (srcData []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	blockMode.CryptBlocks(decryptData, decryptData)
	decryptData, err = PKCS7UnPadding(decryptData, block.BlockSize())
	if err != nil {
		return nil, err
	}
	return decryptData, nil
}

//填充
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//填充
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

//填充
func PKCS7UnPadding(data []byte, blockSize int) ([]byte, error) {
	if blockSize <= 0 {
		return nil, errors.New("invalid block size")
	}
	if len(data)%blockSize != 0 || len(data) == 0 {
		return nil, errors.New("invalid PKCS7 data")
	}
	c := data[len(data)-1]
	n := int(c)
	if n == 0 || n > len(data) {
		return nil, errors.New("invalid padding on input")
	}
	for i := 0; i < n; i++ {
		if data[len(data)-n+i] != c {
			return nil, errors.New("invalid padding on input")
		}
	}
	return data[:len(data)-n], nil
}
