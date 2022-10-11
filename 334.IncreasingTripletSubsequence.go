func increasingTriplet(nums []int) bool {
    var largest,mid int
    
    largest=math.MinInt
    mid=math.MinInt
    
    for i:=len(nums)-1;i>=0;i--{
        if(largest<=nums[i] && mid>nums[i]){
            largest=nums[i]
        }else if (mid<=nums[i]){
            mid=nums[i]
        }else {
            return true
        }
    }
    return false
}
