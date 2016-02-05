Preliminary information:

* Each question has a separate folder and .txt file
* Some questions have a code example in .go file
* Each question has only one correct answer
* All Questions have options from *a* to *d*

Test:

```
for i in */*.txt; do (cat $i && echo ""); done
```

##### What's the output of the following code?

```go
package main

import "fmt"

const (
	a = iota
	b = iota
	c = iota
)

const (
	d, e, f = iota, iota, iota
)

func main() {
	fmt.Println(a, b, c, d, e, f)
}
```

a. 0 1 2 3 4 5
b. 0 1 2 0 1 2
c. 0 1 2 0 0 0
d. 0 0 0 0 1 2

##### What verb should we use with fmt.Printf to print boolean?

a. %v
b. %t
c. %b
d. %s

##### Does Go support type inheritance?

a. Yes
b. No

##### Is it possible to declare multiple types of variables in single declaration in Go?

var a, b, c = 3, 4, "foo"

a. Yes
b. No

##### Is it possible to make package content directly accessible without need to be preceed by "fmt."?

a. Yes. import "fmt"
b. Yes. import _ "fmt"
c. Yes. import . "fmt"
d. No

##### Is it possible to have multiple tag strings in struct field?

```package main

import (
	"fmt"
	"reflect"
)

func main() {
	type User struct {
		Name string `key:"name"`
	}

	u := User{}
	st := reflect.TypeOf(u)
	field := st.Field(0)
	fmt.Println(field.Tag.Get("key"))
}
```

a. Yes. They should be comma-separated: `key:"name",maxlength:"128"`
b. Yes. They should be space-separated: `key:"name" maxlength:"128"`
c. No

##### What is the output of the following code?

```package main

import "fmt"

func main() {
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
}
```

a. 0123
b. 3210
c. 0000
d. 3333

##### What characters is *go fmt* command using for indent?

a. 4 spaces
b. 2 spaces
c. tab character
d. it's configurable

##### Choose the correct statement regarding the output of the following code.

```package main

import "fmt"

func main() {
	tagsViews := map[string]int{
		"unix":   0,
		"python": 1,
		"go":     2,
		"golang": 3,
		"devops": 4,
		"gc":     5,
	}
	for key, views := range tagsViews {
		fmt.Println("There are", views, "views for", key)
	}
}
```

a. Output will be ordered by values
b. Output will be ordered by keys
c. Output will be ordered randomly
d. Output will be ordered by position in the code

##### How can we change the value of GOMAXPROCS in Go?

a. Via environment variable GOMAXPROCS
b. In the code
c. a and b
d. Impossible as it's equal to the number of available CPUs

##### Which is the slowest concatenation method from the list?

```package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	// 1
	s := ""
	for i := 0; i < 5; i++ {
		s += "a"
	}
	fmt.Println(s)

	// 2
	var buffer bytes.Buffer
	for i := 0; i < 5; i++ {
		buffer.WriteString("a")
	}
	fmt.Println(buffer.String())

	// 3
	sl := []string{"a", "a", "a", "a", "a"}
	fmt.Printf(strings.Join(sl, ""))
}
```

a. 1
b. 2
c. 3

##### What's the output of the following code?

```go
package main

import "fmt"

func main() {
	var a int8 = 3
	var b int16 = 4

	sum := a + b

	fmt.Println(sum)
}
```

a. 7
b. 4
c. 3
d. invalid operation: a + b (mismatched types int8 and int16)

##### Where should we use *defer* in the following code?

```go
package main

import "os"

func main() {
	src, err := os.Open("filename")
	// 1
	//defer src.Close()
	if err != nil {
		return
	}

	// 2
	//defer src.Close()

	src.WriteString("Hello")

	// 3
	//defer src.Close()
}
```

a. 1
b. 2
c. 3
d. None of them

##### 2

##### On what data types you can use "for - range" statement?

a. array, slice, map
b. array, slice, map, string
c. slice, map, string
d. slice, map

##### Does Go support optional parameters in functions?

a. Yes
b. No

##### When will init() function be called?

a. Before main() function in main package
b. After importing a package with defined init() function
c. Only when you call it
d. a and b

##### Does Go have a ternary operator?

a. Yes
b. No

##### Is it possible to check if the slices are equal with "==" operator or not?

a. Yes
b. No

##### What is the size of the following struct?

```go
package main

import "fmt"
import "unsafe"

func main() {
	s := struct {
		A float32
		B string
	}{0, "go"}

	fmt.Printf("%T, %d\n", s, unsafe.Sizeof(s))
}
```

a. 12
b. 8
c: 24
d: 32

##### What's the default buffer size of the channel in Go?

a. 0
b. 1
c. No default size

##### What are the default values of these types: string, *string ?

a. "", nil
b. "", ""
c. nil, nil
d. nil, ""

##### Is it possible to define constant of an array type float32?

```package main

import "fmt"

const array = []float32{0.1, 0.2, 0.3}

func main() {
	fmt.Printf("%v", array)
}
```

a. Yes
b. No

##### Can short declaration ":=" be used for defining global variables?

a. Yes
b. No

##### In Go, does a *break* statement exit from a *select* block?

a. Yes
b. No

##### Which one of the following is correct?

a. const Pi = math.Pi
b. const Pi = 3.14
c. both a and b are correct
d. None of the above

##### Arrays are value types. In case of arrays perform as arguments, functions get their copies instead of a reference. Don’t they?

a. Yes
b. No

##### Which of the following variables are exportable from another external package?

```go
package main

var (
	aName string

	BigBro string

	爱 string
)

func main() {

}
```

a. aName, BigBro
b. BigBro
c. aName, BigBro, 爱
d. BigBro, 爱

##### What's the sequence for the output of the following code?

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("1")
		wg.Done()
	}()

	go func() {
		fmt.Println("2")
	}()

	wg.Wait()
	fmt.Println("3")
}
```

a. 3 2 1
b. 1 2 3
c. 2 3 1
d. 2 1 3

##### How to compile the following code to the binary file with name "eight"?

```go
package main

import "fmt"

func main() {
	fmt.Println("Hi")
}
```

a. go build 8.go eight
b. go build -o eight 8.go
c. go build 8.go -o eight

##### Can we set DEBUG=true with *go build*?

```go
package main

import (
	"fmt"
)

var DEBUG bool

func main() {
	fmt.Printf("DEBUG is %t\n", DEBUG)
}
```

a. Yes. go build -ldflags '-X main.DEBUG=true' 9/9.go
b. No

A: b
