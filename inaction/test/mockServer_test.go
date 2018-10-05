// 这个示例程序展示如何内部模仿 HTTP GET 调用
// 与本书之前的例子有些差别
package test

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

// feed 模仿了我们期望接收的 XML 文档
var feed = `<?xml version="1.0" encoding="UTF-8"?>
<rss>
<channel>
<title>Going Go Programming</title>
<description>Golang : https://github.com/goinggo</description>
<link>http://www.goinggo.net/</link>
<item>
<pubDate>Sun, 15 Mar 2015 15:04:00 +0000</pubDate>
<title>Object Oriented Programming Mechanics</title>
<description>Go is an object oriented language.</description>
<link>http://www.goinggo.net/2015/03/object-oriented</link>
</item>
</channel>
</rss>`

// mockServer 返回用来处理请求的服务器的指针
func mockServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/xml")
		fmt.Fprintln(w, feed)
	}
	return httptest.NewServer(http.HandlerFunc(f))
}