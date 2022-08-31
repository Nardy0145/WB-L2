package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	cmd()
}

func cmd() {
	reader := bufio.NewReader(os.Stdin)
	for {
		dir, _ := os.Getwd()
		fmt.Printf("%s terminal>", dir)
		i, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error: invalid input")
		}
		i = strings.TrimSpace(i)
		code, err := handle(i)
		if err != nil {
			fmt.Println(err)
		}
		switch code {
		case 10:
			fmt.Println("stop executing...")
			time.Sleep(1 * time.Second)
			os.Exit(0)
		}
	}
}

func handle(str string) (int, error) {
	if str == "" {
		return 0, nil
	}
	args := strings.Split(str, " ")
	switch args[0] {
	case "exit":
		return 10, nil
	case "cd":
		if len(args) < 2 {
			return 0, fmt.Errorf("no path declared")
		}
		return 0, os.Chdir(args[1])
	case "echo":
		for i := 0; i < len(args); i++ {
			fmt.Println(args[i])
		}
		return 0, nil
	case "pwd":
		dir, _ := os.Getwd()
		fmt.Println("Current working directory is:\t", dir)
		return 0, nil
	case "ps":
		// Как я понял это невозможно сделать на Windows не прибегая к доп. модулям.
		// Реализую под Линукс.
		ex := exec.Command("ps")
		ex.Stderr = os.Stderr
		ex.Stdout = os.Stdout
		ex.Stdin = os.Stdin
		err := ex.Run()
		if err != nil {
			return 0, err
		}
	case "kill":
		if len(args) < 2 {
			return 0, fmt.Errorf("no process pid declared")
		}
		pid, err := strconv.Atoi(args[1])
		if err != nil {
			return 0, err
		}
		process, err := os.FindProcess(pid)
		if err != nil {
			return 0, err
		}
		fmt.Printf("Terminating process PID%d\n", process.Pid)
		err = process.Kill()
		if err != nil {
			return 0, err
		}
		return 0, nil
	case "exec":
		if len(args) < 2 {
			return 0, nil
		}
		ex := exec.Command(args[1])
		ex.Stderr = os.Stderr
		ex.Stdout = os.Stdout
		ex.Stdin = os.Stdin
		err := ex.Run()
		if err != nil {
			return 0, err
		}
		// Реализация под WINDOWS
	case "fork":
		if len(args) < 2 {
			return 0, nil
		}
		path := strings.Join(args[1:], " ")
		fmt.Println(path)
		ex := exec.Command("cmd.exe", "/C", "start", path)
		if err := ex.Run(); err != nil {
			log.Println("Error:", err)
		}
	}
	return 0, nil
}

// ps
