package main

import (
	"flag"
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/os/gfile"
	"net"
	"strings"
)

func CNAME(src string) (dst string, err error)  {
	dst, err = net.LookupCNAME(src)
	return
}

func CnameResult(src string) (c string) {
	a, _ := CNAME(src)
	b := strings.Split(a, ".")
	c = b[len(b)-3] + "." + b[len(b)-2]
	return
}

func CdnParse(src string) (isCdn bool) {
	content := gfile.GetContents("cdn.json")
	result , _ := gjson.LoadContent(content)

	for k, _ := range result.Map(){
		if src == k {
			 fmt.Println(result.Map()[k])
			return true
		}
	}
	return false
}

func main()  {
	var targetURL string
	flag.StringVar(&targetURL, "u", "", "")
	flag.CommandLine.Usage = func() { fmt.Println("使用说明：\n执行命令：./main -u www.baidu.com ") }
	flag.Parse()
	if len(targetURL) == 0{
		fmt.Println("使用说明：\n执行命令：./main -u www.baidu.com ")
	}else {
		a := CnameResult(targetURL)
		fmt.Println(CdnParse(a))
	}
}