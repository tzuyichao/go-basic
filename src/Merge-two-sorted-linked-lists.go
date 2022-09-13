package main
import "fmt"

type SinglyLinkedListNode struct {
    data int32
    next *SinglyLinkedListNode
}

func mergeLists(list1 *SinglyLinkedListNode, list2 *SinglyLinkedListNode) *SinglyLinkedListNode {
    var res *SinglyLinkedListNode
    var curr_res *SinglyLinkedListNode
    var curr_list1 *SinglyLinkedListNode
    var curr_list2 *SinglyLinkedListNode
    
    curr_list1 = list1
    curr_list2 = list2
    
    for curr_list1 != nil || curr_list2 != nil {
        if curr_list1 == nil {
            if res == nil {
                res = &SinglyLinkedListNode{data: curr_list2.data, next: nil}
                curr_res = res
                curr_list2 = curr_list2.next
            } else {
                curr_res.next = &SinglyLinkedListNode{data: curr_list2.data, next: nil}
                curr_list2 = curr_list2.next
                curr_res = curr_res.next
            }
        } else if curr_list2 == nil {
            if res == nil {
                res = &SinglyLinkedListNode{data: curr_list1.data, next: nil}
                curr_res = res
                curr_list1 = curr_list1.next
            } else {
                curr_res.next = &SinglyLinkedListNode{data: curr_list1.data, next: nil}
                curr_list1 = curr_list1.next
                curr_res = curr_res.next
            }
        } else {
            if curr_list1.data >= curr_list2.data {
                if res == nil {
                    res = &SinglyLinkedListNode{data: curr_list2.data, next: nil}
                    curr_res = res
                    curr_list2 = curr_list2.next
                } else {
                    curr_res.next = &SinglyLinkedListNode{data: curr_list2.data, next: nil}
                    curr_list2 = curr_list2.next
                    curr_res = curr_res.next
                }
            } else {
                if res == nil {
                    res = &SinglyLinkedListNode{data: curr_list1.data, next: nil}
                    curr_res = res
                    curr_list1 = curr_list1.next
                } else {
                    curr_res.next = &SinglyLinkedListNode{data: curr_list1.data, next: nil}
                    curr_list1 = curr_list1.next
                    curr_res = curr_res.next
                }
            }
        }
    }
    
    return res
}

func main() {
    var t int
    fmt.Scanf("%d\n", &t)
    for i := 0; i<t; i++ {
        var n int
        var list1 *SinglyLinkedListNode
        var curr1 *SinglyLinkedListNode
        curr1 = list1
        fmt.Scanf("%d\n", &n)
        for j := 0; j < n; j++ {
            var v int32
            
            fmt.Scanf("%d\n", &v)
            if curr1 == nil {
                list1 = &SinglyLinkedListNode{data: v, next: nil}
                curr1 = list1
            } else {
                curr1.next = &SinglyLinkedListNode{data: v, next: nil}
                curr1 = curr1.next
            }
        }
        var m int
        var list2 *SinglyLinkedListNode
        var curr2 *SinglyLinkedListNode
        curr2 = list2
        fmt.Scanf("%d\n", &m)
        for j := 0; j < m; j++ {
            var v int32
            fmt.Scanf("%d\n", &v)
            if curr2 == nil {
                list2 = &SinglyLinkedListNode{data: v, next: nil}
                curr2 = list2
            } else {
                curr2.next = &SinglyLinkedListNode{data: v, next: nil}
                curr2 = curr2.next
            }
        }
        res := mergeLists(list1, list2)
        curr := res
        for curr != nil {
            fmt.Printf("%d ", curr.data)
            curr = curr.next
        }
        fmt.Printf("\n")
    }
}