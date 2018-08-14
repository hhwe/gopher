// 在下面的例子中，我们结合使用了缓冲读取文件和命令行 flag 解析这两项技术。
// 如果不加参数，那么你输入什么屏幕就打印什么。
// 参数被认为是文件名，如果文件存在的话就打印文件内容到屏幕。
// 命令行执行 cat test 测试输出。
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var Num = flag.Bool("n", false, "print newline")

func main() {
	flag.PrintDefaults()
	flag.Parse()
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdout))
	}

	for i := 0; i < flag.NArg(); i++ {
		fmt.Println(flag.Args(), *Num)
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s:error reading from %s: %s\n",
				os.Args[0], flag.Arg(i), err.Error())
			continue
		}
		cat(bufio.NewReader(f))
	}
}

func cat(r *bufio.Reader) {
	numOfLine := 0
	for {
		buf, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if *Num {
			fmt.Fprintf(os.Stdout, "%d %s", numOfLine, buf)
		} else {
			fmt.Fprintf(os.Stdout, "%s", buf)
		}
		numOfLine++
	}
	return
}

func catf(f *os.File) {
	const NBUF = 512
	var buf [NBUF]byte
	for {
		switch nr, err := f.Read(buf[:]); true {
		case nr < 0:
			fmt.Fprintf(os.Stderr, "cat: error reading: %s\n", err.Error())
			os.Exit(0)
		case nr == 0:
			return
		case nr > 0:
			if nw, ew := os.Stdout.Write(buf[0:nr]); nw != nr {
				fmt.Fprintf(os.Stderr, "cat: error writing: %s\n", ew.Error())
			}
		}
	}
}
