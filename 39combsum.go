package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
  singleperm:=make([]int,0)
    output:=make([][]int,0)
    backtrack(candidates,singleperm,&output,target,0,0)

    fmt.Println(output)
    return output

    //goal

}


func backtrack(candidates []int, singlecomb []int ,output *[][]int,target int,total int,index int){//,mapper map[int]bool) {   //goal
if(target==total){
  temp:=make([]int,len(singlecomb))
  copy(temp,singlecomb)
  *output=append(*output,temp)
  return
}

if(total>target || index>=len(candidates)){
  return
}

  singlecomb=append(singlecomb,candidates[index])
  backtrack(candidates,singlecomb,output,target,total+candidates[index],index)
  singlecomb=singlecomb[:len(singlecomb)-1]
  backtrack(candidates,singlecomb,output,target,total,index+1)


}
func main(){
  combinationSum([]int{2,3,5},8)
}
