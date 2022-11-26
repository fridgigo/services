package main

func main() {
	server()
}

// server
func server() {
	r := Routes()
	r.Run(":4000")
}
