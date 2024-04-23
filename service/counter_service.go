package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

	for _, d := range parsedData {
		log.Println(d)
	}
}

type Data struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
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
			Name:            item[1],
			Alias:           item[2],
			Meaning:         item[3],
			Interpretation1: item[4],
			Interpretation2: item[5],
		}
		result = append(result, d)
	}
	return result, nil
}

// getIndex 获取主页
func getIndex() (string, error) {
	b, err := ioutil.ReadFile("./qigua.txt")
	if err != nil {
		return "", err
	}
	return string(b), nil
}
