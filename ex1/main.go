package main

import (
	store2 "StudyYura/store"
	"fmt"
)

func main() {

	userStore := store2.NewUserStore()

	user1 := store2.Document{
		Fields: map[string]store2.DocumentField{
			"key":  {Type: store2.DocumentFieldTypeString, Value: "1"},
			"name": {Type: store2.DocumentFieldTypeString, Value: "Anna"},
			"age":  {Type: store2.DocumentFieldTypeNumber, Value: "23"},
		},
	}
	user2 := store2.Document{
		Fields: map[string]store2.DocumentField{
			"key":  {Type: store2.DocumentFieldTypeString, Value: "2"},
			"name": {Type: store2.DocumentFieldTypeString, Value: "Nastia"},
			"age":  {Type: store2.DocumentFieldTypeNumber, Value: "22"},
		},
	}
	fmt.Println("Create documents:", user1, user2)
	userStore.Create(user1)
	userStore.Create(user2)

	fmt.Println("Get documents:")
	key := "1"
	doc, found := userStore.Get(key)
	if found {
		fmt.Println("user:", doc)
	} else {
		fmt.Println("Not found")
	}

	fmt.Println("Delete documents:")
	//key = "2"
	//doc, found = userStore.Get(key)
	//if found {
	//	deleted := userStore.Delete(key)
	//	if deleted {
	//		fmt.Println("user:", doc.Fields["name"].Value)
	//	} else {
	//		fmt.Println("Not found")
	//	}
	//} else {
	//	fmt.Println("User not found, nothing to delete.")
	//}

	//if userStore.Delete("1") {
	//	fmt.Println("Delete:", doc.Fields["name"].Value)
	//} else {
	//	fmt.Println("Not delete user")
	//}

	deleted := userStore.Delete("1")
	if deleted {
		fmt.Println("Delete:", doc.Fields["name"].Value, deleted)
	} else {
		fmt.Println("Not delete", deleted)
	}

	fmt.Println("List documents:")
	for _, doc := range userStore.List() {
		fmt.Println(doc.Fields)
	}
}
