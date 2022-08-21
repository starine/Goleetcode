/*
* @Author: starine
* @Date:   2022/7/24 10:54
 */

package main

import (
	"fmt"
	"sort"
)

type FoodRatings struct {
	Foods []string
	Cuisines []string
	Ratings []int
	Pos map[string]int
	Me map[string][]int
}


func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	pos := make(map[string]int)
	for i, food := range foods {
		pos[food] = i
	}
	me := make(map[string][]int)
	for i, cuisine := range cuisines {
		me[cuisine] = append(me[cuisine], i)
	}
	return FoodRatings{foods,cuisines,ratings,pos,me}
}


func (this *FoodRatings) ChangeRating(food string, newRating int)  {
	this.Ratings[this.Pos[food]] = newRating

}


func (this *FoodRatings) HighestRated(cuisine string) string {
	s := this.Me[cuisine]
	var res []string
	max := 0
	for _, idx := range s {
		if this.Ratings[idx] > max{
			max = idx
		}
	}
	for _, i2 := range s {
		if this.Ratings[i2] == this.Ratings[max] {
			res = append(res,this.Foods[i2])
		}
	}
	sort.Strings(res)
	return res[0]


}
func countExcellentPairs(nums []int, k int) int64 {
	var res int64
	set := map[int]struct{}{}
	for _, num := range nums {
		set[num] = struct{}{}
	}
	for n1,_ := range set {
		for n2,_ := range set {
			fmt.Println(n1,n2)
			if countOnes(n1,n2) >= k {
				res++
			}
		}
	}
	return res
}
func countOnes(nums1,nums2 int) int{
	or := nums1 | nums2
	an := nums1 & nums2
	res := 0
	for ; or > 0; or &= (or-1){
		res++
	}
	for ; an > 0;an &= (an-1){
		res++
	}
	return res
}
func main() {
	nums := []int{1,2,3,1}
	fmt.Println(countExcellentPairs(nums,3))
}