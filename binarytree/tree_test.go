package binarytree

import (
    "fmt"
    "testing"
)

/*
 *        1
 *     2     3
 *   4   5  6  7
 *
 * 前序： 1 2 4 5 3 6 7
 * 中：   4 2 5 1 6 3 7
 * 后：   4 5 2 6 7 3 1
 */
func TestOrder(t *testing.T) {
    tree := NewTree(1,
        NewTree(2,
            NewTree(4, nil, nil),
            NewTree(5, nil, nil),
        ),
        NewTree(3,
            NewTree(6, nil, nil),
            NewTree(7, nil, nil),
        ),
    )
    tree.PreOrder()
    fmt.Println()
    tree.InOrder()
    fmt.Println()
    tree.PostOrder()
}

func TestNodeCnt(t *testing.T) {
    tree := NewTree(1,
        NewTree(2,
            NewTree(4, nil, nil),
            NewTree(5, nil, nil),
        ),
        NewTree(3,
            NewTree(6, nil, nil),
            NewTree(7, nil, nil),
        ),
    )

    if 7 != tree.NodeCnt() {
        t.Fatal("Wrong node cnt")
    }
    t.Log("The node cnt is 7")
}

func TestTreeHeight(t *testing.T) {
    tree := NewTree(1,
        NewTree(2,
            NewTree(4, nil, nil),
            NewTree(5, nil, nil),
        ),
        NewTree(3,
            NewTree(6, nil, nil),
            NewTree(7, nil, nil),
        ),
    )

    if 3 != tree.TreeHeight() {
        t.Fatal("Wrong tree height")
    }
    t.Log("The tree height is 3")

    tree.Right.Right.Right = NewTree(8, nil, nil)

    if 4 != tree.TreeHeight() {
        t.Fatal("Wrong tree height")
    }
    t.Log("The tree height is 4")
}

func TestOrderNoRecursion(t *testing.T) {
    tree := NewTree(1,
        NewTree(2,
            NewTree(4, nil, nil),
            NewTree(5, nil, nil),
        ),
        NewTree(3,
            NewTree(6, nil, nil),
            NewTree(7, nil, nil),
        ),
    )
    tree.PreOrderNoRecursion()
    fmt.Println()

}