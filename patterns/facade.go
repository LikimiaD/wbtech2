package patterns

import (
	"fmt"
)

/*
	Фасад - унифицированный интерфейс к набору интерфейса в подсистеме.
	Классы CPU, Memory и HardDrive представляют сложную подсистему с различными методами.
	Класс ComputerFacade объединяет сложные методы подсистемы в единый простой метод start, который выполняет последовательность действий.
	Создается объект ComputerFacade и вызывается метод start, который упрощает работу с подсистемой.

	Плюсы:
	Снижение сложности
	Инкапсуляция
	Ослабление связи

	Минусы:
	Снижение производительности
	Единая точка отказа
	Ограниченный функционал
*/

//? Подсистема CPU

type CPU struct {
	Active  bool
	Cores   int
	Threads int
}

func (c *CPU) start() {
	c.Active = true
	fmt.Println("CPU starts...")
}

func (c *CPU) freeze() {
	fmt.Println("CPU freezing...")
}

//? Подсистема Memory

type Memory struct {
	TotalMemory int
	UsageMemory int
	FreeMemory  int
}

func (m *Memory) load(position string) {
	fmt.Printf("Memory loading data at position %s\n", position)
}

//? Подсистема HardDisk

type HardDisk struct {
	Name  string
	Size  int
	Usage int
	Free  int
}

func (h *HardDisk) read(lba, size string) {
	fmt.Printf("HardDrive reading %s bytes from LBA %s\n", size, lba)
}

//? Фасад

type ComputerFacade struct {
	CPU
	Memory
	HardDisk
}

func getComputer() *ComputerFacade {
	cpu := CPU{Active: false, Cores: 8, Threads: 16}
	memory := Memory{TotalMemory: 16 * 1024, UsageMemory: 0, FreeMemory: 16 * 1024}
	hardDisk := HardDisk{Name: "SSD", Size: 510 * 1024, Usage: 0, Free: 510 * 1024}
	return &ComputerFacade{cpu, memory, hardDisk}
}

//? Метод Фасада для упрощения работы с подсистемой

func (c *ComputerFacade) startComputer() {
	c.freeze()
	c.Memory.load("0x00")
	c.start()
	c.HardDisk.read("1000", "4096")
}

func CheckFacade() {
	computer := getComputer()
	computer.startComputer()
}
