# 学习笔记
### 1 基础语法
#### 1.1当标识符（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，如：Group1，那么使用这种形式的标识符的对象就可以被外部包的代码所使用（客户端程序需要先导入这个包），这被称为导出（像面向对象语言中的 public）；标识符如果以小写字母开头，则对包外是不可见的，但是他们在整个包的内部是可见并且可用的（像面向对象语言中的 private ）

#### 1.2 设置包别名：import fm "fmt"
```
package main

import fm "fmt" // alias3

func main() {
   fm.Println("hello, world")
}
```
##### 注意：如果你导入了一个包却没有使用它，则会在构建程序时引发错误

#### 1.3 包的分级声明和初始化 你可以在使用 import 导入包之后定义或声明 0 个或多个常量（const）、变量（var）和类型（type），这些对象的作用域都是全局的（在本包范围内），所以可以被本包中所有的函数调用（如 gotemplate.go 源文件中的 c 和 v），然后声明一个或多个函数（func）。
##### type有几种用法：定义结构体，定义接口， 类型别名， 类型定义， 类型开关

###### 定义结构体
```
//定义一个 Books结构体
type Books struct {
   title string
   author string
   subject string
   book_id int
}

//结构体内内嵌匿名成员变量定义
func main() {
   p := person{"abc",12}
   fmt.Println(p.string,p.int)
}

type person struct {
   string
   int
}
```
####得意于匿名嵌入的特性，我们可以直接访问叶子属性而不需要给出完整的路径：

```
var w Wheel
w.X = 8            // equivalent to w.Circle.Point.X = 8
w.Y = 8            // equivalent to w.Circle.Point.Y = 8
w.Radius = 5       // equivalent to w.Circle.Radius = 5
w.Spokes = 20
```
###### 定义接口

```
//定义电话接口
type Phone interface {
   call()
}
```
###### 自定义类型 ： 
type name string   // 使用 type 基于现有基础类型，结构体，函数类型创建用户自定义类型。 
```
type handle func(str string)  //自定义一个函数func，别名叫做handle，传入一个string参数
```
###### 类型开关
```
func classifier(items ...interface{})  {
   for i,x := range items {
      switch x.(type) {
      case bool:
         fmt.Printf("type #%d is bool",i)
      case float64:
         fmt.Printf("type #%d is float64",i)
      case string:
         fmt.Printf("type #%d is string",i)
      case int:
         fmt.Printf("type #%d is int",i)
      default:
         fmt.Printf("type is unknow")
      }
   }
}
```

#### 函数

```
func functionName()
```
##### 你可以在括号 () 中写入 0 个或多个函数的参数（使用逗号 , 分隔），_**每个参数的名称后面必须紧跟着该参数的类型**_。

##### **main 函数是每一个可执行程序所必须包含的**，一般来说都是在启动后第一个执行的函数（如果有 init() 函数则会先执行该函数）。如果你的 main 包的源代码没有包含 main 函数，则会引发构建错误 undefined: main.main。main 函数既没有参数，也没有返回类型（与 C 家族中的其它语言恰好相反）。如果你不小心为 main 函数添加了参数或者返回类型，将会引发构建错误

##### 程序正常退出的代码为 0 即 Program exited with code 0；如果程序因为异常而被终止，则会返回非零值，如：1。这个数值可以用来测试是否成功执行一个程序

#### 变量
##### 声明变量
```
var a string = "123"

//等于,两个不能同时存在 在相同的代码块中，我们不可以再次对于相同名称的变量使用初始化声明
a := "123"
``` 
##### 同一类型的多个变量可以声明在同一行，如：
```
var a, b, c int
```
##### 多变量可以在同一行进行赋值，如：
```
a, b, c = 5, 7, "abc"
```
##### 声明了变量但