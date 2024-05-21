package patterns

import "fmt"

/*
	Паттерн "Фабричный метод" определяет интерфейс для создания объектов,
	но позволяет подклассам изменять тип создаваемых объектов.

	Плюсы:
	Расширяемость
	Снижение зависимости
	Инкапсуляция создания объектов

	Минусы:
	Сложность кода
	Избыточность
*/

// ? Интерфейс для документов

type Document interface {
	open()
	save()
	close()
}

// ? Реализация текстового документа

type TextDocument struct{}

func (d *TextDocument) open() {
	fmt.Println("Opening text document")
}

func (d *TextDocument) save() {
	fmt.Println("Saving text document")
}

func (d *TextDocument) close() {
	fmt.Println("Closing text document")
}

// ? Реализация PDF документа

type PDFDocument struct{}

func (d *PDFDocument) open() {
	fmt.Println("Opening PDF document")
}

func (d *PDFDocument) save() {
	fmt.Println("Saving PDF document")
}

func (d *PDFDocument) close() {
	fmt.Println("Closing PDF document")
}

// ? Интерфейс фабрики документов

type DocumentFactory interface {
	createDocument() Document
}

// ? Фабрика текстовых документов

type TextDocumentFactory struct{}

func (f *TextDocumentFactory) createDocument() Document {
	return &TextDocument{}
}

// ? Фабрика PDF документов

type PDFDocumentFactory struct{}

func (f *PDFDocumentFactory) createDocument() Document {
	return &PDFDocument{}
}

func CheckFabric() {
	// ? Создаем фабрику текстовых документов
	var textFactory DocumentFactory = &TextDocumentFactory{}
	textDoc := textFactory.createDocument()
	textDoc.open()
	textDoc.save()
	textDoc.close()

	// ? Создаем фабрику PDF документов
	var pdfFactory DocumentFactory = &PDFDocumentFactory{}
	pdfDoc := pdfFactory.createDocument()
	pdfDoc.open()
	pdfDoc.save()
	pdfDoc.close()
}
