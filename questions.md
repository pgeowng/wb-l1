## 01

Какой самый эффективный способ конкатенации строк?

[source](https://shantanubansal.medium.com/golang-how-to-efficiently-concatenate-strings-f2e51564f8d)

При известной конечной длине строки лучше использовать copy.
Иначе эффективнее использовать StringBuilder

```
$ go test -bench=.
  goos: linux
  goarch: amd64
  pkg: example.com/questions/questions/01
  cpu: AMD Ryzen 5 2600 Six-Core Processor
  BenchmarkConcat-12                        781464            167126 ns/op
  BenchmarkSprintF-12                       525726            184798 ns/op
  BenchmarkBuffer-12                      152405841                8.105 ns/op
  BenchmarkCopy-12                        260037604                4.228 ns/op
  BenchmarkStringBuilder-12               163439224                6.574 ns/op
  BenchmarkAppendStringArray-12            6657156               156.4 ns/op
  BenchmarkAppendStringArrayCap-12        53225012                19.40 ns/op
  PASS
  ok      example.com/questions/questions/01      236.771s

```

## 02

Что такое интерфейсы, как они применяются в Go?

Интерфейс - тип данных, который говорит, что с данной переменной может быть вызвано некоторое количество методов.

Интерфейсы является основным механизмом динамического полиморфизма.
Реализуя методы интерфейса, мы сообщаем, что тип может вести себя подобным образом.
Это удобно при организации кода и взаимозаменяемости компонентов.

## 03

Чем отличаются RWMutex от Mutex?

Mutex - разрешает только одному писать или читать.
RWMutex - разрешает только одному писать или сразу многим читать.

## 04

Чем отличаются буферизированные и небуферизированные каналы?

Запись в небуферизированный канал блокирует текущую в goroutine, если никто при этом не читает.
Для того, чтобы оттянуть процесс блокировки и продолжить дальшейшую работу, можно выделить буфер в котором будут храниться данные, ожидая обработки.


## 05

Какой размер у структуры struct{}{}?

Размер структуры зависит от типа данный, которые она содержит.


[source](https://stackoverflow.com/questions/2113751/sizeof-struct-in-go#2115688)

```go
fmt.Printf("sizeof(struct{}{}) = %d\n", unsafe.Sizeof(struct{}{}))
// sizeof(struct{}{}) = 0
```

## 09

Сколько существует способов задать переменную типа slice или map?

```go

// map:
var a map[int]int // nil map - нельзя добавить данные, a == nil -> true
b := map[int]int{}
c := make(map[int]int)     // по умолчанию задана какая-то cap
d := make(map[int]int, 23) // при этом нельзя cap(d)
```


## Good to know

+ [sizeof](https://stackoverflow.com/questions/2113751/sizeof-struct-in-go#2115688)

```
bool, int8/uint8 take 1 byte
int16, uint16 - 2 bytes
int32, uint32, float32 - 4 bytes
int64, uint64, float64, pointer - 8 bytes
string - 16 bytes (2 alignments of 8 bytes)
any slice takes 24 bytes (3 alignments of 8 bytes). So []bool, [][][]string are the same (do not forget to reread the citation I added in the beginning)
array of length n takes n * type it takes of bytes.
```
