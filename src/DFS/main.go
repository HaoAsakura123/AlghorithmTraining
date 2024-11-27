package main

import (
	"fmt"
	"log"
	"slices"
)

func main() {
	var vertex int
	_, _ = fmt.Scan(&vertex)
	var edge int
	_, _ = fmt.Scan(&edge)

	graph := make([][]int, vertex)
	for i := 0; i < vertex; i++ {
		graph[i] = []int{}
	}

	for i := 0; i < edge; i++ {
		var tmp1, tmp2 int
		_, _ = fmt.Scan(&tmp1, &tmp2)
		if tmp1 != tmp2 {
			graph[tmp1-1] = append(graph[tmp1-1], tmp2)
			graph[tmp2-1] = append(graph[tmp2-1], tmp1)
		}
	}
	mapGraph := make(map[int]bool)
	s := Stack{}
	s.Push(1)
	mapGraph[1] = true
	resultat := DFS(&s, mapGraph, graph)
	fmt.Println(len(resultat))
	slices.Sort(resultat)
	for i, value:= range resultat{
		if i <len(resultat) - 1{
			fmt.Printf("%d ",value)
			} else {
				fmt.Printf("%d\n",value)
			}
			
		}
}

type Stack struct {
	content []int
}

func (stack *Stack) Display() {
	log.Print(stack.content)
}

func (stack *Stack) Push(item int) {
	stack.content = append(stack.content, item)
}

func (stack *Stack) Pop() int {
	if len(stack.content) == 0 {
		return -1 
	}
	item := stack.content[len(stack.content)-1]
	stack.content = stack.content[:len(stack.content)-1]
	return item
}

func DFS(stack *Stack, mapGraph map[int]bool, graph [][]int) []int {
	sliceGraph := make([]int, 0)

	for len(stack.content) > 0 {
		vertex := stack.Pop()
		sliceGraph = append(sliceGraph, vertex)
		for _, neighbor := range graph[vertex-1] {
			if !mapGraph[neighbor] {
				stack.Push(neighbor)
				mapGraph[neighbor] = true
			}
		}
	}

	return sliceGraph
}
