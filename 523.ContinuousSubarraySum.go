func checkSubarraySum(nums []int, k int) bool {
    //create the hashmap and init with a starting value
    
    mapper:=make(map[int]int)
    mapper[0]=-1
    
    total,remainder:=0,0
    
    for i,val:=range nums {
        total+=val
        remainder=total%k
        
        _,ok:=mapper[remainder]
        if(!ok){
            mapper[remainder]=i
        }else if (i - mapper[remainder] > 1){
            return true
        }
    }
    
    return false
}
