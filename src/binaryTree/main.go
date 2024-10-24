package main

import "fmt"

type Node struct {
    Value int
    Index int
    Left  *Node
    Right *Node
}


func (n *Node) Insert(value int, index int) {
    if value <= n.Value {
        if n.Left == nil {
            n.Left = &Node{Value: value, Index: index}
        } else {
            n.Left.Insert(value, index)
        }
    } else {
        if n.Right == nil {
            n.Right = &Node{Value: value, Index: index}
        } else {
            n.Right.Insert(value, index)
        }
    }
}

func (n *Node) Search(value int, index int) (bool, int) {
    if n == nil {
        return false, -1
    }
    if n.Value == value && index != n.Index {
        return true, n.Index
    } else if value < n.Value {
        return n.Left.Search(value, index)
    } else {
        return n.Right.Search(value, index)
    }
}

func twoSum(nums []int, target int) []int { // 
   var solution []int
   root := &Node{}
   for i := range nums{
    root.Insert(nums[i], i)
   }
    for i:=0; i<len(nums); i++{
        poisk, sol := root.Search(target - nums[i], i)
        if poisk == true{
            solution = append(solution, i)
            solution = append(solution, sol)
            break
        }
    }

    return solution
}

func main() {
    root := &Node{Value: 10}
	var sulution []int
	for i:=0; i<10;i++{
		sulution = append(sulution, i)
	}

	for i:=0; i<10;i++{
		root.Insert(sulution[i],i)
	}
	first , position := root.Search(7, 7)

    fmt.Println("Дерево содержит 7:", first, position)
}
