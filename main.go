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

    

    //We now need to build a queue of nodes, and if we see that node we need to add the children
    var queue[] *Node


    var curr *Node
    for _, line := range graphOutputLines {
        if len(strings.Split(line," ")) < 2 {
            break
        }
        parentDep := strings.Split(line," ")[0] 
        childDep := strings.Split(line," ")[1]
        if root == nil {
            data := parentDep
            root = &Node{data: data}
            curr = root
            queue = append(queue, insertChildren(root, childDep))
        } else if curr.data == parentDep {
            queue = append(queue, insertChildren(curr, childDep))
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
                insertChildren(curr, childDep)
            }
        }
    }

}

func insertChildren(parentNode *Node, childData string) *Node{
    newNode := &Node{data:childData}
    parentNode.children = append(parentNode.children, newNode)
    fmt.Printf("New curr : %v\n", parentNode.data)
    fmt.Printf("Child curr data : %v\n", newNode.data)
    return newNode
}



