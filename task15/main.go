package main

// Задание 15
// К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.
// var justString string
// func someFunc() {
//   v := createHugeString(1 << 10)
//   justString = v[:100]
// }
// func main() {
//   someFunc()
// }

// Цитата из книги Jon Bodner:
// Хоть очень удобно, что Go позволяет использовать нотацию взятия среза для
// получения подстрок и индексную нотацию для извлечения отдельных элементов строки, и в том и в другом случае нужно быть очень внимательными.
// Поскольку строки являются неизменяемыми, они лишены проблем с модификацией элементов, свойственных срезам срезов. Однако в их случае имеется
// другая проблема: в то время как строка представляет собой последовательность байтов, размер кодовой точки в кодировке UTF-8 может составлять
// от одного до четырех байт. В нашем предыдущем примере все сработало,
// как ожидалось, потому что мы использовали исключительно кодовые точки
// кодировки UTF-8, длина которых равна одному байту. Однако при работе
// с текстами на других языках, помимо английского, или с символами эмоций
// вы будете иметь дело с кодовыми точками кодировки UTF-8, длина которых
// составляет больше одного байта:
// var s string = "Hello 🌞"
// var s2 string = s[4:7]
// var s3 string = s[:5]
// var s4 string = s[6:]
// В данном примере переменной s3, как и раньше, присваивается строка "Hello".
// Переменной s4 присваивается символ эмоций в виде солнца. Однако в переменную s2 вместо строки "o 🌞" заносится строка "o ".
// Это объясняется тем, что мы копируем только первый байт символа эмоций, что дает некорректный результат.
// var s string = "Hello 🌞"
// fmt.Println(len(s))
// Этот код выводит 10, а не 7, потому что для представления символа эмоций
// в виде улыбающегося солнца в кодировке UTF-8 требуется 4 байта.
// Для извлечения из строки подстрок и кодовых точек рекомендуется использовать не выражения среза и индекса,
// а функции из пакетов strings и unicode/ utf8 стандартной библиотеки.

import "golang.org/x/exp/utf8string"
func main() {
   a := utf8string.NewString("ÄÅàâäåçèéêëìîïü")
   s := a.Slice(1, 3)
   println(s == "Åà") // true
}





