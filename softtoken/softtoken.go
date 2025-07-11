package softtoken

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"microsoft/config"
	"net/http"
	"net/url"
	"os"
	"time"
)

func GetMicrosoftToken(softcode string) (string, int, time.Time) {

	// 准备表单数据
	formData := url.Values{
		"client_id":     {"a1b9057a-9fde-4119-b25c-ffd2e59e4f9a"},
		"scope":         {"Files.ReadWrite"},
		"code":          {softcode},
		"redirect_uri":  {"https://login.microsoftonline.com/common/oauth2/nativeclient"},
		"grant_type":    {"authorization_code"},
		"client_secret": {"M6q8Q~Oojx4Z9YhZGj3-txFswUYXRkp5pFZGib0k"}, // 这个值可能会过期
	}

	// 创建请求体
	reqBody := bytes.NewBufferString(formData.Encode())

	// 创建http请求
	req, err := http.NewRequest("POST", config.TargetURL, reqBody)
	if err != nil {
		log.Fatalf("创建请求失败:%v", err)
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 创建HTTP客户端并发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("读取响应失败: %v", err)
	}

	// 输出结果
	fmt.Printf("状态码: %d\n", resp.StatusCode)
	// fmt.Println("响应标头:")
	// for key, values := range resp.Header {
	// 	fmt.Printf("%s %s\n", key, strings.Join(values, ";"))
	// }
	// fmt.Printf("\n响应内容:\n%s\n", body)
	Structbody := &config.Microbody{}
	err = json.Unmarshal([]byte(body), Structbody)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		os.Exit(0)
	}
	return Structbody.Access_token, Structbody.Expires_in, time.Now()
}
