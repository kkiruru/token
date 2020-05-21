package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

//var repo map[string][]string

func LoadTable(tableName string) (interface{}, error) {

	file, err := ioutil.ReadFile("./model/repository/data.json")
	if err != nil {
		fmt.Println(err)
		return "", errors.New("JSON 에러..?")
	}

	var dataMap map[string]interface{}
	err = json.Unmarshal(file, &dataMap)
	return dataMap[tableName], err
}

func LoadRecord(tableName string, attribute string) (string, error) {

	file, err := ioutil.ReadFile("./model/repository/data.json")
	if err != nil {
		fmt.Println(err)
		return "", errors.New("JSON 에러..?")
	}

	var dataMap map[string]interface{}
	err = json.Unmarshal(file, &dataMap) // JSON 문서의 내용을 변환하여 data에 저장

	return "", nil
}

//
//func account(m interface{}) (string, error) {
//	fmt.Println(m)
//	return "", nil
//}
