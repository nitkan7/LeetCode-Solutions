func halvesAreAlike(s string) bool {
    if(len(s)%2!=0){
        return false
    }
    
    arr:=strings.Split(s,"")
    left,right:=0,len(arr)-1
    leftCount,rightCount:=0,0
    for left<right {
        
        if(arr[left]=="A" || arr[left]=="E" || arr[left]=="I" || arr[left]=="O" || arr[left]=="U" || 
           arr[left]=="a" || arr[left]=="e" || arr[left]=="i" || arr[left]=="o" || arr[left]=="u" ){
            leftCount++
        }
        if(arr[right]=="A" || arr[right]=="E" || arr[right]=="I" || arr[right]=="O" || arr[right]=="U" || 
           arr[right]=="a" || arr[right]=="e" || arr[right]=="i" || arr[right]=="o" || arr[right]=="u" ){
            rightCount++
        }
        left++
        right--
    }
    
    return leftCount==rightCount
}
