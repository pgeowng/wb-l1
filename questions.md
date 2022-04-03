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
Если структура содержит 0 полей, соответственно размер 0.

+ [stackoverflow](https://stackoverflow.com/questions/2113751/sizeof-struct-in-go#2115688)
+ [go spec](https://go.dev/ref/spec#Size_and_alignment_guarantees)

```go
fmt.Printf("sizeof(struct{}{}) = %d\n", unsafe.Sizeof(struct{}{}))
// sizeof(struct{}{}) = 0
```

## 06

Есть ли в Go перегрузка методов или операторов?

В Go нет перегрузки операторов, то есть определение функций, которые будут вызываться при использовании +,-,\*.

Также нет перегрузки методов, то есть определение методов с одним и тем же именем, но разными аргументами

## 07

В какой последовательности будут выведены элементы map[int]int?

Пример:
```go
m[0]=1
m[1]=124
m[2]=281
```

Порядок неопределен. Можно считать, что случайный.
Таким образом, если добавлять элементы в map во время итерации, то можно не встретить их в следующих итерациях цикла.

[go spec/for statements with range clause/third entry in list](https://go.dev/ref/spec#For_statements)


## 08

В чем разница make и new?

+ new(T) \*T - выделяет тип Т в памяти и возвращает на него указатель.
+ make(T, n, m) T - специальная функция для выделения slice, map, channel. Так как для них initial value - nil

[go spec/allocation](https://go.dev/ref/spec#Allocation)
[go spec/making slice, maps and channels](https://go.dev/ref/spec#Making_slices_maps_and_channels)

## 09

Сколько существует способов задать переменную типа slice или map?


```go
// slice:
var s []int                   // 1

s := []int{}                  // 2

s := make([]slice)            // 3
s := make([]slice, len)
s := make([]slice, len, cap)

arr := [3]int{1,2,3}
s := arr[:]                   // 4

// map:
var a map[int]int             // 1
// nil map - нельзя добавить данные,
// a == nil -> true

b := map[int]int{}            // 2
c := make(map[int]int)        // 3
d := make(map[int]int, cap)
// при этом нельзя cap(d)
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
