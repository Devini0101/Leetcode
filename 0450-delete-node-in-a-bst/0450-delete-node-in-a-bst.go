/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
    if root == nil {
        return nil
    }

    if key > root.Val {
        root.Right = deleteNode(root.Right, key)
    } else if key < root.Val {
        root.Left = deleteNode(root.Left, key)
    } else {
        // case when theres no sutrees
        if root.Right == nil {
            return root.Left
        } else if root.Left == nil {
            return root.Right
        }

        sub := root.Right
        
        for sub.Left != nil {
            sub = sub.Left
        }

        root.Val = sub.Val
        root.Right = deleteNode(root.Right, sub.Val)
    }

    return root
}