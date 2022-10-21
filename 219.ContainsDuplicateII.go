func containsNearbyDuplicate(nums []int, k int) bool {
    mapper:=make(map[int]int)
    
    for i,val:=range(nums){
        _,ok:=mapper[val]
        
        if ( ok && (i-mapper[val])<=k) {
            return true
        }
        mapper[val]=i
    }
    return false
}
