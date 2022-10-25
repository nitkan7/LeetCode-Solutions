func arrayStringsAreEqual(word1 []string, word2 []string) bool {
    var str1,str2 string
    
    for _,val:=range word1 {
        str1+=val
    }
    
    for _,val:=range word2 {
        str2+=val
    }
    return str1==str2
}
