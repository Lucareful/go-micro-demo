package main

import (
	"log"
	"net/http"
	"github.com/dchest/captcha"
)

func main()  {

	// 获取验证码图片
	http.HandleFunc("/r", func(w http.ResponseWriter, r *http.Request) {
		id :=captcha.NewLen(6)
		if id == "" {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "image/png")
		if err := captcha.WriteImage(w, id, 120, 80); err != nil {
			log.Println("show captcha error", err)
		}

	})

	// 启动服务
	log.Fatal(http.ListenAndServe(":8080", nil))


}
