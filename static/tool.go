package static

func GetGenderStr(gender int64) string {
	genderStr := ""
	switch gender {
	case 1:
		genderStr = "男"
	case 2:
		genderStr = "女"
	default:
		genderStr = "未知"
	}
	return genderStr
}
