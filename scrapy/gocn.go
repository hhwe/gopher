package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

var userAgent = [...]string{
	"Mozilla/5.0 (compatible, MSIE 10.0, Windows NT, DigExt)",
	"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, 360SE)",
	"Mozilla/4.0 (compatible, MSIE 8.0, Windows NT 6.0, Trident/4.0)",
	"Mozilla/5.0 (compatible, MSIE 9.0, Windows NT 6.1, Trident/5.0,",
	"Opera/9.80 (Windows NT 6.1, U, en) Presto/2.8.131 Version/11.11",
	"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, TencentTraveler 4.0)",
	"Mozilla/5.0 (Windows, U, Windows NT 6.1, en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	"Mozilla/5.0 (Macintosh, Intel Mac OS X 10_7_0) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11",
	"Mozilla/5.0 (Macintosh, U, Intel Mac OS X 10_6_8, en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	"Mozilla/5.0 (Linux, U, Android 3.0, en-us, Xoom Build/HRI39) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13",
	"Mozilla/5.0 (iPad, U, CPU OS 4_3_3 like Mac OS X, en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
	"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, Trident/4.0, SE 2.X MetaSr 1.0, SE 2.X MetaSr 1.0, .NET CLR 2.0.50727, SE 2.X MetaSr 1.0)",
	"Mozilla/5.0 (iPhone, U, CPU iPhone OS 4_3_3 like Mac OS X, en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
	"MQQBrowser/26 Mozilla/5.0 (Linux, U, Android 2.3.7, zh-cn, MB200 Build/GRJ22, CyanogenMod-7) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

var urlChannel = make(chan string, 200) //chan中存入string类型的href属性,缓冲200

var atagRegExp = regexp.MustCompile(`<a[^>]+[(href)|(HREF)]\s*\t*\n*=\s*\t*\n*[(".+")|('.+')][^>]*>[^<]*</a>`) //以Must前缀的方法或函数都是必须保证一定能执行成功的,否则将引发一次panic
var titleRegExp = regexp.MustCompile(`<article class="article box_white">`)

/*
<article class="article box_white">
			<div class="row">
				<div>
					<h2>
						<img src="https://static.studygolang.com/gopher/default_project.jpg" alt="jiacrontab" width="36px">
						<a href="/p/jiacrontab" target="_blank" title="jiacrontab">可视化定时任务管理工具 <em>jiacrontab</em></a>
					</h2>
					<p class="text">提供可视化界面的定时任务管理工具。

允许设置每个脚本的超时时间，超时操作可选择邮件通知管理者，或强杀脚本进程。

允许设置脚本的最大并发数。

一台server管理多个client。

每个脚本都可在server端灵活配置，如测试脚本运行，查看日志，强杀进程，停止定时...。

允许添加脚本依赖（支持跨服务器），依赖脚本提供同步和异步的执行模式。

友好的web界面，方便用户操作。

脚本出错时可选择邮箱通知多人。

jiacrontab由server，client两部分构成，两者完全独立通过...<a href="/p/jiacrontab" target="_blank" title="阅读全文">阅读全文</a></p>
				</div>
			</div>
			<div class="row">
				<div class="col-md-8 metatag">
					<i class="glyphicon glyphicon-calendar"></i>
					<span class="date" title="发布日期">2018-07-16 12:30:01</span>
					<i class="glyphicon glyphicon-user"></i>
					<span class="author" title="作者">网友</span>
				</div>
				<div class="col-md-4 metatag text-right">
					<span class="view" title="阅读数">
						<i class="glyphicon glyphicon-eye-open"></i>
						阅读:<span>75</span>次
					</span>
					<a href="/p/jiacrontab#commentForm" class="cmt" target="_blank" title="评论数">
						<i class="glyphicon glyphicon-comment"></i>
						评论:<span>1</span>条
					</a>

					<a href="#" class="like" title="我喜欢" data-objid="841" data-objtype="4" data-flag="0">
						<i class="glyphicon glyphicon-heart-empty"></i>

						<span class="likenum">0</span>人喜欢
					</a>
				</div>
			</div>
		</article>
*/
func GetRandomUserAgent() string {
	return userAgent[r.Intn(len(userAgent))]
}

func GetHref(atag string) (href, content string) {
	inputReader := strings.NewReader(atag)
	decoder := xml.NewDecoder(inputReader)
	for t, err := decoder.Token(); err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
		// 处理元素开始（标签）
		case xml.StartElement:
			for _, attr := range token.Attr {
				attrName := attr.Name.Local
				attrValue := attr.Value
				if strings.EqualFold(attrName, "href") || strings.EqualFold(attrName, "HREF") {
					href = attrValue
				}
			}
		// 处理元素结束（标签）
		case xml.EndElement:
		// 处理字符数据（这里就是元素的文本）
		case xml.CharData:
			content = string([]byte(token))
		default:
			href = ""
			content = ""
		}
	}
	return href, content
}

func crawl(url string, urlChannel chan string) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("[E]", r)
		}
	}()
	fmt.Println(url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("User-Agent", GetRandomUserAgent())
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln("get url wrong")
	}
	fmt.Println(res.StatusCode)
	if res.StatusCode == 200 {
		body := res.Body
		defer body.Close()

		bodyByte, err := ioutil.ReadAll(body)
		if err != nil {
			log.Fatalln(err)
		}
		resStr := string(bodyByte)
		fmt.Println(resStr[:2])
		// atag := atagRegExp.FindAllString(resStr, -1)
		// fmt.Println(atag)
		// for _, a := range atag {
		// 	href, _ := GetHref(a)
		// 	if strings.Contains(href, "article/details/") {
		// 		fmt.Println("☆", href)
		// 	} else {
		// 		fmt.Println("□", href)
		// 	}
		// 	urlChannel <- href
		// }
		titles := titleRegExp.FindAllString(resStr, -1)
		fmt.Println(titles)
	}
	urlChannel <- url
}

func main() {
	args := os.Args
	if len(args) != 2 {
		log.Fatalln("error args, 2 need")
	}

	url := args[1]
	defer close(urlChannel)

	go crawl(url, urlChannel)

	p := <-urlChannel
	fmt.Println(p, url)
}
