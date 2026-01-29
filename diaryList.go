package main

import (
	"fmt"
	"time"
)

var diaryList []Diary //slice = list
var record Diary

func createDiary(title string, content string) {
	record = Diary{ID: len(diaryList) + 1, Title: title, Content: content, CreateAt: time.Now()}
	diaryList = append(diaryList, record)

}

// Vi phạm nguyên tắc: Model chỉ làm việc với dữ liệu, ko routing hay in ấn gì hết
func showAllDiaries() {
	if len(diaryList) == 0 {
		fmt.Printf("No diary")
	}
	for _, diary := range diaryList {
		fmt.Printf(diary.String())
	}
}

func findDiary(id int) int {
	for index, diary := range diaryList {
		if diary.ID == id {
			return index
		}
	}
	return -1
}

func updateDiary(id int, newTitle string, newContent string) {
	diaryList[findDiary(id)].Title = newTitle
	diaryList[findDiary(id)].Content = newContent
	diaryList[findDiary(id)].CreateAt = time.Now()

}

func deleteDiary(id int) {
	var index = findDiary(id)
	diaryList = append(diaryList[:index], diaryList[index+1:]...)

}
