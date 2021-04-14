package main

import (
	"awesomeProject/src/designMode/command"
	"fmt"
	"time"
)

func test03() {
	eventChan := make(chan string)
	go func() {
		events := []string{"start","archive"}
		for _,e := range events {
			eventChan <- e
		}
	}()
	defer close(eventChan)

	commands:= make(chan command.ICommand,1000)
	defer close(commands)

	go func() {
		for {
			event,ok := <-eventChan
			if !ok {
				return
			}
			var c command.ICommand
			switch event {
			case "start":
				c = command.NewStartCommand()
			case "archive":
				c = command.NewArchiveCommand()
			}
			commands <- c
		}
	}()

	for {
		select {
		 case c := <- commands:
			 _ = c.Execute()
		 	case <- time.After(1 * time.Second):
				return
		}
	}
}


func test04() {
	patches := command.Patches{
		Do: func() {
			fmt.Println("开始游戏")
		},
		Undo: func() {
			fmt.Println("结束游戏")
		},
	}
	command.DoCommand(patches)


	patches1 := command.Patches{
		Do: func() {
			fmt.Println("移动到前方10米")
		},
		Undo: func() {
			fmt.Println("移动到后方10米")
		},
	}
	command.DoCommand(patches1)

	time.Sleep(1 * time.Second / 2)

	command.UndoWorker()
	command.UndoWorker()
}
func main() {
-- 分界线

}