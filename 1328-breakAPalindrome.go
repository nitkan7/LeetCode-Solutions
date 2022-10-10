func breakPalindrome(palindrome string) string {
    length:=len(palindrome)
    
    if(length==1){
        return ""
    }
    
    for i:=0;i<length/2;i++{
        if palindrome[i]!='a'{
            return palindrome[:i]+"a"+palindrome[i+1:]
        }
    }
    return palindrome[:length-1]+"b"
}