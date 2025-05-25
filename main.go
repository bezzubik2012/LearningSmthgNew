package main

import "fmt"

func main() {
	//var bookmarks map[string]string
	bookmarks := make(map[string]string)
	fmt.Println("Hello! This is bookmarks application!")
	for {
		fmt.Printf("Choose option that you needed: \n1: View all bookmarks\n2: Add new bookmark\n3: Delete bookmark\n4: Exit\n")
		var input int8
		fmt.Scanln(&input)
		switch input {
		case 1:
			{
				viewBookmarks(bookmarks)
			}
		case 2:
			{
				addNewBookmark(bookmarks)
			}
		case 3:
			{
				deleteBookmark(bookmarks)
			}
		case 4:
			{
				exit()
			}
		default:
			{
				fmt.Println("Invalid option or input. Please enter number from 1 to 4.")
			}
		}
	}
}

func viewBookmarks(bookmarks map[string]string) {
	fmt.Println("All bookmarks will be viewed")
	for key, value := range bookmarks {
		fmt.Printf("%s: %s\n", key, value)
	}
}

func addNewBookmark(bookmarks map[string]string) {
	var key string
	fmt.Println("Print bookmark name that would be added")
	fmt.Scanln(&key)
	var value string
	fmt.Println("Print bookmark link")
	fmt.Scanln(&value)
	bookmarks[key] = value
}

func deleteBookmark(bookmarks map[string]string) {
	fmt.Println("Print bookmark name that would be deleted")
	var key string
	fmt.Scanln(&key)
	delete(bookmarks, key)
}

func exit() {
	fmt.Println("Exiting...")
}
