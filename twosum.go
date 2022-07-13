func twoSum(nums []int, target int) []int {
    out:=make([]int,2)
    mapper:=make(map[int]int)
    
    diff:=0
    for index,value :=range(nums) {
        diff=target-value
        
        _,ok:=mapper[diff]
        
        if(ok){
            out[0]=mapper[diff]
            out[1]=index
            return out
        }
        mapper[value]=index
    }
    
    return out
    }
