package appTypes

import "encoding/json"

// Register 用户注册来源
type Register int

const (
	Email Register = iota // 邮箱验证码注册
	QQ                    // QQ登录注册
)

// MarshalJSON 实现了 json.Marshaler 接口
func (r Register) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

// UnmarshalJSON 实现了 json.Unmarshaler 接口
func (r *Register) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	*r = ToRegister(str)
	return nil
}

// String 方法返回 Register 的字符串表示
func (r Register) String() string {
	var str string
	switch r {
	case Email:
		str = "邮箱"
	case QQ:
		str = "QQ"
	default:
		str = "未知"
	}
	return str
}

// ToRegister 函数将字符串转换为 Register
func ToRegister(str string) Register {
	switch str {
	case "邮箱":
		return Email
	case "QQ":
		return QQ
	default:
		return -1
	}
}
