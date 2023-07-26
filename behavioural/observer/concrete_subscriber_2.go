package observer

type Subscriber2 struct {
}
func(subscriber2 *Subscriber2) Update(filename string) {
	println("\tSub 2: Tracking file: " + filename)
}
func(subscriber2 *Subscriber2) Name() string {
	return "Subscriber2"
}
func(subscriber2 *Subscriber2) ID() int {
	return 2
}