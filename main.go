package main

import "fmt"

type stringMap map[string]string

func main() {
	bookmarks := make(stringMap)
	fmt.Println("Hello! This is bookmarks application!")
mainLoop:
	for {
		input := menu()
		switch input {
		case 1:
			viewBookmarks(bookmarks)
		case 2:
			addNewBookmark(bookmarks)
		case 3:
			deleteBookmark(bookmarks)
		case 4:
			exit()
			break mainLoop
		default:
			fmt.Println("Invalid option or input. Please enter number from 1 to 4.")
		}
	}
}

func viewBookmarks(bookmarks stringMap) {
	fmt.Println("All bookmarks will be viewed:")
	if len(bookmarks) == 0 {
		fmt.Println("No bookmarks found")
		return
	}
	for key, value := range bookmarks {
		fmt.Printf("%s: %s\n", key, value)
	}

}

func addNewBookmark(bookmarks stringMap) {
	var key string
	fmt.Println("Print bookmark name that would be added")
	fmt.Scanln(&key)
	var value string
	fmt.Println("Print bookmark link")
	fmt.Scanln(&value)
	bookmarks[key] = value
}

func deleteBookmark(bookmarks stringMap) {
	fmt.Println("Print bookmark name that would be deleted")
	var key string
	fmt.Scanln(&key)
	delete(bookmarks, key)
}

func exit() {
	fmt.Println("Exiting...")
}

func menu() int8 {
	fmt.Printf("Choose option that you needed: \n1: View all bookmarks\n2: Add new bookmark\n3: Delete bookmark\n4: Exit\n")
	var input int8
	fmt.Scanln(&input)
	return input
}
