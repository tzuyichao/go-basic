package main
import "fmt"

type SinglyLinkedListNode struct {
    data int32
    next *SinglyLinkedListNode
}

func compare_lists(list1 *SinglyLinkedListNode, list2 *SinglyLinkedListNode) int32 {
    for list1 != nil && list2 != nil {
        if list1.data != list2.data {
            return 0
        }
        list1 = list1.next
        list2 = list2.next
    }
    return 1
}

func main() {
    var numTestCases int
    
    fmt.Scanf("%d\n", &numTestCases)
    for idxTestCases := 0; idxTestCases < numTestCases; idxTestCases++ {
        var list1 *SinglyLinkedListNode
        var len1 int
        var list2 *SinglyLinkedListNode
        var len2 int
        
        fmt.Scanf("%d\n", &len1)
        var curr *SinglyLinkedListNode
        for i := 0; i<len1; i++ {
            var v int32
            fmt.Scanf("%d\n", &v)
            if i == 0 {
                list1 = &SinglyLinkedListNode{data: v, next: nil}
                curr = list1
            } else {
                curr.next = &SinglyLinkedListNode{data: v, next: nil}
                curr = curr.next
            }
        }    
        fmt.Scanf("%d\n", &len2)
        for i := 0; i<len2; i++ {
            var v int32
            fmt.Scanf("%d\n", &v)
            if i == 0 {
                list2 = &SinglyLinkedListNode{data: v, next: nil}
                curr = list2
            } else {
                curr.next = &SinglyLinkedListNode{data: v, next: nil}
                curr = curr.next
            }
        }    
        
        if len1 == len2 {
            res := compare_lists(list1, list2)
            fmt.Printf("%d\n", res)
        } else {
            fmt.Printf("0\n");
        }
    }
}