func maxLength(arr []string) int {
    var out int
    dfs(arr,"",&out,0)
    return out
}


func dfs(arr []string, current string, out *int,index int){
    if(index==len(arr) && isUnique(current)>*out){
        *out=len(current)
        return
    }
    if(index==len(arr)){
        return
    }
    if(isUnique(current)==-1){
        return
    }
    dfs(arr,current+arr[index],out,index+1)
    dfs(arr,current,out,index+1)
}

func isUnique(current string)int{
    mapper:=make(map[rune]int)
    for _,c:= range current {
        _,ok:=mapper[c]
        if ok {
            return -1
        }else{
            mapper[c]=1
        }
    }
    return len(current)
}
