// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	// 打开文件
// 	file, err := os.Open("example.txt")
// 	if err != nil {
// 		fmt.Println("文件打开失败：", err)
// 		return
// 	}
// 	defer file.Close()

// 	// 创建一个带缓冲的读取器
// 	reader := bufio.NewReader(file)

// 	// 初始化字符计数器
// 	charCount := 0

// 	// 逐个读取字符
// 	for {
// 		_, _, err := reader.ReadRune()
// 		if err != nil {
// 			// 如果已经到达文件末尾，就退出循环
// 			if err.Error() == "EOF" {
// 				break
// 			}
// 			// 否则打印错误信息并退出程序
// 			fmt.Println("文件读取失败：", err)
// 			return
// 		}
// 		// 如果读取成功，就将字符计数器加一
// 		charCount++
// 	}

// 	// 输出结果
// 	fmt.Println("字符数量：", charCount)
// }
package main

func reverse(b *[]byte, left int, right int) {
	for left < right {
		(*b)[left], (*b)[right] = (*b)[right], (*b)[left]
		left++
		right--
	}
}

func reverseWords(s string) string {
	slowIndex, fastIndex := 0, 0
	b := []byte(s)
	//头部冗余
	for len(b) > 0 && fastIndex < len(b) && b[fastIndex] == ' ' {
		fastIndex++
	}
	//单词间冗余
	for ; fastIndex < len(b); fastIndex++ {
		if fastIndex-1 > 0 && b[fastIndex-1] == b[fastIndex] && b[fastIndex] == ' ' {
			continue
		}
		b[slowIndex] = b[fastIndex]
		slowIndex++
	}
	//尾部冗余
	if slowIndex-1 > 0 && b[slowIndex-1] == ' ' {
		b = b[:slowIndex-1]
	} else {
		b = b[:slowIndex]
	}
	//reverse
	reverse(&b, 0, len(b)-1)
	println(string(b))

	//reverse single word
	slow, fast := 0, 0
	for fast < len(b) {
		if b[fast] == ' ' {
			sw := fast - 1
			for slow < sw {
				b[slow], b[sw] = b[sw], b[slow]
				slow++
				sw--
			}
			slow = fast + 1
		}
		fast++
	}
	sw := fast - 1
	for slow < sw {
		b[slow], b[sw] = b[sw], b[slow]
		slow++
		sw--
	}

	return string(b)
}
func main() {
	s := "the sky is blue"
	println(reverseWords(s))
}
