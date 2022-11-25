package main

func main() {
	server()
}

// server
func server() {
	r := Routes()
	r.Run() // listen and serve on 0.0.0.0:8080
}
