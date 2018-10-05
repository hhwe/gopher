// 这个示例程序展示如何写基础单元测试
package test

import (
	"net/http"
	"testing"
)

//const checkMark = "\u2713"
//const ballotX = "\u2717"

// TestDownload 确认 http 包的 Get 函数可以下载内容
// 并且内容可以被正确地反序列化并关闭
func TestsDownload(t *testing.T) {
	statusCode := http.StatusOK
	server := mockServer()
	defer server.Close()
	t.Log("Given the need to test downloading content.")
	{
		t.Logf("\tWhen checking \"%s\" for status code \"%d\"",
			server.URL, statusCode)
		{
			resp, err := http.Get(server.URL)
			if err != nil {
				t.Fatal("\t\tShould be able to make the Get call.",
					ballotX, err)
			}
			t.Log("\t\tShould be able to make the Get call.",
				checkMark)
			defer resp.Body.Close()
			if resp.StatusCode != statusCode {
				t.Fatalf("\t\tShould receive a \"%d\" status. %v %v",
					statusCode, ballotX, resp.StatusCode)
			}
			t.Logf("\t\tShould receive a \"%d\" status. %v",
				statusCode, checkMark)
		}
	}
}