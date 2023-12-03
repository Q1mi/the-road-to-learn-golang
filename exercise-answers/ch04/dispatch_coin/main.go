package main

/*
分金币。现在共有 50 枚硬币，你需要按如下规则分配给以下几个小朋友，
他们的名字分别是：Matthew、Sarah、Augustus、Heidi、Emilie、Peter、Giana、Adriano、Aaron、Elizabeth。
分配规则如下：

1. 名字中每包含1个'e'或'E'分1枚硬币
2. 名字中每包含1个'i'或'I'分2枚硬币
3. 名字中每包含1个'o'或'O'分3枚硬币
4. 名字中每包含1个'u'或'U'分4枚硬币

编写一个程序，计算每个孩子分到了多少硬币以及最后剩余多少硬币。程序结构如下，请实现 dispatchCoin 函数完成上述需求。

*/
import (
	"fmt"
	"strings"
)

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter",
		"Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

// dispatchCoin 分硬币的函数
func dispatchCoin() int {
	for _, name := range users {
		num := 0
		for _, n := range name {
			switch n {
			case 'e', 'E':
				num += 1
			case 'i', 'I':
				num += 2
			case 'o', 'O':
				num += 3
			case 'u', 'U':
				num += 4
			}
		}
		distribution[name] = num
		coins -= num
	}
	fmt.Println(distribution)
	return coins
}

func dispatchCoin2() int {
	left := coins
	for _, name := range users {
		// strings.ToUpper(name) 转为全大写
		// strings.Count(s, c) 返回字符串s中包含c的数量
		e := strings.Count(strings.ToUpper(name), "E")
		i := strings.Count(strings.ToUpper(name), "I")
		o := strings.Count(strings.ToUpper(name), "O")
		u := strings.Count(strings.ToUpper(name), "U")
		sum := e*1 + i*2 + o*3 + u*4
		distribution[name] = sum
		left -= sum
	}
	fmt.Println(distribution)
	return left
}

func main() {
	//left := dispatchCoin()
	//fmt.Println("剩下：", left)
	left2 := dispatchCoin2()
	fmt.Println("剩下：", left2)
}
