package entities

type ModelDataItem struct {
	ProductID int32
	Rank      float32
}

type ModelDataMaster struct {
	Key         string
	ProductRank []*ModelDataItem
}
