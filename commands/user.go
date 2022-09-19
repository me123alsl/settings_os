package commands

// OS 현재 사용자 이름 가져오기
func GetOSUsername() (string, error) {
	return ExecCommandWithOutput("whoami")
}

// 사용자가 있는지 확인
func IsUserExist(username string) (bool, error) {
	_, err := ExecSudoCommandWithOutput("id " + username)
	if err != nil {
		return false, err
	}
	return true, nil
}

// sudo 권한으로 사용자 생성 명령어 실행
func CreateUser(username, password string) error {
	return ExecSudoCommand("adduser --quiet --disabled-password --shell /bin/bash --home /home/" + username + " --gecos '' " + username)
}

func ChangePassword(username, password string) error {
	return ExecSudoCommand("echo \"" + username + ":" + password + "\" | sudo chpasswd")
}

// sudo 권한으로 사용자 sudo 권한 부여
func AddSudo(username string) error {
	return ExecSudoCommand("usermod -aG sudo " + username)
}
