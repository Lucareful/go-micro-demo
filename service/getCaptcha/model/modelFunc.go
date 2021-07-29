package model

import (
	"github.com/gomodule/redigo/redis"
	"log"
)

func SaveImgCode(code, uuid string) error {
	// 链接数据库
	conn, err := redis.Dial("tcp", "121.199.72.50:6379")
	if err != nil {
		log.Println("redis链接错误:", err)
		return err
	}

	// 延时关闭 redis
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println("redis 链接关闭错误:", err)
		}
	}(conn)

	// 链接认证
	if _, err := conn.Do("AUTH", "luenci"); err != nil {
		log.Println(err)
		return err
	}

	// 验证码写入数据库 -- 设置有效时间
	if _, err := conn.Do("SET", uuid, 60*5, code); err != nil {
		return err
	}

	return nil

}

