func intersect(nums1 []int, nums2 []int) []int {
    mapper:=make(map[int]int)
    length:=len(nums1)
    var output []int
    
    for i:=0;i<length;i++{
        _,ok:=mapper[nums1[i]]
        if(ok){
            mapper[nums1[i]]++

        }else{
            
            mapper[nums1[i]]=1
  
        }
        
    }
    for i:=0;i<len(nums2);i++ {
        if(mapper[nums2[i]]>0){
            output = append(output,nums2[i])
            mapper[nums2[i]]--
            
        }
    }
    return output
}
