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
			fmt.Printf("Root baby %v\n", root.data)
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
	nodePtr := root
	bfs(nodePtr)

}

func insertChildren(parentNode *Node, childData string) {
	newNode := &Node{data: childData}
	parentNode.children = append(parentNode.children, newNode)
}

func bfs(node *Node) {
	var q []*Node
	q = append(q, node.children[:]...)
	fmt.Printf("%v\n", node.data)
	for i := 0; i < len(q); i++ {
		childNode := q[i]
		if node.children != nil {
			q = append(q, childNode.children[:]...)
		}
		fmt.Printf("├──%v\n", childNode.data)
		printTreeBox(childNode.children)
		//The queue also has the same data. So I need way to shrink the queue. for the elements seen?
	}
}
//need a seen map for the nodes. and if its seen
//So now I need an index to pass, and the parent node to print out
func printTreeBox(children []*Node) {
	for i := 0; i < len(children); i++ {
		fmt.Printf("\u2502\n")
		fmt.Printf("\t├── %v\n", children[i].data)
	}
}
