package encryption

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"encoding/json"

	commonConst "github.com/deepfence/ThreatMapper/deepfence_server/constants/common"
	"github.com/deepfence/ThreatMapper/deepfence_server/model"
	pgDb "github.com/deepfence/golang_deepfence_sdk/utils/postgresql/postgresql-db"
)

type AES struct {
	IV  string `json:"aes_iv"`
	Key string `json:"aes_key"`
}

const blockSize = aes.BlockSize

func (a *AES) Encrypt(plaintext string) (string, error) {
	bKey, err := hex.DecodeString(a.Key)
	if err != nil {
		return "", err
	}
	bIV, err := hex.DecodeString(a.IV)
	if err != nil {
		return "", err
	}
	bPlaintext := PKCS5Padding([]byte(plaintext), blockSize, len(plaintext))
	block, err := aes.NewCipher(bKey)
	if err != nil {
		return "", err
	}
	ciphertext := make([]byte, len(bPlaintext))
	mode := cipher.NewCBCEncrypter(block, bIV)
	mode.CryptBlocks(ciphertext, bPlaintext)
	return hex.EncodeToString(ciphertext), nil
}

func (a *AES) Decrypt(cipherText string) (string, error) {
	bKey, err := hex.DecodeString(a.Key)
	if err != nil {
		return "", err
	}
	bIV, err := hex.DecodeString(a.IV)
	if err != nil {
		return "", err
	}
	cipherTextDecoded, err := hex.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(bKey)
	if err != nil {
		return "", err
	}

	mode := cipher.NewCBCDecrypter(block, bIV)
	mode.CryptBlocks([]byte(cipherTextDecoded), []byte(cipherTextDecoded))
	return string(PKCS5UnPadding(cipherTextDecoded)), nil
}

func PKCS5Padding(ciphertext []byte, blockSize int, after int) []byte {
	padding := (blockSize - len(ciphertext)%blockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	if unpadding > length {
		return src
	}
	return src[:(length - unpadding)]
}

func GetAESValueForEncryption(ctx context.Context, pgClient *pgDb.Queries) (json.RawMessage, error) {
	s := model.Setting{}
	aes, err := s.GetSettingByKey(ctx, pgClient, commonConst.AES_SECRET)
	if err != nil {
		return nil, err
	}
	var sValue model.SettingValue
	err = json.Unmarshal(aes.Value, &sValue)
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(sValue.Value)
	if err != nil {
		return nil, err
	}

	return json.RawMessage(b), nil
}
