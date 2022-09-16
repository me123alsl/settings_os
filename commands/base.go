package commands

import (
	"os/exec"
	"strings"
)

// sudo 권한으로 명령어 실행
func ExecSudoCommand(username, command string) error {
	cmd := exec.Command("sudo", "-u", username, "sh", "-c", command)
	return cmd.Run()
}

// 미리 입력하여 sudo 명령어로 실행
func ExecSudoCommandWithInput(username, command, input string) error {
	cmd := exec.Command("sudo", "-u", username, "sh", "-c", command)
	cmd.Stdin = strings.NewReader(input)
	return cmd.Run()
}
