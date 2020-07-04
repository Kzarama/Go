package main

func main() {
	server := new_server(":3000")
	server.Handle("/", HandleRoot)
	server.Handle("/api", server.AddMiddleware(HandleHome, CheckAuth(), Logging()))
	server.Listen()
}
