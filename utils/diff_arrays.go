package utils

// DiffArrays 接受旧数组和新数组，返回新增和删除的元素数组
func DiffArrays(oldArray, newArray []string) (added []string, removed []string) {
	// 创建一个map来存储旧数组的元素，便于快速查找
	oldMap := make(map[string]struct{})
	for _, item := range oldArray {
		oldMap[item] = struct{}{} // 值为空结构体，只关心键
	}

	// 找出新增的元素
	for _, item := range newArray {
		if _, exists := oldMap[item]; !exists {
			added = append(added, item)
		}
	}

	// 找出删除的元素
	newMap := make(map[string]struct{})
	for _, item := range newArray {
		newMap[item] = struct{}{} // 值为空结构体，只关心键
	}

	for _, item := range oldArray {
		if _, exists := newMap[item]; !exists {
			removed = append(removed, item)
		}
	}
	return
}
