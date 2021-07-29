package handler

import (
	"context"
	"encoding/json"
	"github.com/afocus/captcha"
	"image/color"
	"log"
	"luenci.web.com/service/getCaptcha/model"
	getCaptcha "luenci.web.com/service/getCaptcha/proto/getCaptcha"
)

type GetCaptcha struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *GetCaptcha) Call(ctx context.Context, req *getCaptcha.Request, rsp *getCaptcha.Response) error {
	log.Println("Received GetCaptcha.Call request")
	// 生成图片验证码

	// 初始化对象
	c := captcha.New()

	// 设置字体
	err := c.SetFont("./conf/叶立群几何黑体.ttf")
	if err != nil {
		log.Println("字体设置错误:", err)
		return err
	}

	// 设置验证码大小
	c.SetSize(128, 64)

	// 设置干扰强度
	c.SetDisturbance(captcha.MEDIUM)

	// 设置前景色 可以多个 随机替换文字颜色 默认黑色
	c.SetFrontColor(color.RGBA{255, 255, 255, 255})

	// 设置背景色 可以多个 随机替换背景色 默认白色
	c.SetBkgColor(color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})

	// 生成字体图片
	img, str := c.Create(4, captcha.NUM)

	// 存储图片验证码数组到 redis 中
	if err := model.SaveImgCode(str, req.Uuid); err != nil {
		return err
	}

	// 将生成的图片 序列化
	imgBuf, _ := json.Marshal(img)

	rsp.Img = imgBuf

	return nil
}
