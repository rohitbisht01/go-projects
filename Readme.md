✅ Go Concepts

1. Statically typed language, produces binary executables like C++.

2. Go has built in support for concurrency through goroutines and channels. Goroutine are lightweight threads managed by Go runtime, allowing for efficient concurrent execution of code. Channels faciliate communication and synchronization between goroutines, enabling developers to write highly concurrent programs without overhead typically associated with threading in other languages.

3. Go's garbage collector is designed to minimize pause times and memeory overhead. It uses a concurrent, tri-color mark and sweep algorithm, which allows garbage collection to be performed concurrently with the execution of Go code.

4. Cross platform (runs on linux/mac/windows). It means the ability to compile code to run on different os and architecture.

5. Companies that uses Golang: Dropbox, Uber, Alibaba, Paypal

6. Go module is to initialize ur project in go, its like a collection of packages, it defines all the dependencies which we install in our project.

7. Named return parameter and variadic functions (variadic function accpets variable number of arguments of a specific type, it is decalred by using an ellipsis as suffix)

8. Naming a function with uppercase letter meaning it is exported

9. Basic types and zero valuess

   - int (signed integer based on platform 32,64 bits)
   - int8,int16,int32, int64: signed integer
   - uint8,uint16: unsigned integer
   - float32, float64: floating point represnting single and doublee precision
   - complex128, complex64: complex number types with float32 nad float64 parts

   - boolean
   - string
   - composite type: array, slice, map,struct, pointer
   - special types: interface (defines a set of methods, allowing polymorphism), channel (used for commuinicating between goroutinges in concurrent programming)

   - zero values: for numberic types such as int, flaaot the zero value is 0,for boolean it is false, for strings it is "", for slices map and poniter funcitons and interfaces it is nil

10. Type conversion and type inference

    - type conversion: converting one data type into another one. in go we cannot do implicit type conversion so it is done expliicity.
      - numberic conversion
      - string conversion: strvconv.Itoa(), strconv.FormatFloat()
      - type assertion: to extract teh underlying value of interface
      - pointer conversion: pointers can be converted into another type
        type inference: in go intference means the compiler cna figure ouyt what type of vcariblee shoul;d be based on hte value you give it. eg. var x = 10 // here x is of type int var x = "helo" // here it is of type string

11. for loops

    ```
    simple looop
    for initial; condition; update {
        statement;
    }

    range based
    for index,value := range arr/slices/map{

    }

    go does not have the traditional while and do while loop
    ```

12. fmt package: printiing and formating in go

    - print
    - printf
    - println
    - type of => %t
    - pointer address => %p

13. defer statement

    - defer statement defers teh execution of a function until its surrounding function returns. defer function will not run immmediately instead it will run run just before the function in which defer statement is located exits
    - multiple defer statements ina function wiill be executed in LIFO order meaning the last defer statement will run first
    - use case: used for task need to be done at teh end of the fnctiion like closing a file or releasing a resource
    - even iif an error occurs in a function the deferred fucntion still be executed

14. arrays in go

    - var arr [size]datatype
    - arr := [size]datatype{values,values}
    - arrays in go are passed by value meaning when u assign one array tro another array each item from first array is copied to another array resulting two independent array

15. slices in go: dynamic arrays

    - dynamic in length (holds the pointer to underlying array)
    - var sliicess [] datatype => here it is pointint to nil and size = capacity = 0
    - sliice1 := [] int {1,3,4,5,6}
    - slice1 := arr[startInd:endInd], including startInd and exluding endInd
    - we cna use make fucntion to create a slice => make([]datatype,length,capacity) here capacity is optional
    - diff between capacity and size : size is the actual number of elements in an slice and capacity is the total number of element that the slice can hold before needing to grow. And once u hit the capacity it will increase its capacity to twice its capacity like 5 to 10 and 10 to 20
    - diff between array and slice => fixed size vs dyanmic size, length and capacity are fixied for array adn not for slices, memory allocation for array happens at compile time adn in slice memeory is reallocated dynamically as the slice grows

16. structs

    - a composite data type that groups together variables (fields) under a single name.
    - type struct name {field datatype field datatype}
    - we can access and modify the fields of struct
    - struct methods like func (u User) Greet(){} => here User is struct within this function u can access u User Struct
    - nested structs : composing one struct type into another one
    - anonymuos structs : structs creted without names liike address := struct{first, last string }{ value1 ,valuee2}
    - comparing two structs: return true or false value

17. maps in go

    - unordered collection of key value pairs => var m map[string]int or make(map[string]int)
    - maps are reference types meaning u can pass map to a funcaiton and it behaves like a pointer, changes in the fucntion affecets the original map
    - delete a ket using delete(map_name,key)
    - length of map => len(map)
    - nil maps vs empty map => initializing map in simple manner will be a nil map as assigning somehtying will give u an error, use make function oto create map

