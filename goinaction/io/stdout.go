// 这个示例程序展示来自不同标准库的不同函数是如何
// 使用 io.Writer 接口的
package main

import (
	"bytes"
	"fmt"
	"os"
)

// main 是应用程序的入口
func main() {
	// 创建一个 Buffer 值,并将一个字符串写入 Buffer
	// 使用实现 io.Writer 的 Write 方法
	var b bytes.Buffer
	b.Write([]byte("Hello "))
	// 使用 Fprintf 来将一个字符串拼接到 Buffer 里
	// 将 bytes.Buffer 的地址作为 io.Writer 类型值传入
	fmt.Fprintf(&b, "World!")
	// 将 Buffer 的内容输出到标准输出设备
	// 将 os.File 值的地址作为 io.Writer 类型值传入
	b.WriteTo(os.Stdout)
}
