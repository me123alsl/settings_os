package commands

// 파이썬 3.8 설치
func InstallPython38(username string) error {
	return ExecSudoCommand(username, "sudo apt-get install python3.8")
}

// 파이썬 3.8 링크 생성 (overwrite)
func LinkPython38(username string) error {
	return ExecSudoCommand(username, "sudo ln -sf /usr/bin/python3.8 /usr/bin/python")
}

// "/usr/bin/python" 링크가 python3인지 확인
func CheckPythonLinkm(username string) error {
	return ExecSudoCommand(username, "python --version")
}

// 패키지 list 설치 yes로 모두 설치
func InstallPackageList(username string, packageList []string) error {
	for _, pkg := range packageList {
		err := ExecSudoCommand(username, "sudo apt-get install -y "+pkg)
		if err != nil {
			return err
		}
	}
	return nil
}
