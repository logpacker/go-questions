Each folder contains one question related to Golang, some of them have example code.

Rules:

* Each question has a separate folder and .txt file
* Some questions have a code example in .go file that should be pasted to the question
* Each question has only one correct answer
* All questions have options from *a* to *d*

Questions:

```
for i in */*.txt; do (echo "---"$i"---" && cat $i && echo ""); done
```


```---1/1.txt---
Q: What's the output of the following code?

a. 0 1 2 3 4 5
b. 0 1 2 0 1 2
c. 0 1 2 0 0 0
d. 0 0 0 0 1 2

A: c

---10/10.txt---
Q: What's correct verb should we use with fmt.Printf to print boolean?

a. %v
b. %t
c. %b
d. %s

A: b

---11/11.txt---
Q: Does Go support type inheritance?

a. Yes
b. No

A: b

---12/12.txt---
Q: Is it possible to declare multiple types of variables in single declaration in Go?

var a, b, c = 3, 4, "foo"

a. Yes
b. No

A: a

---13/13.txt---
Q: Is it possible to make content of the package to be accessed directly, without the need for it to be preceded with "fmt."?

a. Yes. import "fmt"
b. Yes. import _ "fmt"
c. Yes. import . "fmt"
d. No

A: c

---14/14.txt---
Q: Is it possible to have multiple tag strings of struct field?

a. Yes. They should be comma-separated: `key:"name",maxlength:"128"`
b. Yes. They should be space-separated: `key:"name" maxlength:"128"`
c. No

A: b

---15/15.txt---
Q: What is the output of the following code?

a. 0123
b. 3210
c. 0000
d. 3333

A: b

---16/16.txt---
Q: What characters are used by *go fmt* as indent?

a. 4 spaces
b. 2 spaces
c. tab character
d. it's configurable

A: a

---17/17.txt---
Q: Choose the correct statement about the output of the following code.

a. Output will be ordered by values
b. Output will be ordered by keys
c. Output will be ordered randomly
d. Output will be ordered by position in the code

A: c

---18/18.txt---
Q: How to change the value of GOMAXPROCS in Go?

a. Via environment variable GOMAXPROCS
b. On runtime in code
c. a and b
d. Not possible, equal to the number of CPUs available

A: c

---19/19.txt---
Q: Which of the following concatenation methods is slowest?

a. 1
b. 2
c. 3

A: a

---2/2.txt---
Q: What's the output of following code?

a. 7
b. 4
c. 3
d. invalid operation: a + b (mismatched types int8 and int16)

A: d

---20/20.txt---
Q: Where should we use defer in the following code?

a. 1
b. 2
c. 3
d. None of them

Q: 2

---21/21.txt---
Q: On what data types you can use "for - range" statement?

a. array, slice, map
b. array, slice, map, string
c. slice, map, string
d. slice, map

A: b

---22/22.txt---
Q: Does Go support optional parameters in functions?

a. Yes
b. No

A: b

---23/23.txt---
Q: When init() function will be called?

a. Before main() function in main package
b. For any non-main package after import
c. Only when you call it
d. a and b

A: d

---24/24.txt---
Q: Does Go have a ternary operator?

a. Yes
b. No

A: b

---25/25.txt---
Q: Is it possible to find that slices are equal with "==" operator?

a. Yes
b. No

A: b

---26/26.txt---
Q: What size of the following struct?

a. 12
b. 8
c: 24
d: 32

A: c

---27/27.txt---
Q: What's the default buffer size of channel in Go?

a. 0
b. 1
c. No default size

A: a

---28/28.txt---
Q: What are the default values for types: string, *string ?

a. "", nil
b. "", ""
c. nil, nil
d. nil, ""

A: a

---29/29.txt---
Q: Is it possible to define constant with array of type float32?

a. Yes
b. No

A: b

---3/3.txt---
Q: Short declaration ":=" can be used for defining global variables, isn't it?

a. Yes
b. No

A: b

---30/30.txt---
Q: In Go, does a break statement break from a select?

a. Yes
b. No

A: a

---4/4.txt---
Q: Which one of the following is correct?

a. const Pi = math.Pi
b. const Pi = 3.14
c. both a and b are correct
d. None of the above

A: b

---5/5.txt---
Q: Arrays are value types. In case of arrays as arguments, functions get their copies instead of a reference. isn't it?

a. Yes
b. No

A: a

---6/6.txt---
Q: Which of the following variables are exportable from another external package?

a. aName, BigBro
b. BigBro
c. aName, BigBro, 爱
d. BigBro, 爱

A: d

---7/7.txt---
Q: What's the output sequence of the following code?

a. 3 2 1
b. 1 2 3
c. 2 3 1
d. 2 1 3

A: d

---8/8.txt---
Q: How to build the following code to the binary file with custom name "eight"?

a. go build 8.go eight
b. go build -o eight 8.go
c. go build 8.go -o eight

A: b

---9/9.txt---
Q: Can we set DEBUG=true with go build?

a. Yes. go build -ldflags '-X main.DEBUG=true' 9/9.go
b. No

A: b
