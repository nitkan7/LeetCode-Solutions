package main

import "fmt"
import "math"


func maxSubArray(nums []int) int {
    globalMaximum:=math.MinInt
    localMaximum:=0
    length:=len(nums)
    for i:=0;i<length;i++ {
        localMaximum=int(math.Max(float64(nums[i]),float64(nums[i]+localMaximum)))
        
        if(localMaximum>globalMaximum){
            globalMaximum=localMaximum
        }
        
    }
    return globalMaximum
}
