func largestOverlap(img1 [][]int, img2 [][]int) int {
    output:=0
    
    for i:= -len(img1[0])+1; i<len(img1[0]);i++{
        for j:= -len(img1[0])+1; j<len(img1[0]);j++{
            output=int(math.Max(float64(output),float64(helper(img1,img2,i,j))))
         }     
    } 
    return output
}

func helper(A [][]int, B [][]int, rowOffset int, colOffset int)int {
    count:=0
    for i:=0;i<len(A[0]);i++{
     for j:=0;j<len(A[0]);j++{
         if(i+rowOffset < 0 || i+rowOffset >=len(A[0]) || j+colOffset <0 || j+colOffset >=len(A[0]) ){
             continue
         }
         if(A[i][j] + B[i+rowOffset][j+colOffset] == 2){
             count++
         }
        }   
    }
    
    return count
}
