package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	f, err := os.Create("fzu.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} //创建文件
	defer f.Close()
	for i := 0; i < 12; i++ { //12
		url := "https://www.fzu.edu.cn/index/fdyw/" + strconv.Itoa(i+3) + ".htm"
		fmt.Printf("url: %v\n", url)
		fetchHttm(request(url))

	} //福大要闻在53和54页左右爬取的不再是完整链接，暂时跳过一部分内容
	for j := 0; j < 9; j++ { //9
		url := "https://www.fzu.edu.cn/index/fdyw/" + strconv.Itoa(j+65) + ".htm"
		fmt.Printf("url: %v\n", url)
		fetchHttm(request(url))

	}
}
func request(url string) *goquery.Document {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.71 Safari/537.36 Core/1.94.188.400 QQBrowser/11.4.5225.400")
	req.Header.Add("Cookie", "_ga=GA1.3.791916269.1662601994; _gscu_1331749010=7037965318u87b20; _gscbrs_1331749010=1; JSESSIONID=FCAA93E887DB3AF1C9C9E2810FDFE3B1")
	resp, _ := client.Do(req)
	dom, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("解析失败")
	}
	return dom

}

func fetchHttm(dom *goquery.Document) {
	dom.Find("body > section > section.n_container > div > div > div.n_right.fr > section.n_list > ul > li > a").Each(func(i int, s *goquery.Selection) {
		http, ok := s.Attr("href")
		if ok {
			fmt.Printf("http: %v\n", http)
			if strings.Count(http, "news") == 2 { //观察到旧链接有两个news，新链接则只有一个，以此区分不同解析法
				num := http[39 : len(http)-4]
				Parse1(request(http), num)
			} else {
				num := http[34 : len(http)-4]
				Parse2(request(http), num)
			}
		}

	})
	fmt.Println("-----------------------------------")
}
func Parse1(dom *goquery.Document, num string) {
	file, _ := os.OpenFile("fzu.txt", os.O_RDWR|os.O_APPEND, 0775)
	dom.Find("#main > div.right > form").Each(func(i int, s *goquery.Selection) {
		title := s.Find("div.detail_main_content > p:nth-child(1)").Text()
		date := s.Find("#fbsj").Text()
		author := s.Find("#author").Text()
		article := s.Find(".v_news_content").Text()
		file.WriteString("标题:" + string(title) + "\n")
		file.WriteString("作者:" + string(author) + "\n")
		file.WriteString("日期:" + string(date) + "\n")
		file.WriteString("正文" + string(article) + "\n")
		// fmt.Printf("标题: %v\n", title)
		// fmt.Printf("作者: %v\n", author)
		// fmt.Printf("日期: %v\n", date)
		// fmt.Printf("正文: %v\n", article)
	})
	readApi := `https://news.fzu.edu.cn/system/resource/code/news/click/dynclicks.jsp?clickid=` + num + `&owner=1744991928&clicktype=wbnews`
	//fmt.Println(readApi)
	readcount := request(readApi)
	file.WriteString("阅读数" + string(readcount.Text()) + "\n")
	err := file.Close()
	fmt.Printf("err: %v\n", err)
	//fmt.Printf("阅读数: %v\n", readcount.Text())

}
func Parse2(dom *goquery.Document, num string) {
	file, _ := os.OpenFile("fzu.txt", os.O_RDWR|os.O_APPEND, 0775)
	dom.Find("body > section > section.n_container > div > div.n_right.fr > section > form > div").Each(func(i int, s *goquery.Selection) {
		title := s.Find("div.nav01 > h3").Text()
		date := s.Find("div.nav01 > h6 > span:nth-child(1)").Text()
		author := s.Find("div.nav01 > h6 > span:nth-child(2)").Text()
		article := s.Find(".v_news_content").Text()
		// fmt.Printf("%v\n", title)
		// fmt.Printf("%v\n", author)
		// fmt.Printf("%v\n", date)
		// fmt.Printf("正文: %v\n", article)
		file.WriteString(string(title) + "\n")
		file.WriteString(string(author) + "\n")
		file.WriteString(string(date) + "\n")
		file.WriteString("正文" + string(article) + "\n")
	})
	readApi := `https://news.fzu.edu.cn/system/resource/code/news/click/dynclicks.jsp?clickid=` + num + `&owner=1779559075&clicktype=wbnews`
	//fmt.Println(readApi)
	readcount := request(readApi)
	//fmt.Printf("阅读数: %v\n", readcount.Text())
	file.WriteString("阅读数" + string(readcount.Text()) + "\n")
	err := file.Close()
	fmt.Printf("err: %v\n", err)
}
