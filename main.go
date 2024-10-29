package main

import "fmt"

type bookmarkMap = map[string]string

func main() {
	bookmarks := bookmarkMap {}
	fmt.Println("Bookmark app")
Menu:
	for {
		option := getMenu()
		switch option {
		case 1:
			printBookmarks(bookmarks)
		case 2:
			addBookmark(bookmarks)
		case 3:
			deleteBookmark(bookmarks)
		case 4:
			break Menu
		}
	}
}

func getMenu() int {
	var option int
	fmt.Println("Choose an option ")
	fmt.Println("1. See bookmarks")
	fmt.Println("2. Add bookmark")
	fmt.Println("3. Delete bookmark")
	fmt.Println("4. Exit")
	fmt.Scan(&option)
	return option
}

func printBookmarks(bookmarks bookmarkMap) {
	if len(bookmarks) == 0 {
		fmt.Println("No bookmarks")
	}
	for key, value := range bookmarks {
		fmt.Println(key, ": ", value)
	} 
}

func addBookmark(bookmarks bookmarkMap) {
	var newBookmarkKey string
	var newBookmarkValue string
	fmt.Print("Enter the name: ")
	fmt.Scan(&newBookmarkKey)
	fmt.Print("Enter the link: ")
	fmt.Scan(&newBookmarkValue)
	bookmarks[newBookmarkKey] = newBookmarkValue
}

func deleteBookmark(bookmarks bookmarkMap) {
	var bookmarkKeyToDelete string
	fmt.Print("Enter the name: ")
	fmt.Scan(&bookmarkKeyToDelete)
	delete(bookmarks, bookmarkKeyToDelete)
}
