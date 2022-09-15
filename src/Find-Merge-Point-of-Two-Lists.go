package main
import "fmt"

type SinglyLinkedListNode struct {
    data int32
    next *SinglyLinkedListNode
}

func findMergeNode(list1 *SinglyLinkedListNode, list2 *SinglyLinkedListNode) int32 {
    curr1 := list1
    for curr1 != nil {
        curr2 := list2
        for curr2 != nil {
            if curr1 == curr2 {
                return curr1.data
            }
            curr2 = curr2.next
        }
        curr1 = curr1.next
    }
    return 0
}

func main() {
    var t int
    fmt.Scanf("%d\n", &t)
    for idx:=0; idx < t; idx++ {
        var ind int
        fmt.Scanf("%d\n", &ind)
        index := &SinglyLinkedListNode{data: -1, next: nil}
        var list1 *SinglyLinkedListNode
        var curr *SinglyLinkedListNode
        var c1 int
        fmt.Scanf("%d\n", &c1)
        for i:=0; i<c1; i++ {
            var v int32
            fmt.Scanf("%d\n", &v)
            if i == ind {
                index.data = v
                if curr != nil {
                    curr.next = index
                    curr = curr.next
                } else {
                    list1 = index
                    curr = list1
                }
            } else {
                if curr != nil {
                    curr.next = &SinglyLinkedListNode{data: v, next: nil}
                    curr = curr.next
                } else {
                    list1 = &SinglyLinkedListNode{data: v, next: nil}
                    curr = list1
                }
            }
        }
        var list2 *SinglyLinkedListNode
        curr = nil
        fmt.Scanf("%d\n", &c1)
        for i:=0; i<c1; i++ {
            var v int32
            fmt.Scanf("%d\n", &v)
            if curr != nil {
                curr.next = &SinglyLinkedListNode{data: v, next: nil}
                curr = curr.next
            } else {
                list2 = &SinglyLinkedListNode{data: v, next: nil}
                curr = list2
            }
        }
        curr.next = index
        res := findMergeNode(list1, list2)
        fmt.Printf("%d\n", res)
    }
}