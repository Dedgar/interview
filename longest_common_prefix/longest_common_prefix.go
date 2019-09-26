package main

import (
	"encoding/json"
	"fmt"
)

var (
	Longest = prefix{letters: "", nodeCount: 0, instanceCount: 0}
)

type Node struct {
	Name      string
	Path      string
	Instances int
	Nodes     map[string]*Node
}

type prefix struct {
	letters       string
	nodeCount     int
	instanceCount int
}

func (n *Node) existNode(name string) bool {
	for _, v := range n.Nodes {
		if v.Name == name {
			return true
		}
	}
	return false
}

func (n *Node) addNode(nodeName, nodePath string) {
	if !n.existNode(nodeName) {
		n.Nodes[nodeName] = newNode(nodeName, nodePath)
	} else {
		fmt.Println("path of existing node", nodePath, n.Instances)
		n.Nodes[nodeName].Instances++
	}
}

func newNode(name, path string) *Node {
	startInstances := 1
	if name == "" {
		startInstances = 0
	} else if path == "" {
		path = name[0:1]
	}
	return &Node{name, path, startInstances, make(map[string]*Node)}
}

func (n *Node) getNode(name string) *Node {
	if nextN, ok := n.Nodes[name]; ok {
		return nextN
	} else if n.Name == name {
		return n
	} else {
		return &Node{}
	}
}

func (n *Node) json() {
	j, err := json.MarshalIndent(n, "", "   ")
	if err != nil {
		fmt.Println("Error printing node json", err)
	}

	fmt.Println(string(j))
}

func (n *Node) getList() (result string) {

	fmt.Printf("Longest starts at %+v\n", Longest)

	for _, c := range n.Nodes {
		if c.Path == "" {
			c.Path = c.Name
		}
		//if len(c.Path) > Longest.nodeCount { // || c.Path == "" {
		//
		if len(c.Path) >= len(Longest.letters) || Longest.nodeCount <= 1 && len(c.Path) < len(Longest.letters) {
			if len(c.Nodes) >= Longest.nodeCount || c.Instances >= Longest.instanceCount {
				if c.Instances >= Longest.instanceCount {
					/*if len(c.Nodes) <= 1 && len(n.Nodes) <= 1 {
						fmt.Println(c.Name, c.Path, "has 1 node")
					} else if len(c.Nodes) >= 2 {
						fmt.Println(c.Name, "has more than 1 node")
					}*/

					Longest.letters = c.Path
					Longest.nodeCount = len(c.Nodes)
					Longest.instanceCount = c.Instances
				} else {
					fmt.Printf("%v is not longer than longest instance count %v\n", c.Instances, Longest.instanceCount)
				}
			} else {
				fmt.Printf("%v has fewer nodes than longest %v\n", len(c.Nodes), Longest.nodeCount)
				fmt.Printf("lenth of c Path %v is less than longest letters len %v\n", c.Path, Longest.letters)
			}
		} else {
			fmt.Printf("well, %v with path %q aint long enough because its %v nodes is less than %v of %q\n", c.Name, c.Path, len(c.Nodes), len(Longest.letters), Longest.letters)
		}

		//fmt.Printf("This node: %v %v %v\n", v.Path, v.Name, len(v.Nodes))

		//for _, pNode := range v.Nodes {
		//	fmt.Printf("%v\n", pNode.Name)
		//}

		c.getList()
	}
	//fmt.Printf("Longest ends at %+v\n", Longest)
	return
}

func longestCommonPrefix(strs []string) string {
	RootNode := newNode("", "")

	Longest.letters = ""
	Longest.nodeCount = 0
	Longest.instanceCount = 0

	startLetter := ""

	if len(strs) == 1 {
		return strs[0]
	}

	// range over each word in the string slice
	for _, w := range strs {
		if w == "" {
			fmt.Println("w is empty")
			RootNode.addNode("", "")
			return ""
		}

		tmpNode := RootNode

		if startLetter == "" {
			startLetter = w[0:1]
		}

		// if we find something with a different first letter, return immediately
		if w[0:1] != startLetter {
			fmt.Println("Encountered a different prefix during tree build")
			return ""
		}

		// range over each character in the word,
		// prewords becomes the path of the word so far, including this character
		for j, c := range w {
			prewords := w[0 : j+1]

			//if prewords == "" {
			//	prewords = w
			//}

			// try to add a new node named after the character, with the current
			// prefix as the path. e.g. for "cake", node k has the prefix "cak"
			tmpNode.addNode(string(c), prewords)
			tmpNode = tmpNode.getNode(string(c))
		}
	}

	if len(RootNode.Nodes) < 1 {
		fmt.Println("RootNode has zero starting branches")
		return ""
	}
	//RootNode.Instances = 0

	RootNode.json()

	//for _, a := range RootNode.Nodes {
	//	fmt.Println("no child nodes")
	//	if len(a.Nodes) == 0 {
	//		return a.Name
	//	}
	//}

	/*for _, t := range RootNode.Nodes {
		fmt.Println("rootnode debug", t.Name)
		if len(t.Nodes) == 0 {
			return ""
		}
	}*/
	RootNode.getList()

	return Longest.letters
}

// The expected usage is to create a tree structure containing all the letters from the string slice
// The common prefix is selected by the node with the largest number of child nodes, because if there is
// any deviation from that point, then it means some words do not start with the same letters from there on.
// A really basic 3 word slice of ["testing", "tester", "tested"] would be built into a tree like this:
//   t
//   |
//   e
//   |
//   s
//   |
//   t
//  / \
//  i  e
//  | / \
//  n r d
//  |
//  g
func main() {
	fmt.Println("expecting 'a'", longestCommonPrefix([]string{"a", "aa", "aaa"}))
	fmt.Println("expecting 'a'", longestCommonPrefix([]string{"a"}))
	fmt.Println("expecting 'aa'", longestCommonPrefix([]string{"aa", "aa"}))
	fmt.Println("expecting 'a'", longestCommonPrefix([]string{"aa", "a"}))
	fmt.Println("expecting 'a'", longestCommonPrefix([]string{"aa", "ab"}))
	fmt.Println("expecting 'aaa'", longestCommonPrefix([]string{"aaa", "aaa"}))
	fmt.Println("expecting 'fl'", longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println("expecting ''", longestCommonPrefix([]string{""}))
	fmt.Println("expecting ''", longestCommonPrefix([]string{"", "b"}))
	fmt.Println("expecting 'c'", longestCommonPrefix([]string{"c", "c"}))
	fmt.Println("expecting 'test'", longestCommonPrefix([]string{"testa", "testb"}))
	fmt.Println("expecting 'test'", longestCommonPrefix([]string{"tester", "testing", "tested", "testament", "testo", "testz"}))
	fmt.Println("expecting prefix warning", longestCommonPrefix([]string{"nothing", "matches", "here"}))
}
