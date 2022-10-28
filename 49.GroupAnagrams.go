func groupAnagrams(strs []string) [][]string {
    
    mapper:=make(map[string] []string)
    
    result:=make([][]string,0) 
    
    for _,val := range strs {
        key:=strings.Split(val,"")
        sort.Strings(key)
        finalKey:=strings.Join(key,"") 
        mapper[finalKey]=append(mapper[finalKey],val)
    }
    
    for _,val := range mapper {
        result=append(result,val)
    } 
    return result
}
