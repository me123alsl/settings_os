package commands

// 다른 사용자로 ssh 키 생성
func CreateSSHKey(username string) error {
	return ExecSudoCommand(username, "ssh-keygen -t rsa -b 4096 -f ~/.ssh/id_rsa -q -N \"\"")
}

// ssh키를 다른 서버들로 복사 with server list
func CopySSHKey(username string, serverList []string) error {
	for _, server := range serverList {
		err := ExecSudoCommand(username, "ssh-copy-id -i ~/.ssh/id_rsa.pub "+server)
		if err != nil {
			return err
		}
	}
	return nil
}
