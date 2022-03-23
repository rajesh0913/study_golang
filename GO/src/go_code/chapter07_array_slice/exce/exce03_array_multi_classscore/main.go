package main

import "fmt"

func main() {
	/*
		定义二维数组存取三个班成绩。每班五名学生，求出班级平均分以及所有班平均分
	*/
	var scores [3][6]float64
	for i := 0; i < len(scores); i++ {
		for j := 0; j < len(scores[i])-1; j++ {
			fmt.Printf("请输入%v班%v学生成绩：", i+1, j+1)
			fmt.Scanln(&scores[i][j])
			scores[i][len(scores[i])-1] += scores[i][j]
		}
		scores[i][len(scores[i])-1] /= float64(len(scores[i]) - 1)
		fmt.Printf("%v班平均成绩：%.2f\n", i+1, scores[i][len(scores[i])-1])
	}
	fmt.Printf("三个班平均总成绩：%.2f\n", (scores[0][5]+scores[1][5]+scores[2][5])/3)
}
