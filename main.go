package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

type Node struct {
	data     string
	children []*Node
}

func main() {
	cmd := exec.Command("go", "mod", "graph")

	var out, stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()

	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
	if err != nil {
		fmt.Println("Error occurred")
		fmt.Println(err)
	}

	var root *Node
	regex, err := regexp.Compile("\n\n")
	if err != nil {
		return
	}
	output := regex.ReplaceAllString(out.String(), "\n")
	graphOutputLines := strings.Split(output, "\n")

	//We now need to build a queue of nodes, and if we see that node we need to add the children
	var queue []*Node

	var curr *Node
	for _, line := range graphOutputLines {
		if len(strings.Split(line, " ")) < 2 {
			break
		}
		parentDep := strings.Split(line, " ")[0]
		childDep := strings.Split(line, " ")[1]
		if root == nil {
			root = &Node{data: parentDep}
			curr = root
			//fmt.Printf("Root baby %v\n", root.data)
			insertChildren(root, childDep)
			queue = append(queue, root.children[len(root.children)-1])
		} else if curr.data == parentDep {
			insertChildren(curr, childDep)
			queue = append(queue, curr, curr.children[len(curr.children)-1])
		} else {
			if len(queue) > 0 {
				//We encountered a depenency with no dependencies of its own
				// In that case we dequeue
				// until we find one matching go mod graph output
				for _, elm := range queue {
					if elm.data == parentDep {
						curr = elm
					}
				}
				insertChildren(curr, childDep)
			}
		}
	}

	//bfs(nodePtr)
	printTreeBox([]*Node{root},0)

}

func insertChildren(parentNode *Node, childData string) {
	newNode := &Node{data: childData}
	parentNode.children = append(parentNode.children, newNode)
}

func bfs(node *Node) {
	var q []*Node
	q = append(q, node)
 	seen := make(map[string]int)
	//depth := 0
	for i := 0; i < len(q); i++ {
		childNode := q[i]
		if childNode.children != nil {
			q = append(q, childNode.children[:]...)
			fmt.Printf("parent %v\n", childNode.data)
			//printTreeBox([]*Node{childNode}, depth) //A little bit of a hack but if we pass the parentNode
													// as a list we can print it using the same loginc
			//fmt.Printf("Parent├──%v\n", childNode.data
			//We shouldn't increase the depth until we exhaust all the children??
			//Some of these are all children of 1. element
			// If we  seen that nodes parent increase the depth else continue
			// We increase the depth bc if it has a child we want it to be at the
			// same level as its
			
		
	
			
		}

		if _, prs := seen[childNode.data]; prs {
			seen[childNode.data] = 1
			//fmt.Println(seen)
		} else {
			seen[childNode.data]++
		}
		//I want to print
		 // i thnk we fix this. Purge data from queue that are contained in the
											//children you;ve aeen
		//I need to add the seen to the children, and
		printTreeBox(childNode.children, seen[childNode.data])
		//The queue also has the same data. So I need way to shrink the queue. for the elements seen?
	}
}
//need a seen map for the nodes. and if its seen
//So now I need an index to pass, and the parent node to print out
func printTreeBox(children []*Node, depth int) {
	for i := 0; i < len(children); i++ {
		printSpaces(depth)
		fmt.Printf("Depth %v", depth)
		fmt.Printf("\u2502\n")
		// We're either printing the first element
		// or we're printing everything but the last
		// 
		if depth == 0 {
			fmt.Printf("── %v\n", children[i].data)
		} else if i < len(children) - 1 { 
			printSpaces(depth * 2)
			//fmt.Printf("\u2502\n")
			//printSpaces(depth * 2)
			fmt.Printf("   ├── %v\n", children[i].data)
		} else {
			//"\u2514"
			printSpaces(depth * 2)
			//fmt.Printf("\u2502\n")
			//printSpaces(depth * 2)
			fmt.Printf("   └── %v\n", children[i].data)
		}

		if len(children[i].children) > 0 {
			printTreeBox(children[i].children, depth + 1)
		}
		
	}
}


func printSpaces(depth int) {
	for j := 0 ; j < depth * 2; j++  {
		fmt.Printf(" ")
	}
}
