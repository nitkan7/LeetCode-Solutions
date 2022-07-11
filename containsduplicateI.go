func containsDuplicate(nums []int) bool {
       //to check whether the given array is distinct or not
       //how??
    //by going through the elements and storing them in a map,
    //if a key is already present,then array contains a duplicate value.
    
    mapper:=make(map[int]int)
    for i:=0;i<len(nums);i++{
        _,present:=mapper[nums[i]]
        if(present){
            return true
        }
        mapper[nums[i]]=i
    }
    return false
}
