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

func TestPreOrderNoRecursion(t *testing.T) {
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

func TestInOrderNoRecursion(t *testing.T) {
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
    tree.InOrderNoRecursion()
    fmt.Println()
}

func TestInOrderPostRecursion(t *testing.T) {
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
    tree.PostOrderNoRecursion()
    fmt.Println()
}

func TestChange2Child(t *testing.T) {
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
    tree.Change2Child()
    tree.PreOrder()
    fmt.Println()
    tree.InOrder()
    fmt.Println()
    tree.PostOrder()
}

func TestChange2ChildNoRecursion(t *testing.T) {
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
    tree.Change2ChildNoRecursion()
    tree.PreOrder()
    fmt.Println()
    tree.InOrder()
    fmt.Println()
    tree.PostOrder()
}

func TestCheckSubTree(t *testing.T) {
    tree1 := NewTree(8,
        NewTree(8,
            NewTree(9, nil, nil),
            NewTree(2,
                NewTree(4, nil, nil),
                NewTree(7, nil, nil),
            ),
        ),
        NewTree(7, nil, nil),
    )
    tree2 := NewTree(8,
        NewTree(9, nil, nil),
        NewTree(2, nil, nil),
    )

    fmt.Println(tree1.CheckSubTree(tree2))
}

func TestMirrorTree(t *testing.T) {
    tree := NewTree(8,
        NewTree(6,
            NewTree(5, nil, nil),
            NewTree(7, nil, nil),
        ),
        NewTree(10,
            NewTree(9, nil, nil),
            NewTree(11, nil, nil),
        ),
    )


    tree.PreOrder()
    fmt.Println()
    tree.MirrorTree()
    tree.PreOrder()
}

func TestCheckSymmetry(t *testing.T) {
    tree1 := NewTree(8,
        NewTree(6,
            NewTree(5, nil, nil),
            NewTree(7, nil, nil),
        ),
        NewTree(6,
            NewTree(7, nil, nil),
            NewTree(5, nil, nil),
        ),
    )

    tree2 := NewTree(8,
        NewTree(6,
            NewTree(5, nil, nil),
            NewTree(7, nil, nil),
        ),
        NewTree(9,
            NewTree(7, nil, nil),
            NewTree(5, nil, nil),
        ),
    )

    tree3 := NewTree(8,
        NewTree(7,
            NewTree(7, nil, nil),
            NewTree(7, nil, nil),
        ),
        NewTree(7,
            NewTree(7, nil, nil),
            nil,
        ),
    )

    fmt.Println(tree1.CheckSymmetry())
    fmt.Println(tree2.CheckSymmetry())
    fmt.Println(tree3.CheckSymmetry())
}