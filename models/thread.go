package models

type Thread struct {
	Id       int
	Title    string
	Category string
	URL      string
}

func (Thread) TableName() string {
	return "forocoches_top_threads"
}
