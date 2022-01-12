package server

type Item struct {
	Title string
	Body  string
}

var database []Item

// implement custom type
// this is used to convert functions to methods as it is the requirement og rpc to have methods
type API int

func (a *API) GetDB(empty string, reply *[]Item) error {
	*reply = database
	return nil
}

func (a *API) AddItem(item Item, reply *Item) error {
	database = append(database, item)
	*reply = item
	return nil
}

func (a *API) GetByName(title string, reply *Item) error {
	var getItem Item

	for _, item := range database {
		if item.Title == title {
			getItem = item
			break
		}
	}
	*reply = getItem
	return nil
}

func (a *API) CreateItem(item Item, reply *Item) error {
	database = append(database, item)
	*reply = item
	return nil
}

func (a *API) EditItem(edit Item, reply *Item) error {
	var changedItem Item

	for idx, item := range database {
		if item.Title == edit.Title {
			database[idx] = edit
			changedItem = edit
			break
		}
	}
	*reply = changedItem
	return nil
}

func (a *API) DeleteItem(toDelete Item, reply *Item) error {
	var deletedItem Item

	for idx, item := range database {
		if item.Title == toDelete.Title && item.Body == toDelete.Body {
			database = append(database[:idx], database[idx+1:]...)
			deletedItem = toDelete
		}
	}
	*reply = deletedItem
	return nil
}
