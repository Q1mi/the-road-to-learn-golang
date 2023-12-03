package split

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) { // 函数名必须以Test开头并接收一个*testing.T类型参数
	got := Split("a:b:c", ":")         // 程序输出的结果
	want := []string{"a", "b", "c"}    // 期望的结果
	if !reflect.DeepEqual(want, got) { // slice不能比较直接，需要借助反射包中的方法
		t.Errorf("expected:%v, got:%v", want, got) // 测试失败输出错误提示
	}
}

