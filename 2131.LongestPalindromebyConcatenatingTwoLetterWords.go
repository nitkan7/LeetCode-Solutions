func longestPalindrome(words []string) int {
    length:=len(words)
    countMap:=make(map[string]int,length)
    stringLength:=0
    
    
    for _,word:=range words {
        reversedString:=reverse(word)
        
        if value,ok:=countMap[reversedString];ok{
            stringLength+=4
            if value==1 {
                delete(countMap,reversedString)
            }else{
                countMap[reversedString]--
            }
        }else{
            countMap[word]++
        }
    }
    
    for key,_:=range countMap{
        if isPair(key){
            stringLength+=2
            break
        }
    }
    return stringLength
    
}

func reverse(s string)string{
    reversed:=make([]byte,2)
    reversed[0],reversed[1]=s[1],s[0]
    return string(reversed)
}

func isPair(s string)bool{
    return s[0] == s[1]
}
