package main
import "fmt"

type SinglyLinkedListNode struct {
    data int
    next *SinglyLinkedListNode
}

func insertNodeAtHead(head *SinglyLinkedListNode, val int) *SinglyLinkedListNode {
    tmp := SinglyLinkedListNode{data: val, next: head}
    return &tmp
}

func main() {
    var n int
    var v int
    
    fmt.Scanf("%d\n", &n)
    fmt.Scanf("%d\n", &v)
    head := &SinglyLinkedListNode{data: v, next: nil}
    for i := 1; i<n; i++ {
        fmt.Scanf("%d\n", &v)
        head = insertNodeAtHead(head, v)
    }
    curr := head
    for curr != nil {
       fmt.Printf("%d\n", curr.data)
       curr = curr.next 
    }
}