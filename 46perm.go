package main

import "fmt"

func main(){

fmt.Println(permute([]int{1,2,3,4,5,6}))
}

func permute(nums []int) [][]int {
  //  mapper:=make(map[int]bool)
    singleperm:=make([]int,0)
    output:=make([][]int,0)

    //backtrack(nums,singleperm,&output,mapper)
dfs(nums,singleperm,&output,0)
    return output
}


//goal

//constraint

//choice

func backtrack(nums []int,singleperm []int,output *[][]int,mapper map[int]bool){
    //goal
    if(len(singleperm)==len(nums)){
        *output=append(*output,singleperm)
        return
    }

    for _,val:=range(nums){
        //constraints
        if(mapper[val]==false){
           //make the choice
            mapper[val]=true
            singleperm=append(singleperm,val)

            //call backtrack
            backtrack(nums,singleperm,output,mapper)

            mapper[val]=false
            if len(singleperm) > 0 {
             singleperm = singleperm[:len(singleperm)-1]
            }
        }
    }
  }

    func dfs(nums []int,singleperm []int,output *[][]int,index int){
      if(len(nums)==len(singleperm)){
        temp:=make([]int,len(singleperm))
        copy(temp,singleperm)
        *output=append(*output,temp)
        return
      }

      if(index>len(singleperm)){
        return
      }

      singleperm=append(singleperm,nums[index])
      dfs(nums,singleperm,output,index+1)
      singleperm=singleperm[:len(singleperm) -1]

    }
