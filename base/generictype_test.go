package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

// 匿名结构体不支持泛型

// 一个泛型类型的结构体。可用 int 或 sring 类型实例化
type MyStruct[T int | string] struct {
	Name string
	Data T
}

// 一个泛型接口
type IPrintData[T int | float32 | string] interface {
	Print(data T)
}

// 类型形参是可以互相套用的，如下
type WowStruct[T int | float32, S []T] struct {
	Data     S
	MaxValue T
	MinValue T
}

// 一个泛型通道，可用类型实参 int 或 string 实例化
type MyChan[T int | string] chan T

// 给泛型类型添加方法
type MySlice[T int | float32] []T

func (s MySlice[T]) Sum() T {
	var sum T
	for _, value := range s {
		sum += value
	}
	return sum
}

type Slice[T int | float32 | float64] []T

func (s Slice[T]) Len() int {
	return len(s)
}

func (s Slice[T]) Sort() {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

func FunctionGenericTypeToAdd[T int | float64](a, b T) T {
	return a + b
}

func TestSlice(t *testing.T) {

	if ret := FunctionGenericTypeToAdd(1, 2); ret != 3 {
		t.Errorf("FunctionGenericTypeToAdd should be 3, got %v", ret)
	}
	if ret := FunctionGenericTypeToAdd(1.3, 2.8); ret != 4.1 {
		t.Errorf("FunctionGenericTypeToAdd should be 4.1, got %v", ret)
	}
	s := Slice[int]{3, 2, 1, 5, 4}
	expected := Slice[int]{1, 2, 3, 4, 5}
	s.Sort()
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected %v, got %v", expected, s)
	}

	s2 := Slice[float32]{3.14, 2.71, 1.618, 1.414, 1.732}
	expected2 := Slice[float32]{1.414, 1.618, 1.732, 2.71, 3.14}
	s2.Sort()
	if !reflect.DeepEqual(s2, expected2) {
		t.Errorf("Expected %v, got %v", expected2, s2)
	}

	// 这里传入了类型实参int，泛型类型Slice[T]被实例化为具体的类型 Slice[int]
	var a Slice[int] = []int{1, 2, 3}
	fmt.Printf("Type Name: %T", a) //输出：Type Name: Slice[int]

	// 传入类型实参float32, 将泛型类型Slice[T]实例化为具体的类型 Slice[string]
	var b Slice[float32] = []float32{1.0, 2.0, 3.0}
	fmt.Printf("Type Name: %T", b) //输出：Type Name: Slice[float32]

	// ✗ 错误。因为变量a的类型为Slice[int]，b的类型为Slice[float32]，两者类型不同
	//a = b

	// ✗ 错误。string不在类型约束 int|float32|float64 中，不能用来实例化泛型类型
	//var c Slice[string] = []string{"Hello", "World"}

	// ✗ 错误。Slice[T]是泛型类型，不可直接使用必须实例化为具体的类型
	//var x Slice[T] = []int{1, 2, 3}

}
