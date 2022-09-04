package utils

func GetPage(pageSize, pageNum int) int {

	if pageNum == 0 {
		return pageNum
	}
	return (pageNum - 1) * pageSize
}
