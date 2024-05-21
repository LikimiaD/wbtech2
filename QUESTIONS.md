### Что выведет программа?

```go
package main

import (
    "fmt"
)

func main() {
    a := [5]int{76, 77, 78, 79, 80}
    var b []int = a[1:4]
    fmt.Println(b)
}
```

Ответ: 77, 78, 79

### Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и порядок их вызовов.

```go
package main

import (
    "fmt"
)

func test() (x int) {
    defer func() {
    x++
    }()
    x = 1
    return
}


func anotherTest() int {
    var x int
    defer func() {
    x++
    }()
    x = 1
    return x
}


func main() {
    fmt.Println(test())
    fmt.Println(anotherTest())
}
```

Ответ:
```cmd
2
1
```

Основное различие заключается в том, что в первой функции используется именованное возвращаемое значение, которое изменяется отложенным вызовом, тогда как во второй функции возвращаемое значение копируется до выполнения отложенного вызова.

### Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
    "fmt"
    "os"
)

func Foo() error {
    var err *os.PathError = nil
    return err
}

func main() {
    err := Foo()
    fmt.Println(err)
    fmt.Println(err == nil)
}
```

Ответ:

```cmd
<nil>
false
```

Так как сравнивается динамический тип и динамическое значение.

### Что выведет программа? Объяснить вывод программы.

```go
package main
 
func main() {
    ch := make(chan int)
    go func() {
        for i := 0; i < 10; i++ {
            ch <- i
        }
    }()
 
    for n := range ch {
        println(n)
    }
}
```

Deadlock, так как канал не закрывается, а for loop будет ждать результата пока не получит сигнал о закрытии.

### Что выведет программа? Объяснить вывод программы.

```go
package main
 
type customError struct {
     msg string
}
 
func (e *customError) Error() string {
    return e.msg
}
 
func test() *customError {
     {
         // do something
     }
     return nil
}
 
func main() {
    var err error
    err = test()
    if err != nil {
        println("error")
        return
    }
    println("ok")
}
```

Ответ: вернет error, так как динамический тип это указатель на customError

### Что выведет программа? Объяснить вывод программы. Рассказать про внутреннее устройство слайсов и что происходит при передаче их в качестве аргументов функции

```go
package main
 
import (
  "fmt"
)
 
func main() {
  var s = []string{"1", "2", "3"}
  modifySlice(s)
  fmt.Println(s)
}
 
func modifySlice(i []string) {
  i[0] = "3"
  i = append(i, "4")
  i[1] = "5"
  i = append(i, "6")
}
```

Ответ: изменится только те элементы, которые изначально были в capacity, а добавление ничего не изменит. Когда передаем в функцию slice то мы передаем SliceHeader, который знает где начинается массив и capacity, если мы сделаем `[:5]` то увидим остальные.

### Что выведет программа? Объяснить вывод программы.

```go
package main
 
import (
    "fmt"
    "math/rand"
    "time"
)
 
func asChan(vs ...int) <-chan int {
   c := make(chan int)
 
   go func() {
       for _, v := range vs {
           c <- v
           time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
      }
 
      close(c)
  }()
  return c
}
 
func merge(a, b <-chan int) <-chan int {
   c := make(chan int)
   go func() {
       for {
           select {
               case v := <-a:
                   c <- v
              case v := <-b:
                   c <- v
           }
      }
   }()
 return c
}
 
func main() {
 
   a := asChan(1, 3, 5, 7)
   b := asChan(2, 4 ,6, 8)
   c := merge(a, b )
   for v := range c {
       fmt.Println(v)
   }
}
```

Ответ: не правильный вывод, так как `merge` не обрабатывает закрытие канала.