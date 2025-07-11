package main

import (
	"encoding/json"
	"fmt"
	"log"
	"microsoft/chtoken"
	"microsoft/config"
	"net/http"
	"time"
)

func getMcToken(w http.ResponseWriter, r *http.Request) {
	// 验证 Basic Auth
	username, password, ok := r.BasicAuth()
	if !ok || username != config.AuthUsername || password != config.AuthPassword {
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		http.Error(w, "认证错误", http.StatusUnauthorized)
		return
	}

	// 只处理GET请求
	if r.Method != http.MethodGet {
		http.Error(w, "该接口只支持GET方法", http.StatusMethodNotAllowed)
		return
	}

	// 设置响应头
	w.Header().Set("Content-Type", "application/json")

	// 返回JSON格式的值
	response := map[string]interface{}{
		"currenToken": chtoken.CurrentToken.Mctoken,
	}

	json.NewEncoder(w).Encode(response)
}

func main() {
	// 定时取新的token存储到通道上
	c1 := time.Tick(time.Minute * 30)
	go func() {
		chtoken.JobStoreToken()
		chtoken.CurrentToken = <-chtoken.Chtoken
		for range c1 {
			chtoken.JobStoreToken()
		}
	}()

	// 定时校验token的过期时间，如果过期了就从通道里拿个新的数据
	c2 := time.Tick(time.Minute * 2)
	go func() {
		for range c2 {
			chtoken.JobGetToken()
		}
	}()

	// 编写HTTP接口
	http.HandleFunc("/api/mctoken", getMcToken)
	fmt.Println("服务器启动成功,接口地址为127.0.0.1:1435..")
	log.Fatal(http.ListenAndServe(":1435", nil))
	chtoken.Wg.Wait()
}
