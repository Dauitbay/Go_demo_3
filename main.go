package main

import "fmt"

func main() {
	bookmarks := map[string]string {}
	fmt.Println("Bookmark app")
Menu:
	for {
		option := getMenu()
		switch option {
		case 1:
			printBookmarks(bookmarks)
		case 2:
			bookmarks = addBookmark(bookmarks)
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

func printBookmarks(bookmarks map[string]string) {
	if len(bookmarks) == 0 {
		fmt.Println("No bookmarks")
	}
	for key, value := range bookmarks {
		fmt.Println(key, ": ", value)
	} 
}

func addBookmark(bookmarks map[string]string) map[string]string {
	var newBookmarkKey string
	var newBookmarkValue string
	fmt.Print("Enter the name: ")
	fmt.Scan(&newBookmarkKey)
	fmt.Print("Enter the link: ")
	fmt.Scan(&newBookmarkValue)
	bookmarks[newBookmarkKey] = newBookmarkValue
	return bookmarks
}

func deleteBookmark(bookmarks map[string]string) map[string]string{
	var bookmarkKeyToDelete string
	fmt.Print("Enter the name: ")
	fmt.Scan(&bookmarkKeyToDelete)
	delete(bookmarks, bookmarkKeyToDelete)
	return bookmarks
}
