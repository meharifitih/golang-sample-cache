package main

import "fmt"

const MaxSize = 5

type Node struct {
	Value string
	Left  *Node
	Right *Node
}

type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}

type Cache struct {
	Queue Queue
	Hash  Hash
}

type Hash map[string]*Node

func NewCache() Cache {
	return Cache{
		Queue: NewQueue(),
		Hash:  Hash{},
	}
}

func  NewQueue() Queue {
	head := &Node{}
	tail := &Node{}

	head.Right = tail
	tail.Left = head

	return Queue{
		Head: head,
		Tail: tail,
	}
}

func (c *Cache) Check(str string) {
	node := &Node{}

	if val, ok := c.Hash[str]; ok {
		node = c.Remove(val)
	} else {
		node = &Node{Value: str}
	}

	c.Add(node)
	c.Hash[str] = node
}

func (c *Cache) Remove(n *Node) *Node {
	fmt.Printf("remove %s\n", n.Value)

	left := n.Left
	right := n.Right

	left.Right = right
	right.Left = left

	c.Queue.Length--
	return n
}

func (c *Cache) Add(n *Node) {
	fmt.Printf("add %s\n", n.Value)

	tmp := c.Queue.Head.Right

	c.Queue.Head.Right = n
	n.Left = c.Queue.Head
	n.Right = tmp
	tmp.Left = n

	c.Queue.Length++

	if c.Queue.Length > MaxSize {
		c.Remove(c.Queue.Tail.Left)
	}

}

func (c *Cache) Display() {
	c.Queue.Display()
}

func (q *Queue) Display() {
	node := q.Head.Right
	fmt.Printf("%d - [", q.Length)

	for i := 0; i < q.Length; i++ {
		fmt.Printf("{%s}", node.Value)
		if i < q.Length-1 {
			fmt.Printf("<-->")
		}
		node = node.Right
	}
	fmt.Println("]")
}

func main() {
	cache := NewCache()

	for _, word := range []string{"potato", "tomato", "avocado", "lemon",
		"banana", "orange", "kiwi", "grape", "peach", "plum"} {
		cache.Check(word)
		cache.Display()
	}
}
