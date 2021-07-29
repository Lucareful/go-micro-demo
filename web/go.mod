module luenci.web.com/web

go 1.15

replace luenci.web.com/web/proto/getCaptcha => ./proto/getCaptcha // 本地包相对路径导入替换

require (
	github.com/afocus/captcha v0.0.0-20191010092841-4bd1f21c8868
	github.com/alibabacloud-go/darabonba-env v1.0.0
	github.com/alibabacloud-go/darabonba-openapi v0.0.9
	github.com/alibabacloud-go/darabonba-string v1.0.0
	github.com/alibabacloud-go/darabonba-time v1.0.0
	github.com/alibabacloud-go/dysmsapi-20170525/v2 v2.0.1
	github.com/alibabacloud-go/tea v1.1.11
	github.com/alibabacloud-go/tea-console v1.0.0
	github.com/alibabacloud-go/tea-utils v1.3.8
	github.com/dchest/captcha v0.0.0-20200903113550-03f5f0333e1f
	github.com/gin-gonic/gin v1.7.2
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/golang/protobuf v1.5.0
	github.com/gomodule/redigo v1.8.5
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/consul/v2 v2.9.1
	google.golang.org/protobuf v1.27.1
	gorm.io/driver/mysql v1.1.1
	gorm.io/gorm v1.21.11

)
