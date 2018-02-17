## 正誤表が充実した（thx yurakawa, ruicc）

https://github.com/LambdaNote/errata-gosyspro-1-1/issues

## Open Data Structure をGoに移植してみているけどうまくうごかない

```go
type Node struct {
	x int
	next *Node
}

var head Node = Node {next:nil}
var tail Node = Node {next:nil}
var n int

func fmtsslist (u Node) {
	for i := n; i > 0; i-- {
		fmt.Println(u.x)
		u = *u.next
	}
}

func main() {
	push(0)
	push(1)
	push(2)
	fmtsslist(head)
}

func push(x int) int {
	u := Node {next:nil}
	w := head
	u.x = x
	u.next = &head
	head = u
	if (n == 0) {
		tail = u
	}
	n++
	return x
}
```
