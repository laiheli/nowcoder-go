package problems

import (
	"testing"
)

var (
	arr       = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	list      = NewListNode(arr)
	cycleList = NewCycleListNode(arr, 10)
	tree      = NewTreeNode(arr, 0)
)

// NC1 大数加法
func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(`999999999999999999`, `123456789123456789`)
	}
}

// NC2 重排链表
func BenchmarkReorderList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reorderList(list)
	}
}

// NC3 链表中环的入口结点
func BenchmarkEntryNodeOfLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		entryNodeOfLoop(list)
	}
}

// NC4 判断链表中是否有环
func BenchmarkHasCycle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hasCycle(cycleList)
	}
}

// NC5 二叉树根节点到叶子节点的所有路径和
func BenchmarkSumNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumNumbers(tree)
	}
}

// NC6 二叉树中的最大路径和
func BenchmarkMaxPathSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxPathSum(tree)
	}
}

// NC7 买卖股票的最好时机(一)
func BenchmarkMaxProfit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxProfit(arr)
	}
}

// NC8 二叉树中和为某一值的路径(二)
func BenchmarkFindPath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findPath(tree, 22)
	}
}

// NC9 二叉树中和为某一值的路径(一)
func BenchmarkHasPathSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hasPathSum(tree, 22)
	}
}

// NC10 大数乘法
func BenchmarkSolveMulti(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solveMulti(`9999999999`, `9999999999`)
	}
}

// NC11 将升序数组转化为平衡二叉搜索树
func BenchmarkSortedArrayToBST(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortedArrayToBST(arr)
	}
}

// NC12 重建二叉树
func BenchmarkReConstructBinaryTree(b *testing.B) {
	x, y := []int{1, 2, 4, 7, 3, 5, 6, 8}, []int{4, 7, 2, 1, 5, 3, 8, 6}
	for i := 0; i < b.N; i++ {
		reConstructBinaryTree(x, y)
	}
}

// NC13 二叉树的最大深度
func BenchmarkMaxDepth(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxDepth(tree)
	}
}

// NC14 按之字形顺序打印二叉树
func BenchmarkPrintTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		printTree(tree)
	}
}

// NC15 求二叉树的层序遍历
func BenchmarkLevelOrder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		levelOrder(tree)
	}
}

// NC16 对称的二叉树
func BenchmarkIsSymmetrical(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isSymmetrical(tree)
	}
}

// NC17 最长回文子串
func BenchmarkGetLongestPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getLongestPalindrome(`abbba`)
	}
}

// NC18 顺时针旋转矩阵
func BenchmarkRotateMatrix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rotateMatrix([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}, 4)
	}
}

// NC19 连续子数组的最大和
func BenchmarkFindGreatestSumOfSubArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findGreatestSumOfSubArray(arr)
	}
}
