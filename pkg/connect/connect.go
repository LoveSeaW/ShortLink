package connect

import (
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"time"
)

// Client 全局HTTP客户端
var client = &http.Client{
	Transport: &http.Transport{
		DisableKeepAlives: true,
	},
	Timeout: 2 * time.Second,
}

// Get 判断传入的url是否能正确访问
func Get(url string) bool {
	resp, err := client.Get(url)
	if err != nil {
		logx.Error("connect client.Get fail: ", err)
		return false
	}
	resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}
