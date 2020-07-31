package main

func main() {
	engine, cleanup := InitializeEngine()
	defer cleanup()
	engine.Run()
}
