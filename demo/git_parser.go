package demo

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"syscall"
)

func GetExistedGitUrl(path string) string {
	if IsFileExists(path) {
		_, gitUrl, err := RunCmd("cd " + path + " && git remote get-url origin")
		if err!= nil {
			fmt.Println(err)
			return "";
		}
		return gitUrl
	}
	return ""
}

//can judge if a file or dir is existed
func IsFileExists(filePath string) bool {
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		return true
	}
	return false
}

//wait until command return
func RunCmd(cmd string) (pid int, output string, err error) {
	//out, err := exec.Command("sh", "-c", cmd).Output()
	command := exec.Command("sh", "-c", cmd)
	out, err := command.CombinedOutput()
	if err != nil {
		return 0, "", err
	}
	pid = command.Process.Pid
	status := command.ProcessState.Sys().(syscall.WaitStatus)
	exitStatus := status.ExitStatus()
	if exitStatus != 0 {
		err = errors.New("exit code: " + strconv.Itoa(exitStatus))
	}
	return pid, string(out), err
}

//run cmd under particular user's privilege
func RunCmdUser(cmd string, username string) (pid int, output string, err error) {
	//u, err := user.Lookup(username)
	if err != nil {
		return 0, "", err
	}
	//uid, err := strconv.Atoi(u.Uid)
	//gid, err := strconv.Atoi(u.Gid)
	command := exec.Command("sh", "-c", cmd)
	command.SysProcAttr = &syscall.SysProcAttr{}
	//Credential only can be used in non-windows enviroment
	//command.SysProcAttr.Credential = &syscall.Credential{Uid: uint32(uid), Gid: uint32(gid)}

	out, err := command.CombinedOutput()
	if err != nil {
		return 0, "", err
	}
	pid = command.Process.Pid
	status := command.ProcessState.Sys().(syscall.WaitStatus)
	exitStatus := status.ExitStatus()
	if exitStatus != 0 {
		err = errors.New("exit code: " + strconv.Itoa(exitStatus))
	}
	return pid, string(out), err
}
