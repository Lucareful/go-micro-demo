package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/afocus/captcha"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/registry/consul/v2"
	"image/png"
	"log"
	"luenci.web.com/web/model"
	getCaptchaMicro "luenci.web.com/web/proto/getCaptcha"
	userMicro "luenci.web.com/web/proto/user"
	"luenci.web.com/web/utils"
	"net/http"
)

// GetSession api 获取 session 信息
func GetSession(ctx *gin.Context) {
	// 初始化错误返回的 map
	resp := make(map[string]string)

	resp["error"] = utils.RECODE_SESSIONERR
	resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)

	ctx.JSON(http.StatusOK, resp)
}

// GetImageCd api 获取图片信息
func GetImageCd(ctx *gin.Context) {
	uuid := ctx.Param("uuid")

	// 指定 consul 的服务发现
	consulReg := consul.NewRegistry()
	consulSer := micro.NewService(
		micro.Registry(consulReg), // 注册服务发现
	)

	// 初始化客户端
	microClient := getCaptchaMicro.NewGetCaptchaService("go.micro.service.getCaptcha", consulSer.Client())

	// 使用客户端调用远程函数
	resp, err := microClient.Call(context.TODO(), &getCaptchaMicro.Request{Uuid: uuid})
	if err != nil {
		log.Println("未找到远程服务 getCaptcha ...", err)
		return
	}

	var img captcha.Image
	if err := json.Unmarshal(resp.Img, &img); err != nil {
		log.Println("反序列化出错", err)
		return
	}

	if err := png.Encode(ctx.Writer, img); err != nil {
		log.Println("解析图片出错", err)
		return
	}

}

// GetSmsCd api 获取短信验证码
func GetSmsCd(ctx *gin.Context) {
	// 声明返回信息的结构和状态

	phone := ctx.Param("phone")

	// 拆分 GET 请求中的URL
	// 资源路径 ?k=v&k=v&k=v
	imgCode := ctx.Query("text")
	uuid := ctx.Query("id")

	//fmt.Println(phone, imgCode, uuid)

	// 检验验证码
	isCheckimg := model.CheckImgCode(uuid, imgCode)

	// 指定 consul 的服务发现
	consulReg := consul.NewRegistry()
	consulSer := micro.NewService(
		micro.Registry(consulReg), // 注册服务发现
	)

	// 初始化客户端
	microClient := userMicro.NewUserService("go.micro.service.user", consulSer.Client())

	// 使用客户端调用远程函数
	resp, err := microClient.SendSMS(context.TODO(), &userMicro.Request{Phone: phone, ImgCode: imgCode, Uuid: uuid, IsCheckImg: isCheckimg})
	if err != nil {
		log.Println("未找到远程服务 user ...", err)
		return
	}
	//delete(resp,"HttpStatus")

	ctx.JSON(int(resp.HttpStatus), resp)
}

func CreateUser(ctx *gin.Context) {
	// 无法获取数据
	// phone := ctx.PostForm("mobile")
	// password := ctx.PostForm("password")
	// sms_code := ctx.PostForm("sms_code")
	// fmt.Println(phone, password, sms_code)

	// 获取数据
	var reqData struct {
		Mobile   string `json:"mobile"`
		PassWord string `json:"password"`
		SmsCode  string `json:"sms_code"`
	}

	if err := ctx.Bind(&reqData); err != nil {
		fmt.Println("数据 bind 出错:", err)
		return
	}
	fmt.Println("获取的数据为：", reqData)

}
