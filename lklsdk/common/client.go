package common

import (
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	ran "math/rand"
	"time"

	"github.com/black1552/base-common/utils"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gconv"
)

// Client 拉卡拉SDK客户端
type Client[T any] struct {
	config   *Config
	response T
	ctx      context.Context
}

// NewClient 创建拉卡拉SDK客户端
func NewClient[T any](ctx context.Context, cfgJson string) *Client[T] {
	var config *Config
	err := gconv.Struct(cfgJson, &config)
	if err != nil {
		return nil
	}
	return &Client[T]{
		config: config,
		ctx:    ctx,
	}
}
func (c *Client[T]) generateNonceStr() string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyz0123456789"
	nonce := make([]byte, 12)
	for i := range nonce {
		nonce[i] = letterBytes[ran.Intn(len(letterBytes))]
	}
	return string(nonce)
}

// generateSign 生成签名
func (c *Client[T]) generateSign(request []byte) (string, error) {
	// 生成随机字符串
	nonceStr := c.generateNonceStr()
	// 获取当前时间戳（秒）
	timestamp := fmt.Sprintf("%d", time.Now().Unix())

	// 构建待签名报文
	signData := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n", c.config.AppId, c.config.SerialNo, timestamp, nonceStr, string(request))

	// 计算签名
	hashed := sha256.Sum256([]byte(signData))

	privateKey, err := c.loadPrivateKey()
	if err != nil {
		return "", err
	}
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		return "", err
	}
	signatureBase64 := base64.StdEncoding.EncodeToString(signature)

	// 构建Authorization头
	authorization := fmt.Sprintf("LKLAPI-SHA256withRSA appid=\"%s\",serial_no=\"%s\",timestamp=\"%s\",nonce_str=\"%s\",signature=\"%s\"",
		c.config.AppId, c.config.SerialNo, timestamp, nonceStr, signatureBase64)
	return authorization, nil
}

func (c *Client[T]) loadPrivateKey() (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(gfile.GetBytes(c.config.PrivateKey))
	if block == nil {
		return nil, fmt.Errorf("无法解码私钥PEM数据")
	}
	// 解析PKCS#8格式私钥
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return privateKey.(*rsa.PrivateKey), nil
}

// DoRequest 发送HTTP请求
func (c *Client[T]) DoRequest(url string, reqData interface{}) (*T, error) {
	jsonStr, err := gjson.EncodeString(reqData)
	if err != nil {
		return nil, fmt.Errorf("请求参数转JSON字符串失败: %v", err)
	}
	reqJson, err := CleanJSON(jsonStr)
	if err != nil {
		return nil, fmt.Errorf("清理JSON空值失败: %v", err)
	}
	// 序列化为JSON
	jsonData, err := json.Marshal(reqJson)
	if err != nil {
		return nil, fmt.Errorf("序列化请求数据失败: %v", err)
	}

	auth, err := c.generateSign(jsonData)
	if err != nil {
		return nil, fmt.Errorf("生成签名失败: %v", err)
	}
	header := map[string]string{
		"Authorization": auth,
		"Content-Type":  "application/json",
		"Accept":        "application/json",
	}
	// 设置其他必要的请求头

	return utils.NewClient[T](jsonData, url, header).Post(c.ctx)
}
