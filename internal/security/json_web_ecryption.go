package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"

	"github.com/go-jose/go-jose/v3"
)

// 定义加密密钥 (在生产环境应从环境变量加载)
var sharedEncryptionKey = []byte("this-is-a-32-byte-long-secret-key!!") // AES-256 需要 32 字节

// EncryptPayload 将对象加密为 JWE 字符串
func EncryptPayload(payload interface{}) (string, error) {
	// 1. 创建加密器 (使用 AES-GCM 算法)
	encrypter, err := jose.NewEncrypter(
		jose.A256GCM,
		jose.Recipient{Algorithm: jose.DIRECT, Key: sharedEncryptionKey},
		nil,
	)
	if err != nil {
		return "", err
	}

	// 2. 序列化数据
	data, _ := json.Marshal(payload)

	// 3. 执行加密
	object, err := encrypter.Encrypt(data)
	if err != nil {
		return "", err
	}

	return object.FullSerialize(), nil
}

// DecryptPayload 将 JWE 字符串解密到对象
func DecryptPayload(jweString string, dest interface{}) error {
	object, err := jose.ParseEncrypted(jweString)
	if err != nil {
		return err
	}

	// 使用密钥解密
	decrypted, err := object.Decrypt(sharedEncryptionKey)
	if err != nil {
		return err
	}

	return json.Unmarshal(decrypted, dest)
}

// EncryptWithGCM 使用AES-GCM加密数据
func EncryptWithGCM(plaintext []byte) (string, error) {
	block, err := aes.NewCipher(sharedEncryptionKey)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// DecryptWithGCM 使用AES-GCM解密数据
func DecryptWithGCM(ciphertext string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(sharedEncryptionKey)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertextData := data[:nonceSize], data[nonceSize:]
	return gcm.Open(nil, nonce, ciphertextData, nil)
}