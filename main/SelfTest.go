package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}

func countLines2() { //统计文件的重复行
	counts := make(map[string]int)
	for _, fileName := range os.Args[1:] {
		data, err := os.ReadFile(fileName) //读取文件 本来是ioUtil.ReadFile ioUtil废弃了，改用os.ReadFile了 ,这里得到的data 是byte切片类型的
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") { // 先将data转成string 类型 ,再按换行符 \n 切分
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++

	}
}
