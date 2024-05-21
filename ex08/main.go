package main

import (
	"bufio"
	"fmt"
	"github.com/shirou/gopsutil/process"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"
)

var (
	currentDirectory string
)

func helpMenu() {
	clearTerminal()
	println("Help menu:")
	println("cd <args> -> смена директории")
	println("pwd -> показать путь до текущего каталога")
	println("echo <args> -> вывод аргумента в STDOUT")
	println("kill <args> -> завершить процесс по PID")
	println("clear -> очистка терминала")
	println("ls -> вывод названия папок в текущем каталоге")
	println("ps -> вывод информации о процессах")
}

func clearTerminal() {
	print("\033[H\033[2J")
}

func getCurrentPath() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("error while getting current directory: %s\n", err.Error())
		os.Exit(1)
	}
	currentDirectory = dir
}

func printCurrentProcess() {
	processes, err := process.Processes()
	if err != nil {
		fmt.Printf("Error fetching processes: %s", err.Error())
	}

	// Создаем новый табличный писатель
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintln(w, "PID\tTIME\tCMD")

	for _, p := range processes {
		pid := p.Pid
		name, err := p.Name()
		if err != nil {
			name = "Unknown"
		}

		createTime, err := p.CreateTime()
		var startTime time.Time
		if err == nil {
			startTime = time.Unix(createTime/1000, 0)
		}

		upTime := time.Since(startTime)
		upTimeFormatted := fmt.Sprintf("%02d:%02d:%02d", int(upTime.Hours()), int(upTime.Minutes())%60, int(upTime.Seconds())%60)

		fmt.Fprintf(w, "%d\t%s\t%s\n", pid, upTimeFormatted, name)
	}

	w.Flush()
}

func printCurrentDirectory() {
	println(currentDirectory)
}

func readDirectory() {
	entries, err := os.ReadDir(currentDirectory)
	if err != nil {
		fmt.Printf("error while reading current directory: %s\n", err.Error())
		return
	}

	for _, entry := range entries {
		if entry.IsDir() {
			println("Directory: ", entry.Name())
		} else {
			println("File: ", entry.Name())
		}
	}
}

func killProcess(pidString string) {
	pid, err := strconv.Atoi(pidString)
	if err != nil {
		fmt.Printf("error parsing pid to int: %s", err.Error())
		return
	}

	processes, err := process.Processes()
	if err != nil {
		fmt.Printf("error fetching processes: %s", err.Error())
	}

	for _, p := range processes {
		if p.Pid == int32(pid) {
			if err := p.Kill(); err != nil {
				fmt.Printf("error while killing process: %s", err.Error())
			}
			return
		}
	}

	println("process not found")
}

func changeDirectory(path string) {
	if path == "" {
		println("directory name can't be empty")
		return
	}

	var newPath string
	if filepath.IsAbs(path) {
		newPath = path
	} else {
		newPath = filepath.Join(currentDirectory, path)
	}

	if err := os.Chdir(newPath); err != nil {
		fmt.Printf("error while changing directory: %s\n", err.Error())
		return
	}
	getCurrentPath()
}

func echoPrint(info string) {
	println(info)
}

func executeCommand(command string, args []string) {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		fmt.Printf("error executing command: %s\n", err.Error())
	}
}

func main() {
	clearTerminal()
	getCurrentPath()

	scanner := bufio.NewScanner(os.Stdin)
LOOP:
	for {
		print("> ")
		scanner.Scan()
		text := scanner.Text()

		textArr := strings.Fields(text)
		if len(textArr) == 0 {
			continue
		}

		switch textArr[0] {
		case "pwd":
			printCurrentDirectory()
		case "help":
			helpMenu()
		case "ls":
			readDirectory()
		case "clear":
			clearTerminal()
		case "ps":
			printCurrentProcess()
		case "cd":
			if len(textArr) > 1 {
				changeDirectory(textArr[1])
			} else {
				fmt.Println("directory name can't be empty")
			}
		case "kill":
			if len(textArr) > 1 {
				killProcess(textArr[1])
			} else {
				fmt.Println("no pid")
			}
		case "echo":
			echoPrint(strings.Join(textArr[1:], " "))
		case "exit":
			break LOOP
		default:
			executeCommand(textArr[0], textArr[1:])
		}
	}
}
