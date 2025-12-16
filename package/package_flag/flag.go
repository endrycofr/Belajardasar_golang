package main

import "flag"

func main() {
	var host *string = flag.String("host", "localhost", "Hostname for the server")
	var username *string = flag.String("username", "admin", "Username for authentication")
	var password *string = flag.String("password", "password", "Password for authentication")
	var port *int = flag.Int("port", 8080, "Port number for the server")

	flag.Parse()

	println("Host:", *host)
	println("Username:", *username)
	println("Password:", *password)
	println("Port:", *port)

}
