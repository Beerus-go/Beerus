package controller

import (
	"github/yuyenews/beerus/network/http/commons"
	"github/yuyenews/beerus/web/route"
	"io/ioutil"
)

func CreateRoute() {
	route.GET("/downLoad/file", func(req *commons.BeeRequest, res *commons.BeeResponse) {
		file, err := ioutil.ReadFile("/Users/yeyu/Downloads/goland-2021.2.4.dmg")
		if err == nil {

		}
		res.SendStream("哈哈哈哈.dmg", file)
	})
}
