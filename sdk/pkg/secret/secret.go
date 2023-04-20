package secret

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/big"
	"math/rand"
	"net/url"
)

const (
	modulus = "00e0b509f6259df8642dbc35662901477df22677ec152b5ff68ace615bb7" +
		"b725152b3ab17a876aea8a5aa76d2e417629ec4ee341f56135fccf695280" +
		"104e0312ecbda92557c93870114af6c9d05c4f7f0c3685b7a46bee255932" +
		"575cce10b424d813cfe4875d3e82047b97ddef52741d546b8e289dc6935b" +
		"3ece0462db0a22b8e7"
	publicKey = "010001"
	nonce     = "0CoJUm6Qyw8W8jud"
)

// EncryptJSONBytes encrypts provided request.
func EncryptRequest(data []byte) (string, error) {
	secret := randHexStr(16)

	params, err := aesEncrypt(data, []byte(nonce))
	if err != nil {
		return "", err
	}
	params, err = aesEncrypt(params, secret)
	if err != nil {
		return "", err
	}

	encSecKey, err := rsaEncrypt(secret, publicKey, modulus)
	if err != nil {
		return "", err
	}

	body := url.Values{
		"params":    []string{string(params)},
		"encSecKey": []string{encSecKey},
	}
	return body.Encode(), nil
}

func reverseBytes(input []byte) []byte {
	length := len(input)
	reversed := make([]byte, length)
	for i, b := range input {
		reversed[length-1-i] = b
	}
	return reversed
}

func rsaEncrypt(text []byte, pubkey, modulus string) (string, error) {
	rsaText := reverseBytes(text)

	rsaPubkey := new(big.Int)
	rsaPubkey.SetString(pubkey, 16)

	rsaModulus := new(big.Int)
	rsaModulus.SetString(modulus, 16)

	rsaTextInt := new(big.Int)
	rsaTextInt.SetBytes(rsaText)

	rsaResult := new(big.Int).Exp(rsaTextInt, rsaPubkey, rsaModulus)
	return fmt.Sprintf("%0256x", rsaResult), nil
}

func aesEncrypt(text, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	pad := 16 - len(text)%16
	text = append(text, bytes.Repeat([]byte{byte(pad)}, pad)...)

	blockMode := cipher.NewCBCEncrypter(block, []byte("0102030405060708"))
	ciphertext := make([]byte, len(text))
	blockMode.CryptBlocks(ciphertext, text)
	encoded := make([]byte, base64.StdEncoding.EncodedLen(len(ciphertext)))
	base64.StdEncoding.Encode(encoded, ciphertext)
	return encoded, nil
}

func randHexStr(size int) []byte {
	key := make([]byte, size/2)
	rand.Read(key)
	return []byte(hex.EncodeToString(key))
}

// HexStr32 returns random hex string of length 32.
func HexStr32() string {
	return string(randHexStr(32))
}

// HexStr4 returns random hex string of length 4.
func HexStr8() string {
	return string(randHexStr(8))
}
