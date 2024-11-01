package reflect_test

import (
	"regexp"
	"testing"
)

func TestBase(t *testing.T) {
	re := regexp.MustCompile(`\s+`)
	fff := re.ReplaceAllString("你好啊啊啊我的 我的啊我的外和发和iedfhiwad[](dawdad)", " ")
	t.Log(fff)
}

//反射获取
