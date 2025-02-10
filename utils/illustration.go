package utils

import "regexp"

func FindIllustrations(text string) ([]string, error) {
	// 定义正则表达式，匹配 Markdown 图片语法
	// 格式: ![alt text](image_url)
	regex := `!\[([^\]]*)\]\(([^)]+)\)`

	// 编译正则表达式
	re, err := regexp.Compile(regex)
	if err != nil {
		return nil, err
	}

	// 查找所有符合正则表达式的匹配项
	matches := re.FindAllStringSubmatch(text, -1)

	// 存储匹配到的所有图片链接
	var illustrations []string

	// 提取每个匹配项中的图片链接部分（即第二个捕获组）
	for _, match := range matches {
		if len(match) > 2 {
			illustrations = append(illustrations, match[2])
		}
	}

	return illustrations, nil
}
