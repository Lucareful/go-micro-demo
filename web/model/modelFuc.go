package model

import (
	"github.com/gomodule/redigo/redis"
	"log"
	utils "luenci.web.com/web/utils"
)

// CheckImgCode 校验图片验证码
func CheckImgCode(uuid, imgCode string) bool {
	// 链接 redis 数据库
	conn, err := utils.GetRedisConn()
	if err != nil {
		log.Println(err)
		return false
	}

	defer conn.Close()

	code, err := redis.String(conn.Do("get", uuid))
	if err != nil {
		log.Fatalln("查询错误:", err)
		return false
	}

	return code == imgCode
}
