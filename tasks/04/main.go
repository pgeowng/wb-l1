package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"time"
)

// Горутина, которая читает input пока может
// и в конце выполнения отправляет в exit свой id
func worker(id int, input <-chan int, exit chan<- int) {
	for input := range input {
		fmt.Printf("(%d): %d\n", id, input)
	}
	exit <- id
}

// Спавнер запускает N воркеров и сообщает о законченной работe.
// Учет завершенных воркеров ведется при помощи канала exitCount
// Таким образом, все данные будут обработаны.
func spawner(workerCount int, worker func(id int, exit chan<- int)) <-chan struct{} {
	processed := make(chan struct{})
	exitCount := make(chan int)
	for i := 0; i < workerCount; i++ {
		go worker(i, exitCount)
	}

	// Считаем, количество выполненных. По завершению сообщаем, что мы закончили
	go func() {
		defer close(exitCount)
		defer close(processed)
		count := 0
		for id := range exitCount {
			fmt.Printf("Worker %d just stopped\n", id)
			count++
			if count == workerCount {
				return
			}
		}
	}()

	return processed
}

func main() {
	// Можем задать количество воркеров, размер входного буфера и макс. задержку между входными данными
	workerCount := flag.Int("n", 5, "number of workers")
	bufferSize := flag.Int("b", 8, "size of input buffer")
	delayMs := flag.Int("d", 0, "delay between sends")
	flag.Parse()

	if *workerCount <= 0 {
		fmt.Println("worker count must be positive integer:", *workerCount)
		return
	}

	if *bufferSize < 0 {
		fmt.Println("buffer size must be non-negative integer:", *bufferSize)
		return
	}

	if *delayMs < 0 {
		fmt.Println("delay must be non-negative integer:", *delayMs)
		return
	}

	rand.Seed(time.Now().UnixNano())

	// Выход программы осуществляется через ожидание сигнала SIGINT от операционной системы.
	// При данном сигнале ожидается нормальное завершение работы программы.
	// По сигналу закрывается канал done.
	// Из-за того, что из закрытого канала можно бесконечно читать данные - нулевое значение типа канала,
	// становится удобно сообщать о произошедшем событии.
	// В данном случае, мы оповещаем о прекращении своей работы всех, кто заинтересован.
	done := make(chan struct{})
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	defer func() {
		signal.Stop(interrupt)
	}()
	go func() {
		<-interrupt
		fmt.Println("[Exiting...]")
		close(done)
	}()

	input := make(chan int, *bufferSize)
	processed := spawner(*workerCount, func(id int, exit chan<- int) {
		worker(id, input, exit)
	})

	// Чтобы поддержать инвариант цикла(получение задержки отправкой), создаем канал, который сразу закрыт.
	// Так, из delayChBase можно получить данные всегда.
	// Присваивание здесь необходимо из-за возвращаемого типа time.After() <-chan time.Time
	// При попытке использовать make(<-chan time.Time), нельзя закрыть канал, который только на чтение.
	delayChBase := make(chan time.Time)
	var delayCh <-chan time.Time = delayChBase
	close(delayChBase)

	// Постоянно отсылаем в input данные. Если вдруг смогли из done считать, значит приступаем к завершению.
	// 1. Для воркеров завершение работы происходит с закрытием канала input.
	//    Это связано с работой for range для каналов. При закрытии, происходит выход из цикла.
	//    Закрыт ли канал так же можно при помощи val, ok := <-ch, где ok bool
	// 2. Задержка реализована через канал, который по истечению времени можно считать.
	// 3. Для выходы из цикла используется метка mainloop.
	//    Это связано с тем, что break внутри select/for/switch выходит из блока, в котором он находится.
mainloop:
	for {
		select {
		case <-done:
			close(input)
			break mainloop
		case <-delayCh:
			val := rand.Intn(1000000000)
			fmt.Println("pushing", val)
			input <- val
		}
		if *delayMs > 0 {
			r := rand.Intn(*delayMs) + 1
			delayCh = time.After(time.Duration(r) * time.Millisecond)
		}
	}

	// Ожидаем заверщения всех воркеров и выходим.
	<-processed
	fmt.Println("Done!")
}
