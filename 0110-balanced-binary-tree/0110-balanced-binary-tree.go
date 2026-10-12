/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func Height (root *TreeNode) int {
    if root == nil {
        return 0
    }

    return 1 + max(Height(root.Left), Height(root.Right))
}

func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }    

    heightLeft := Height(root.Left)
    heightRight := Height(root.Right)

    if heightLeft - heightRight > 1 || heightLeft - heightRight < -1 {
        return false
    }

    return isBalanced(root.Left) && isBalanced(root.Right)
}