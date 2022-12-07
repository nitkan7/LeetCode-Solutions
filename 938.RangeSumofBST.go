 var ans int
func rangeSumBST(root *TreeNode, low int, high int) int {
    ans=0
    dfs(root,low,high)
    return ans
}

func dfs(root *TreeNode, low int, high int){
    if(root!=nil){
        if(low<=root.Val && root.Val<=high){
            ans+=root.Val
        }
        if(low < root.Val){
            dfs(root.Left,low,high)
        }
        if(root.Val < high){
            dfs(root.Right,low,high)
        }
    }
}
