package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/thiennguyenluonghoang/small-dairy/internal/services"
)

var globalReader = bufio.NewReader(os.Stdin)

func main() {
	// var op []func() //Slice chứa function
	// op = append(op, createNewDiary_main)
	// op = append(op, showAllDiaries_main)
	// op = append(op, updateDiary_main)
	// op = append(op, deleteDiary_main)
	// for {
	// 	if v, err := strconv.ParseInt(displayMenu_main(), 10, 64); v <= int64(len(op)) && err == nil {
	// 		op[v-1]()
	// 	} else {
	// 		return
	// 	}
	// }

	for {
		switch displayMenu_main() {
		case "1":
			createNewDiary_main()
		case "2":
			showAllDiaries_main()
		case "3":
			updateDiary_main()
		case "4":
			deleteDiary_main()
		default:
			return
		}
	}

}

// ! Tách ra
func displayMenu_main() string {
	defer fmt.Printf("--------------------------\n")
	fmt.Printf("Diary App\n")
	fmt.Printf("1. Create new diary\n")
	fmt.Printf("2. Show all diaries\n")
	fmt.Printf("3. Update diary by Id\n")
	fmt.Printf("4. Delete existed diary\n")
	fmt.Printf("5. Exit app\n")
	fmt.Printf("Enter choices (1-5): ")
	choices, _ := globalReader.ReadString('\n')
	return strings.TrimSpace(choices)
}

// Chưa xử lý ngoại lệ ở trong này -> Cần 1 lớp chuyên cho nhập xuất nữa, ko đc bỏ hết logic hiển thị trong controller, controller chỉ routing thôi!
func createNewDiary_main() {
	defer fmt.Printf("--------------------------\n")
	fmt.Printf("Create new diary\n")
	fmt.Printf("Enter title: ")
	title, _ := globalReader.ReadString('\n')
	fmt.Printf("Enter content: ")
	content, _ := globalReader.ReadString('\n')

	services.CreateDiary(title, content)
	fmt.Printf("Add successful\n")
}

func showAllDiaries_main() {
	fmt.Printf("List diaries: ")
	defer fmt.Printf("--------------------------\n")
	services.ShowAllDiaries()
}

func updateDiary_main() {
	defer fmt.Printf("--------------------------\n")
	var id int
	fmt.Printf("Enter id of diary u want to find: ")
	fmt.Scan(&id)

	if services.FindDiary(id) == -1 {
		fmt.Printf("No diary\n")
		return
	}
	fmt.Printf("Enter new title: ")
	newTitle, _ := globalReader.ReadString('\n')
	fmt.Printf("Enter new content: ")
	newContent, _ := globalReader.ReadString('\n')
	services.UpdateDiary(id, newTitle, newContent)
	fmt.Printf("Diary updated!\n")
}

func deleteDiary_main() {
	defer fmt.Printf("--------------------------\n")
	var id int
	fmt.Printf("Enter id of diary u want to find: ")
	fmt.Scan(&id)

	if services.FindDiary(id) == -1 {
		fmt.Printf("No diary\n")
		return
	}

	services.DeleteDiary(id)
	fmt.Printf("Diary with id %d is deleted", id)

}
