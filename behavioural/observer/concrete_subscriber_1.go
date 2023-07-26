package observer

type Subscriber1 struct {
}
func(subscriber1 *Subscriber1) Update(filename string) {
	println("\tSub 1: Woking on file: " + filename)
}
func(subscriber1 *Subscriber1) Name() string {
	return "Subscriber1"
}
func(subscriber1 *Subscriber1) ID() int {
	return 1
}