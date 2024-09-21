package handler

import (
	"captcha/logic"
	pb "captcha/proto/captcha"
	"context"
	"fmt"
	"image/color"

	"github.com/mojocn/base64Captcha"
)

type Captcha struct{}

func (e *Captcha) GenerateCaptcha(ctx context.Context, req *pb.GenerateCaptchaRequest, res *pb.GenerateCaptchaResponse) error {
	// todo
	fmt.Println("generate--------", req)
	driverString := base64Captcha.DriverString{
		Height:          int(req.GetHeight()),
		Width:           int(req.GetWidth()),
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		Length:          int(req.GetLength()),
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		Fonts: []string{"3Dumb.ttf"},
	}
	var driver base64Captcha.Driver = driverString.ConvertFonts()
	c := base64Captcha.NewCaptcha(driver, logic.CaptchaStore{})
	id, b64s, answer, err := c.Generate()

	res.Id = id
	res.B64S = b64s
	res.Answer = answer

	return err
}

func (e *Captcha) VerifyCaptcha(ctx context.Context, req *pb.VerifyCaptchaRequest, res *pb.VerifyCaptchaResponse) error {
	// todo
	fmt.Println("verify--------", req)
	res.VerifyPass = logic.CaptchaStore{}.Verify(req.Id, req.VerifyString, true)
	return nil
}
