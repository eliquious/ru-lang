
```
    package main

    import "io"
    import "fmt"
    import "net/http"

    func main()
        router := http.Default()
        router.GET("/ping", func(w http.ResponseWriter, req http.Request) {
            w.String(200, "pong")
        })
        router.Run(":8080")
    end

    func generic [T] () string
    end

    // Person represents a single individual.
    type Person struct
        Name string
        Age  uint8
    end

    enum Gender
        MALE, FEMALE
    end

    type Employee struct
        extends Person

        Company string
        Dept    string
    end

    greet := func(p Person) string {
        `Hi, ${p.Name}`
    }

    interface Introducer [T]
        Introduce(a, b T) string
    end

    impl interface Introducer [Person] for Person
        func Introduce(a, b Person)
            if len(a.Name) > 0 and len(b.Name) > 0
                "Hi, ${a.Name}. This is ${b.Name}."
            else if len(a.Name) > 0 and len(b.Name) == 0
                "Hi, ${a.Name}. Have you guys met before?"
            else if len(a.Name) == 0 and len(b.Name) > 0
                "Hey, this is ${b.Name}."
            else
                "Have you guys met before?"
            end
        end
    end

    max := Person{Name: "Max"}
    max.Introduce(Person{Name: "Thomas Edison"}, Person{Name: "Nikola Tesla"})

    macro PrintWhen [T] (T, T, string)

    impl macro PrintWhen [Number] (a, b Number, string)
        if a > b
            fmt.Println(msg)
        end
    end

    impl macro PrintWhen [string] (a, b, msg string)
        if len(a) > len(b)
            fmt.Println(msg)
        end
    end

    binaryop > (a, b, string) bool
        len(a) > len(b)
    end

    alias Number
        uint8, uint16, uint32, uint64,
        int8, int16, int32, int64,
        float32, float64,
    end

    alias uint uint32
    alias int  int32

    define macro times(n uint, *prog) as
        for i := 0; i < n; i++
            &prog
        end
    end

    greet2 := symbol {
        p := Person{Name: "Max"}
        p.Greet()
    }

    symbol greetMax
        p := Person{Name: "Max"}
        p.Greet()
    end

    times(5, @greetMax)
    times(5, @greet2)

    binaryop == (a, b Person) bool
        a.Name == b.Name
    end

    unaryop ++ (n Number) Number
        n + 1
    end

    func caller(done func(Error))
        done(Nil)
    end

    async caller(func(e Error) {
        fmt.Println(e.Error())
    })()

    // Channels

    c := make(chan string, 1)

    async func(input <-chan int)
        for var i int = 0, i < 5; ++i
            in <- i
        end
        err := close(input)
    end

    for i := range c
        fmt.Println(i)
    end

    type [T] List struct
        Array []T
    end

    alias [T] List as []T

    // Errors

    try
        // success
    then
        // success
    then
        // failure
    then
        // does not execute
    catch
        // handles error
    finally
        // after error handling
    end

```