package main

func main() {
	server := InitializedServer()

	server.MigrateDB()
	server.RunServer()

}
