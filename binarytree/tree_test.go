package binarytree

import "testing"

/*
 *        1
 *     2     3
 *   4   5  6  7
 *
 * 前序： 1 2 4 5 3 6 7
 */
func TestPreOrder(t *testing.T) {
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
}