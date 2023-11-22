package model

type CartItem struct {
	Preset
	UserID     int `json:"user_id"`
	GoodsId    int `json:"goods_id"`
	GoodsCount int `json:"goods_count"`
}
