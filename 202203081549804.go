package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"

	//"io/ioutil"
	"net"
	//"net/http"
	"time"

	//"fmt"
	"io/ioutil"
	"net/http"
	//"time"
)

func hmacSha256(stringToSign string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(stringToSign))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// Sign 发送钉钉消息
func Sign() string {
	secret := "SEC5f337113b3d0b2a32369d984f223e6fac7b4237bf2ec8d3d49e9d81932f38bc7"
	webhook := "https://oapi.dingtalk.com/robot/send?access_token=2c40e4bf1e42e1402b4fa762f6c599ff3d70360372733cdf6f52a766458cb878"
	timestamp := time.Now().UnixNano() / 1e6
	stringToSign := fmt.Sprintf("%d\n%s", timestamp, secret)
	sign := hmacSha256(stringToSign, secret)
	url := fmt.Sprintf("%s&timestamp=%d&sign=%s", webhook, timestamp, sign)
	return url
}

func dingToInfo(s, url, phone string) bool {
	content, data, at_user := make(map[string]string), make(map[string]interface{}), make(map[string]string)
	content["content"] = s
	at_user["atMobiles"] = phone
	data["msgtype"] = "text"
	data["text"] = content
	data["at"] = at_user
	b, _ := json.Marshal(data)
	//fmt.Println(data)//打印调试信息
	resp, err := http.Post(url,
		"application/json",
		bytes.NewBuffer(b))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	return true
}

func tcping(servername, ip, phone string) {
	var err error
	var conn net.Conn
	conn, err = net.DialTimeout("tcp", ip, 5*time.Second)
	if err != nil {
		//TG推送
		msg := fmt.Sprintf("https://tg.piqiucloud.workers.dev/?text=%s服务告警！！！%s无法ping通了，请检查服务状态！", servername, ip)
		fmt.Printf("%s服务告警！！！%s无法ping通了，请检查服务状态！\n", servername, ip)
		resp, err := http.Get(msg)
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()
		//body, err := ioutil.ReadAll(resp.Body)
		//if err != nil {
		// 	fmt.Println(err)
		// }
		//fmt.Println(string(body))
		//钉钉推送
		//msg := fmt.Sprintf("%s服务告警！！！%s无法ping通了，请检查服务状态！", servername, ip)
		//url := Sign()
		//dingToInfo(msg, url, phone)
	} else {
		conn.Close()
		fmt.Printf("亲，%v服务正常！！！%v端口通信正常！\n", servername, ip)
	}
}

//测试服务是否可用  IP:PORT
func main() {
	fmt.Println("安监系统端口在线检测工具，每60S检测一次端口存活情况！！！")
	fmt.Println("正在为您执行检测，请稍等...")
	ticker := time.Tick(time.Second)
	for range ticker {
		//钉钉推送
		// tcping("安监系统", "10.0.137.180:8086", "[13219511229]")
		//tcping("湿地大屏", "10.0.137.180:8084", "[13219511229]")
		tcping("领导驾驶舱", "10.0.137.179:80", "[15082631514]")
		tcping("领导驾驶舱", "10.0.137.179:1947", "[15082631514]")
		tcping("领导驾驶舱", "10.0.137.179:5985", "[15082631514]")
		tcping("领导驾驶舱", "10.0.137.179:8090", "[15082631514]")
		tcping("领导驾驶舱", "10.0.137.179:23302", "[15082631514]")
		tcping("领导驾驶舱", "10.0.137.179:47001", "[15082631514]")
		tcping("领导驾驶舱", "10.0.137.179:49152", "[15082631514]")
		tcping("领导驾驶舱", "10.0.137.179:49153", "[15082631514]")
		tcping("领导驾驶舱", "10.0.137.179:49154", "[15082631514]")
		tcping("领导驾驶舱", "10.0.137.179:49155", "[15082631514]")
		tcping("领导驾驶舱", "10.0.137.179:49156", "[15082631514]")
		tcping("领导驾驶舱", "10.0.137.179:49157", "[15082631514]")
		tcping("领导驾驶舱", "10.0.137.179:49158", "[15082631514]")
		tcping("领导驾驶舱", "10.0.137.179:49158", "[15082631514]")
		tcping("安监应用服务器", "10.0.137.180:80", "[15082631514]")
		tcping("安监应用服务器", "10.0.137.180:567", "[15082631514]")
		tcping("安监应用服务器", "10.0.137.180:1191", "[15082631514]")
		tcping("安监应用服务器", "10.0.137.180:1801", "[15082631514]")
		tcping("安监应用服务器", "10.0.137.180:2103", "[15082631514]")
		tcping("安监应用服务器", "10.0.137.180:2103", "[15082631514]")
		tcping("安监应用服务器", "10.0.137.180:2107", "[15082631514]")
		tcping("安监应用服务器", "10.0.137.180:5985", "[15082631514]")
		//tcping("安监应用服务器", "10.0.137.180:8081", "[15082631514]")
		tcping("安监应用服务器", "10.0.137.180:8082", "[15082631514]")
		tcping("安监应用服务器", "10.0.137.180:8083", "[15082631514]")
		tcping("安监应用服务器", "10.0.137.180:23302", "[15082631514]")
		tcping("安监应用服务器", "10.0.137.180:47001", "[15082631514]")
		tcping("安监应用服务器", "10.0.137.180:49152", "[15082631514]")
		tcping("安监应用服务器", "10.0.137.180:49153", "[15082631514]")
		tcping("安监应用服务器", "10.0.137.180:49154", "[15082631514]")
		tcping("安监应用服务器", "10.0.137.180:49155", "[15082631514]")
		tcping("安监应用服务器", "10.0.137.180:49156", "[15082631514]")
		tcping("安监应用服务器", "10.0.137.180:49161", "[15082631514]")
		tcping("安监应用服务器", "10.0.137.180:49162", "[15082631514]")
		//tcping("安监应用服务器", "218.89.67.36:1191", "[15082631514]") //APP公网映射
		tcping("安监应用服务器", "10.0.137.181:1433", "[15082631514]")
		tcping("安监数据服务器", "10.0.137.181:2383", "[15082631514]")
		tcping("安监数据服务器", "10.0.137.181:5985", "[15082631514]")
		tcping("安监数据服务器", "10.0.137.181:8391", "[15082631514]")
		tcping("安监数据服务器", "10.0.137.181:23302", "[15082631514]")
		tcping("安监数据服务器", "10.0.137.181:47001", "[15082631514]")
		tcping("安监数据服务器", "10.0.137.181:49152", "[15082631514]")
		tcping("安监数据服务器", "10.0.137.181:49153", "[15082631514]")
		tcping("安监数据服务器", "10.0.137.181:49154", "[15082631514]")
		tcping("安监数据服务器", "10.0.137.181:49155", "[15082631514]")
		tcping("安监数据服务器", "10.0.137.181:49156", "[15082631514]")
		tcping("安监数据服务器", "10.0.137.181:49157", "[15082631514]")
		tcping("安监数据服务器", "10.0.137.181:49159", "[15082631514]")
		tcping("安监数据服务器", "10.0.137.181:49160", "[15082631514]")
		tcping("安监数据服务器", "10.0.137.181:49161", "[15082631514]")
		tcping("安监数据服务器", "10.0.137.181:49385", "[15082631514]")

	}

}
