/*
Эта простая программа читает файлы, указанные в
аргументах командной строки, и подсчитывает число вхождений
каждого найденного в них слова. В завершение она выводит
список слов, встречающихся более одного раза.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

type words struct { // извлекаемые слова помещаются в структуру
	sync.Mutex
	found map[string]int
}

func main() {
	var wg sync.WaitGroup // группа ожидания для мониторинга сопрограмм
	w := newWords()
	for _, f := range os.Args[1:] {
		wg.Add(1)
		go func(file string) {
			if err := tallyWords(file, w); err != nil {
				fmt.Println(err.Error())
			}
			wg.Done()
		}(f)
	}
	wg.Wait()
	fmt.Println("Words that appear more than once:") // вывод результатов поиска
	w.Lock()                                         // заблокировать объект, изменить и разблокировать в defer
	defer w.Unlock()
	for word, count := range w.found {
		if count > 1 {
			fmt.Printf("%s: %d\n", word, count)
		}
	}
}

func newWords() *words { // создание нового экземпляра слова
	return &words{found: map[string]int{}}
}

func (w *words) add(word string, n int) { // фиксирование количества входа слова
	w.Lock() // заблокировать объект, изменить и разблокировать в defer
	defer w.Unlock()
	count, ok := w.found[word]
	if !ok { // если слово не зафиксировано, то добавляем
		w.found[word] = n
		return
	}
	w.found[word] = count + n // в противном случае увеличим счетчик
}

func tallyWords(filename string, dict *words) error { // открытие файла анализ содержимого и подсчет слов
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file. error:%v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file) // инструмент сканирования
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		dict.add(word, 1)
	}
	return scanner.Err()
}

// go run race.go exampledata/*   or // go run race.go exampledata/example1/example2
