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
