package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
	"wxcloudrun-golang/db/dao"
)

// JsonResult 返回结构
type JsonResult struct {
	Code     int         `json:"code"`
	ErrorMsg string      `json:"errorMsg,omitempty"`
	Data     interface{} `json:"data"`
}

// IndexHandler 计数器接口
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	data, err := getIndex()
	if err != nil {
		fmt.Fprint(w, "内部错误")
		return
	}
	fmt.Fprint(w, data)
}

// QiguaHandler 起卦接口
func QiguaHandler(w http.ResponseWriter, r *http.Request) {
	content, _ := dao.Imp.GetQiguaData()

	parsedData, err := parseJSONData(content)
	if err != nil {
		log.Fatal("Error parsing JSON: ", err)
	}
	//三钱起卦过程
	sanQianGua, bianGua, coinResult := QiGua()
	//返回结果
	for _, d := range parsedData {
		if d.GuaKey == sanQianGua {
			resp := JsonResult{
				Code: 0,
				Data: map[string]interface{}{
					"id":              d.ID,
					"gua_key":         d.GuaKey,
					"shang_gua":       d.ShangGua,
					"xia_gua":         d.XiaGua,
					"alias":           d.Alias,
					"meaning":         d.Meaning,
					"interpretation1": d.Interpretation1,
					"interpretation2": d.Interpretation2,
					"bianGua":         bianGua,
					"coin_result":     coinResult,
				},
			}
			jsonResp, err := json.Marshal(resp)
			if err != nil {
				log.Fatalf("Error happened in JSON marshal. Err: %s", err)
			}
			w.Write(jsonResp)
			return
		}
	}

	return
}

type Data struct {
	ID              string `json:"id"`
	GuaKey          string `json:"gua_key"`
	ShangGua        string `json:"shang_gua"`
	XiaGua          string `json:"xia_gua"`
	Alias           string `json:"alias"`
	Meaning         string `json:"meaning"`
	Interpretation1 string `json:"interpretation1"`
	Interpretation2 string `json:"interpretation2"`
}

func parseJSONData(jsonData string) ([]Data, error) {
	var data [][]string
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		return nil, err
	}

	var result []Data
	for _, item := range data {
		d := Data{
			ID:              item[0],
			GuaKey:          item[1],
			ShangGua:        item[2],
			XiaGua:          item[3],
			Alias:           item[4],
			Meaning:         item[5],
			Interpretation1: item[6],
			Interpretation2: item[7],
		}
		result = append(result, d)
	}
	return result, nil
}

func QiGua() (string, bool, string) {
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
		coinResult += fmt.Sprintf("%d%d%d", coin1, coin2, coin3)
		// 输出每个硬币的正反面情况
		//fmt.Printf("第%d次投掷：硬币1：%d, 硬币2：%d, 硬币3：%d，起卦结果：%s\n", i+1, coin1, coin2, coin3, results[0])
	}

	// 输出六次的起卦结果
	guaKey := ""
	bianGua := false
	for _, result := range results {
		if result == "老阴" || result == "老阳" {
			bianGua = true
		}
		guaKey += guaMap[result]
	}
	return guaKey, bianGua, coinResult
}

// getIndex 获取主页
func getIndex() (string, error) {
	b, err := ioutil.ReadFile("./qigua.txt")
	if err != nil {
		return "", err
	}
	return string(b), nil
}
