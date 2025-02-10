package utils

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"server/global"
	"strings"

	"github.com/jordan-wright/email"
)

// Email 发送电子邮件
func Email(To, subject string, body string) error {
	to := strings.Split(To, ",") // 将收件人邮箱地址按逗号拆分成多个地址
	return send(to, subject, body)
}

// send 执行邮件发送操作
func send(to []string, subject string, body string) error {
	emailCfg := global.Config.Email // 获取全局配置中的邮件设置

	from := emailCfg.From
	nickname := emailCfg.Nickname
	secret := emailCfg.Secret
	host := emailCfg.Host
	port := emailCfg.Port
	isSSL := emailCfg.IsSSL

	// 使用 PlainAuth 创建认证信息
	auth := smtp.PlainAuth("", from, secret, host)

	// 创建新的电子邮件对象
	e := email.NewEmail()
	if nickname != "" {
		// 如果设置了昵称，则格式化发件人地址为 "昵称 <邮箱>"
		e.From = fmt.Sprintf("%s <%s>", nickname, from)
	} else {
		// 否则直接使用发件人邮箱
		e.From = from
	}

	// 设置收件人、主题和邮件内容
	e.To = to
	e.Subject = subject
	e.HTML = []byte(body)

	// 定义错误变量
	var err error
	// 构建邮件服务器的地址，格式为 host:port
	hostAddr := fmt.Sprintf("%s:%d", host, port)

	// 根据配置的是否使用 SSL 来选择邮件发送方法
	if isSSL {
		// 使用带 TLS 的邮件发送
		err = e.SendWithTLS(hostAddr, auth, &tls.Config{ServerName: host})
	} else {
		// 使用普通的邮件发送
		err = e.Send(hostAddr, auth)
	}

	return err
}
