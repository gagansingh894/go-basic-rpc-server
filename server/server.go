package server

type Item struct {
	title string
	body  string
}

var database []Item

func GetByName(title string) Item {
	var getItem Item

	for _, item := range database {
		if item.title == title {
			getItem = item
			break
		}
	}
	return getItem
}

func CreateItem(item Item) Item {
	database = append(database, item)
	return item
}

func EditItem(edit Item) Item {
	var changedItem Item

	for idx, item := range database {
		if item.title == edit.title {
			database[idx] = edit
			changedItem = edit
			break
		}
	}
	return changedItem
}

func DeleteItem(toDelete Item) Item {
	var deletedItem Item

	for idx, item := range database {
		if item.title == toDelete.title && item.body == toDelete.body {
			database = append(database[:idx], database[idx+1:]...)
			deletedItem = toDelete
		}
	}
	return deletedItem
}
