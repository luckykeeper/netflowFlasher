// netflowFlasher 刷下行流量
package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// 下载链接
// 星铁pc，星铁apk，崩2，崩3PC完整，原pc，原apk
var downloadList = []string{"https://api-takumi.mihoyo.com/event/download_porter/link/hkrpg_cn/official/pc_default",
	"https://api-takumi.mihoyo.com/event/download_porter/link/hkrpg_cn/official/android_default",
	"https://static.benghuai.com/Download/v10_1/Original.StripResource_10.1.8_313_utw.shell.apk",
	"https://bundle.bh3.com/ptpublic/rel/20230416151857_AsBJm4PVPKKR38YI/PC/BH3_v6.6.0_4ed7d53313df.7z",
	"https://ys-api.mihoyo.com/event/download_porter/link/ys_cn/official/pc_default",
	"https://ys-api.mihoyo.com/event/download_porter/link/ys_cn/official/android_default"}

// 下载限速，单位KB
const (
	datachunk = 3 * 1024 * 1024 // 下载限速
	timelapse = 1               // per second
)

func main() {
	i := 0
	for {
		i = i + 1
		log.Println("第", i, "轮下载开始")
		for n, url := range downloadList {
			timeSleep := time.Duration(rand.Intn(10)) * time.Second
			log.Println("第", n, "个下载结束，等待", timeSleep)
			time.Sleep(timeSleep)
			log.Println("开始下载：", url)
			resp, err := http.Get(url)
			if err != nil {
				log.Println("Get failed:", err)
			} else {
				defer resp.Body.Close()
				contentLength := resp.ContentLength / 1024 / 1024

				file := io.Discard

				var alreadyDown int64

				for range time.Tick(timelapse * time.Second) {
					n, err := io.CopyN(file, resp.Body, datachunk)
					if err != nil {
						if err == io.EOF {
							log.Println(url, "下载完成")
						} else {
							log.Println("写入失败:", err)
						}
						break
					}
					alreadyDown = alreadyDown + n/1024/1024

					log.Println("已下载" + fmt.Sprint(alreadyDown) + "兆，完成百分之" + fmt.Sprint(alreadyDown*10/contentLength*10))
				}
			}
		}
	}
}
