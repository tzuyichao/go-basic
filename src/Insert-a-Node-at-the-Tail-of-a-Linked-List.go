package main
import "fmt"

type SinglyLinkedListNode struct {
    data int
    next *SinglyLinkedListNode
}

func main() {
    var n int
    var i int
    var d int
    
    fmt.Scanf("%d\n", &n)
    fmt.Scanf("%d\n", &d)
    head := &SinglyLinkedListNode{data: d, next: nil}
    current := head
    
    for i = 1; i < n; i++ {
        fmt.Scanf("%d\n", &d)
        current.next = &SinglyLinkedListNode{data: d, next: nil}
        current = current.next
    }
    current = head
    for current != nil {
        fmt.Printf("%d\n", current.data)
        current = current.next
    }
}
