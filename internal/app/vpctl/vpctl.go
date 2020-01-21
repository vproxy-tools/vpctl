package vpctl

import (
	vp "../../../pkg/vproxy_config"
	"os"
)
import "../../pkg/versions"
import "fmt"

func Main(args []string) int {
	if len(args) == 0 {
		fmt.Println("vpctl: " + versions.VPCTL)
		return 0
	}

	cmd := args[0]
	switch cmd {
	case "help":
		fallthrough
	case "--help":
		fallthrough
	case "-help":
		fallthrough
	case "-h":
		fmt.Println(HelpStr)
		return 0
	case "version":
		fallthrough
	case "--version":
		fallthrough
	case "-version":
		fmt.Println(versions.VPCTL)
		return 0
	case "apply":
		if len(args) < 2 {
			_, _ = fmt.Fprintln(os.Stderr, "too few arguments for command 'apply'")
			return 1
		}
		second := args[1]
		if second != "-f" {
			_, _ = fmt.Fprintln(os.Stderr, "unknown option for command 'apply': "+second)
			return 1
		}
		if len(args) < 3 {
			_, _ = fmt.Fprintln(os.Stderr, "at least one filename should be provided")
			return 1
		}
		filenames := args[2:]
		for _, f := range filenames {
			todoList, err := vp.ApplyByFile(f)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "failed when processing file %s: %v\n", f, err)
				return 1
			}
			err = vp.RunTodoListAndPrint(todoList)
			if err != nil {
				_, _ = fmt.Fprintln(os.Stderr, err.Error())
				return 1
			}
		}
		return 0
	case "delete":
		if len(args) < 2 {
			_, _ = fmt.Fprintln(os.Stderr, "too few arguments for command 'delete'")
			return 1
		}
		second := args[1]
		switch second {
		case "-f":
			if len(args) < 3 {
				_, _ = fmt.Fprintln(os.Stderr, "at least one filename should be provided")
				return 1
			}
			filenames := args[2:]
			for _, f := range filenames {
				todoList, err := vp.DeleteByFile(f)
				if err != nil {
					_, _ = fmt.Fprintf(os.Stderr, "failed when processing file %s: %v\n", f, err)
					return 1
				}
				err = vp.RunTodoListAndPrint(todoList)
				if err != nil {
					_, _ = fmt.Fprintln(os.Stderr, err.Error())
					return 1
				}
			}
			return 0
		default:
			if len(args) < 3 {
				_, _ = fmt.Fprintln(os.Stderr, "too few arguments for command 'delete "+second+"', name should be provided")
				return 1
			}
			if len(args) > 3 {
				_, _ = fmt.Fprintln(os.Stderr, "too many arguments for command 'delete'")
				return 1
			}
			todoList, err := vp.DeleteOne(second, args[2])
			if err != nil {
				_, _ = fmt.Fprintln(os.Stderr, err.Error())
				return 1
			}
			err = vp.RunTodoListAndPrint(todoList)
			if err != nil {
				_, _ = fmt.Fprintln(os.Stderr, err.Error())
				return 1
			}
			return 0
		}
	case "get":
		if len(args) < 2 {
			_, _ = fmt.Fprintln(os.Stderr, "too few arguments for command 'get'")
			return 1
		}
		typ := args[1]
		format := "wide"
		name := ""
		if len(args) > 2 {
			third := args[2]
			if third == "-o" {
				if len(args) < 4 {
					_, _ = fmt.Fprintln(os.Stderr, "missing print format after '-o'")
					return 1
				}
				format = args[3]
				if format != "wide" && format != "yaml" {
					_, _ = fmt.Fprintln(os.Stderr, "unknown print format: "+format)
					return 1
				}
				if len(args) > 4 {
					_, _ = fmt.Fprintln(os.Stderr, "too many arguments for command 'get'")
					return 1
				}
			} else {
				name = third
				if len(args) > 3 {
					if args[3] != "-o" {
						_, _ = fmt.Fprintln(os.Stderr, "unknown option for command 'get "+typ+" "+name+"'")
						return 1
					}
					if len(args) < 5 {
						_, _ = fmt.Fprintln(os.Stderr, "missing print format after '-o'")
						return 1
					}
					format = args[4]
					if format != "wide" && format != "yaml" {
						_, _ = fmt.Fprintln(os.Stderr, "unknown print format: "+format)
						return 1
					}
					if len(args) > 5 {
						_, _ = fmt.Fprintln(os.Stderr, "too many arguments for command 'get'")
						return 1
					}
				}
			}
		}
		yaml := false
		if format == "yaml" {
			yaml = true
		}
		if name == "" {
			err := vp.UtilGetAndPrintList(typ, yaml)
			if err != nil {
				_, _ = fmt.Fprintln(os.Stderr, err.Error())
				return 1
			}
		} else {
			err := vp.UtilGetAndPrintOne(typ, name, yaml)
			if err != nil {
				_, _ = fmt.Fprintln(os.Stderr, err.Error())
				return 1
			}
		}
	default:
		_, _ = fmt.Fprintln(os.Stderr, "unknown command '"+cmd+"'")
	}
	return 0
}
