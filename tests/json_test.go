package tests

import (
	"GinMall/model"
	"GinMall/serializer/response"
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonMarshal(t *testing.T) {
	user := &model.User{
		Preset:   model.Preset{ID: 1},
		Username: "XXXXX",
		Password: "XXXXX",
	}
	resp := response.SuccessData(user)
	out, _ := json.Marshal(resp)
	fmt.Println(string(out))
}
