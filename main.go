package main

func main() {
	initQueue()
	dispatcher := newDispatcher()
	dispatcher.run()
	start()
}
