package model

import "task-api/internal/constant"

type RequestItem struct {
	Title    string  `binding:"required" json:"title"`
	Amount   float64 `binding:"gte=5" json:"amount"`
	Quantity uint    `binding:"gte=1" json:"quantity"`
}

type RequestFindItem struct {
	Statuses constant.ItemStatus `form:"status"`
}

type RequestUpdateItem struct {
	Status constant.ItemStatus
}

type RequestLogin struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
}
