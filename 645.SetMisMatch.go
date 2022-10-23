func findErrorNums(nums []int) []int {
    out:=make([]int,2)
    mapper:=make(map[int]int)
    dup:=-1
    missing:=-1
    
    for _,val:=range(nums){
        _,ok:=mapper[val]
        if ok {
            mapper[val]++
        }else{
            mapper[val]=1
        }
    }
    
    for i:=1;i<=len(nums);i++{
        _,ok:=mapper[i]
        if(ok){
            if(mapper[i]==2){
                dup=i
            }
        }else{
            missing=i
        }
    }
 out[1]=missing
    out[0]=dup
    return out
}
