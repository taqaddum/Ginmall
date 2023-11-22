package tests

import (
	"GinMall/model"
	"GinMall/utils"
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	var banner = model.Banner{
		Url: "./const",
	}
	res, _ := utils.CopyStruct(&banner).(*model.Banner)
	fmt.Println(res.Url)
}
