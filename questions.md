## 09

Сколько существует способов задать переменную типа slice или map?

```go

// map:
var a map[int]int // nil map - нельзя добавить данные, a == nil -> true
b := map[int]int{}
c := make(map[int]int)     // по умолчанию задана какая-то cap
d := make(map[int]int, 23) // при этом нельзя cap(d)
```
