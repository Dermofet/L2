package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/mitchellh/go-ps"
)

var (
	errCd   = errors.New("cd должен иметь 1 аргумент")
	errPwd  = errors.New("pwd не должен иметь аргументов")
	errEcho = errors.New("echo должен иметь 1+ аргумент")
	errKill = errors.New("kill должен иметь 1+ аргумент")
	errExec = errors.New("exec должен иметь 1+ аргумент")
	errPs   = errors.New("ps не должен иметь аргументов")
)

func main() {
	sh := NewShell(os.Stdout, os.Stdin)
	err := sh.Run()
	if err != nil {
		return
	}
}

// Shell представляет собой основную структуру программы с конфигурацией.
type Shell struct {
	Out      io.Writer
	In       io.Reader
	Pipe     bool
	PipeBuff *bytes.Buffer
}

// NewShell создает новый экземпляр Shell с заданными выходным и входным потоками.
func NewShell(w io.Writer, r io.Reader) *Shell {
	return &Shell{Out: w, In: r}
}

// Run запускает оболочку.
func (s *Shell) Run() error {
	if err := s.GetLines(); err != nil {
		if _, err := fmt.Fprintln(s.Out, err); err != nil {
			log.Fatalln(err)
		}
	}
	return nil
}

// cd изменяет текущую директорию.
func (s *Shell) cd(arg string) error {
	err := os.Chdir(arg)
	if err != nil {
		return err
	}
	return nil
}

// pwd выводит полный путь текущей директории.
func (s *Shell) pwd() error {
	path, err := os.Getwd()
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(s.Out, path)
	if err != nil {
		return err
	}
	return nil
}

// echo выводит аргументы в выходной поток.
func (s *Shell) echo(args []string, fullLine string) error {
	var line string
	start := args[0]
	end := args[len(args)-1]

	if start[0] == '"' && end[len(end)-1] == '"' {
		line = strings.TrimPrefix(fullLine, "echo ")
		line = strings.TrimLeft(line, `"`)
		line = strings.TrimRight(line, `"`)
	} else {
		line = strings.Join(args, " ")
	}

	_, err := fmt.Fprintln(s.Out, line)
	return err
}

// kill завершает процесс по его ID.
func (s *Shell) kill(pid []string) []error {
	var errs []error
	for _, value := range pid {
		id, err := strconv.Atoi(value)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		process, err := os.FindProcess(id)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		if err := process.Signal(os.Kill); err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

// ps выводит список запущенных процессов.
func (s *Shell) ps() error {
	processList, err := ps.Processes()
	if err != nil {
		return err
	}
	for _, process := range processList {
		_, err := fmt.Fprintf(s.Out, "%v\t%v\t%v\n", process.Pid(), process.PPid(), process.Executable())
		if err != nil {
			return err
		}
	}
	return nil
}

// GetLines считывает строки из входного потока.
func (s *Shell) GetLines() error {
	src := bufio.NewScanner(s.In)
	fmt.Fprint(s.Out, "$ ")
	for src.Scan() && (src.Text() != `\quit`) {
		line := src.Text()
		err := s.Fork(line)
		if err != nil {
			return err
		}
		fmt.Fprint(s.Out, "$ ")
	}
	if src.Err() != nil {
		os.Exit(0)
	}
	return nil
}

// Exec поддерживает выполнение exec-команд.
func (s *Shell) Exec(line []string) error {
	cmd := exec.Command(line[0], line[1:]...)
	cmd.Stdout = s.Out
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// CaseShell выбирает команду для выполнения.
func (s *Shell) CaseShell(line string) error {
	commandAndArgs := strings.Fields(line)
	if len(commandAndArgs) != 0 {
		switch commandAndArgs[0] {
		case "cd":
			if len(commandAndArgs) == 2 {
				err := s.cd(commandAndArgs[1])
				if err != nil {
					_, err := fmt.Fprintln(s.Out, err)
					if err != nil {
						return err
					}
				}
			} else {
				return errCd
			}
		case "ps":
			if len(commandAndArgs) == 1 {
				if err := s.ps(); err != nil {
					return err
				}
			} else {
				return errPs
			}
		case "pwd":
			if len(commandAndArgs) != 1 {
				return errPwd
			}
			err := s.pwd()
			if err != nil {
				return err
			}

		case "echo":
			if len(commandAndArgs) != 1 {
				err := s.echo(commandAndArgs[1:], line)
				if err != nil {
					return err
				}
			} else {
				return errEcho
			}
		case "kill":
			if len(commandAndArgs) != 1 {
				errs := s.kill(commandAndArgs[1:])
				for _, err := range errs {
					fmt.Fprintln(s.Out, err)
				}
			} else {
				return errKill
			}
		case "exec":
			if len(commandAndArgs) != 1 {
				err := s.Exec(commandAndArgs[1:])
				if err != nil {
					return err
				}
			} else {
				return errExec
			}
		default:
			if _, err := fmt.Fprintf(s.Out, "неизвестная команда '%v'\n", commandAndArgs[0]); err != nil {
				return err
			}
		}
	}
	return nil
}

// CheckPipes проверяет строку на наличие пайпов.
func (s *Shell) CheckPipes(line string) error {
	strCmd := strings.Split(line, "|")
	if len(strCmd) > 1 {
		s.Pipe = true
		s.PipeBuff = new(bytes.Buffer)
		for index, value := range strCmd {
			if index != 0 {
				comm1 := strings.Fields(value)
				if len(comm1) > 1 {
					comm1New := make([]string, 2)
					comm1New[0], comm1New[1] = comm1[0], s.PipeBuff.String()
					comm1 = comm1New
				} else {
					comm1 = append(comm1, s.PipeBuff.String())
				}
				value = strings.Join(comm1, " ")
			}
			s.PipeBuff.Reset()
			if index == len(strCmd)-1 {
				s.Pipe = false
			}
			if err := s.CaseShell(value); err != nil {
				fmt.Fprintln(s.Out, err)
			}
		}
	} else {
		if err := s.CaseShell(line); err != nil {
			fmt.Fprintln(s.Out, err)
		}
	}
	return nil
}

// Fork обрабатывает команду форка.
func (s *Shell) Fork(str string) error {
	str = strings.TrimRight(str, " ")
	if strings.Contains(str, "&") {
		cmd := exec.Command("cmd", "/C", str)
		err := cmd.Start()
		if err != nil {
			return err
		}
		fmt.Fprintf(s.Out, "[%v]\t%v\n", cmd.Process.Pid, str)
		go func() {
			err := cmd.Wait()
			if err != nil {
				fmt.Fprintln(s.Out, err)
			}
			fmt.Fprintf(s.Out, "[%v]+\tЗавершено\n", cmd.Process.Pid)
		}()
	} else {
		return s.CheckPipes(str)
	}
	return nil
}
