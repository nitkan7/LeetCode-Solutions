package main

import "fmt"
import "sort"


func main(){
fmt.Println(permuteUnique([]int{1,2,1}))
}

func permuteUnique(nums []int) [][]int {

	output:=make([][]int,0)
singleperm:=make([]int,0)
checker:=make([]bool,len(nums))

sort.Ints(nums)
backtrack(nums,&output,singleperm,checker)
return output
}

func backtrack(nums []int,output *[][]int,singleperm []int,checker []bool){
	if(len(singleperm)==len(nums)){
		temp:=make([]int,len(singleperm))
		copy(temp,singleperm)
        *output=append(*output,temp)
		//append
		return
	}

	if(len(singleperm)>len(nums)){
		return
	}

	for i,val:=range(nums){
		//for duplicate entries
		if(checker[i] || (i>0 && (val==nums[i-1]) && !checker[i-1] )){
			continue
		}

		checker[i]=true
		singleperm=append(singleperm,val)

		backtrack(nums,output,singleperm,checker)

		checker[i]=false
		singleperm=singleperm[:len(singleperm)-1]

	}
}

