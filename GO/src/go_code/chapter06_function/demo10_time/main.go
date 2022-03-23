package main

import (
	"fmt"
	"time"
)

func main() {
	//获取当前时间
	now := time.Now()
	fmt.Printf("time.Now() value = %v,type = %T\n", now, now)
	//通过now变量获取当前年月日时分秒
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	//格式化1
	fmt.Printf("year: %v\nmonth: %v\nday: %v\nhour: %v\nminute: %v\nsecond: %v\n", year, int(month), day, hour, minute, second)
	fmt.Printf("当前年月日： %d-%d-%d %d:%d:%d\n", year, month, day, hour, minute, second)
	dateStr := fmt.Sprintf("当前年月日： %d-%d-%d %d:%d:%d\n", year, month, day, hour, minute, second)
	fmt.Println("dateStr = ", dateStr)
	//格式化2
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("15:04:05"))
	fmt.Println(now.Format("15"))

	//时间常量 type:duration
	k := 0
	for {
		k++
		fmt.Println(k)
		//time.Sleep(time.Second)
		time.Sleep(time.Microsecond * 100)

		if k == 100 {
			break
		}
	}

	//时间戳 Unix UnixNano
	fmt.Printf("Unix时间戳 value = %v,UnixNano时间戳 = %v\n", now.Unix(), now.UnixNano())

}
