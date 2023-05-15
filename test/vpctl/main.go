package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
)

const (
	Separator = "***********"
)

func main() {
	// first, this test program reads a text file
	// each line is a filename to a sample file
	// then, the program reads the series of samples
	// the first line is the arguments to vpctl
	// the following lines are the stdout messages
	// multiple samples in the same file are separated with a separator

	byteArray, err := ioutil.ReadFile("./test/samples/samples.txt")
	if err != nil {
		panic(err)
	}
	fileListStr := string(byteArray)
	fileList := strings.Split(fileListStr, "\n")
	for _, filename := range fileList {
		filename = strings.TrimSpace(filename)
		if filename == "" {
			continue
		}
		if strings.HasPrefix(filename, "#") {
			continue
		}
		byteArray, err = ioutil.ReadFile(path.Join("./test/samples/", filename))
		if err != nil {
			panic(err)
		}
		str := string(byteArray)
		lines := strings.Split(str, "\n")
		lines = append(lines, Separator)
		cmd := ""
		stdout := ""
		state := 0
		// 0 -> init, expecting command
		// 1 -> stdout, reading expected stdout
		for _, line := range lines {
			if strings.HasPrefix(line, "#") {
				continue
			}
			if line == Separator {
				if cmd != "" {
					err = execute(filename, cmd, stdout)
					if err != nil {
						errMsg := fmt.Sprintf("test failed in file: %s, cmd: %s, err: %v", filename, cmd, err)
						panic(errMsg)
					}
				}
				state = 0
				cmd = ""
				stdout = ""
				continue
			}
			switch state {
			case 0:
				line = strings.TrimSpace(line)
				if line == "" {
					continue
				}
				cmd = line
				state = 1
				break
			case 1:
				line = strings.Trim(line, "\r")
				stdout += line + "\n"
				break
			default:
				panic("should not reach here: state=" + strconv.FormatInt(int64(state), 10))
			}
		}
	}
	fmt.Println("OK")
	os.Exit(0)
}

func execute(filename string, cmd string, expectedStdout string) error {
	fmt.Printf("[%s] testing %s\n", filename, cmd)
	c := exec.Command("../../vpctl", strings.Split(cmd, " ")...)
	c.Env = append(os.Environ(), "VPCTL_WORKING_DIRECTORY=/tmp/vpctl_test")
	c.Dir = "./test/samples"
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	c.Stdout = &stdout
	c.Stderr = &stderr
	err := c.Run()
	if err != nil {
		return fmt.Errorf("%v\n%s", err, stderr.String())
	}
	actualStdout := stdout.String()
	actualStdout = strings.ReplaceAll(actualStdout, "\r\n", "\n")
	if strings.TrimSpace(expectedStdout) != strings.TrimSpace(actualStdout) {
		return fmt.Errorf("\nexpecting: %v\n%s\n%s\nactual: %v\n%s", []byte(expectedStdout), expectedStdout, Separator, []byte(actualStdout), actualStdout)
	}
	return nil
}
