package main

import (
	"fmt"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	_ "github.com/taosdata/driver-go/v3/taosRestful"
	_ "github.com/taosdata/driver-go/v3/taosWS"
	_ "go.uber.org/automaxprocs"
	"io/ioutil"
	"net/http"
	"sviwo/internal/cmd"
	_ "sviwo/internal/logic"
	_ "sviwo/internal/packed"
)

func main() {
	cmd.Main.Run(gctx.New())
	//sim()
}

func sim() {
	url := "https://api.linksfield.net/cube/v4/sims/8984012211500000019/remaining_data"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Accept-Language", "zh-CN")
	req.Header.Add("Authorization", "")
	req.Header.Add("timestamp", gtime.Now().TimestampMilliStr())
	req.Header.Add("nonce", gconv.String(grand.Intn(100)))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
