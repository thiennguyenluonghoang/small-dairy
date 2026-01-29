# 📁 PROJECT EXPORT FOR LLMs

## 📊 Project Information

- **Project Name**: `GO`
- **Generated On**: 2026-01-29 06:50:55 (Asia/Bangkok / GMT+07:00)
- **Total Files Processed**: 4
- **Export Tool**: Easy Whole Project to Single Text File for LLMs v1.1.0
- **Tool Author**: Jota / José Guilherme Pandolfi

### ⚙️ Export Configuration

| Setting | Value |
|---------|-------|
| Language | `en` |
| Max File Size | `1 MB` |
| Include Hidden Files | `false` |
| Output Format | `both` |

## 🌳 Project Structure

```
├── 📄 diary.go (354 B)
├── 📄 diaryList.go (986 B)
├── 📄 go.mod (63 B)
└── 📄 main.go (2.08 KB)
```

## 📑 Table of Contents

**Project Files:**

- [📄 diary.go](#📄-diary-go)
- [📄 diaryList.go](#📄-diarylist-go)
- [📄 main.go](#📄-main-go)

---

## 📈 Project Statistics

| Metric | Count |
|--------|-------|
| Total Files | 4 |
| Total Directories | 0 |
| Text Files | 3 |
| Binary Files | 1 |
| Total Size | 3.45 KB |

### 📄 File Types Distribution

| Extension | Count |
|-----------|-------|
| `.go` | 3 |
| `.mod` | 1 |

## 💻 File Code Contents

### <a id="📄-diary-go"></a>📄 `diary.go`

**File Info:**
- **Size**: 354 B
- **Extension**: `.go`
- **Language**: `go`
- **Location**: `diary.go`
- **Relative Path**: `root`
- **Created**: 2026-01-28 10:00:20 (Asia/Bangkok / GMT+07:00)
- **Modified**: 2026-01-29 06:29:20 (Asia/Bangkok / GMT+07:00)
- **MD5**: `27b6ca1948624d45e2c2d860527a72ba`
- **SHA256**: `a17689c7b26b543ffeebce9c7433d5a9b24c730179b003e1a3a6fa555a921901`
- **Encoding**: UTF-8

**File code content:**

```go
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

```

---

### <a id="📄-diarylist-go"></a>📄 `diaryList.go`

**File Info:**
- **Size**: 986 B
- **Extension**: `.go`
- **Language**: `go`
- **Location**: `diaryList.go`
- **Relative Path**: `root`
- **Created**: 2026-01-28 10:05:15 (Asia/Bangkok / GMT+07:00)
- **Modified**: 2026-01-29 06:38:24 (Asia/Bangkok / GMT+07:00)
- **MD5**: `5297d46915fdea0d2e07605630fc1f28`
- **SHA256**: `895e6c965537c41f6f23b37b4bb8eebcb5756bd4611cbfc06c386c2692f60bac`
- **Encoding**: UTF-8

**File code content:**

```go
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

```

---

### <a id="📄-main-go"></a>📄 `main.go`

**File Info:**
- **Size**: 2.08 KB
- **Extension**: `.go`
- **Language**: `go`
- **Location**: `main.go`
- **Relative Path**: `root`
- **Created**: 2026-01-28 07:31:18 (Asia/Bangkok / GMT+07:00)
- **Modified**: 2026-01-29 06:42:14 (Asia/Bangkok / GMT+07:00)
- **MD5**: `84a634c9cf851f1acd46409398236cc1`
- **SHA256**: `cc1f5ff37c0cbbc4880c6fed2976528fa7a2938434e9b04fec8e62e21d9c2e51`
- **Encoding**: UTF-8

**File code content:**

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var globalReader = bufio.NewReader(os.Stdin)

func main() {

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

	createDiary(title, content)
	fmt.Printf("Add successful\n")
}

func showAllDiaries_main() {
	fmt.Printf("List diaries: ")
	defer fmt.Printf("--------------------------\n")
	showAllDiaries()
}

func updateDiary_main() {
	defer fmt.Printf("--------------------------\n")
	var id int
	fmt.Printf("Enter id of diary u want to find: ")
	fmt.Scan(&id)

	if findDiary(id) == -1 {
		fmt.Printf("No diary\n")
		return
	}
	fmt.Printf("Enter new title: ")
	newTitle, _ := globalReader.ReadString('\n')
	fmt.Printf("Enter new content: ")
	newContent, _ := globalReader.ReadString('\n')
	updateDiary(id, newTitle, newContent)
	fmt.Printf("Diary updated!\n")
}

func deleteDiary_main() {
	defer fmt.Printf("--------------------------\n")
	var id int
	fmt.Printf("Enter id of diary u want to find: ")
	fmt.Scan(&id)

	if findDiary(id) == -1 {
		fmt.Printf("No diary\n")
		return
	}

	deleteDiary(id)
	fmt.Printf("Diary with id %d is deleted", id)

}

```

---

## 🚫 Binary/Excluded Files

The following files were not included in the text content:

- `go.mod`

