package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	resp, err := http.Get("https://api6.ipify.org?format=json")
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read from resp.Body failed, err:%v\n", err)
		return
	}
	//json转struct,获取本地IPv6地址
	//声明一个叫ipv6Addr的结构体
	type ipv6 struct {
		Ip string `json:"ip"`
	}
	var addr ipv6
	error := json.Unmarshal(body, &addr)
	if err != nil {
		fmt.Println("json err:", error)
	}
	Ip_addr := addr.Ip //IP地址
	fmt.Println(Ip_addr)
	//与Dnspod中的域名记录中的IPv6地址进行比较
	//获取DnsPod记录，net/http包没有封装直接使用请求带header的get或者post方法，
	//所以，要想请求中带header或设置cookie(下面列子中有设置cookie的方法)，
	//只能使用NewRequest方法(使用该方法时需要先对client实例化:client := &http.Client{})

	data := url.Values{} //结合url.Values发送post请求
	dnspodId := "276813"
	dnspodToken := "d3f687bb0094cfb084ac5168744f5411"
	fmt.Sprintf("%s,%s", dnspodId, dnspodToken)
	data.Add("login_token", fmt.Sprintf("%s,%s", dnspodId, dnspodToken))
	data.Add("format", "json")
	data.Add("domain", "997744.xyz")
	data.Add("sub_domain", "nas")
	request, err = http.NewRequest("POST", "https://dnsapi.cn/Record.List", strings.NewReader(data.Encode()))
	if err != nil {
		err = errors.New("request对象创建失败,err:" + err.Error())
		return
	}

	//如果与记录相同，不做更新
	//如果与记录不通，把本地的ipv6地址更新到dnspod

}
