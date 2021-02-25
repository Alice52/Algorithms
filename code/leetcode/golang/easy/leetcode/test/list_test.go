package test

import (
    "fmt"
    "golang/common/api/model"
    "golang/easy/leetcode/list"
    "testing"
)

func TestGetIntersectionNode(t *testing.T) {

    nodeA1 := &model.ListNode{
        Val:  1,
        Next: nil,
    }

    nodeA2 := &model.ListNode{
        Val:  2,
        Next: nil,
    }

    nodeA3 := &model.ListNode{
        Val:  3,
        Next: nil,
    }

    nodeA4 := &model.ListNode{
        Val:  4,
        Next: nil,
    }

    nodeA5 := &model.ListNode{
        Val:  5,
        Next: nil,
    }

    nodeB1 := &model.ListNode{
        Val:  11,
        Next: nil,
    }
    nodeB2 := &model.ListNode{
        Val:  12,
        Next: nil,
    }

    nodeA1.Next = nodeA2
    nodeA2.Next = nodeA3
    nodeA3.Next = nodeA4
    nodeA4.Next = nodeA5

    nodeB2.Next = nodeA2
    nodeB1.Next = nil

    result := list.GetIntersectionNode(nodeA1, nodeB1)
    if result == nil {
        fmt.Printf("intersection node value: NULL\n")
    } else {
        fmt.Printf("intersection node value: %v\n", result.Val)
    }
}

func TestAddTwoSum(t *testing.T) {

    a := &model.ListNode{
        Val: 2,
        Next: &model.ListNode{
            Val: 4,
            Next: &model.ListNode{
                Val:  3,
                Next: nil,
            },
        },
    }

    b := &model.ListNode{
        Val: 5,
        Next: &model.ListNode{
            Val: 6,
            Next: &model.ListNode{
                Val:  4,
                Next: nil,
            },
        },
    }

    result := list.AddTwoSum(a, b)
    for result != nil {
        fmt.Print(result.Val)
        result = result.Next
    }

    fmt.Println()

    a2 := &model.ListNode{
        Val: 9,
        Next: &model.ListNode{
            Val: 9,
            Next: &model.ListNode{
                Val:  9,
                Next: nil,
            },
        },
    }

    b2 := &model.ListNode{
        Val:  1,
        Next: nil,
    }

    result2 := list.AddTwoSum(a2, b2)
    for result2 != nil {
        fmt.Print(result2.Val)
        result2 = result2.Next
    }
    fmt.Println()
}

func TestMaxArea(t *testing.T) {
    nums := []int{1, 2, 3}
    result := list.MaxArea(nums)
    fmt.Println(result)
}

func TestRemoveDuplicateOfSortedList(t *testing.T) {
    node := &model.ListNode{
        Val: 1,
        Next: &model.ListNode{
            Val: 2,
            Next: &model.ListNode{
                Val: 3,
                Next: &model.ListNode{
                    Val:  3,
                    Next: nil,
                },
            },
        },
    }

    result := list.RemoveDuplicateOfSortedList(node)
    model.PrintListNode(result)
}

func TestRemoveDuplicateOfSortedListII(t *testing.T) {
    node := &model.ListNode{
        Val: 1,
        Next: &model.ListNode{
            Val: 2,
            Next: &model.ListNode{
                Val: 3,
                Next: &model.ListNode{
                    Val:  3,
                    Next: nil,
                },
            },
        },
    }

    result := list.RemoveDuplicateOfSortedListII(node)
    model.PrintListNode(result)
}
