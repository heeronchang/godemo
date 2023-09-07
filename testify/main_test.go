package main

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAssert(t *testing.T) {
	// Equal 断言
	assert.Equal(t, 4, Add(1, 3))

	sl1 := []int{1, 2, 3}
	sl2 := []int{1, 2, 3}
	sl3 := []int{2, 3, 4}
	assert.Equal(t, sl1, sl2, "sl1 should equal to sl2 ")

	p1 := &sl1
	p2 := &sl2
	assert.Equal(t, p1, p2, "the content which p1 point to should equal to which p2 point to")

	err := errors.New("demo error")
	assert.EqualError(t, err, "demo error")

	// assert.Equal(t, 1, 2) // FAIL

	// assert.Exactly(t, int32(123), int64(123)) // failed! Types expected to match exactly int32 != int64

	// 布尔断言
	assert.True(t, 1+1 == 2, "1+1 == 2 should be true")
	assert.Contains(t, "Hello World", "World")
	assert.Contains(t, []string{"Hello", "World"}, "World")
	assert.Contains(t, map[string]string{"Hello": "World"}, "Hello")
	assert.ElementsMatch(t, []int{1, 3, 2, 3}, []int{1, 3, 3, 2})

	// 反向断言
	assert.NotEqual(t, 4, Add(2, 3), "The result should not be 4")
	assert.NotEqual(t, sl1, sl3, "sl1 should not equal to sl3 ")
	assert.False(t, 1+1 == 3, "1+1 == 3 should be false")
	assert.Never(t, func() bool { return false }, time.Second, 10*time.Millisecond) //1秒之内condition参数都不为true，每10毫秒检查一次
	assert.NotContains(t, "Hello World", "Go")

	// and so on

	// !注：assert.Equal底层实现使用的是reflect.DeepEqual。

	// 简化每次传入 t(*testing.T)
	assert := assert.New(t)
	assert.Equal(1, 2)

	// require包可以理解为assert包的“姊妹包”，require包实现了assert包提供的所有导出的断言函数
	require.Equal(t, 1, 2)
	require := require.New(t)
	require.Equal(1, 1, []string{"1", "1", "err"})

	// !require包可以理解为assert包的“姊妹包”，require包实现了assert包提供的所有导出的断言函数
}
