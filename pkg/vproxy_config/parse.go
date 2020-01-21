package vproxy_config

import (
	"bufio"
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"os"
	"strings"
)

func ParseFile(fileName string) ([]Config, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	//noinspection GoUnhandledErrorResult
	defer f.Close()
	reader := bufio.NewReader(f)
	ret := make([]Config, 0)
	s := ""
	parse := func() error {
		if strings.TrimSpace(s) == "" {
			s = ""
			return nil
		}
		base := Base{}
		err := yaml.Unmarshal([]byte(s), &base)
		if err != nil {
			return fmt.Errorf("parse the input file %s failed %s", fileName, err.Error())
		}
		err = base.Validate()
		if err != nil {
			return err
		}
		conf, err := NewEntity(base.Kind)
		if err != nil {
			return err
		}
		err = yaml.Unmarshal([]byte(s), conf)
		if err != nil {
			return fmt.Errorf("parse the input file %s failed %s", fileName, err.Error())
		}
		ret = append(ret, conf)
		s = ""
		return nil
	}
	for ; ; {
		lineBytes, _, err := reader.ReadLine()
		if err != nil && err != io.EOF {
			return nil, err
		}
		if lineBytes == nil {
			if err = parse(); err != nil {
				return nil, err
			}
			break
		}
		line := string(lineBytes)
		trim := strings.TrimSpace(line)
		if trim == "---" || trim == "..." {
			if err = parse(); err != nil {
				return nil, err
			}
			continue
		}
		s += line + "\n"
	}
	return ret, nil
}
