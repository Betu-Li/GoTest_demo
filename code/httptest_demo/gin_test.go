package httptest_demo

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_HelloHandler(t *testing.T) {
	// 定义两个测试用例
	tests := []struct {
		name     string
		param    string
		expected string
	}{
		{"base case", `{"name": "go!"}`, "hello go!"},
		{"bad case", "", "we need a name"},
	}

	r := SetupRouter()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// mock一个测试请求
			req := httptest.NewRequest(
				"POST",                        // 请求方法
				"/hello",                      // 请求URL
				strings.NewReader(test.param), // 请求参数
			)

			// mock一个响应记录器
			w := httptest.NewRecorder()

			// 让server端处理mock请求并记录返回的响应内容
			r.ServeHTTP(w, req)

			// 校验状态码是否符合预期
			assert.Equal(t, http.StatusOK, w.Code)

			//  解析并断言返回的响应内容是否符合预期
			var resp map[string]string
			err := json.Unmarshal(w.Body.Bytes(), &resp)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, resp["msg"])
		})
	}
}
