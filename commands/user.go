package commands

// OS 현재 사용자 이름 가져오기
func GetOSUsername() (string, error) {
	return ExecSudoCommandWithOutput("whoami")
}

// sudo 권한으로 사용자 생성 명령어 실행
func CreateUser(username, password string) error {
	return ExecSudoCommand("useradd -m -p " + password + " " + username)
}

// sudo 권한으로 사용자 sudo 권한 부여
func AddSudo(username string) error {
	return ExecSudoCommand("usermod -aG sudo " + username)
}
