package logic

import "time"

// 参考官网：https://github.com/mojocn/base64Captcha
// 将生成的验证码保存在redis里面，需要自己实现官方的接口

// Implement Store interface or use build-in memory store
/*
type Store interface {
	// Set sets the digits for the captcha id.
	Set(id string, value string)

	// Get returns stored digits for the captcha id. Clear indicates
	// whether the captcha must be deleted from the store.
	Get(id string, clear bool) string

    //Verify captcha's answer directly
	Verify(id, answer string, clear bool) bool
}
*/

type CaptchaStore struct{}

var (
	Prefix_string = "captcha:"
)

func (s CaptchaStore) Set(id, value string) error {
	return Rset(Prefix_string+id, value, time.Minute*2) // 验证码2分钟过期
}

func (s CaptchaStore) Get(id string, clear bool) string {
	// todo
	v, err := RGet(Prefix_string + id)
	if err != nil {
		return ""
	}
	if clear {
		RDel(Prefix_string + id)
	}
	return v
}

func (s CaptchaStore) Verify(id, answer string, clear bool) bool {
	// todo
	return s.Get(id, clear) == answer
}
