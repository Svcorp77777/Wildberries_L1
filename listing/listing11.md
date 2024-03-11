Что выведет данная программа и почему?

```go
func main() {
  wg := sync.WaitGroup{}
  for i := 0; i < 5; i++ {
     wg.Add(1)
     go func(wg sync.WaitGroup, i int) {
        fmt.Println(i)
        wg.Done()
     }(wg, i)
  }
  wg.Wait()
  fmt.Println("exit")
}
```

Ответ:
```
Программа выводит числа от 0 до 4 в случайном порядке и происходит 
deadlock, так как в анонимную функцию передается копия WaitGroup.
```