package commands

import (
	"os/exec"
	"strings"

	utils "github.com/me123alsl/settings_os/utils"
)

var Log = utils.Log

// sudo 권한으로 명령어 실행
func ExecSudoCommand(command string) error {
	cmd := exec.Command("sudo", "sh", "-c", command)
	return cmd.Run()
}

// 다른 사용자로 sudo 명령어 실행
func ExecSudoCommandAsOtherUser(username, command string) error {
	cmd := exec.Command("sudo", "-u", username, "sh", "-c", command)
	return cmd.Run()
}

// 미리 입력하여 sudo 명령어로 실행
func ExecSudoCommandWithInput(username, command, input string) error {
	cmd := exec.Command("sudo", "-u", username, "sh", "-c", command)
	cmd.Stdin = strings.NewReader(input)
	return cmd.Run()
}

// sudo 권한으로 명령어 실행 후 결과값 가져오기
func ExecSudoCommandWithOutput(command string) (string, error) {
	cmd := exec.Command("sudo", "-u", "root", "sh", "-c", command)
	out, err := cmd.Output()
	return string(out), err
}
