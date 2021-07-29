package main

import (
	"fmt"
	"luenci.web.com/web/utils"
	"time"
)

func main() {
	// 链接redis数据库
	utils.PoolInitRedis("121.199.72.50:6379", "luenci")
	for i := 0; i < 5; i++ {
		conn, err := utils.GetRedisConn()
		fmt.Println(conn)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(conn)
	}

	// 回复助手类函数，---- 确定程序具体的数据类型
	//reply, _ := redis.String(conn.Do("set", "name", "Lynn"))
	//fmt.Println(reply)

	//connChan := make(chan redis.Conn, 5)
	//for i := 0; i < 5; i++ {
	//	connChan <- coon.Get()
	//}
	//for i := 0; i < 5; i++ {
	//	conn := <-connChan
	//	fmt.Printf("%#v\n", conn)
	//	conn.Close()
	//}

	time.Sleep(time.Second * 5)
	////下次是怎么取出来的？？
	//b1 := pool.Get()
	//b2 := pool.Get()
	//b3 := pool.Get()
	//if b4 := pool.Get(); b4 == nil{
	//	fmt.Printf("redis 连接超出限制 %#v\n", b4)
	//}
	//
	//fmt.Printf(" %#v\n, %#v\n, %#v\n", b1, b2, b3)
	//time.Sleep(time.Second * 5)
	//b1.Close()
	//b2.Close()
	//b3.Close()
	//
	////redis目前一共有多少个连接？？
	//for {
	//	fmt.Println("主程序运行中....")
	//	time.Sleep(time.Second * 1)
	//}

	//
	//fmt.Println(r, e)
}
