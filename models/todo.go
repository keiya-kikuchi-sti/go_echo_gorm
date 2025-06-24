package models

// GORM では、テーブル名は struct の名前をスネークケースにして s をつけたもの
type Todo struct {
	Id   int
	Text string
	Done bool
}
