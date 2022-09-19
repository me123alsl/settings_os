package commands

import (
	"os/exec"
	"strings"

	utils "github.com/me123alsl/openmsa_installer/utils"
)

var Log = utils.Log

// sudo 권한으로 명령어 실행
func ExecSudoCommand(command string) error {
	cmd := exec.Command("sudo", "sh", "-c", command)
	return cmd.Run()
}

// 암호입력을 포함한 sudo 다른 계정 명령어 실행
func ExecSudoCommandWithPassword(username, password, command string) error {
	cmd := exec.Command("sudo", "-S", "su", "-", username, "-c", command)
	Log.Debug(cmd.String())
	cmd.Stdin = strings.NewReader(password)

	/* 입력결과*/
	output, err := cmd.CombinedOutput()
	if err != nil {
		Log.Info("output:", string(output))
	}
	return err
}

// 다른 사용자로 sudo 명령어 실행
func ExecSudoCommandAsOtherUser(username, password, command string) error {
	cmd := exec.Command("echo", "-n", password, "|", "sudo", "-S", "su", "-", username, "-c", command)
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

func ExecCommandWithOutput(command string) (string, error) {
	cmd := exec.Command("sh", "-c", command)
	out, err := cmd.Output()
	return string(out), err
}

// file exists
func FileExists(filename string) bool {
	_, err := exec.Command("sudo", "sh", "-c", "test -e "+filename).Output()
	return err == nil
}
