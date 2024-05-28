package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/teris-io/shortid"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
	"xuanfeng_gin/pkg/log"
)

// ArrIsContainValueInt
/**
 *  @Description: 数组是否包含 int
 *  @param arr
 *  @param val
 *  @return res
 */
func ArrIsContainValueInt(arr []int, val int) (res bool) {
	res = false
	for _, v := range arr {
		if val == v {
			res = true
			return
		}
	}
	return
}

func TodayLastSecond() int64 {
	year, month, day := time.Now().Date()
	location, _ := time.LoadLocation("Asia/Shanghai") // 这一步把错误忽略了，时区用Shanghai是因为没有Beijing
	return time.Date(year, month, day, 23, 59, 59, 0, location).Unix() - time.Now().Unix()
}

func DecimalStringToFloat(str string) (f float64) {
	d, _ := decimal.NewFromString(str)
	f, _ = d.Float64()
	return
}

func GetUserMineCode() string {
	var code string
	var e error
	for i := 0; i < 3; i++ {
		code, e = shortid.Generate()
		if e != nil {
			log.L.Error(e.Error())
			continue
		}
		if strings.Contains(code, "-") || strings.Contains(code, "_") {
			continue
		} else {
			return code
		}
	}
	return code
}

// UniqueArr
/**
 *  @Description: 数组清除重复
 *  @param arr
 *  @return arrRes
 */
func UniqueArr(arr []string) (arrRes []string) {
	//Create a   dictionary of values for each element
	dict := make(map[string]int)
	for _, v := range arr {
		dict[v] += 1
	}
	for k, _ := range dict {
		arrRes = append(arrRes, k)
	}
	return
}

func HttpPost(reqData map[string]string, httpUrl string) (res string, err error) {
	// 发起请求

	paramJSONByte, err := json.Marshal(reqData)
	if err != nil {
		return
	}

	reader := bytes.NewReader(paramJSONByte)
	log.L.Info("HttpPost curl " + httpUrl)
	request, err := http.NewRequest("POST", httpUrl, reader)
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.L.Error(err.Error())
		return
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.L.Error(err.Error())
		return
	}
	res = string(respBytes)
	log.L.Info(fmt.Sprintf("HttpPost url:%s,res:%s,req:%v", httpUrl, res, reqData))
	return
}

// ArrIsContainValue
/**
 *  @Description: 数组是否包含
 *  @param arr
 *  @param val
 *  @return res
 */
func ArrIsContainValue(arr []string, val string) (res bool) {
	res = false
	for _, v := range arr {
		if val == v {
			res = true
			return
		}
	}
	return
}

func Struct2Interface(s any) (res []interface{}) {
	d := reflect.ValueOf(s)
	t := reflect.TypeOf(s)
	count := d.NumField()
	for i := 0; i < count; i++ {
		f := d.Field(i)
		f1 := t.Field(i)
		switch f.Kind() {
		case reflect.String:
			res = append(res, f1.Name, f.String())
		case reflect.Int, reflect.Int64:
			res = append(res, f1.Name, f.Int())
		case reflect.Bool:
			res = append(res, f1.Name, f.Bool())
		case reflect.Float32, reflect.Float64:
			res = append(res, f1.Name, f.Float())
		default:
			//res = append(res, f1.Name, f.String())
		}
	}
	return
}

func Struct2RedisHashInterface(s any) (res []interface{}) {
	d := reflect.ValueOf(s)
	t := reflect.TypeOf(s)
	count := d.NumField()
	for i := 0; i < count; i++ {
		f := d.Field(i)
		f1 := t.Field(i)
		name := f1.Name
		tagStr := f1.Tag.Get("json")
		if tagStr != "" {
			name = strings.Split(tagStr, ",")[0]
		}
		switch f.Kind() {
		case reflect.String:
			res = append(res, name, f.String())
		case reflect.Int, reflect.Int64:
			res = append(res, name, f.Int())
		case reflect.Bool:
			res = append(res, name, f.Bool())
		case reflect.Float32, reflect.Float64:
			res = append(res, name, f.Float())
		default:
			//res = append(res, f1.Name, f.String())
		}
	}
	return
}

func GetTimeYMD(t time.Time) (year, month, day int) {
	tStr := t.Format("2006-01-02")
	timeArr := strings.Split(tStr, "-")
	year, _ = strconv.Atoi(timeArr[0])
	month, _ = strconv.Atoi(timeArr[1])
	day, _ = strconv.Atoi(timeArr[2])
	return
}

func GetArrCommon(arr1, arr2 []string) (res []string) {
	if len(arr1) == 0 || len(arr2) == 0 {
		return
	}
	arr1Map := make(map[string]int)
	for _, v := range arr1 {
		arr1Map[v] = 1
	}
	for _, v := range arr2 {
		if _, ok := arr1Map[v]; ok {
			res = append(res, v)
		}
	}
	return
}

func CalcAgeByBirthday(birthday string) (age int) {
	if birthday == "" {
		return
	}
	arr := strings.Split(birthday, "-")
	birYear, _ := strconv.Atoi(arr[0])
	age = time.Now().Year() - birYear
	return
}

func DealPhoneB3A4(phone string) string {
	phoneStr := phone
	if len(phoneStr) > 8 {
		phoneStr = phoneStr[:3] + "*****" + phoneStr[8:]
	} else if len(phoneStr) > 3 {
		phoneStr = phoneStr[:3] + "*****"
	} else {
		phoneStr = ""
	}
	return phoneStr
}

// UniqueIntArr
/**
 *  @Description: 数组清除重复
 *  @param arr
 *  @return arrRes
 */
func UniqueIntArr(arr []int) (arrRes []int) {
	//Create a   dictionary of values for each element
	dict := make(map[int]int)
	for _, v := range arr {
		dict[v] += 1
	}
	for k, _ := range dict {
		arrRes = append(arrRes, k)
	}
	return
}

func DealName(name string) (result string) {
	nameRune := []rune(name)
	lens := len(nameRune)
	if lens <= 1 {
		result = "***"
	} else if lens == 2 {
		result = string(nameRune[:1]) + "*"
	} else if lens == 3 {
		result = string(nameRune[:1]) + "*" + string(nameRune[2:3])
	} else if lens == 4 {
		result = string(nameRune[:1]) + "**" + string(nameRune[lens-1:lens])
	} else {
		result = string(nameRune[:2]) + "***" + string(nameRune[lens-2:lens])
	}
	return
}
func DealIdCard(idCard string) string {
	if len(idCard) == 0 {
		return idCard
	}
	return idCard[:3] + "******" + idCard[len(idCard)-4:]
}
