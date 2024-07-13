package main

func twoSum(nums []int, target int) []int {
    resultList := []int {}
    for i:=0; i<len(nums); i++ {
        for j:=i+1; j<len(nums); j++ {
            if nums[i]+nums[j] == target {
                resultList = append(resultList, i, j)
            }
        }
    }
    return resultList
}