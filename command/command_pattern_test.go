package command

import "testing"

type RemoteController struct {
	OnCommand   []Command
	OffCommand  []Command
	UndoCommand Command
}

func NewRemoteController(slot int) RemoteController {
	if slot <= 0 {
		slot = 7
	}
	rc := RemoteController{
		OnCommand:  make([]Command, slot, slot),
		OffCommand: make([]Command, slot, slot),
	}
	for i := 0; i < slot; i++ {
		rc.OnCommand[i] = &FakeCommand{}
		rc.OffCommand[i] = &FakeCommand{}
	}
	return rc
}

func (rc *RemoteController) SetCommand(slot int, onCommand, offCommand Command) {
	rc.OnCommand[slot] = onCommand
	rc.OffCommand[slot] = offCommand
}

func (rc *RemoteController) PressOnButton(slot int) {
	rc.OnCommand[slot].Execute()
	rc.UndoCommand = rc.OnCommand[slot]

}

func (rc *RemoteController) PressOffButton(slot int) {
	rc.OffCommand[slot].Execute()
	rc.UndoCommand = rc.OffCommand[slot]
}

func (rc *RemoteController) PressUndo() {
	rc.UndoCommand.Undo()
}

func TestLightOffCommand_execute(t *testing.T) {
	var light1 = &Light{Name: "卧室"}
	var light2 = &Light{Name: "客厅"}
	var light3 = &Light{Name: "厨房"}
	var light4 = &Light{Name: "卫生间"}
	var allLight = []*Light{light4, light3, light2, light1}

	rc := NewRemoteController(6)
	rc.SetCommand(0, &LightOnCommand{light1}, &LightOffCommand{light1})
	rc.SetCommand(1, &LightOnCommand{light2}, &LightOffCommand{light2})
	rc.SetCommand(2, &LightOnCommand{light3}, &LightOffCommand{light3})
	rc.SetCommand(3, &LightOnCommand{light4}, &LightOffCommand{light4})
	rc.SetCommand(4, &AllLightOnCommand{allLight}, &AllLightOffCommand{allLight})
	rc.SetCommand(5, &FakeCommand{}, &FakeCommand{})
	rc.PressOffButton(1)
	rc.PressUndo()
	rc.PressOffButton(2)
	rc.PressUndo()
	rc.PressOffButton(3)
	rc.PressUndo()
	rc.PressOffButton(0)
	rc.PressUndo()
	rc.PressOnButton(4)
	rc.PressUndo()
	rc.PressOnButton(5)
	rc.PressUndo()
}
