package main

import (
    "strings"
)

type StackNode struct {
    value int
    next *StackNode
}

func createStackNode(value int) *StackNode{
    return &StackNode{value,nil}
}

func pushStack(oldStack *StackNode, value int) *StackNode{
    head := oldStack
    newNode := createStackNode(value)
    if head == nil {
        head = newNode
        return head
    } else {
        newNode.next = head
        oldStack = newNode
    }
    return oldStack
}

func popStack(stack *StackNode) *StackNode{
    head := stack
    if head == nil {
        return head
    } else {
        stack = head.next
        head = nil
    }
    return stack
}

func topStack(stack *StackNode) int{
    head := stack
    if head == nil {
        return -1
    } 
    return head.value
}

func reverseString(oldString string, start, stop int) string{
    stop_temp := stop
    temp := strings.Split(oldString, "")
    result := ""
    for i:=start; i<=stop; i++ {
        temp[i] = string(oldString[stop_temp]) 
        stop_temp--
    }
    for j:=0; j<len(temp); j++ {
        result += temp[j]
    }
    return result
}

func reverseParentheses(s string) string {
    var stack *StackNode
    var value string
    var topPos int
    result := ""
    temp := s
    for i:=0; i<len(s); i++{
        value = string(s[i])
        if value == "(" {
            stack = pushStack(stack, i)
        }
        if value == ")"{
            topPos = topStack(stack)
            temp = reverseString(temp, topPos, i)
            stack = popStack(stack)
        }
    }
    for j:=0; j<len(s); j++ {
        value = string(temp[j])
        if value == "(" || value == ")" {
            continue
        } else {
            result += value
        }
    }
    return result
}