package main

type ListNode struct {
  Val int
  Next *ListNode
}

func createNewNode(Val int) *ListNode{
    return &ListNode{Val,nil}
}

func addNewNode(newNode *ListNode, oldList *ListNode) *ListNode{
    var head *ListNode = oldList
    if head == nil {
        head = newNode
        return head
    }
    for true {
        if head.Next == nil {
            head.Next = newNode
            break
        }
        head = head.Next
    }
    return oldList
}

func addAllNode(head *ListNode, resultList *ListNode, carry int) *ListNode{
    var tempCarry int = carry 
    var newVal int
    for true {
        if head == nil {
            if tempCarry != 0 {
                newNode := createNewNode(tempCarry)
                resultList = addNewNode(newNode, resultList)
            }
            break
        } else {
            newVal = head.Val+tempCarry
            tempCarry = 0
            if newVal > 9 {
                newVal = newVal-10
                tempCarry++
            }
            newNode := createNewNode(newVal)
            resultList = addNewNode(newNode, resultList)
        }
        head = head.Next
    }
    return resultList
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var resultListNode *ListNode
    var head1 *ListNode = l1
    var head2 *ListNode = l2
    var newNode *ListNode
    carry := 0
    for true {
        if head1 == nil || head2 == nil {
            if head1 != nil {
                resultListNode = addAllNode(head1, resultListNode, carry)
            } else if head2!= nil {
                resultListNode = addAllNode(head2, resultListNode, carry)
            } else if carry != 0 {
                newNode = createNewNode(carry)
                resultListNode = addNewNode(newNode,resultListNode)
            }
            break
        }
        newVal := head1.Val+head2.Val+carry
        carry = 0
        if newVal > 9 {
            newVal = newVal - 10
            carry++
        } 
        newNode = createNewNode(newVal)
        resultListNode = addNewNode(newNode,resultListNode)
        head1 = head1.Next
        head2 = head2.Next
    }
    return resultListNode
}