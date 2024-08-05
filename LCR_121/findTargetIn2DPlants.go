package _121

/*
m*n 的二维数组 plants 记录了园林景观的植物排布情况，具有以下特性：

每行中，每棵植物的右侧相邻植物不矮于该植物；
每列中，每棵植物的下侧相邻植物不矮于该植物。

请判断 plants 中是否存在目标高度值 target。

输入：plants = [[2,3,6,8],[4,5,8,9],[5,9,10,12]], target = 8

输出：true
*/

func findTargetIn2DPlants(plants [][]int, target int) bool {
	if len(plants) == 0 || len(plants[0]) == 0 {
		return false
	}
	i, j := 0, len(plants[0])-1
	for i < len(plants) && j >= 0 {
		if plants[i][j] == target {
			return true
		}
		if plants[i][j] > target {
			j--
			continue
		}
		if plants[i][j] < target {
			i++
		}
	}
	return false
}
