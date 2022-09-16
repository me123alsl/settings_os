package main

func main() {

	// sudo 권한으로 명령어 실행
	ExecSudoCommand("root", "echo hello")
	// 다른 사용자로 명령어 실행
	ExecCommand("root", "echo hello")

}
