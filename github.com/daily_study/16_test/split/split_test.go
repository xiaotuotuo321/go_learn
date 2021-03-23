package split

import (
	"reflect"
	"testing"
	"time"
)

func TestSplit(t *testing.T) {	// 测试函数名必须以Test开头，必须接收一个*testing.T类型参数
	got := Split("a:b:c", ":")	//	程序输出的结果

	want := []string{"a", "b", "c"}		// 期望的结果

	if !reflect.DeepEqual(want, got){	// 因为slice不能直接比较，借助反射包中的方法比较
		t.Errorf("excepted: %v, %T; got:%v, %T", want, want, got, got)	// 测试失败输出错误提示
	}
}

func TestMoreSplit(t *testing.T) {
	got := Split("abcd", "bc")
	want := []string{"a", "d"}
	if !reflect.DeepEqual(want, got){
		t.Errorf("excepted:%v, got: %v", want, got)
	}
}

func TestSplit2	(t *testing.T){
	// 定义一个测试用例类型
	type test struct{
		input string
		sep string
		want []string
	}

	// 定义一个存储测试用例的切片
	tests := []test{
		{input:"a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		{input:"a:b:c", sep: ",", want: []string{"a:b:c"}},
		{input:"abcd", sep: "bc", want: []string{"a", "d"}},
		{input:"沙河有沙又有河", sep: "沙", want: []string{"", "河有", "又有河"}},
	}
	// 遍历切片，逐一执行测试用例
	for name, tc := range tests {
		got := Split(tc.input, tc.sep)
		if !reflect.DeepEqual(got, tc.want){
			t.Errorf("name: %v, excepted: %#v, got: %#v", name, tc.want, got)
		}
	}
}

// 基准测试
func BenchmarkSplit(b *testing.B) {
	time.Sleep(5 * time.Second)		// 假设需要做一些耗时的无关操作
	b.ResetTimer()		// 重置计时器
	for i := 0; i < b.N; i++{
		Split("沙河有沙又有河", "沙")
	}
}

// 并行测试
func BenchmarkSplitParallel(b *testing.B){
	// b.SetParallelism(1)	// 设置使用的CPU数
	b.RunParallel(func (pb *testing.PB) {
		for pb.Next() {
			Split("沙河有沙又有河", "沙")
		}
	})
}
