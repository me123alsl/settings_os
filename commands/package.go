package commands

import (
	"strings"

	"github.com/me123alsl/settings_os/utils"
)

var RequrePackageList_client = []string{
	"openssh-client",
	"curl",
	"wget",
}

var RequrePackageList_server = []string{
	"openssh-server",
	"curl",
	"wget",
}

// 파이썬 3.8 설치
func InstallPython38() error {
	return ExecSudoCommand("apt-get install -y python3.8")
}

// 파이썬 3.8 링크 생성 (overwrite)
func LinkPython38() error {
	return ExecSudoCommand("ln -sf /usr/bin/python3.8 /usr/bin/python3")
}

// python 버전의 x.x 버전만 가져오기
func GetPythonVersion() (float32, error) {
	out, err := ExecSudoCommandWithOutput("python3 -c \"import sys; print('.'.join(map(str, sys.version_info[:2])))\"")
	if err != nil {
		return 0, err
	}

	return utils.StringToFloat32(strings.TrimSpace(out))
}

// 패키지 list 설치 yes로 모두 설치
func InstallPackageList(packageList []string) error {
	for _, pkg := range packageList {
		err := ExecSudoCommand("sudo apt-get install -y " + pkg)
		if err != nil {
			return err
		}
	}
	return nil
}

// client 패키지 설치
func InstallClientPackage() error {
	return InstallPackageList(RequrePackageList_client)
}

// server 패키지 설치
func InstallServerPackage() error {
	return InstallPackageList(RequrePackageList_server)
}

// ///////////////////////////////////////////////////////////////////////
func InstallPython3IfNeed() error {
	// python3이 존재하는지 확인
	if !utils.IsFileExist("/usr/bin/python3") {
		// python3이 없으면 설치
		err := InstallPython38()
		if err != nil {
			Log.Info("not exist file : /usr/bin/python3")
			Log.Info("error : ", err)
		}
	}

	// python3 버전 3.8이하면 python3.8 설치하고 링크 생성
	pythonVersion, err := GetPythonVersion()
	Log.Info("pythonVersion : ", pythonVersion)
	if err != nil {
		Log.Info("error : ", err)
	}

	if pythonVersion < 3.8 {
		err = InstallPython38()
		if err != nil {
			Log.Info(err)
		}
	}

	// /usr/bin/python3 링크 생성
	err = LinkPython38()
	if err != nil {
		return err
	}

	return nil
}
