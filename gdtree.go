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

	var root *Node
	regex, err := regexp.Compile("\n\n")
	if err != nil {
		return
	}
	output := regex.ReplaceAllString(out.String(), "\n")
	graphOutputLines := strings.Split(output, "\n")

	// We now need to build a queue of nodes based on go mod graph output.
	// For every package on the left side of the output, we create a parent node
	// and insert append children based on the pacakges from the right side of the output.
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
				// We search for the new parent from the child nodes we've queued up matching go mod graph output
				// These are the nodes that are connected to root.
				for _, elm := range queue { 
					if elm.data == parentDep {
						curr = elm
					}
				}

				insertChildren(curr, childDep)
			}
		}
	}

	printTreeBox([]*Node{root}, 0)

}

func insertChildren(parentNode *Node, childData string) {
	newNode := &Node{data: childData}
	parentNode.children = append(parentNode.children, newNode)
}

func printTreeBox(children []*Node, depth int) {
	for i := 0; i < len(children); i++ {

		// This is here so we can print a line representing the parent tree
		// and the child that has its own children
		// parent
		//  └── child
		//  |     └── children
		if depth > 1 {
			for j := 0; j < depth; j++ {

				fmt.Printf("\u2502")
				printSpaces(depth - 1)
			}
		}

		if depth == 0 {
			fmt.Printf("%v\n", children[i].data)
		} else if i < len(children)-1 {
			fmt.Printf("├── %v\n", children[i].data)
		} else {
			fmt.Printf("└── %v\n", children[i].data)
		}

		if len(children[i].children) > 0 {
			printTreeBox(children[i].children, depth+1)
		}

	}
}

func printSpaces(depth int) {
	for j := 0; j < 4; j++ {
		fmt.Printf(" ")
	}
}
