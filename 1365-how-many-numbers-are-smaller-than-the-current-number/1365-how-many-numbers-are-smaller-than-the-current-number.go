func smallerNumbersThanCurrent(nums []int) []int {
    res := []int{}
    for i := range nums{
        count := 0
        for j:= range nums{
            if nums[i]>nums[j]{
                count +=1
            }
        }
        res = append(res, count)
    }
    return res
}