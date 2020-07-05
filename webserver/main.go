package main

func main() {
	server := new_server(":3000")
	server.Handle("GET", "/", HandleRoot)
	server.Handle("POST", "/create", PostRequest)
	server.Handle("POST", "/api", server.AddMiddleware(HandleHome, CheckAuth(), Logging()))
	server.Listen()
}
