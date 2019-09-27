package main

import (
	"fmt"
	"io"
)

type put_in struct{
	start int		//开始页数
	end int			//结束页数
	leng int		//每页长度
	type int		//分页方式
	out string		//输出文件名
	filename string	//输入文件名
}

var putin put_in

func main() {
	args := os.Args
	putin.start = 1
	putin.end = 1
	putin.filename = ""
	putin.out = ""
	putin.leng = 20 //默认20行一页
	putin.type = 'l'
	getputin(args)
	run()
}

func error() {
	fmt.Fprintf(os.Stderr, "command error\n")
}

func getputin(args []string) {
	if len(args) < 3 {
		error()
		os.Exit(1)
	}
	if args[1][0] != '-' || args[1][1] != 's' {
		error()
		os.Exit(1)
	}

	begin, err1 := strconv.Atoi(args[1][2:])//读取开始页数
	if begin < 1 || err1 != nil {
        error()
        os.Exit(1)
    }
    putin.begin = begin

    if args[2][0] != '-' || args[2][1] != 'e' {
    	error()
    	os.Exit(1)
    }
    end, err2 := strconv.Atoi(args[2][2:])//读取结束页数
    if end < 1 || end < begin || err2 != nil {
    	error()
    	os.Exit(1)
    }
    putin.end = end

    count := len(args)//获取指令长度
    num := 3
    for num = 3; num < count; num++ {
    	if args[num][0] != '-' {
			break
		}
		command := args[num][1]
		if command == 'l' {
			s, _ := strconv.Atoi(args[num][2:])
			if s < 1 {
				error()
				os.Exit(1)
			}
			putin.leng = s
			num++
		} else if command == 'f' {
			if len(args[num]) > 2 {
				error()
				os.Exit(1)
			}
			putin.type = 'f'
			num++
		} else if command == 'd' {
			if len(args[num]) <= 2 {
				error()
				os.Exit(1)
			}
			putin.out = args[num][2:]
			num++
		} else {
			error()
			os.Exit(1)
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
	line_num, page_num := 1, 1			//初始化页数和行数

	if putin.out != "" {
		cmd = exec.Command("bash", "-c", putin.out)
		cmd_in, _ = cmd.StdinPipe()
		cmd_out, _ = cmd.StdoutPipe()
		cmd.Start()
	}

	if putin.filename != "" {
		file, err := os.Open(putin.filename)
		if err != nil {
			error()
			os.Exit(1)
		}

		f := bufio.NewReader(file)

		for {
			m, _, err := f.ReadLine()
			if err == io.EOF {
				break
			}
			if err != nil {
				error()
				os.Exit(1)
			}
			if page_num >= putin.start && page_num <= putin.end {
				if putin.out == "" {
					fmt.Println(string(m))
				} else {
					fmt.Fprintln(cmd_in, string(m))		//读入文件中
				}
			}
			line_num++						//行数加一
			if putin.type == 'l' {
				if line_num > putin.leng {	//分页
					line_num = 1
					page_num++
				}
			} else if string(m) == "\f" {
				page_num++
			}
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
		write_m := ""
		for read_m.Scan() {
			text := read_m.Text()	//从输入读取
			if page_num >= putin.start && page_num <= putin.end {
				write_m += text
				write_m += "\n"
			}
			line_num++
			if putin.type == 'l' {
				if line_num > putin.leng {	//下一页
					line_num = 1
					page_num++
				}
			} else if string(text) == "\f" {
				page_num++
			}
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
			//等待command退出
			cmd.Wait()
		}
	}
}