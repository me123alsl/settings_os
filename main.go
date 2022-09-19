package main

import (
	"flag"

	cmd "github.com/me123alsl/openmsa_installer/commands"
	utils "github.com/me123alsl/openmsa_installer/utils"
)

func main() {
	Log := cmd.Log
	utils.SetLogLevel("debug")

	username := flag.String("user", "openmsa", "username")
	password := flag.String("password", "openmsa", "password")
	server_type := flag.String("type", "client", "server type [server/client]")
	serverListPath := flag.String("path", "/usr/local/openmsa_hosts", "serverlist file path [default : /usr/local/openmsa_hosts]")
	flag.Parse()
	currentUser, _ := cmd.GetOSUsername()

	//serverlist path check
	if *serverListPath == "" {
		*serverListPath = "/usr/local/openmsa_hosts"
	}
	// show flags
	Log.Info("username:", *username)
	Log.Info("password:", *password)
	Log.Info("server_type:", *server_type)
	if *server_type == "server" {
		Log.Info("server_path:", *serverListPath)
	}
	Log.Info("currentUser:", currentUser)

	// create user
	if IsUserExist, _ := cmd.IsUserExist(*username); IsUserExist {
		Log.Info("user already exist")
	} else {
		err := cmd.CreateUser(*username, *password)
		if err != nil {
			Log.Info(err)
			panic(err)
		}
		err = cmd.ChangePassword(*username, *password)
		if err != nil {
			Log.Info(err)
			panic(err)
		}
	}

	// add sudo
	err := cmd.AddSudo(*username)
	if err != nil {
		Log.Info(err)
		panic(err)
	}

	// package_server install
	if *server_type == "server" {
		err = cmd.InstallServerPackage()
		if err != nil {
			Log.Info(err)
			panic(err)
		}
	} else {
		err = cmd.InstallClientPackage()
		if err != nil {
			Log.Info(err)
			panic(err)
		}
	}

	// check python version
	err = cmd.InstallPython3IfNeed()
	if err != nil {
		Log.Info(err)
		panic(err)
	}

	if *server_type == "server" {
		// ssh keygen
		Log.Info("ssh keygen")
		err = cmd.CreateSSHKey(*username, *password)
		if err != nil {
			Log.Info(err)
			panic(err)
		}

		//read serverlist
		serverList, err := utils.ReadFileText(*serverListPath)
		if err != nil {
			Log.Info(err)
			panic(err)
		}

		Log.Info("ssh key copy")
		err = cmd.CopySSHKey(*username, *password, serverList)
		if err != nil {
			Log.Info(err)
			panic(err)
		}
	}
}
