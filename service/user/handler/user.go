package handler

import (
	"context"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"log"
	utils "luenci.web.com/service/utils"
	"math/rand"
	"net/http"
	"time"
	user "user/proto/user"
)

type User struct{}

// SendSMS is a single request handler called via client.Call or the generated client code
func (e *User) SendSMS(ctx context.Context, req *user.Request, rsp *user.Response) error {
	log.Println("Received User.Call request")
	if req.IsCheckImg {
		if SendMessage(req.Phone) {
			rsp.Errorn = utils.RECODE_OK
			rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)
			rsp.HttpStatus = http.StatusOK
		} else {
			rsp.Errorn = utils.RECODE_SMSERR
			rsp.Errmsg = utils.RecodeText(utils.RECODE_SMSERR)
			rsp.HttpStatus = http.StatusForbidden
		}
	} else {
		rsp.Errorn = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
		rsp.HttpStatus = http.StatusForbidden
	}

	return nil
}


// SendMessage 发送短信
func SendMessage(phone string) bool {
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", "<accessKeyId>", "<accessSecret>")

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.PhoneNumbers = phone
	request.SignName = "luenci"
	request.TemplateCode = "test"

	// 生成一个随机六位数，做短信验证码
	rand.Seed(time.Now().UnixNano()) // 播种随机数种子
	// 生成随机数
	smsCode := fmt.Sprintf("%06d", rand.Int31n(1000000))

	request.TemplateParam = `{"code":` + smsCode + `"}`

	response, err := client.SendSms(request)
	if err != nil {
		fmt.Print(err.Error())
		return false
	}

	return response.IsSuccess()

}

