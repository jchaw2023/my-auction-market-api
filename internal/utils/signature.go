package utils

import (
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// VerifySignature 验证以太坊签名
// message: 原始消息
// signature: 签名（十六进制字符串，0x开头）
// expectedAddress: 期望的地址
func VerifySignature(message, signature, expectedAddress string) (bool, error) {
	// 移除签名前缀
	sig := strings.TrimPrefix(signature, "0x")
	if len(sig) != 130 {
		return false, errors.New("invalid signature length")
	}

	// 解析签名
	sigBytes, err := hex.DecodeString(sig)
	if err != nil {
		return false, fmt.Errorf("failed to decode signature: %w", err)
	}

	// 验证签名长度
	if len(sigBytes) != 65 {
		return false, errors.New("invalid signature length")
	}

	// 恢复签名中的 v 值（最后一个字节）
	v := sigBytes[64]
	if v < 27 {
		v += 27
	}

	// 构建消息哈希（EIP-191 标准）
	// personal_sign 格式: "\x19Ethereum Signed Message:\n" + len(message) + message
	messageHash := accounts.TextHash([]byte(message))

	// 恢复公钥
	sigBytes[64] = v - 27
	pubKey, err := crypto.SigToPub(messageHash, sigBytes)
	if err != nil {
		return false, fmt.Errorf("failed to recover public key: %w", err)
	}

	// 从公钥获取地址
	recoveredAddress := crypto.PubkeyToAddress(*pubKey)

	// 比较地址（不区分大小写）
	expectedAddr := common.HexToAddress(expectedAddress)
	if !strings.EqualFold(recoveredAddress.Hex(), expectedAddr.Hex()) {
		return false, nil
	}

	return true, nil
}

// VerifySignatureWithPublicKey 使用公钥验证签名（备用方法）
func VerifySignatureWithPublicKey(message, signature string, publicKey *ecdsa.PublicKey) (bool, error) {
	// 移除签名前缀
	sig := strings.TrimPrefix(signature, "0x")
	if len(sig) != 130 {
		return false, errors.New("invalid signature length")
	}

	// 解析签名
	sigBytes, err := hex.DecodeString(sig)
	if err != nil {
		return false, fmt.Errorf("failed to decode signature: %w", err)
	}

	// 构建消息哈希
	messageHash := accounts.TextHash([]byte(message))

	// 恢复签名中的 v 值
	v := sigBytes[64]
	if v < 27 {
		v += 27
	}

	// 恢复公钥
	sigBytes[64] = v - 27
	recoveredPubKey, err := crypto.SigToPub(messageHash, sigBytes)
	if err != nil {
		return false, fmt.Errorf("failed to recover public key: %w", err)
	}

	// 比较公钥
	if recoveredPubKey.X.Cmp(publicKey.X) != 0 || recoveredPubKey.Y.Cmp(publicKey.Y) != 0 {
		return false, nil
	}

	return true, nil
}

