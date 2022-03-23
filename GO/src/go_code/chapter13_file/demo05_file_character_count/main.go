package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"unicode"
)

type CharCount struct {
	Character_Count int
	Number_Count    int
	Space_Count     int
	Other_Count     int
	Chinese_Count   int
}

func main() {
	file_path01 := "F:/GO/project/src/go_code/chapter13_file/demo05_file_character_count/develop.txt"
	file_path02 := "F:/GO/project/src/go_code/chapter13_file/demo05_file_character_count/out.txt"
	file01, err := os.Open(file_path01)
	if err != nil {
		fmt.Println("open file fail: ", err)
		return
	}
	defer file01.Close()
	reader := bufio.NewReader(file01)

	file02, err := os.OpenFile(file_path02, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("open file fail: ", err)
		return
	}
	defer file02.Close()
	writer := bufio.NewWriter(file02)

	var count CharCount

	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		str = string([]rune(str)) //处理中文
		for _, v := range str {
			if v >= 'A' && v <= 'z' {
				count.Character_Count++
			}
			if v == ' ' || v <= '\t' {
				count.Space_Count++
			}
			if v >= '0' && v <= '9' {
				count.Number_Count++
			}
			if unicode.Is(unicode.Han, v) {
				count.Chinese_Count++
			} else {

				count.Other_Count++
			}
		}
	}
	fmt.Println(count.Chinese_Count, count.Character_Count, count.Number_Count, count.Space_Count, count.Other_Count)
	str := "文件中文字符数为： " + strconv.FormatInt(int64(count.Chinese_Count), 10) + "\n文件英文字符数为： " + strconv.FormatInt(int64(count.Character_Count), 10) + "\n文件数字字符数为： " + strconv.FormatInt(int64(count.Number_Count), 10) +
		"\n文件空格字符数为： " + strconv.FormatInt(int64(count.Space_Count), 10) + "\n文件其他字符数为： " + strconv.FormatInt(int64(count.Other_Count), 10)
	_, err = writer.WriteString(str)
	if err != nil {
		fmt.Println("write file fail: ", err)
		return
	} else {
		writer.Flush()
		fmt.Println("count done!")
	}

}
