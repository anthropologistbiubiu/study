package main

//积跬步，至千里。
/*
输入一个非负整数数组numbers，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。
例如输入数组[3，32，321]，则打印出这三个数字能排成的最小数字为321323。
1.输出结果可能非常大，所以你需要返回一个字符串而不是整数
2.拼接起来的数字可能会有前导 0，最后结果不需要去掉前导 0
*/
//JZ45 把数组排成最小的数

func PrintMinNumber(numbers []int) string {
	// write code here
	return "c"
}

func getPos(nums []int,fisrt,last int) int {


	for first <= last {

		var key = nums[first]
		for first < last && key <= nums[last] {
			last--
		}
		nums[first], nums[last] = nums[last], nums[first]
		for first < last && key > nums[first] {
			first++
		}
		nums[first], nums[last] = nums[last], nums[first]

	}
	return first
}
func QuickSort(nums []int) {
	var fisrt ,last = 0
	if 
	var pos = getPos(nums[],0,len(nums)-1)
	getPos(nums,0,pos)
	getPos(nums,post+1,pos)
}
