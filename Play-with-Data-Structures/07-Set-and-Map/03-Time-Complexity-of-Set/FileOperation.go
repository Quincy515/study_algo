package _3_Time_Complexity_of_Set

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// 文件相关操作-读取文件名称为filename中的内容，并将其中包含的所有词语放进words中
func ReadFile(filename string) []string {
	var words []string

	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 文件读取
	reader := bufio.NewReader(file)
	for {
		// 简单分词，没有考虑很多文本处理中的特殊问题，这里只做demo使用
		line, err := reader.ReadString('\n')

		if err != nil || io.EOF == err {
			break
		}

		wordSlice := strings.Fields(line)
		for _, word := range wordSlice {
			if word = extractStr(strings.ToLower(word)); word != "" {
				words = append(words, word)
			}
		}
	}
	return words
}

func extractStr(str string) string {
	var res []rune
	for _, letter := range str {
		if letter >= 'a' && letter <= 'z' {
			res = append(res, letter)
		}
	}
	return string(res)
}
