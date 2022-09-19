package commands

import (
	"fmt"
)

// 다른 사용자로 ssh 키 생성
func CreateSSHKey(username, password string) error {
	sshkey := fmt.Sprintf("/home/%s/.ssh/id_rsa.pub", username)
	Log.Debug("create sshkey : " + sshkey)
	if FileExists(sshkey) {
		Log.Debug("ssh key already exist")
		return nil
	}
	return ExecSudoCommandWithPassword(username, password, "ssh-keygen -b 2048 -t rsa -f ~/.ssh/id_rsa -q -N ''")
}

// ssh키를 다른 서버들로 복사 with server list
func CopySSHKey(username, password string, serverList []string) error {
	Log.Debug("CopySSHKey")
	for _, server := range serverList {
		Log.Debug("Copy SSH Key server:", server)
		err := ExecSudoCommandWithPassword(username, password, "\"sshpass -p "+password+" sudo ssh-copy-id -o StrictHostKeyChecking=no "+username+"@"+server+" -p 22\"")
		if err != nil {
			Log.Error(err)
		}
	}
	return nil
}
