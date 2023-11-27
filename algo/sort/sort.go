package sort

// SelectionSort 选择排序，每轮从未排序区间选择最小元素，将其放到已排序区间末尾
func SelectionSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

// BubbleSort 遍历所有元素，每轮选择一个最大的放在数组末尾
func BubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		flag := false
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = true
			}
		}
		if !flag {
			return
		}
	}
}

// InsertionSort 像洗牌一样，找到合适的位置插入。遍历每个元素，与前面有序区间最后一个开始逐个比较，比前面的小就交换
func InsertionSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j > 0; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
}

// QuickSort1 快速排序【非就地排序】
// 选一个值，作为中间值
// 小于中间值的放一个小值数组中，大于中间值的放一个大值数组中
// 拼接小值数组、中间值、大值数组
// 并递归对小值数组和大值数组作同样的操作，递归至数组中剩一个元素
func QuickSort1(nums []int) []int {
	return quickSort1(nums)
}

func quickSort1(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	l := make([]int, 0)
	r := make([]int, 0)
	mid := nums[0] // 随机避免反向适应性
	for i := 1; i < len(nums); i++ {
		if nums[i] <= mid {
			l = append(l, nums[i])
		} else {
			r = append(r, nums[i])
		}
	}

	return append(quickSort1(l), append([]int{mid}, quickSort1(r)...)...)
}

// QuickSort2 快速排序【就地排序】
// 选择数组中的某个元素作为“基准数”，将所有小于基准数的元素移到其左侧，而大于基准数的元素移到其右侧
// 选取数组最左端元素作为基准数，初始化两个指针 i 和 j 分别指向数组的两端。
// 设置一个循环，在每轮中使用 i（j）分别寻找第一个比基准数大（小）的元素，然后交换这两个元素。
// 循环执行步骤 2. ，直到 i 和 j 相遇时停止，最后将基准数交换至两个子数组的分界线
func QuickSort2(nums []int) {
	left, right := 0, len(nums)-1
	quickSort2(nums, left, right)
}

func quickSort2(nums []int, left, right int) {
	if left >= right {
		return
	}

	pivot := partition(nums, left, right)
	partition(nums, left, pivot-1)
	partition(nums, pivot+1, right)
}

func partition(nums []int, left, right int) int {
	i, j := left, right
	for i < j {
		for i < j && nums[j] >= nums[left] {
			j--
		}
		for i < j && nums[i] <= nums[left] {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[left], nums[i] = nums[i], nums[left]

	return i
}

// MergeSort 归并排序
// 递归地在中间位置划分数组
// 子数组长度为1时停止划分，开始合并
// 合并的同时排序
func MergeSort(nums []int) {
	mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	mid := (left + right) / 2
	mergeSort(nums, left, mid)
	mergeSort(nums, mid+1, right)

	merge(nums, left, mid, right)
}

func merge(nums []int, left, mid, right int) {
	tmp := make([]int, right-left+1)
	i, j, k := left, mid+1, 0

	for i <= mid && j <= right {
		if nums[i] <= nums[j] {
			tmp[k] = nums[i]
			i++
		} else {
			tmp[k] = nums[j]
			j++
		}
		k++
	}
	for i <= mid {
		tmp[k] = nums[i]
		i++
		k++
	}

	for j <= right {
		tmp[k] = nums[j]
		j++
		k++
	}

	for k := 0; k < len(tmp); k++ {
		nums[left+k] = tmp[k]
	}
}
