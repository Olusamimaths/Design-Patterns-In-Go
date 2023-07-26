package observer

type IEventListener interface {
	ID() int
	Update(filename string)
	Name() string
}
