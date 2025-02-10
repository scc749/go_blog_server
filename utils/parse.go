package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// ParseDuration 解析持续时间字符串为 time.Duration。
// 持续时间字符串应由数字值和时间单位组成，单位可以是 "d" 表示天，"h" 表示小时，"m" 表示分钟，"s" 表示秒。
// 例如，"1d2h30m" 会被解析为 1 天、2 小时和 30 分钟。
// 如果字符串为空或格式无效，则返回错误。
func ParseDuration(d string) (time.Duration, error) {
	d = strings.TrimSpace(d) // 去除字符串两端的空格
	if len(d) == 0 {
		return 0, fmt.Errorf("empty duration string")
	}

	// 定义每个单位及其对应的持续时间值
	unitPattern := map[string]time.Duration{
		"d": time.Hour * 24, // "d" 对应 24 小时
		"h": time.Hour,      // "h" 对应 1 小时
		"m": time.Minute,    // "m" 对应 1 分钟
		"s": time.Second,    // "s" 对应 1 秒
	}

	var totalDuration time.Duration // 总持续时间
	// 遍历 "d"、"h"、"m"、"s" 单位
	for _, unit := range []string{"d", "h", "m", "s"} {
		// 提取所有以当前单位结尾的部分
		for strings.Contains(d, unit) {
			// 找到单位的位置
			unitIndex := strings.Index(d, unit)
			// 提取单位前面的部分
			part := d[:unitIndex]
			if part == "" {
				part = "0" // 如果部分为空，默认为 0
			}
			// 将部分转换为整数值
			val, err := strconv.Atoi(part)
			if err != nil {
				return 0, fmt.Errorf("invalid duration part: %v", err)
			}
			// 将该部分的持续时间累加到总持续时间
			totalDuration += time.Duration(val) * unitPattern[unit]
			// 从字符串中移除已处理的部分
			d = d[unitIndex+len(unit):]
		}
	}

	// 检查是否有剩余未处理的字符
	if len(d) > 0 {
		return 0, fmt.Errorf("unrecognized duration format")
	}

	// 返回总的持续时间
	return totalDuration, nil
}
