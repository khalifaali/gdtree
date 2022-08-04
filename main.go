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

	for i := 0; i < len(q); i++ {
		node := q[i]
		if node.children != nil {
			q = append(q, node.children[:]...)
		}
		fmt.Printf("Node data part 2.... :  %v\n", node.data)

	}
}

func printTreeBox(node *Node, depth int) {
	for i := 0; i < depth; i++ {
		fmt.Println("|")
	}
	fmt.Printf("\u251c")
}
