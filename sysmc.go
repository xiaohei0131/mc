package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"io/ioutil"
	"log"
	mo "mc/monitor"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
	"github.com/xiaohei0131/plock"
)

var Interval int
var URL string

const CT string = "application/json"

var logger *log.Logger

func init() {
	file, err := os.OpenFile("mc.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("fail to create mc.log file!")
	}
	logger = log.New(file, "", log.Llongfile)
	logger.SetFlags(log.LstdFlags)
}
func main() {
	plock.Lock()
	defer plock.UnLock()
	flag.IntVar(&Interval, "i", 5, "数据采集间隔(单位s)")
	flag.StringVar(&URL, "server", "", "数据中心地址")
	flag.Parse()
	if len(URL) == 0 {
		logger.Panicln("server参数为空")
		//syscall.Exit(-1)
	}
	logger.Println("**********服务已启动**********")
	logger.Println("数据中心地址", URL)
	logger.Println("数据采集间隔时间", Interval, "秒")
	ticker := time.NewTicker(time.Second * time.Duration(Interval))
	for _ = range ticker.C {
		start()
	}
}

func start() {
	//logger.Println("采集数据")
	monitor := map[string]interface{}{}

	monitor["tips"] = "Memory or storage in megabytes (Mib)"
	wg := sync.WaitGroup{}

	wg.Add(6)
	var sysInfo ,ip  string
	var memory,cpu,disk,gpu interface{}
	go func() {
		defer  wg.Done()
		sysInfo = mo.GetSysInfo(logger)
	}()
	go func() {
		defer  wg.Done()
		ip = mo.GetLocalIP(logger)
	}()
	go func() {
		defer  wg.Done()
		memory = mo.MemInfo(logger)
	}()
	go func() {
		defer  wg.Done()
		cpu = mo.CpuInfo(logger)
	}()
	go func() {
		defer  wg.Done()
		disk = mo.DiskMonitor(logger)
	}()
	go func() {
		defer  wg.Done()
		gpu = mo.GetGpuInfo(logger)
	}()
	wg.Wait()
	//logger.Println("上报数据")
	monitor["system"] = sysInfo
	monitor["ip"] = ip
	monitor["memory"] = memory
	monitor["cpu"] = cpu
	monitor["disk"] = disk
	monitor["gpu"] = gpu
	_, err := post(URL, monitor, CT)
	if err != nil {
		logger.Println("上报数据失败", err.Error())
	} /*else {
		logger.Println("上报数据成功")
	}*/
}

// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容
func post(url string, data interface{}, contentType string) (string, error) {
	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, _ := json.Marshal(data)
	//logger.Println(string(jsonStr))
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	re := string(result)
	if strings.Contains(re, "404") {
		return re, errors.New(re)
	}
	return re, nil
}
