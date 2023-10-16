package utils

type timeDetail struct {
	StartTime string `json:"startTime" bson:"startTime"`
	EndTime   string `json:"endTime" bson:"endTime"`
}

// 注意时间从小到大排序
var blackList []timeDetail = []timeDetail{
	{
		StartTime: "2022-01-04 10:00:00",
		EndTime:   "2022-02-01 00:00:00",
	},
	{
		StartTime: "2022-11-04 12:00:00",
		EndTime:   "2022-11-12 11:00:00",
	},
	{
		StartTime: "2022-11-17 00:00:00",
		EndTime:   "2022-11-21 00:00:00",
	},
	{
		StartTime: "2022-11-24 08:00:00",
		EndTime:   "2022-11-26 15:00:00",
	},
	{
		StartTime: "2022-12-06 08:00:00",
		EndTime:   "2022-12-10 14:00:00",
	},
	{
		StartTime: "2022-12-24 12:00:00",
		EndTime:   "2023-01-06 12:00:00",
	},
	{
		StartTime: "2023-01-10 12:00:00",
		EndTime:   "2023-01-14 10:00:00",
	},
	{
		StartTime: "2023-01-10 12:00:00",
		EndTime:   "2023-01-14 10:00:00",
	},
	{
		StartTime: "2023-02-03 14:00:00",
		EndTime:   "2023-02-08 10:00:00",
	},
	{
		StartTime: "2023-02-17 08:00:00",
		EndTime:   "2023-02-19 10:00:00",
	},
	{
		StartTime: "2023-02-26 12:00:00",
		EndTime:   "2023-03-01 16:00:00",
	},
	{
		StartTime: "2023-03-04 16:00:00",
		EndTime:   "2023-03-10 12:00:00",
	},
	{
		StartTime: "2023-03-17 12:00:00",
		EndTime:   "2023-03-22 12:00:00",
	},
}

// IsValidTime IsValTime 判断是否在该时间段内
func IsValidTime(startTime, endTime string) bool {

	if startTime == "" || endTime == "" {
		return false
	}

	//黑名单为空
	if len(blackList) <= 0 {
		return false
	}

	if endTime < blackList[0].StartTime {
		return false
	} else if startTime > blackList[len(blackList)-1].EndTime {
		return false
	}

	//不能有时间交集
	for _, value := range blackList {

		if startTime > value.StartTime && startTime < value.EndTime {
			return true
		} else if endTime > value.StartTime && endTime < value.EndTime {
			return true
		}

	}

	return false
}
