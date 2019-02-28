# 读取用户的输入
## 从键盘和标准输入os.Stdin读取输入,最简单的办法是使用fmt提供的Scan和Sscan开头的函数。
```
package main
import (
	"fmt"
)
var (
	firstName,lastName,s string
	i  int
	f float32
	input = "56.12 / 5212 / Go"
	format = "%f / %d / %s"
)
func main() {
	fmt.Println("Please enter your full name: ")
	fmt.Scanln(&firstName,&lastName)
	fmt.Printf("Hi, %s %s!\n",firstName,lastName)
	fmt.Sscanf(input,format,&f,&i,&s)
	fmt.Println("From the string we read:",f,i,s)
}
```
### Scanln扫描来自标准输入的文本，将空格分隔的值以此存放到后续的参数内，直到换行。Scanf与其类似，除了Scanf的第一个参数用作格式字符串，用来决定如何读取。Sscan和以Sscan开头的函数则是从字符串读取，除此外，与Scanf相同。
## 你也可以使用bufio包提供的缓冲读取(buffered reader)来读取数据。
```go
package main
import (
	"fmt"
	"bufio"
	"os"
)
var inputReadr *bufio.Reader
var input string
var err error

func main() {
	inputReadr = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input: ")
	input,err = inputReadr.ReadString('\n')
	if err == nil {
		fmt.Println("The input was: %s\n",input)
	}

}
```
### inputReader是一个指向bufio.Reader的指针。inputReader := bufio.NewReader(os.Stdin)这一行代码。