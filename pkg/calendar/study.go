package calendar

// 紀錄當天學習成果，以網頁為單位
type Record struct {
	index       int
	url         string
	description string

	tags uint8
}
