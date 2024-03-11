Что выведет данная программа и почему?

```go
func main() {
  slice := []string{"a", "a"}

  func(slice []string) {
     slice = append(slice, "a")
     slice[0] = "b"
     slice[1] = "b"
     fmt.Print(slice)
  }(slice)
  fmt.Print(slice)
}
```

Ответ:
```
Программа выведет [b b a][a a] так как когда в анонимной функции происходит 
действие slice = append(slice, "a") создается новая копия в памяти с другим 
адресом и так как функция ничего не возвращает это будет актуально только в 
её границах.
```