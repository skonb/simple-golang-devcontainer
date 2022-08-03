package main

import (
	"flag" //コマンドライン引数を適当に解析してくれるやつ
	"fmt"
	"os"
	"strconv"
)

func main() {
	var help = false
	//var help bool //ゼロ値で初期化されるのでfalseが入る
	//help := false

	var format = "kanji"
	//var format string

	Roman := [10]string{"0", "Ⅰ", "Ⅱ", "Ⅲ", "Ⅳ", "Ⅴ", "Ⅵ", "Ⅶ", "Ⅷ", "Ⅸ"}
	Kanji := [10]string{"零", "壱", "弐", "参", "肆", "伍", "陸", "質", "捌", "玖"}

	dict := map[string][10]string{
		"kanji": Kanji,
		"roman": Roman,
	}

	flag.BoolVar(&help, "h", false, "show help")
	flag.StringVar(&format, "f", "kanji", "set kanji or roman")
	flag.Parse()

	if help {
		flag.Usage()
		return
	}

	for i, arg := range flag.Args() {
		atoiArg, err := strconv.Atoi(arg)
		if err == nil {
			if _, err := dict[format]; err != true {
				//panicでも落ちるけど雑すぎ，失格．丁寧にやる
				fmt.Println("Format Okasii??????????????")
				os.Exit(0)
			}
			fmt.Printf("%d %s\n", i, dict[format][atoiArg])
		} else {
			fmt.Println("Atoi Error")
			os.Exit(0)
		}
	}
}
