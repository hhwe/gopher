package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://www.zhihu.com/"
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	request.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")
	request.Header.Set("cookie", `_zap=5c6477f9-796d-4933-a0d2-723a146d8595; _xsrf=e4e5c213-5b47-49e5-aa47-cc735d10ab63; d_c0="ALBmz8qL_Q2PToFSUcE8FXPHsDqBA2skW7g=|1533133487"; q_c1=8c5c636abba244fe9fab4035e77c531f|1533133489000|1533133489000; l_n_c=1; n_c=1; l_cap_id="N2JjOGJlODlkODA3NDllY2JjYzRmYzc3ZmQ5YjRhMTU=|1534565780|c2fa725cfcfbf1a476671420a4f7bc7894d29e24"; r_cap_id="ODZhNzAzZDUzZTA2NGQ3NzhkOWE5ZWQ4ZTIwZmQ0NjY=|1534565780|47c41f7a00e414980c8d7f4511543f9c97a88450"; cap_id="OTEwMzYwNDA4NjRjNDZjYThhYjk4NmRiMmU0NjdiYzI=|1534565780|14cd994bffe6d3038cb21197ae5efeffdc9ce53c"; capsion_ticket="2|1:0|10:1534568696|14:capsion_ticket|44:NmJlN2ZhMGQ1MTg5NDQ5ZGJhZGQyOWEyZDgxMDNjN2Y=|b858876694f94e2bd7f1527a99029b5a75bbaf23192726d68f291e74b0a00ab6"; z_c0="2|1:0|10:1534568699|4:z_c0|92:Mi4xZWNjMEFBQUFBQUFBc0diUHlvdjlEU1lBQUFCZ0FsVk4tX1prWEFDNy1sMklnSl91YnpHa0wySzk1eERIRWNyY21R|65836ca126e7e9bde59fa08a691d5e73fb281754996f9a9224251f5cef956099"; tgw_l7_route=170010e948f1b2a2d4c7f3737c85e98c`)

	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Status)

	html, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("zhihu.html", html, 0666)
	if err != nil {
		panic(err)
	}
}
