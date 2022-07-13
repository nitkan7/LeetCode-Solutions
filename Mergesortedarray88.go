func merge(nums1 []int, m int, nums2 []int, n int)  {
    first,second:=m-1,n-1
    position:=m+n-1
    for ;first+1>0 && second+1>0; {
        
        if(nums1[first]>nums2[second]){
            nums1[position]=nums1[first]
            first--
        }else{
            nums1[position]=nums2[second]
            second--
        }
        position--
    }
    
    for ;second+1>0;{
        nums1[position]=nums2[second]
        second--
        position--
        
    }
    
}
