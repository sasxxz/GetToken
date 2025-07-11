package code

import (
	"context"
	"log"
	"microsoft/config"
	"regexp"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

func GetMicrosoftCode() string {
	// 配置浏览器选项
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true), // 禁用无头模式
		chromedp.Flag("hide-scrollbars", false),
		chromedp.Flag("mute-audio", false),
		chromedp.Flag("window-size", "1200,800"),
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// 无头浏览器网页交互
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	// 超时时间
	ctx, cancel = context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()

	var result string
	err := chromedp.Run(ctx,
		// 访问登录页面
		chromedp.Navigate(config.Url),
		// 输入用户名
		chromedp.Sleep(5*time.Second),
		chromedp.SendKeys(`#i0116`, config.Username, chromedp.ByID),
		// 点击下一步
		chromedp.Sleep(5*time.Second),
		chromedp.Click(`#idSIButton9`, chromedp.ByID),
		// 输入密码
		chromedp.Sleep(5*time.Second),
		chromedp.SendKeys(`#i0118`, config.Password, chromedp.ByID),
		// 点击登录
		chromedp.Sleep(5*time.Second),
		chromedp.Click(`#idSIButton9`, chromedp.ByID),
		// 点击不保持登录状态
		chromedp.Sleep(5*time.Second),
		// chromedp.WaitReady(`#idBtn_Back`, chromedp.ByID),
		chromedp.Click(`#idBtn_Back`, chromedp.ByID),
		chromedp.Sleep(5*time.Second),

		// 获得当前界面的url
		chromedp.Location(&result),
	)
	if err != nil {
		// log.Fatal("自动化操作失败:", err)
		return err.Error()
	}
	log.Println("登录成功！！")
	// 通过正则取到code
	re := regexp.MustCompile(`(code=[^&]+)`)
	index := re.FindString(result)
	Code := strings.Split(index, "=")
	if len(Code) < 2 {
		return "code get error" // 判断是否能正常拿到code
	}
	return Code[1]
}
