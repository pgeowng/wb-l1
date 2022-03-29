package main

import (
	"fmt"
	"sync"
)

// Используем RWMutex, чтобы разрешить
// только одному записывать в переменную или
// многим читать из переменной
type MapMutex struct {
	sync.RWMutex
	data map[int]int
}

func (mm *MapMutex) Set(idx, val int) {
	mm.Lock()
	defer mm.Unlock()
	mm.data[idx] = val
}

func (mm *MapMutex) Get(idx int) (val int, ok bool) {
	mm.RLock()
	defer mm.RUnlock()
	val, ok = mm.data[idx]
	return
}

// Создаем конструктор, чтобы не забыть проинициализировать map.
// Так как по умолчанию переменная типа map[K]V является nil map,
// в которую нельзя записать новое значение по ключу.
//   var m map[int]int     // nil map[int]int
//   fmt.Println(m == nil) // true
//   m[1] = 2              // panic: assignment to entry in nil map
func NewMapMutex() *MapMutex {
	return &MapMutex{
		data: map[int]int{},
	}
}

// Для проверки race condition можно использовать go run -race
func main() {
	// Используем wg для ожидания исполнения всех горутин перед выходом.
	wg := sync.WaitGroup{}
	mm := NewMapMutex()

	// Вызов горутин для записи.
	// Без передачи индекса внутри функции будет сохранена reference на i.
	// Данная переменная постоянно меняется и в конце цикла будет равна 100.
	// Тем самым, внутри функции будет использоваться некорректное значение индекса.
	// Чтобы этого избежать, создадим в новую переменную с нужным значением.
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			//mm[i] = i
			mm.Set(i, i)
		}(i)
	}

	// Вызов горутин для чтения.
	// Чтобы не блокировать исполнение выводом данных, запустим еще горутину.
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			// val, ok := map[i]
			val, ok := mm.Get(i)

			wg.Add(1)
			go func() {
				defer wg.Done()
				fmt.Printf(" %2d | %2d | %v \n", i, val, ok)
			}()
		}(i)
	}

	wg.Wait()
}
