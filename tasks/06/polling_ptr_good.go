package main

import (
	"fmt"
	"sync"
	"time"
)

// Используем mutex, чтобы исключить race condition.
// Недопускаем одновременное чтение и запись в переменную.
// В более сложном случае обращение к тем же данным безконтрольно
// может приводить к разным результатам выполнения.
type Stop struct {
	sync.Mutex
	stop bool
}

func (d *Stop) Done() {
	d.Lock()
	defer d.Unlock()
	d.stop = true
}

func (d *Stop) Cond() bool {
	d.Lock()
	defer d.Unlock()
	return d.stop
}

// Выход, если условие сработало
func routine(cond *Stop) {
	defer fmt.Println("stopped")
	fmt.Println("working")

	// постоянно вызывает функцию cond
	// решением может быть интервал между запросами
	for {
		if cond.Cond() {
			return
		}
	}
}

func main() {
	s := &Stop{}
	go routine(s)

	time.Sleep(time.Second)
	s.Done()
	time.Sleep(time.Second)
}
