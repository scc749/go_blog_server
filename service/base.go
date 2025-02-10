package service

import (
	"server/global"
	"server/utils"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type BaseService struct {
}

func (baseService *BaseService) SendEmailVerificationCode(c *gin.Context, to string) error {
	verificationCode := utils.GenerateVerificationCode(6)
	expireTime := time.Now().Add(5 * time.Minute).Unix()

	// 将验证码、验证邮箱、过期时间存入会话中
	session := sessions.Default(c)
	session.Set("verification_code", verificationCode)
	session.Set("email", to)
	session.Set("expire_time", expireTime)
	_ = session.Save()

	subject := "您的邮箱验证码"
	body := `亲爱的用户[` + to + `]，<br/>
<br/>
感谢您注册` + global.Config.Website.Name + `的个人博客！为了确保您的邮箱安全，请使用以下验证码进行验证：<br/>
<br/>
验证码：[<font color="blue"><u>` + verificationCode + `</u></font>]<br/>
该验证码在 5 分钟内有效，请尽快使用。<br/>
<br/>
如果您没有请求此验证码，请忽略此邮件。
<br/>
如有任何疑问，请联系我们的支持团队：<br/>
邮箱：` + global.Config.Email.From + `<br/>
<br/>
祝好，<br/>` +
		global.Config.Website.Title + `<br/>
<br/>`

	_ = utils.Email(to, subject, body)

	return nil
}
