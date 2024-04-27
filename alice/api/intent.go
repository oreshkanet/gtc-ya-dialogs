package api

type Intents struct {
	Confirm    *EmptyObj         `json:"YANDEX.CONFIRM"`
	Reject     *EmptyObj         `json:"YANDEX.REJECT"`
	CreateList *IntentCreateList `json:"list_create"`
	ListLists  *EmptyObj         `json:"list_lists"`
	Cancel     *EmptyObj         `json:"cancel"`
}

type IntentCreateList struct {
	Slots IntentCreateListSlots `json:"slots"`
}

type IntentCreateListSlots struct {
	ListName string `json:"listName"`
}
