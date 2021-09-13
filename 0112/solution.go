package s0112



func hasPathSumRecursive(root *TreeNode, targetSum int) bool {
    if root == nil {
        return false
    }
    
    if root.Left == nil && root.Right == nil {
        return targetSum == root.Val
    }
    
    if root.Left != nil {
        if hasPathSum(root.Left, targetSum - root.Val) {
            return true
        }
    }
    
    if root.Right != nil {
        if hasPathSum(root.Right, targetSum - root.Val) {
            return true
        }
    }
    
    return false
}
