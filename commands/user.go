package commands

// sudo 권한으로 사용자 생성 명령어 실행
func CreateUser(username, password string) error {
	return ExecSudoCommand(username, "useradd -m -p "+password+" "+username)
}

// sudo 권한으로 사용자 sudo 권한 부여
func AddSudo(username string) error {
	return ExecSudoCommand(username, "usermod -aG sudo "+username)
}
