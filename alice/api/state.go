package api

type State string

const (
	StateInit          State = ""
	StateCreateReqName State = "CREATE_REQ_NAME"
	StateDelReqConfirm State = "DELETE_REQ_CNFRM"
)

type StateData struct {
	State State
	//ListID   model.TODOListID
	ListName string
	ItemText string
	//ItemID   model.ListItemID
}

func (s *StateData) GetState() State {
	if s == nil {
		return StateInit
	}
	return s.State
}
