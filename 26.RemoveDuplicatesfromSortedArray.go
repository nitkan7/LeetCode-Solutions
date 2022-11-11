func removeDuplicates(nums []int) int {
    left:=1
    for  i:=1;i<len(nums);i++ {
        if(nums[i-1]!=nums[i]){
             nums[left]=nums[i]
            left++
        }
    }
    return left
}
