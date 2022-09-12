package main

import "fmt"
import "strconv"

func main(){
	nums:=[]int{0,2,3,4,6,8,9}
	
	//need an output slice of strings to store the output
	output:=make([]string,0)

	var iter string
//	var left int
	for i,val:=range(nums){
		if(i==0){
		iter=fmt.Sprintf("%s",strconv.Itoa(val))
		//left=0
		continue			
		}
		if(val-nums[i-1]==1){
			inter:=fmt.Sprintf("->%s",strconv.Itoa(val))
			iter=iter+inter
			fmt.Println(val)
		}else{
		//	left=i
			//inter:=fmt.Sprintf("->%s",val)
			//iter=iter+inter
			output=append(output,iter)		
			iter=fmt.Sprintf("%s",strconv.Itoa(val))
		}


		
	}
	output=append(output,iter)
fmt.Println(output)
}