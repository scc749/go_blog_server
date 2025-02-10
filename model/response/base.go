package response

type Captcha struct {
	CaptchaID string `json:"captcha_id"`
	PicPath   string `json:"pic_path"`
}
