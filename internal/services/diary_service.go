package services

import (
	"fmt"
	"time"

	"github.com/thiennguyenluonghoang/small-dairy/internal/models"
)

var diaryList []models.Diary //slice = list
var record models.Diary

func CreateDiary(title string, content string) {
	record = models.Diary{ID: len(diaryList) + 1, Title: title, Content: content, CreateAt: time.Now()}
	diaryList = append(diaryList, record)

}

// Vi phạm nguyên tắc: Model chỉ làm việc với dữ liệu, ko routing hay in ấn gì hết
func ShowAllDiaries() {
	if len(diaryList) == 0 {
		fmt.Printf("No diary")
	}
	for _, diary := range diaryList {
		fmt.Printf(diary.String())
	}
}

func FindDiary(id int) int {
	for index, diary := range diaryList {
		if diary.ID == id {
			return index
		}
	}
	return -1
}

func UpdateDiary(id int, newTitle string, newContent string) {
	diaryList[FindDiary(id)].Title = newTitle
	diaryList[FindDiary(id)].Content = newContent
	diaryList[FindDiary(id)].CreateAt = time.Now()

}

func DeleteDiary(id int) {
	var index = FindDiary(id)
	diaryList = append(diaryList[:index], diaryList[index+1:]...)

}
