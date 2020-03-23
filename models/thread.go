package models

type Thread struct {
	Id    int
	Title string
	URL   string
}

func (Thread) TableName() string {
	return "forocoches_top_threads"
}
