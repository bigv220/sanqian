package service

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestQigua(t *testing.T) {
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())

	var results []string
	var coinResult string
	guaMap := map[string]string{
		"老阴": "1",
		"老阳": "0",
		"少阴": "0",
		"少阳": "1",
	}

	// 模拟起卦过程6次
	for i := 0; i < 6; i++ {
		coin1 := rand.Intn(2) // 0表示反面，1表示正面
		coin2 := rand.Intn(2)
		coin3 := rand.Intn(2)

		sum := coin1 + coin2 + coin3

		// 根据组合结果输出对应的起卦名称
		switch sum {
		case 0:
			results = append([]string{"老阴"}, results...)
		case 1:
			results = append([]string{"少阴"}, results...)
		case 2:
			results = append([]string{"少阳"}, results...)
		case 3:
			results = append([]string{"老阳"}, results...)
		}
		coinResult += fmt.Sprintf("%d%d%d,", coin1, coin2, coin3)
		// 输出每个硬币的正反面情况
		fmt.Printf("第%d次投掷：硬币1：%d, 硬币2：%d, 硬币3：%d，起卦结果：%s\n", i+1, coin1, coin2, coin3, results[0])
	}

	// 输出六次的起卦结果
	fmt.Println("六次的起卦结果（从下到上）：", results)
	guaKey := ""
	for _, result := range results {
		guaKey += guaMap[result]
	}
	fmt.Println(guaKey)
	fmt.Println(coinResult)
}
