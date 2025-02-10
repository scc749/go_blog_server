package utils

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// GenerateVerificationCode 生成一个指定长度的随机验证码
func GenerateVerificationCode(length int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%0*d", length, r.Intn(int(math.Pow10(length))))
}
