package main

// https://studygolang.com/pkgdoc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type myIp struct {
	Address string `json.address`
	Proto   string `json.proto`
}

type IpResStatus struct {
	Code       string `json.code`
	Message    string `json.message`
	Created_at string `json.created_at`
}
type IPRES struct {
	Status *IpResStatus `json.status`
}

const (
	getipurl string = "http://v4v6.ipv6-test.com/json/widgetdata.php"

	setipurl string = "https://dnsapi.cn/Record.Modify"
)

var (
	params map[string]string = map[string]string{
		"login_token": "182442,6479127c66c1c23eb594cd9a10abdce6",
		"format":      "json",
		"domain":      "yuanmeizi.tk",
		"record_id":   "665003013",
		"sub_domain":  "nas",
		"record_line": "默认",
		"record_type": "AAAA",
		"value":       "",
	}

	myip  myIp  = myIp{}
	ipres IPRES = IPRES{}
)

func getMyIp() myIp {
	resp, err := http.Get(getipurl)
	if err != nil {
		fmt.Println("获取本地ip地址错误", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	s := strings.TrimLeft(string(body), "(")
	s = strings.TrimRight(s, ")")
	err = json.Unmarshal([]byte(s), &myip)
	if err != nil {
		fmt.Println("获取本机ip的json解析错误", err)
	}
	return myip
}

func main() {
	var ipaddr myIp = getMyIp()
	params["value"] = ipaddr.Address
	if ipaddr.Proto == "ipv6" {
		fmt.Println("本地有ipv6地址,开始设置")
	} else {
		fmt.Println("本地没有ipv6地址,准备开始设置ipv4解析")
		params["record_type"] = "A"
	}
	data := url.Values{}
	for k, v := range params {
		data.Set(k, v)
	}
	resp, err := http.PostForm(setipurl, data)
	if err != nil {
		fmt.Println("通过api设置解析错误", err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(b, &ipres)
	if err != nil {
		fmt.Println("设置域名解析的json解析错误", err)
	}
	if ipres.Status.Code == "1" {
		fmt.Println("域名解析设置成功,本次设置时间为:", ipres.Status.Created_at)
	}
}
