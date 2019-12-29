package func5_test

import (
	"strings"
	"testing"
	"fmt"
)

/**
测试，数据 跟  操作的分离
 */

func stringProccess(list []string, chain []func(string) string) {
	for index, str := range list  {
		result := str

		for _, proc := range chain  {
			result = proc(result)
		}

		list[index] = result
	}
}

func removePrefix(str string) string  {
	return strings.TrimPrefix(str, "go")
}

func TestOptData(t *testing.T)  {
	// 待处理的字符串列表
	list := []string{
		"go scanner",
		"go parser",
		"go compiler",
		"go printer",
		"go formater",
	}

	chain := []func(string) string{
		removePrefix,
		strings.TrimSpace,
		strings.ToUpper,
	}

	stringProccess(list, chain)

	for _, str := range list  {
		fmt.Println(str)
	}
}
