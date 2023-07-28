package common

import (
	"os"
	"math/rand" // 随机数生成器
	"time" // 时间操作
)

var (
	// is debug
	IS_DEBUG_MODE bool
	// socks
	SOCKS_URL  string
	SOCKS_USER string
	SOCKS_PWD  string
	// user token
	USER_TOKEN_ENV_NAME_PREFIX = "Go_Proxy_BingAI_USER_TOKEN"
	USER_TOKEN_LIST            []string
	// 访问权限密钥，可选
	AUTH_KEY             string
	AUTH_KEY_COOKIE_NAME = "BingAI_Auth_Key"
)

func init() {
	initEnv()
	initUserToken()
}

func initEnv() {
	// is debug
	IS_DEBUG_MODE = os.Getenv("Go_Proxy_BingAI_Debug") != ""
	// socks
	SOCKS_URL = os.Getenv("Go_Proxy_BingAI_SOCKS_URL")
	SOCKS_USER = os.Getenv("Go_Proxy_BingAI_SOCKS_USER")
	SOCKS_PWD = os.Getenv("Go_Proxy_BingAI_SOCKS_PWD")
	// auth
	AUTH_KEY = os.Getenv("Go_Proxy_BingAI_AUTH_KEY")
}

func initUserToken() {
	// 定义一个包含所有可能的字符的字符串
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	// 定义一个随机数生成器，使用当前时间作为种子
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 定义一个生成随机字符串的函数，接受长度作为参数
	randomString := func(length int) string {
		// 创建一个字节切片，长度为参数指定的长度
		b := make([]byte, length)
		// for循环遍历字节切片
		for i := range b {
			// 从charset中随机选择一个字符，并赋值给字节切片的对应位置
			b[i] = charset[rng.Intn(len(charset))]
		}
		// 将字节切片转换为字符串并返回
		return string(b)
	}
	// for循环生成1万条随机token
	for i := 0; i < 100000; i++ {
		// 调用randomString函数，生成长度为15的随机字符串
		token := randomString(15)
		// 将随机字符串追加到用户令牌列表中
		USER_TOKEN_LIST = append(USER_TOKEN_LIST, token)
	}
}
