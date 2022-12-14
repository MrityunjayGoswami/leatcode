func kidsWithCandies(candies []int, extraCandies int) []bool {
    greatest := 0
    res := make([]bool, len(candies))
    for _, v := range candies {
        if v > greatest { greatest = v }
    }
    for i, v := range candies {
        if v + extraCandies >= greatest { res[i] = true }
    }
    return res
}