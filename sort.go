package golib
//
//import (
//"sort"
//)
//
////分类列表信息排序
//type SortCategoryListSlice []mredis.SignInCumulativeMonthTop
//
//func (c SortCategoryListSlice) Len() int {
//	return len(c)
//}
//
//func (c SortCategoryListSlice) Swap(i, j int) {
//	c[i], c[j] = c[j], c[i]
//}
//
//func (c SortCategoryListSlice) Less(i, j int) bool {
//	return c[i].Count > c[j].Count
//}
//
//func Sort(lists []mredis.SignInCumulativeMonthTop) ([]mredis.SignInCumulativeMonthTop) {
//	data := SortCategoryListSlice(lists)
//	sort.Sort(data)
//	return []mredis.SignInCumulativeMonthTop(data)
//}
//
