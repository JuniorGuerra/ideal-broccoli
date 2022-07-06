package pipes

type NamesOrders struct {
	NameList string `json:"nameList" binding:"required"`
}

type StringsFriends struct {
	X string `json:"x" binding:"required"`
	Y string `json:"y" binding:"required"`
}
