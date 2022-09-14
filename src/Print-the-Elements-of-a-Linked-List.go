package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

type SinglyLinkedListNode struct {
    data int32
    next *SinglyLinkedListNode
}

type SinglyLinkedList struct {
    head *SinglyLinkedListNode
    tail *SinglyLinkedListNode
}

func (singlyLinkedList *SinglyLinkedList) insertNodeIntoSinglyLinkedList(nodeData int32) {
    node := &SinglyLinkedListNode {
        next: nil,
        data: nodeData,
    }

    if singlyLinkedList.head == nil {
        singlyLinkedList.head = node
    } else {
        singlyLinkedList.tail.next = node
    }

    singlyLinkedList.tail = node
}

// Complete the printLinkedList function below.

/*
 * For your reference:
 *
 * SinglyLinkedListNode {
 *     data int32
 *     next *SinglyLinkedListNode
 * }
 *
 */
func printLinkedList(head *SinglyLinkedListNode) {
    curr := head
    for curr != nil {
        fmt.Printf("%d\n", curr.data)
        curr = curr.next
    }
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    llistCount, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)

    llist := SinglyLinkedList{}
    for i := 0; i < int(llistCount); i++ {
        llistItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
        checkError(err)
        llistItem := int32(llistItemTemp)
        llist.insertNodeIntoSinglyLinkedList(llistItem)
    }

    printLinkedList(llist.head)
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
