// 4. Median of Two Sorted Arrays
// Runtime: 12 ms, faster than 85.82% of Go online submissions for Median of Two Sorted Arrays.
// Memory Usage: 6 MB, less than 42.28% of Go online submissions for Median of Two Sorted Arrays.

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    var len1 = len(nums1)
    var len2 = len(nums2)
    var concat []int = make([]int, len1+len2)
    i := 0
    j := 0
    for idx:=0; idx < len(concat); idx++ {
        if i < len1 && j < len2 {
            if nums1[i] < nums2[j] {
                concat[idx] = nums1[i]
                i++
            } else {
                concat[idx] = nums2[j]
                j++
            }
        } else {
            if i < len1 && j >= len2 {
                concat[idx] = nums1[i]
                i++
            } else if i >= len1 && j < len2 {
                concat[idx] = nums2[j]
                j++
            }
        }
    }
    //fmt.Println(concat)
    
    if len(concat) % 2 == 0 {
        idx := len(concat)/2
        return (float64(concat[idx-1])+float64(concat[idx]))/2.0
    } else {
        return float64(concat[len(concat)/2]) 
    }
}