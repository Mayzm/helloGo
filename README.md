# 学习笔记
## 1 基础语法
### 1.1当标识符（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，如：Group1，那么使用这种形式的标识符的对象就可以被外部包的代码所使用（客户端程序需要先导入这个包），这被称为导出（像面向对象语言中的 public）；标识符如果以小写字母开头，则对包外是不可见的，但是他们在整个包的内部是可见并且可用的（像面向对象语言中的 private ）

###1.2 设置包别名：import fm "fmt"
```
package main

import fm "fmt" // alias3

func main() {
   fm.Println("hello, world")
}
```
#### 注意：如果你导入了一个包却没有使用它，则会在构建程序时引发错误

### 1.3包的分级声明和初始化 你可以在使用 import 导入包之后定义或声明 0 个或多个常量（const）、变量（var）和类型（type），这些对象的作用域都是全局的（在本包范围内），所以可以被本包中所有的函数调用（如 gotemplate.go 源文件中的 c 和 v），然后声明一个或多个函数（func）。
#### type有几种用法：定义结构体，定义接口， 类型别名， 类型定义， 类型开关
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