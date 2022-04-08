package command

import "fmt"

type Command interface {
	Execute()
	Undo()
}

type Light struct {
	Name string
}

func (l Light) LightOn() {
	fmt.Printf("light %s is On\n", l.Name)
}
func (l Light) LightOff() {
	fmt.Printf("light %s is Off\n", l.Name)
}

type LightOnCommand struct {
	Light *Light
}

type FakeCommand struct {
}

func (f *FakeCommand) Execute() {
	fmt.Println("this is a fake Command")
}

func (f *FakeCommand) Undo() {
	fmt.Println("this is a fake Command")
}

type LightOffCommand struct {
	Light *Light
}

func (l *LightOnCommand) Undo() {
	l.Light.LightOff()
}

func (l *LightOnCommand) Execute() {
	l.Light.LightOn()
}

func (l *LightOffCommand) Execute() {
	l.Light.LightOff()
}

func (l *LightOffCommand) Undo() {
	l.Light.LightOn()
}

type AllLightOnCommand struct {
	Light []*Light
}

func (l *AllLightOnCommand) Execute() {
	for _, v := range l.Light {
		v.LightOn()
	}
}

func (l *AllLightOnCommand) Undo() {
	for _, v := range l.Light {
		v.LightOff()
	}
}

type AllLightOffCommand struct {
	Light []*Light
}

func (l *AllLightOffCommand) Execute() {
	for _, v := range l.Light {
		v.LightOff()
	}
}

func (l *AllLightOffCommand) Undo() {
	for _, v := range l.Light {
		v.LightOff()
	}
}

