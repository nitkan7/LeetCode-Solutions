func minimumRounds(tasks []int) int {
    mapper:=make(map[int]int)
    
    for _,val :=range(tasks){
        _,ok:=mapper[val]
        
        if(ok){
            mapper[val]++
            continue
        }
        
        mapper[val]=1
    }
    
    
    result:=0
    for _,val:=range(mapper){
        if val==1 {
            return -1
        }
        
        result+=(val+2)/3
    }
    return result
    
}