18. Pointers in go

    - pointer is a variable that stores memory address of another variable

    ```
    x := 10
    p := &x // p is a pointer to x

    fmt.Println(x)  // 10
    fmt.Println(p)  // address, e.g., 0xc000018090
    fmt.Println(*p) // 10 (value at that address)
    ```

    - &x => address of x, \*p means value at the address stored in p
    - why use pointers => to modify variables inside a function. without a pointer go passses by value so u only change the copy not the original. Seecond thing is performance => Passing a pointer is cheaper than copying large structs or slices.

19. continue, break and goto:

    - continue: skip to the next iteration
    - break: exits the loop
    - goto: jump to a label elsewhere in the function

20. Select statement in go

    - select lets goroutine wait on mutliple channel operations. Its like a switch for channels.
    - It picks the first available case that can proceed — if multiple cases are ready, one is chosen randomly.
    - when to use select : Listening on multiple channels at once, implenting timeouts, Writing non-blocking channel logic, cancelling goroutines.

21. channels in go

    - chan datatype
    - synchronous by default (sender waits until the receiver is ready)
    - channel can be send-only,receive-only or bidirectional
    - channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine.
    - there are two operation: send and receive
      - send operation: (mychannel <- element) sends data from one goroutine to another goroutine with the help of channel. Values like int, float and bool are safet and easy to send it through channel as they are copied so no risk of accidental concurrent access of the same value, strings are safe as they are immutable. But sending pointers or referenec like slice, map are not safe as the value of pointer or reference may change by sending goroutine or by receiving goroutine at the same time and result is unpredictable. So when sending pointers or refrence in channel make sure only one goroutine can access at a time
      - receive operation (element := <-mychannel or when the result is of no use do this <-mychannel) recieves the data sent by the send operator.
    - close() function to close a channel
    - check if a channel is open or close by this => ele, ok:= <- Mychannel
    - we cna iterate over channel using range method
    - channel types:
      - Unbuffered channel (default) : there is no internal storage so it requires the sender and receiver to be ready at the same time to transfer a value. there is strict sync between goroutines
      - Buffered channel : it has internal storage, so the sender can send values without waiting, and the receiver can receive later, up to the buffer's capacity.
      - Directional channel : the channel can be used to both send and receive values (Used commonly when you create and use a channel in the same function or pass it to another function for both read/write.)
      - send-only:
      - receive-only:

22. goroutine in go

    - goroutine leets ur functions run concurrently (concurrent meaning task are started and managed at the same time)
    - goroutiner is a light weight thread managed by go runtime
    - cheaper than os threads
    - in simple terms, You can run multiple functions at the same time, independently, without waiting for one to finish before starting the other.
    - goroutines lets u run multiplee functions at the same time (concurrently), which help build fast, non-blocking programs like servers, background workeers and API handlers etc
    - goroutines lifecycle => created, running, blocked, terminated
    - resource management strategies for goroutine: explicit termination using context,channel based termination
    - concurrency patterns for managing goroutine resources: worker pool pattern, fan-in/fan-out pattern, semaphore pattern

23. Context in go
    https://medium.com/@jamal.kaksouri/the-complete-guide-to-context-in-golang-efficient-concurrency-management-43d722f6eaea

    - helps in managing concurrent operations in go.
    - With context, you can create a hierarchy of goroutines and pass important information down the chain.
    - context provides a mechanism to control the lifecycle, cancellation, and propagation of requests across multiple goroutintes
    - use casees :
      - managing concurrent API requests: Consider a scenario where you need to fetch data from multiple APIs concurrently. By using context, you can ensure that all the API requests are canceled if any of them exceeds a specified timeout.
      - timeout and deadlines: avoid infinite waits and sla's
      - requesting scopes: propagation of values (context allows passing request-scoped values like user authentication info, request ids etc across api boundaries and goroutines) and consistency (maintain consistent state across diff parts of a program by sharing relevant data through context)
    - creating a context: to create context use context.Backgorund() that returs empty, non cancellable context as the root of context tree. we can also create context with specific timeout or deadline using context.WithTimeout() and context.WithDeadline()
    - propagating context: we can propagate it to downstream functions or goroutines by passing it as an argument. allowing relatded operations to share the same context
    - retrieve values from context
    - cancelling context: allows you to gracefully terminate operations and propagate cancellation signals to related goroutines. By canceling a context, you can avoid resource leaks and ensure the timely termination of concurrent operations.
    - timeout and deadlines: It ensures that operations complete within a specified timeframe and prevents potential bottlenecks or indefinite waits.
    - context in https: It allows you to control request cancellation, timeouts, and pass important values to downstream handlers.
    - context in database: allows you to manage query cancellations, timeouts, and pass relevant data within the database transactions

24. OS package in go

    - getting env variables
    - creating directory
    - create a file
    - read/ write a file
    - remove a file etc

===================================

1. allocaiton with new keyword and make function:
   - neew keyword:
     - allocates memory for any type and returns a pointer to that type
     - use when u wnat zero-initialized values and pointer to it
     - applies to struct, ints, arrays etc
   - make function:
     - allocates and initailizes memory for slice, map or channel type
     - use it when u want to create and initialize map, slice or channel
     - returns the value itself and not pointers
     - applies to only slice, map and channel

https://boldlygo.tech/categories/spec/
