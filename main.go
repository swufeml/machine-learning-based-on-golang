package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 定义输入变量
var dataInX = [][]float64{
	[]float64{3.0, 3.0},
	[]float64{4.0, 3.0},
	[]float64{1.0, 1.0},
}
var dataInY []float64 = []float64{1.0, 1.0, -1.0}
var studyRate float64 = 0.01
var maxIteration int = 100 //设置最大迭代次数
var iteration int = 10     //连续多少次没有误分，就退出迭代

// 定义输出变量
var w = make([]float64, len(dataInX[0]), len(dataInX[0]))
var b float64

// 判断是否误分，返回bool，如果误分，返回true
func isMistake(testx []float64, testy float64) bool {
	var sumx float64 = 0.0
	for index, value := range testx {
		sumx += value * w[index]
	}
	sumx += b
	// fmt.Printf("%d",sumx)
	// fmt.Println(sumx * testy)
	if sumx*testy <= 0 {
		return true
	}
	return false
}

// 计算误分点的梯度，返回更新后的参数
func sgd(testx []float64, testy float64) ([]float64, float64) {
	sumx := 0.0
	for _, value := range testx {
		sumx = value * testy
	}
	for index, value := range w {
		w[index] = value + studyRate*sumx
	}
	b += studyRate * testy
	return w, b

}

// 随机选取x的方式，训练模型
func fitRandom(x [][]float64, y []float64) ([]float64, float64) {
	lenx := len(x)
	j := 0
	cycles:=0
	for i := 0; i < maxIteration; i++ {
		rand.Seed(time.Now().UnixNano())
		randx := rand.Intn(lenx)
		testx := x[randx]
		testy := y[randx]
		if isMistake(testx, testy) {
			j = 0
			w, b = sgd(testx, testy)
		} else {
			j++
		}
		if j == iteration {
			break
		}
		cycles++

	}
	fmt.Println(cycles)
	return w, b
}

// 判断是否误分
func fit(x [][]float64, y []float64) ([]float64, float64) {
	lenx := len(x)
	var j int
	for i := 0; i < maxIteration; i++ {
		j = 0
		for k := 0; k < lenx; k++ {
			testx := x[k]
			testy := y[k]
			anwser := isMistake(testx, testy)
			fmt.Println(anwser)
			if anwser {
				j = 0
				w, b = sgd(testx, testy)
			} else {
				j++
			}
			if j == lenx {
				goto label
			}

		}
	}
label:
	fmt.Println("结束")

	return w, b
}

func main() {
	// 给出输入变量，初始化权值与偏置
	w, b := fit(dataInX, dataInY)
	fmt.Println(w, b)
	w1, b1 := fitRandom(dataInX, dataInY)
	fmt.Println(w1, b1)
}
