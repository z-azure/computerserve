package main
import (
	"fmt"
	"os"
	"io"
	"os/exec"
	"io/ioutil"
	"strconv"
	//"flag"
	"bufio"
	//"github.com/spf13/pflag"
)
type put_in struct {
	start int
	end int
	leng int
	type_ int
	out string
	filename string
}

var putin put_in
/*
func getarg(args *put_in) {
	pflag.IntVarP(&(args.start), "start","s",-l,"get start page")
	pflag.IntVarP(&(args.end), "end","e",-l,"get end page")
	pflag.IntVarP(&(args.leng),"length","l",30,"get length")
	pflag.StringVarP(&(args.out),"outfilename","d","","get outfilename")
	pflag.BoolVarP(&(args.type_),"pagetype","f",false,"get pagetype")
	pflag.Parse()

	file_ := pflag.Args()
	if len(file_) > 0 {
		args.filename = string(file_[0])
	} else {
		args.filename = ""
	}
}*/

func main() {
	//fmt.Println("wait for input")
	//var args put_in
	//getarg(&putin)
	args := os.Args
	putin.start = 1
	putin.end = 1
	putin.filename = ""
	putin.out = ""
	putin.leng = 20
	putin.type_ = 'l'
	getputin(args)
	run()
}

func error_() {
	fmt.Fprintf(os.Stderr, "command error\n")
	os.Exit(1)
}
/*
func Usage() {

}

func getputin(args []string) {
	progname := "asd"
	//参数数量不够
	if len(args) < 3 {
		fmt.Fprintf(os.Stderr, "%s: not enough arguments\n", progname)
		Usage()
		os.Exit(1)
	}
	//处理第一个参数
	if args[1][0] != '-' || args[1][1] != 's' {
		fmt.Fprintf(os.Stderr, "%s: 1st arg should be -sstart_page\n", progname)
		Usage()
		os.Exit(1)
	}
	//提取开始页数
	sp, _ := strconv.Atoi(args[1][2:])
	if sp < 1 {
		fmt.Fprintf(os.Stderr, "%s: invalid start page %d\n", progname, sp)
		Usage()
		os.Exit(1)
	}
	putin.start = sp
	//处理第二个参数
	if args[2][0] != '-' || args[2][1] != 'e' {
		fmt.Fprintf(os.Stderr, "%s: 2nd arg should be -eend_page\n", progname)
		Usage()
		os.Exit(1)
	}
	//提取结束页数
	ep, _ := strconv.Atoi(args[2][2:])
	if ep < 1 || ep < sp {
		fmt.Fprintf(os.Stderr, "%s: invalid end page %d\n", progname, ep)
		Usage()
		os.Exit(1)
	}
	putin.end = ep

	//其他参数处理
	argindex := 3
	argcount := len(args)
	for {
		if argindex > argcount-1 || args[argindex][0] != '-' {
			break
		}
		switch args[argindex][1] {
		case 'l':
			//获取一页的长度
			pl, _ := strconv.Atoi(args[argindex][2:])
			if pl < 1 {
				fmt.Fprintf(os.Stderr, "%s: invalid page length %d\n", progname, pl)
				Usage()
				os.Exit(1)
			}
			putin.leng = pl
			argindex++
		case 'f':
			if len(args[argindex]) > 2 {
				fmt.Fprintf(os.Stderr, "%s: option should be \"-f\"\n", progname)
				Usage()
				os.Exit(1)
			}
			putin.type_ = 'f'
			argindex++
		case 'd':
			if len(args[argindex]) <= 2 {
				fmt.Fprintf(os.Stderr, "%s: -d option requires a printer destination\n", progname)
				Usage()
				os.Exit(1)
			}
			putin.out = args[argindex][2:]
			argindex++
		default:
			fmt.Fprintf(os.Stderr, "%s: unknown option", progname)
			Usage()
			os.Exit(1)
		}
	}

	if argindex <= argcount-1 {
		putin.filename = args[argindex]
	}
}*/

func getputin(args []string) {
	if len(args) < 3 {
		error_()
	}
	if args[1][0] != '-' || args[1][1] != 's' {
		error_()
	}

	begin, _ := strconv.Atoi(args[1][2:])
	if begin < 1 {
		error_()
	}
	putin.start = begin

	if args[2][0] != '-' || args[2][1] != 'e' {
		error_()
	}
	end, _ := strconv.Atoi(args[2][2:])
	if end < 1 || end < begin {
		error_()
	}
	putin.end = end

	count := len(args)
	num := 3
	for {
		if num > count - 1 || args[num][0] != '-' {
			break
		}
		command := args[num][1]
		if command == 'l' {
			s, _ := strconv.Atoi(args[num][2:])
			if s < 1 {
				error_()
			}
			putin.leng = s
			num++
			putin.type_ = 'l'
		} else {
			if command == 'f' {
				if len(args[num]) > 2 {
					error_()
				}
				putin.type_ = 'f'
				num++
			} else if command == 'd' {
				if len(args[num]) <= 2 {
					error_()
				}
				putin.out = args[num][2:]
				num++
			} else {
				error_()
			}
		}
	}
	if num <= count - 1 {
		putin.filename = args[num]
	}
}

func run() {
	var cmd *exec.Cmd
	var cmd_in io.WriteCloser
	var cmd_out io.ReadCloser
	line_num, page_num := 1, 1

	if putin.out != "" {
		cmd = exec.Command("bash", "-c", putin.out)
		cmd_in, _ = cmd.StdinPipe()
		cmd_out, _ = cmd.StdoutPipe()
		cmd.Start()
	}

	if putin.filename != "" {
		file, err := os.Open(putin.filename)
		if err != nil {
			error_()
		}

		f := bufio.NewReader(file)

		for {
			m, _, err := f.ReadLine()
			if err == io.EOF {
				break
			}
			if err != nil {
				error_()
			}
			if page_num >= putin.start && page_num <= putin.end {
				if putin.out == "" {
					fmt.Println(string(m))
				} else {
					fmt.Fprintln(cmd_in, string(m))
				}
			}
			line_num++
			if putin.type_ == 'l' {
				if line_num > putin.leng {
					line_num = 1
					page_num++
				}
			} else if string(m) == "\f" {
				page_num++
			}
			//fmt.Println("over1")
		}

		if putin.out != "" {
			cmd_in.Close()
			cmdBytes, err := ioutil.ReadAll(cmd_out)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Print(string(cmdBytes))
			cmd.Wait()
		}
	} else {
		read_m := bufio.NewScanner(os.Stdin)
		line_num = 1
		page_num = 1
		write_m := ""
		for read_m.Scan() {
			text := read_m.Text()
			text += "\n"
			if page_num >= putin.start && page_num <= putin.end {
				write_m += text
			}
			line_num++
			if putin.type_ == 'l' {
				if line_num > putin.leng {
					line_num = 1
					page_num++
				}
			} else if string(text) == "\f" {
				page_num++
			}
			//fmt.Println("over2")
		}
		if putin.out == "" {
			fmt.Print(write_m)
		} else {
			fmt.Fprint(cmd_in, write_m)
			cmd_in.Close()
			cmdBytes, err := ioutil.ReadAll(cmd_out)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Print(string(cmdBytes))
			cmd.Wait()
		}
	}
}

