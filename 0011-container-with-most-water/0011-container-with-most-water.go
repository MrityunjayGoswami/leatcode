func maxArea(height []int) int {
    left := 0
    right := len(height) - 1
    maxArea := 0

    for left < right {
        width := right - left

        var h int
        if height[left] < height[right] {
            h = height[left]
        } else {
            h = height[right]
        }

        area := h * width
        if area > maxArea {
            maxArea = area
        }

        if height[left] < height[right] {
            left++
        } else {
            right--
        }
    }

    return maxArea
}