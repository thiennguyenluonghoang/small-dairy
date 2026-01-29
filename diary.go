package main

import (
	"fmt"
	"time"
)

type Diary struct {
	//Tên + kiểu dữ liệu + struct tag (metadata)
	ID       int
	Title    string
	Content  string
	CreateAt time.Time
}

func (d Diary) String() string {
	return fmt.Sprintf("\nID: %d\nTitle: %sContent: %sDate: %v\n",
		d.ID, d.Title, d.Content, d.CreateAt.Format("02/01/2006 15:04:05"))
}
