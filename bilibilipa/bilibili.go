package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/tidwall/gjson"
)

func main() {
	f, err := os.Create("comment.txt")
	file, _ := os.OpenFile("comment.txt", os.O_RDWR|os.O_APPEND, 0775)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} //创建文件
	defer f.Close()
	url := `https://api.bilibili.com/x/v2/reply/main?&jsonp=jsonp&next=0&type=1&oid=21071819`
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.71 Safari/537.36 Core/1.94.188.400 QQBrowser/11.4.5225.400")
	req.Header.Add("Cookie", "_uuid=87E4B1C7-C254-10F67-105E1-3F13EB6910421010876infoc; b_nut=1660492611; buvid3=545BC5FC-3AA4-8F7B-33D3-49091AAD1DCA12548infoc; i-wanna-go-back=-1; buvid_fp_plain=undefined; DedeUserID=13528580; DedeUserID__ckMd5=aaab2234010ffb1b; hit-dyn-v2=1; b_ut=5; LIVE_BUVID=AUTO4016606616954647; nostalgia_conf=-1; CURRENT_BLACKGAP=0; blackside_state=1; buvid4=3B0E9AF3-D9D3-64F2-F71D-7D8EE094FCE012548-022081423-Vk7oLekZ8O%2BXf1iUIja6HA%3D%3D; fingerprint3=0afcbd7c052fad6b3200206dd28ff0a5; is-2022-channel=1; hit-new-style-dyn=0; fingerprint=b1c5f5ba701327881eb0b3ea5a27cd45; CURRENT_FNVAL=4048; rpdid=|(u))kkYu|~u0J'uYY)l))lRY; buvid_fp=b1c5f5ba701327881eb0b3ea5a27cd45; CURRENT_QUALITY=64; SESSDATA=1cd7a43b%2C1686206438%2Cc709a%2Ac2; bili_jct=c044c1004c9c8e589809c2e1ea9289f1; innersign=0; b_lsid=3A108A51E_184FBA1746D; bp_video_offset_13528580=737995975713357800; sid=5g9onudx; PVID=6")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("err11", err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	result := string(bodyText) //将resp转化为字符串
	//fmt.Println(result)
	if err != nil {
		fmt.Println("io err", err)
	}
	reply1 := gjson.Get(result, "data.replies.0.content.message") //主评论
	fmt.Println(reply1.String())
	file.WriteString(reply1.String() + "\n")
	reply2 := gjson.Get(result, "data.replies.0.replies.#.content.message") //前三条子评论
	fmt.Println(reply2.String())
	file.WriteString(reply2.String() + "\n")
	for i := 0; i < 22; i++ {
		url1 := `https://api.bilibili.com/x/v2/reply/reply?&jsonp=jsonp&pn=` + strconv.Itoa(i+1) + `&type=1&oid=21071819&ps=10&root=2459609812&_=1670727453733`
		client := &http.Client{}
		req, _ := http.NewRequest("GET", url1, nil)
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.71 Safari/537.36 Core/1.94.188.400 QQBrowser/11.4.5225.400")
		req.Header.Add("Cookie", "_uuid=87E4B1C7-C254-10F67-105E1-3F13EB6910421010876infoc; b_nut=1660492611; buvid3=545BC5FC-3AA4-8F7B-33D3-49091AAD1DCA12548infoc; i-wanna-go-back=-1; buvid_fp_plain=undefined; DedeUserID=13528580; DedeUserID__ckMd5=aaab2234010ffb1b; hit-dyn-v2=1; b_ut=5; LIVE_BUVID=AUTO4016606616954647; nostalgia_conf=-1; CURRENT_BLACKGAP=0; blackside_state=1; buvid4=3B0E9AF3-D9D3-64F2-F71D-7D8EE094FCE012548-022081423-Vk7oLekZ8O%2BXf1iUIja6HA%3D%3D; fingerprint3=0afcbd7c052fad6b3200206dd28ff0a5; is-2022-channel=1; hit-new-style-dyn=0; fingerprint=b1c5f5ba701327881eb0b3ea5a27cd45; CURRENT_FNVAL=4048; rpdid=|(u))kkYu|~u0J'uYY)l))lRY; buvid_fp=b1c5f5ba701327881eb0b3ea5a27cd45; CURRENT_QUALITY=64; SESSDATA=1cd7a43b%2C1686206438%2Cc709a%2Ac2; bili_jct=c044c1004c9c8e589809c2e1ea9289f1; innersign=0; b_lsid=3A108A51E_184FBA1746D; bp_video_offset_13528580=737995975713357800; sid=5g9onudx; PVID=6")
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("err11", err)
		}
		defer resp.Body.Close()
		bodyText, err := ioutil.ReadAll(resp.Body)
		result := string(bodyText) //将resp转化为字符串
		if err != nil {
			fmt.Println("io err", err)
		}
		reply3 := gjson.Get(result, "data.replies.#.content.message")
		fmt.Println(reply3.String())
		file.WriteString(reply3.String() + "\n")
	}

}
