import os
import requests


url = 'https://www.zhihu.com/signup'
login_url = "https://zhihu-web-analytics.zhihu.com/api/v1/logs/batch"
headers = {
    "accept": "*/*",
    "accept-encoding": "gzip, deflate, br",
    "accept-language": "zh-CN,zh;q=0.9",
    "content-encoding": "gzip",
    "content-length": "465",
    "content-type": "application/x-protobuf",
    "origin": "https://www.zhihu.com",
    "referer": "https://www.zhihu.com/signup?next=%2F",
    "user-agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko)Chrome/68.0.3440.106 Safari/537.36",
    "x-za-batch-size": "1",
    "x-za-log-version": "2.3.91",
    "x-za-platform": "1",
    "x-za-product": "1",
}


resp = requests.post(login_url, headers=headers, data=data)
# resp.raise_for_status()
print(resp.status_code)
resp.encoding == resp.apparent_encoding
print(resp.text.encode(resp.encoding).decode())
