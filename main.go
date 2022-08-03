package main

import (
    "fmt"
    "os/exec"
    "bytes"
    "strings"
    "regexp"
)


type Node struct {
    data string
    children []*Node
}


func main() {
    cmd := exec.Command("go", "mod", "graph")
    
    var out bytes.Buffer
    cmd.Stdout = &out

    err := cmd.Run()

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
    graphOutputLines := strings.Split(output,"\n")

    
    //Now we got the root.

    //We now need to build a queue of nodes, and if we see that node we need to add the children
    var queue[] *Node

    //I think I may need to scan the children of the current node because I need to still find the keys of the node

    //We could have a map[string]Node, and you can add it in that way
    var curr *Node
    for _, line := range graphOutputLines {
        if len(strings.Split(line," ")) < 2 {
            break
        }
        parentDep := strings.Split(line," ")[0] 
        childDep := strings.Split(line," ")[1]
        if root == nil {
            data := parentDep
            children := childDep

            root = &Node{data: data}
            child := &Node{data: children}

            root.children = append(root.children, child)
            queue = append(root.children, child) 
            fmt.Printf("Root node %v\n", root.data)
            fmt.Printf("Child node %v\n", child.data)
            curr = root
        } else if curr.data == parentDep {
            newNode := &Node{data:childDep}
            curr.children = append(curr.children, newNode)
            fmt.Printf("Parent Node data : %v\n", curr.data)
            fmt.Printf("Child Node data : %v\n", newNode.data)
            queue = append(queue, newNode)
        } else {
            if len(queue) > 0 { 
                fmt.Println("Changed parent\n")
                //We encountered a depenency with no dependencies of its own
                // In that case we dequeue
                // until we find one matching go mod graph output

                for _, elm := range queue  {
                    if elm.data == parentDep {
                        curr = elm
                    }
                }
                newNode := &Node{data:childDep}
                curr.children = append(curr.children, newNode)
                fmt.Printf("New curr : %v\n", curr.data)
                fmt.Printf("Child curr data : %v\n", newNode.data)
            }else {
                break
            }
        }
    }

}
