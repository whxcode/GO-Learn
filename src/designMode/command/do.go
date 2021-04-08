package command

import "container/list"

type Patches struct {
	Do func()
	Undo func()
}
var undoStack = list.New()
var redoStack = list.New()

func DoCommand(c Patches) {
	Do := c.Do
	do := NewWrapperFunc(Do)
	undoStack.PushBack(&c)
	do.Execute()
}
func UndoWorker() {
	if(undoStack.Len() == 0) {
		return
	}
	do := undoStack.Back()

	undoStack.Remove(do)
	_ = do.Value.(*Patches).Undo

}
func RedoWorker()  {
	if(redoStack.Len() == 0) {
		return
	}

	do := redoStack.Back()
	redoStack.Remove(do)
	_ = do.Value.(ICommand).Execute()
}