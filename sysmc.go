package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	mo "mc/monitor"
	"net/http"
	"strings"
	"syscall"
	"time"
)
var Interval int
var URL string

const CT string = "application/json"
func main()  {
	flag.IntVar(&Interval, "i", 5, "数据采集间隔(单位s)")
	flag.StringVar(&URL, "server", "", "上报服务器地址")
	flag.Parse()
	if len(URL)==0{
		fmt.Errorf("server参数为空")
		syscall.Exit(-1)
	}
	ticker := time.NewTicker(time.Second * time.Duration(Interval))
	for _ = range ticker.C {
		start()
	}
}

func start()  {
	log.Println("采集数据")
	monitor := map[string]interface{}{}
	//monitor["disk"] = diskMonitor()

	monitor["system"] = mo.GetSysInfo()
	monitor["ip"] = mo.GetLocalIP()
	monitor["memory"] = mo.MemInfo()
	monitor["cpu"] = mo.CpuInfo()
	monitor["disk"] = mo.DiskMonitor()

	log.Println("上报数据")
	re,err := post(URL,monitor,CT)
	if err != nil{
		log.Println("上报数据失败",re)
	}else{
		log.Println("上报数据成功",re)
	}
}

// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容
func post(url string, data interface{}, contentType string) (string,error) {
	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		return "",err
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil || strings.Contains(string(result),"404"){
		return string(result),err
	}
	return string(result),nil
}




