package main

import(
	"fmt"
)
func Max(a,b int)int{
	if a >b {
		return a
	}
	return b
}
func main(){
	var n int
	fmt.Scan(&n)
	var arr  =make([]int,n)
	var dp  =make([]int,n)
	for k:=0;k<n;k++{
		fmt.Scan(&arr[k])
		dp[k]=1
	}
	maxlen :=dp[0]
	for i:=1;i<n;i++{
		for j:=0;j<i;j++{
			if arr[i] > arr[j]{
				dp[i]= Max(dp[j]+1,dp[i])
			}
		}
		maxlen = Max(maxlen,dp[i])
	}
	fmt.Println("maxlen",maxlen)
}