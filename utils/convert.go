package utils

import (
	"allmusic-api-server/consts"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

func ConvertChannelId(channelId string) (channel string) {
	channel = "o"
	switch channelId {
	case consts.ChannelPublic:
		channel = "p"
	case consts.ChannelOPPO:
		channel = "o"
	case consts.ChannelVIVO:
		channel = "v"
	case consts.ChannelXiaoMi:
		channel = "x"
	}
	return channel
}

const (
	fullLayout   = "2006-01-02 15:04:05"
	simpleLayout = "2006-01-02"
)

/**
支持的Time格式 2006-01-02T15:04:05.000Z07:00
*/
func TimeStrToTimeStamp(strTime string) (intTime int64) {
	//时间字符串转时间戳
	t, _ := time.Parse(consts.DateTimeFormat, "2006-01-02T15:04:05.000+00:00")
	datetime_str_to_timestamp := t.Unix()
	return datetime_str_to_timestamp
}

func TimeStamp2TimeStrFull(timeStamp int64) (strTime string) {
	return time.Unix(timeStamp, 0).Format(fullLayout)
}

func TimeStamp2TimeStrSimple(timeStamp int64) (strTime string) {
	return time.Unix(timeStamp, 0).Format(simpleLayout)
}

func DateTOStr(currentTime time.Time) string {
	const dateFmt = "2006-01-02"
	dateStr := currentTime.Format(dateFmt)
	return dateStr
}

func RemoveDuplicate(in []int64) []int64 {
	resultMap := map[int64]bool{}
	for _, v := range in {
		resultMap[v] = true
	}
	var result []int64
	for k := range resultMap {
		result = append(result, k)
	}
	return result
}

func RemoveDuplicateString(in []string) []string {
	resultMap := map[string]bool{}
	for _, v := range in {
		resultMap[v] = true
	}
	var result []string
	for k := range resultMap {
		result = append(result, k)
	}
	return result
}

func TransToStr(count int64, lng string) (result string) {
	if strings.Contains(lng, "zh") {
		// 中文
		if count <= 10000 {
			return strconv.Itoa(int(count))
		}
		if count > 10000 && count < 100000000 {
			return fmt.Sprintf("%s 万", FormatFloat(float64(count)/10000))
		}
		if count >= 100000000 {
			return fmt.Sprintf("%s 亿", FormatFloat(float64(count)/(10000*10000)))
		}
	} else {
		// 英文
		if count <= 1000 {
			return strconv.Itoa(int(count))
		}
		if count > 1000 && count < 100*10000 {
			return fmt.Sprintf("%s K", FormatFloat(float64(count)/(1000)))
		}
		if count >= 100*10000*10000 && count < 1000*100*10000*10000 {
			return fmt.Sprintf("%s M", FormatFloat(float64(count)/(100*10000*10000)))
		}
		if count >= 1000*100*10000 {
			return fmt.Sprintf("%s B", FormatFloat(float64(count)/(1000*100*10000)))
		}
	}

	return
}

// 去掉小数点无效位
func FormatFloat(num float64) string {
	num += 0.05
	decimal := 1
	d := float64(1)
	if decimal > 0 {
		d = math.Pow10(decimal)
	}
	return strconv.FormatFloat(math.Trunc(num*d)/d, 'f', -1, 64)
}
