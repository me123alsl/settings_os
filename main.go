package main

import (
	"flag"

	cmd "github.com/me123alsl/settings_os/commands"
)

func main() {
	Log := cmd.Log

	username := flag.String("user", "opnemsa", "username")
	password := flag.String("password", "openmsa", "password")
	server_type := flag.String("type", "client", "server type [server/client]")
	serverlist := flag.Args()
	flag.Parse()
	currentUser, _ := cmd.GetOSUsername()

	// show flags
	Log.Info("username:", *username)
	Log.Info("password:", *password)
	Log.Info("server_type:", *server_type)
	Log.Info("serverlist:", serverlist)
	Log.Info("currentUser:", currentUser)

	// // create user
	// err := cmd.CreateUser(*username, *password)
	// if err != nil {
	// 	Log.Info(err)
	// 	panic(err)
	// }

	// // add sudo
	// err = cmd.AddSudo(*username)
	// if err != nil {
	// 	Log.Info(err)
	// 	panic(err)
	// }

	// // package_server install
	// if *server_type == "server" {
	// 	err = cmd.InstallServerPackage()
	// 	if err != nil {
	// 		Log.Info(err)
	// 		panic(err)
	// 	}
	// } else {
	// 	err = cmd.InstallClientPackage()
	// 	if err != nil {
	// 		Log.Info(err)
	// 		panic(err)
	// 	}
	// }

	// check python version
	// err := cmd.InstallPython3IfNeed()
	// if err != nil {
	// 	Log.Info(err)
	// 	panic(err)
	// }

}
