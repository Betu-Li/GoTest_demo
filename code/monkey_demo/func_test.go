package monkey_demo

import (
	"github.com/agiledragon/gomonkey/v2"
	"strings"
	"testing"
)

// func_test.go

func TestMyFunc(t *testing.T) {
	// 对 varys.GetInfoByUID 进行打桩
	// 无论传入的uid是多少，都返回 &varys.UserInfo{Name: "monkey_name"}, nil

	// 用monkey打桩
	//monkey.Patch(varys.GetInfoByUID, func(int64) (*varys.UserInfo, error) {
	//	return &varys.UserInfo{Name: "monkey_name"}, nil
	//})

	// 用gomonkey打桩
	patch := gomonkey.ApplyFunc(varys.GetInfoByUID, func(int64) (*varys.UserInfo, error) {
		return &varys.UserInfo{Name: "monkey_name"}, nil
	})

	defer patch.Reset()

	// 执行测试
	ret := MyFunc(123)
	if !strings.Contains(ret, "monkey_name") {
		t.Fatal()
	}
}
