package gock_demo

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
	"testing"
)

func TestGetResultByAPI(t *testing.T) {
	tests := []struct {
		name     string
		x        int
		y        int
		expected int
	}{
		{"case1", 1, 1, 101},
		{"case2", 2, 1, 201},
	}

	defer gock.Off() // 测试执行后刷新挂起的mock

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// mock 请求外部api时传参x=n返回n*100
			gock.New("http://your-api.com").
				Post("/post").
				MatchType("json").
				JSON(map[string]int{"x": test.x}).
				Reply(200).
				JSON(map[string]int{"value": test.x * 100})

			// 调用GetResultByAPI
			result := GetResultByAPI(test.x, test.y)
			// 校验结果是否符合预期
			assert.Equal(t, result, test.expected)
		})
	}

	assert.True(t, gock.IsDone()) // 断言mock被触发
}
