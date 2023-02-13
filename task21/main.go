package main

// Задание 21
// Реализовать паттерн «адаптер» на любом примере.

import "fmt"

// Шаблон проектирования адаптера поможет вам соответствовать потребностям двух частей кода, которые
// на первый взгляд несовместимы. Это ключ к тому, чтобы иметь в виду при принятии
// решения о том, подходит ли шаблон адаптера для вашей проблемы – два интерфейса, которые несовместимы, но которые
// должны работать вместе, являются хорошими кандидатами на шаблон адаптера (но они также могут использовать
// шаблон фасада, например).

// В нашем примере у нас будет старый интерфейс принтера и новый. Пользователи нового
// интерфейса не ожидают подписи, которая есть у старого, и нам нужен адаптер, чтобы
// пользователи все еще могли использовать старые реализации при необходимости (например, для работы с некоторым устаревшим кодом)

// Устаревший интерфейс, называемый Legacy Printer, имеет метод печати, который принимает строку и
// возвращает сообщение
type LegacyPrinter interface {
	Print(s string) string
}

type MyLegacyPrinter struct{}

// Наша структура MyLegacyPrinter реализует интерфейс
// LegacyPrinter и изменяет переданную строку, добавляя текстовый префикс Legacy Printer:
// После изменения текста структура MyLegacyPrinter печатает текст на консоли, а затем возвращает его.
func (l *MyLegacyPrinter) Print(s string) (newMsg string) {
	newMsg = fmt.Sprintf("Legacy Printer: %s\n", s)
	println(newMsg)
	return
}

// Теперь мы объявим новый интерфейс, который нам придется адаптировать:
// В этом случае новый метод Print Stored не принимает никакой строки в качестве аргумента,
// поскольку она должна быть сохранена в implementers заранее.
type ModernPrinter interface {
	PrintStored() string
}

// Как упоминалось ранее, адаптер PrinterAdapter должен иметь поле для хранения строки для печати.
// В нем также должно быть поле для хранения экземпляра адаптера LegacyPrinter
type PrinterAdapter struct{
	OldPrinter LegacyPrinter
	Msg string
}

func(p *PrinterAdapter) PrintStored() (newMsg string) {
	if p.OldPrinter != nil {
		newMsg = fmt.Sprintf("Adapter: %s", p.Msg)
		newMsg = p.OldPrinter.Print(newMsg)
	} else {
		newMsg = p.Msg
	}
   return
}

func main() {
	msg := "Hello World!"
	adapter := PrinterAdapter{
		OldPrinter: &MyLegacyPrinter{},
		Msg:        msg,
	}
	fmt.Println(adapter.PrintStored())
}
