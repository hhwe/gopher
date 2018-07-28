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