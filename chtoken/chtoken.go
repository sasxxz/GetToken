package chtoken

// 并发处理token取值的方法

import (
	"fmt"
	"microsoft/code"
	"microsoft/softtoken"
	"sync"
	"time"
)

var (
	Wg           sync.WaitGroup
	CurrentToken Tokenftime
	McToken      Tokenftime
	Chtoken      = make(chan Tokenftime, 10)
)

// 定义token+过期时间的类型
type Tokenftime struct {
	Mctoken string
	Mctime  time.Time
}

func JobStoreToken() {
	Wg.Add(1)
	var indexcode string
	var sum int = 0
	for {
		sum++
		indexcode = (code.GetMicrosoftCode())
		if indexcode == "code get error" { // 判断code是否正常返回，如果没有一直重新尝试，重试20次依旧不行就直接返回
			fmt.Println(indexcode)
		} else if sum == 20 {
			fmt.Println("code获取异常,请检查获取code网页的网络环境..")
			return
		} else {
			break
		}
	}
	lastToken, mctime, mcnow := softtoken.GetMicrosoftToken(indexcode) // 获得token和对应的过期时间
	McToken = Tokenftime{
		Mctoken: lastToken,
		Mctime:  mcnow.Add(time.Duration(mctime-300) * time.Second),
	}

	Chtoken <- McToken // 存放token with timeout
	Wg.Done()
}

// 定时校验token的过期时间
func JobGetToken() {
	Wg.Add(1)
	if time.Now().After(McToken.Mctime) {
		CurrentToken = <-Chtoken
	} else {
		return
	}
	Wg.Done()
}
