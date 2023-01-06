func maxIceCream(costs []int, coins int) int {
    sort.Ints(costs)
    
    if(coins<costs[0]){
        return 0
    }
    count:=0
    for _,val:=range(costs){
        if(val<=coins){
            coins-=val
            count++
        }
    }
    return count
}
