package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"go-randpick/internal/randpick"
)

func main() {
	n := flag.Int("n", 1, "挑几项")
	noDup := flag.Bool("no-dup", false, "不重复挑同一项")
	sorted := flag.Bool("sorted", false, "按字母序排输出（从文件读时才有意义）")
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		picked, err := randpick.Pick(args, *n, *noDup)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		for _, v := range picked {
			fmt.Println(v)
		}
		return
	}

	// 从标准输入读行
	var lines []string
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line != "" {
			lines = append(lines, line)
		}
	}
	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "读输入失败:", err)
		os.Exit(1)
	}
	picked, err := randpick.Pick(lines, *n, *noDup)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if *sorted {
		randpick.SortLines(picked)
	}
	for _, v := range picked {
		fmt.Println(v)
	}
}

func usage() {
	fmt.Fprint(os.Stderr, `go-randpick — 随机挑

用法:
  go-randpick [选项] [候选项...]      不给候选项就读标准输入

选项:
  -n N         挑几项，默认 1
  -no-dup      不重复挑同一项
  -sorted      输出前按字母序排

例子:
  go-randpick -n 3 -no-dup alice bob carol dave eve
  cat names.txt | go-randpick -n 2 -no-dup
  go-randpick -n 1000 -no-dup -sorted  < biglist.txt
`)
}
