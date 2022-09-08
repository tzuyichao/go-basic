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

func printSinglyLinkedList(node *SinglyLinkedListNode, sep string, writer *bufio.Writer) {
    for node != nil {
        fmt.Fprintf(writer, "%d", node.data)

        node = node.next

        if node != nil {
            fmt.Fprintf(writer, sep)
        }
    }
}

/*
 * Complete the 'insertNodeAtPosition' function below.
 *
 * The function is expected to return an INTEGER_SINGLY_LINKED_LIST.
 * The function accepts following parameters:
 *  1. INTEGER_SINGLY_LINKED_LIST llist
 *  2. INTEGER data
 *  3. INTEGER position
 */

/*
 * For your reference:
 *
 * SinglyLinkedListNode {
 *     data int32
 *     next *SinglyLinkedListNode
 * }
 *
 */

func insertNodeAtPosition(llist *SinglyLinkedListNode, data int32, position int32) *SinglyLinkedListNode {
    var idx int32
    idx = 0
    curr := llist
    var prev *SinglyLinkedListNode
    prev = nil
    for curr != nil {
        if idx == position {
            node := SinglyLinkedListNode{data: data, next: curr}
            if prev == nil {
                return &node
            } else {
                prev.next = &node
                return llist
            }
        }
        idx = idx + 1
        prev = curr
        curr = curr.next
    }
    return llist
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    llistCount, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)

    llist := SinglyLinkedList{}
    for i := 0; i < int(llistCount); i++ {
        llistItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
        checkError(err)
        llistItem := int32(llistItemTemp)
        llist.insertNodeIntoSinglyLinkedList(llistItem)
    }

    dataTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    data := int32(dataTemp)

    positionTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    position := int32(positionTemp)

    llist_head := insertNodeAtPosition(llist.head, data, position)

    printSinglyLinkedList(llist_head, " ", writer)
    fmt.Fprintf(writer, "\n")

    writer.Flush()
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
