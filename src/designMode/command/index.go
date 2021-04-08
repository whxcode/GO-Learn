package command

import "fmt"

// 命令触发器
type ICommand interface {
	Execute() error
}

type WrapperFunc struct {
	f func ()
}

func (w *WrapperFunc) Execute()  error{
	w.f()
	return nil
}

func NewWrapperFunc(f func()) *WrapperFunc {
	return &WrapperFunc{f}
}

type StartCommand struct {}



func NewStartCommand() *StartCommand{
	return &StartCommand{  }
}

func (c *StartCommand) Execute() error {
	fmt.Println("game start:")
	return nil
}


type ArchiveCommand struct {

}

func (a *ArchiveCommand) Execute() error {
	fmt.Println("game archive:")
	return nil
}

func NewArchiveCommand() *ArchiveCommand{
	return &ArchiveCommand{}
}
