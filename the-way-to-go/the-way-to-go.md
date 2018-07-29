# The way to go

### 常量

常量的值必须是能够在编译时就能够确定的；你可以在其赋值表达式中涉及计算过程，但是所有用于计算的值必须在编译期间就能获得。

    正确的做法：const c1 = 2/3
    错误的做法：const c2 = getNumber() // 引发构建错误: getNumber() used as value

因为在编译期间自定义函数均属于未知，因此无法用于常量的赋值，但内置函数可以使用，如：len()。

``` go
const (
    Sunday = iota   // 每次调用一次加一
    Monday          // Monday = iota
    Tuesday         // Tuesday = iota
    Wednesday = 1   // const中赋值为继承以前一个表达式
    Thursday        // Thursday = 1
    Friday          // Friday = 1
    Saturday        // Saturday = 1
)   // 0, 1, 2, 1, 1, 1
```

数字型的常量是没有大小和符号的，并且可以使用任何精度而不会导致溢出：


goto:

    特别注意 使用标签和 goto 语句是不被鼓励的：它们会很快导致非常糟糕的程序设计，而且总有更加可读的替代方案来实现相同的需求。

    一个建议使用 goto 语句的示例会在第 15.1 章的 simple_tcp_server.go 中出现：示例中在发生读取错误时，使用 goto 来跳出无限读取循环并关闭相应的客户端链接。

    定义但未使用标签会导致编译错误：label … defined and not used。

    如果您必须使用 goto，应当只使用正序的标签（标签位于 goto 语句之后），但注意标签和 goto 语句之间不能出现定义新变量的语句，否则会导致编译失败。



function:

    Go是编译型语言，所以函数编写的顺序是无关紧要的；鉴于可读性的需求，最好把 main() 函数写在文件的前面，其他函数按照一定逻辑顺序进行编写（例如函数被调用的顺序）。

    函数重载（function overloading）指的是可以编写多个同名函数，只要它们拥有不同的形参与/或者不同的返回值，在 Go 里面函数重载是不被允许的。这将导致一个编译错误：
        funcName redeclared in this book, previous declaration at lineno
    Go 语言不支持这项特性的主要原因是函数重载需要进行多余的类型匹配影响性能；

    目前 Go 没有泛型（generic）的概念，也就是说它不支持那种支持多种类型的函数。不过在大部分情况下可以通过接口（interface），特别是空接口与类型选择（type switch，参考 第 11.12 节）与/或者通过使用反射（reflection，参考 第 6.8 节）来实现相似的功能。使用这些技术将导致代码更为复杂、性能更为低下，所以在非常注意性能的的场合，最好是为每一个类型单独创建一个函数，而且代码可读性更强。

    如果你希望函数可以直接修改参数的值，而不是对参数的副本进行操作，你需要将参数的地址（变量名前面添加&符号，比如 &variable）传递给函数，这就是按引用传递，比如 Function(&arg1)，此时传递给函数的是一个指针。如果传递给函数的是一个指针，指针的值（一个地址）会被复制，但指针的值所指向的地址上的值不会被复制；我们可以通过这个指针的值来修改这个值所指向的地址上的值。`（*译者注：指针也是变量类型，有自己的地址和值，通常指针的值指向一个变量的地址。所以，按引用传递也是按值传递。*）`
    
    命名返回值作为结果形参（result parameters）被初始化为相应类型的零值，当需要返回的时候，我们只需要一条简单的不带参数的return语句。需要注意的是，即使只有一个命名返回值，也需要使用 () 括起来:
        func getX2AndX3_2(input int) (x2 int, x3 int) {
            x2 = 2 * input
            x3 = 3 * input
            // return x2, x3
            return
        }

    任何一个非命名返回值（使用非命名返回值是很糟的编程习惯）在 return 语句里面都要明确指出包含返回值的变量或是一个可计算的值（就像上面警告所指出的那样）。

    尽量使用命名返回值：会使代码更清晰、更简短，同时更加容易读懂。

    内存缓存的技术在使用计算成本相对昂贵的函数时非常有用（不仅限于例子中的递归），譬如大量进行相同参数的运算。这种技术还可以应用于纯函数中，即相同输入必定获得相同输出的函数。

array slice map:

    把一个大数组传递给函数会消耗很多内存。有两种方法可以避免这种现象：
        传递数组的指针
        使用数组的切片

new() 和 make() 的区别
    看起来二者没有什么区别，都在堆上分配内存，但是它们的行为不同，适用于不同的类型。

    new(T) 为每个新的类型T分配一片内存，初始化为 0 并且返回类型为*T的内存地址：这种方法 返回一个指向类型为 T，值为 0 的地址的指针，它适用于值类型如数组和结构体（参见第 10 章）；它相当于 &T{}。
    make(T) 返回一个类型为 T 的初始值，它只适用于3种内建的引用类型：切片、map 和 channel（参见第 8 章，第 13 章）。

切片和垃圾回收

    切片的底层指向一个数组，该数组的实际容量可能要大于切片所定义的容量。只有在没有任何切片指向的时候，底层的数组内存才会被释放，这种特性有时会导致程序占用多余的内存。

    示例 函数 FindDigits 将一个文件加载到内存，然后搜索其中所有的数字并返回一个切片。

    var digitRegexp = regexp.MustCompile("[0-9]+")

    func FindDigits(filename string) []byte {
        b, _ := ioutil.ReadFile(filename)
        return digitRegexp.Find(b)
    }
    这段代码可以顺利运行，但返回的 []byte 指向的底层是整个文件的数据。只要该返回的切片不被释放，垃圾回收器就不能释放整个文件所占用的内存。换句话说，一点点有用的数据却占用了整个文件的内存。

    想要避免这个问题，可以通过拷贝我们需要的部分到一个新的切片中：

    func FindDigits(filename string) []byte {
    b, _ := ioutil.ReadFile(filename)
    b = digitRegexp.Find(b)
    c := make([]byte, len(b))
    copy(c, b)
    return c
    }
    事实上，上面这段代码只能找到第一个匹配正则表达式的数字串。要想找到所有的数字，可以尝试下面这段代码：

    func FindFileDigits(filename string) []byte {
    fileBytes, _ := ioutil.ReadFile(filename)
    b := digitRegexp.FindAll(fileBytes, len(fileBytes))
    c := make([]byte, 0)
    for _, bytes := range b {
        c = append(c, bytes...)
    }
    return c
    }