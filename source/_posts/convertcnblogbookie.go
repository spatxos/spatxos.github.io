package main

import (
	"fmt"
	"flag"
	"strings"
)
var cookie string
func init() {
    flag.StringVar(&cookie,"cookie","","cnblog of cookie")
}
func main() {
	flag.Parse()//暂停获取参数
    println(cookie)
	if(len(cookie)>0){
		newcookie := strings.Replace(cookie,";","_semicolon_",-1)
		newcookie = strings.Replace(newcookie,"|","_vertical_",-1)
		newcookie = strings.Replace(newcookie,"(","_frontbracket_",-1)
		newcookie = strings.Replace(newcookie,")","_backbracket_",-1)
		newcookie = strings.Replace(newcookie," ","_space_",-1)
		fmt.Printf("\r\n newcookie:%s \n", newcookie)
	}else{
		fmt.Printf("未输入正确的CNBLOGS_COOKIE，请到github项目/Settings/Secrets/Actions下通过New repository secret添加一个CNBLOGS_COOKIE并填入正确的cookie值")
	}
}