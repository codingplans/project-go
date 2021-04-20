package _32

import (
	"math"
	"math/rand"
	"testing"
)

// 220. 存在重复元素 III

// 给你一个整数数组 nums 和两个整数 k 和 t 。请你判断是否存在 两个不同下标 i 和 j，使得 abs(nums[i] - nums[j]) <= t ，同时又满足 abs(i - j) <= k 。
//
// 如果存在则返回 true，不存在返回 false。
//
//
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/contains-duplicate-iii
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func Test_220(t *testing.T) {

	// nums := []int{1, 5, 9, 1, 5, 9}
	// nums := []int{0}
	nums := []int{1, 2, 1, 1}

	rs := containsNearbyAlmostDuplicate(nums, 1,
		0)
	println(rs)
}

func containsNearbyAlmostDup22licate(nums []int, k int, t int) bool {
	l := len(nums) - 1
	if l == 0 {
		return false
	}
	for n1 := range nums {

		// 尾部
		n2 := l - n1
		if n2-n1 <= k && (n2-n1)*-1 <= k {
			if nums[n1]-nums[n2] <= t && (nums[n1]-nums[n2])*-1 <= t {
				return true
			}
		}
		println(k, t, n1-n2, nums[n1], nums[n2], nums[n1]-nums[n2])
	}
	return false
}

type node struct {
	ch       [2]*node
	priority int
	val      int
}

func (o *node) cmp(b int) int {
	switch {
	case b < o.val:
		return 0
	case b > o.val:
		return 1
	default:
		return -1
	}
}

func (o *node) rotate(d int) *node {
	x := o.ch[d^1]
	o.ch[d^1] = x.ch[d]
	x.ch[d] = o
	return x
}

type treap struct {
	root *node
}

func (t *treap) _put(o *node, val int) *node {
	if o == nil {
		return &node{priority: rand.Int(), val: val}
	}
	d := o.cmp(val)
	o.ch[d] = t._put(o.ch[d], val)
	if o.ch[d].priority > o.priority {
		o = o.rotate(d ^ 1)
	}

	return o
}

func (t *treap) put(val int) {
	t.root = t._put(t.root, val)
}

func (t *treap) _delete(o *node, val int) *node {
	if d := o.cmp(val); d >= 0 {
		o.ch[d] = t._delete(o.ch[d], val)
		return o
	}
	if o.ch[1] == nil {
		return o.ch[0]
	}
	if o.ch[0] == nil {
		return o.ch[1]
	}
	d := 0
	if o.ch[0].priority > o.ch[1].priority {
		d = 1
	}
	o = o.rotate(d)
	o.ch[d] = t._delete(o.ch[d], val)
	return o
}

func (t *treap) delete(val int) {
	t.root = t._delete(t.root, val)
}

func (t *treap) lowerBound(val int) (lb *node) {
	for o := t.root; o != nil; {
		switch c := o.cmp(val); {
		case c == 0:
			lb = o
			o = o.ch[0]
		case c > 0:
			o = o.ch[1]
		default:
			return o
		}
	}
	return
}

func containsNearbyAlmostDuplicate(nums []int, k, t int) bool {
	for i := range nums {
		for j := i + 1; j <= i+k && j < len(nums); j++ {
			if math.Abs(float64(nums[i]-nums[j])) <= float64(t) {
				return true
			}
		}
	}

	// set := &treap{}
	// for i, v := range nums {
	// 	if lb := set.lowerBound(v - t); lb != nil && lb.val <= v+t {
	// 		return true
	// 	}
	// 	set.put(v)
	// 	if i >= k {
	// 		set.delete(nums[i-k])
	// 	}
	// }
	return false
}
