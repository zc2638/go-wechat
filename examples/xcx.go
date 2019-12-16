package main

import (
	"fmt"
	"github.com/zc2638/wechat"
	"github.com/zc2638/wechat/xcx"
)

/**
 * Created by zc on 2019/12/10.
 */
func main() {

	w := wechat.NewWeChat("", "")

	code := xcx.Code{}
	w.Exec(&code)
	if code.Err != nil {
		fmt.Println(code.Result)
	}
	if code.Result.ErrCode != 0 {
		fmt.Println(code.Result.ErrMsg)
	}
}