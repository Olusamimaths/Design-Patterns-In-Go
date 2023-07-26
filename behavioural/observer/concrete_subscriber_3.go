package observer

type Subscriber3 struct {
}
func(subscriber3 *Subscriber3) Update(filename string) {
	println("\tSub 3: Syncing file: " + filename)
}
func(subscriber3 *Subscriber3) Name() string {
	return "Subscriber3"
}
func(subscriber3 *Subscriber3) ID() int {
	return 3
}