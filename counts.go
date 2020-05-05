package main

import (
	"fmt"
	"math/rand"
)

var countss int

func main() {
	countss = 0
	// a, b := many(3, 0, 0)
	// a, b := many(0, 0, 8)
	// println(a, b, countss)

	// maap()

	s := []int{11, 11 - 1, 9, 0, 6, 5, 8, 2, 1, 7, 4, 3}
	for countss < 50 {
		s = append(s, rand.Intn(80))
		countss++
	}
	fmt.Println(s[1:])
	HeapSort(s)
	fmt.Println(s[1:])
}

func maap() {
	bb := new([11]int)
	aa := make([]int, 40, 40)
	cc := make(map[int]int, 3)
	aa = append(aa, 2)
	for k, _ := range aa {
		if k >= 5 {
			// bb[k] = 3
			aa[k] = 1
			cc[k] = 2
		}
	}
	fmt.Println(aa, cc)

	for k, _ := range cc {
		println(k)
	}
	// bb = append(bb,11)
	fmt.Printf("%+v￿,,,%+v", aa, bb)
}
func many(n, m, g int) (x, y int) {

	if n/3 < 1 && m < 1 && g < 4 {
		println("已没有满足兑换的空瓶子，返回结果", n, m)
		return n, m
	}
	g += m
	println("可兑换的空瓶子数量：", n, "已兑换的酒数量：", m, "已经喝的酒总数：", countss, g)

	countss += m
	n += m
	return many(n%3, n/3+g/4, g%4)

}

// 堆排序
// s[0]不用，实际元素从角标1开始
// 父节点元素大于子节点元素
// 左子节点角标为2*k
// 右子节点角标为2*k+1
// 父节点角标为k/2
func HeapSort(s []int) {
	N := len(s) - 1 // s[0]不用，实际元素数量和最后一个元素的角标都为N
	// 构造堆
	// 如果给两个已构造好的堆添加一个共同父节点，
	// 将新添加的节点作一次下沉将构造一个新堆，
	// 由于叶子节点都可看作一个构造好的堆，所以
	// 可以从最后一个非叶子节点开始下沉，直至
	// 根节点，最后一个非叶子节点是最后一个叶子
	// 节点的父节点，角标为N/2
	for k := N / 2; k >= 1; k-- {
		sink(s, k, N)
	}
	// 下沉排序
	for N > 1 {
		swap(s, 1, N) // 将大的放在数组后面，升序排序
		N--
		sink(s, 1, N)
	}
}

// 下沉（由上至下的堆有序化）
func sink(s []int, k, N int) {
	for {
		i := 2 * k
		if i > N { // 保证该节点是非叶子节点
			break
		}
		if i < N && s[i+1] > s[i] { // 选择较大的子节点
			i++
		}
		if s[k] >= s[i] { // 没下沉到底就构造好堆了
			break
		}
		swap(s, k, i)
		k = i
	}
}

func swap(s []int, i int, j int) {
	s[i], s[j] = s[j], s[i]
}
