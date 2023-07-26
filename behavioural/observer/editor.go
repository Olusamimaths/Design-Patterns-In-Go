package observer

// Editor (Uses the publisher interface)
type IEditor interface {
	OpenFile(filename string)
	SaveFile()
	DeleteFile()
}
type Editor struct {
	eventManager *EventManager
}

func NewEditor() *Editor {
	return &Editor{NewEventManager()}
}

func(editor *Editor) OpenFile(filename string) {
	editor.eventManager.Notify("open", filename)
}

func(editor *Editor) SaveFile() {
	editor.eventManager.Notify("save", "")
}

func(editor *Editor) DeleteFile() {
	editor.eventManager.Notify("delete", "")
}

func ObserverPatternExample() {
	editor := NewEditor()

	subscriber1 := &Subscriber1{}
	subscriber2 := &Subscriber2{}

	editor.eventManager.Subscribe("open", subscriber1)
	editor.eventManager.Subscribe("open", subscriber2)
	editor.eventManager.ListListeners()

	editor.OpenFile("test.txt")

	editor.eventManager.Unsubscribe("open", subscriber1)
	editor.eventManager.Unsubscribe("open", subscriber2)

	editor.OpenFile("test.txt")

	editor.eventManager.ListListeners()

	editor.OpenFile("test2.txt")

	editor.eventManager.Subscribe("save", subscriber1)
	editor.SaveFile()

	editor.eventManager.Unsubscribe("open", subscriber2)
	editor.DeleteFile()
}