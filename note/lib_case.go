package note

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func SortCase01() {
	fmt.Println("Case1: 对基本类型切片排序")
	nums := []int{3, 6, 2, 8, 1, 0, 3}
	sort.Ints(nums)
	fmt.Println(nums)

	// 字符串排序 (按字典序)
	strs := []string{"banana", "apple", "cherry"}
	sort.Strings(strs)
	fmt.Println(strs)
}

func SortCase02() {
	fmt.Println("Case2: 自定义排序")
	nums := []int{3, 6, 2, 8, 1, 0, 3}

	// 降序
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	fmt.Println(nums)

	strs := []string{"banana", "apple", "cherry"}

	// 按字符串长度升序排序
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})
	fmt.Println(strs)

	type Person struct {
		ID   int
		Name string
	}
	people := []Person{{10000, "Nameless"}, {10086, "namelyzz"}, {12345, "无名客"}}

	// 对结构体按 ID 编号降序排列
	sort.Slice(people, func(i, j int) bool {
		return people[i].ID > people[j].ID
	})
	fmt.Println(people)
}

func SortCase03() {
	fmt.Println("Case3: 稳定排序")

	// 对 pairs 按第一个元素排序
	pairs := [][2]int{{1, 5}, {2, 3}, {1, 2}}
	sort.SliceStable(pairs, func(i, j int) bool {
		return pairs[i][0] < pairs[j][0]
	})

	// 稳定排序后，{1,5} 保证仍在 {1,2} 的前面
	fmt.Println(pairs) // 输出: [[1 5] [1 2] [2 3]]
}

func SortCase04() {
	fmt.Println("Case4: 二分查找，二分查找的前提是：切片必须是已经按升序排序好的。")
	nums := []int{1, 3, 5, 7, 9} // 必须已排序

	index := sort.SearchInts(nums, 5)
	fmt.Println(index) // 输出: 2 (找得到，索引为2)

	index = sort.SearchInts(nums, 6)
	fmt.Println(index) // 输出: 3 (没找到，返回第一个大于6的元素‘7’的索引)

	target := 6
	index = sort.Search(len(nums), func(i int) bool {
		return nums[i] >= target // 可以通过修改 f 函数来实现查找第一个 > target、最后一个 < target 等复杂逻辑
	})
	fmt.Println(index) // 输出: 3 (nums[3] = 7 >= 6)

	// 判断是否存在 target 的方式
	if index < len(nums) && nums[index] == target {
		fmt.Println("Found at index:", index)
	} else {
		fmt.Println("Not found, could be inserted at index:", index)
	}
}

func SlicesCase01() {
	fmt.Println("Case1: 排序与二分查找")
	nums := []int{3, 1, 4, 1, 5}
	slices.Sort(nums)
	fmt.Printf("原地修改后的 nums: %v\n", nums)

	// 也可以对字符串、结构体等任何可排序的类型排序
	strs := []string{"banana", "apple", "cherry"}
	slices.Sort(strs)
	fmt.Println(strs) // 输出: [apple banana cherry]

	// 检查是否已排序
	fmt.Printf("是否已排序：%t\n", slices.IsSorted(nums))

	// 二分查找
	idx, ok := slices.BinarySearch(nums, 3)
	fmt.Println(idx, ok) // 输出: 2 true

	idx, ok = slices.BinarySearch(nums, 2)
	fmt.Println(idx, ok) // 输出: 2 false (因为2应该插入在索引2的位置，但没找到)

	slices.Reverse(nums)
	fmt.Println(nums)
}

func SlicesCase02() {
	fmt.Println("Case2: 切片比较")
	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{1, 2, 4}

	// 字典序比较 (s1 < s2: -1; s1 == s2: 0; s1 > s2: 1)
	fmt.Println(slices.Compare(s1, s2)) // 输出: 0, 说明两者相等
	fmt.Println(slices.Compare(s1, s3)) // 输出: -1, 因为 s1 最后是 3，s3 最后是 4
	fmt.Println(slices.Compare(s3, s1)) // 输出: 1

	// 判断是否相等
	fmt.Println(slices.Equal(s1, s2))
	fmt.Println(slices.Equal(s1, s3))
}

func SlicesCase03() {
	fmt.Println("Case3: 元素查找")
	nums := []int{1, 2, 5, 3, 2}

	fmt.Printf("判断元素是否存在（存在为 true）： %t\n", slices.Contains(nums, 3))
	fmt.Printf("判断元素是否存在（不存在为 false）： %t\n", slices.Contains(nums, 99))

	fmt.Printf("查找第一个元素的位置(找到会输出位置)： %d\n", slices.Index(nums, 5))
	fmt.Printf("查找第一个元素的位置(找不到返回-1)： %d\n", slices.Index(nums, 99))

	// 查找第一个偶数
	idx := slices.IndexFunc(nums, func(i int) bool {
		return nums[i]%2 == 0
	})
	fmt.Println(idx) // 输出: 1 (元素2的索引)

	fmt.Printf("切片的最大值：%d\n", slices.Max(nums))
	fmt.Printf("切片的最小值：%d\n", slices.Min(nums))
	// 字符串切片也可以找最大值最小值，按字典序比较
}

func StrconvCase() {
	// *** 字符串 -> 整数 ***

	// 1. Atoi (ASCII to Integer) - 首选, 解析一个十进制字符串
	s := "123"
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d, %T\n", num, num) // 输出: 123, int

	// 2. ParseInt - 功能更强，可指定进制和位宽
	// base: 进制 (2 到 36, 0表示根据字符串前缀判断，如"0x"是16进制)
	// bitSize: 目标整数位宽 (0, 8, 16, 32, 64 -> 对应 int, int8, int16, int32, int64)
	s64 := "-FF"
	num64, err := strconv.ParseInt(s64, 16, 64)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d\n", num64) // 输出: -255

	// *** 整数 -> 字符串 ***

	// 1. Itoa (Integer to ASCII) - 首选, 将int转换为它的十进制字符串形式
	i := 456
	str := strconv.Itoa(i)
	fmt.Printf("%s, %T\n", str, str) // 输出: 456, string

	// 2. FormatInt - 功能更强，可指定进制
	var i64 int64 = 255
	strHex := strconv.FormatInt(i64, 16) // 转换为16进制字符串
	fmt.Println(strHex)                  // 输出: ff
	strBin := strconv.FormatInt(i64, 2)  // 转换为2进制字符串
	fmt.Println(strBin)                  // 输出: 11111111//

	// *** 字符串 -> 浮点数 ***
	fStr := "3.14159"
	fVal, err := strconv.ParseFloat(fStr, 64) // 解析为float64
	if err != nil {
		panic(err)
	}
	fmt.Printf("%f, %T\n", fVal, fVal) // 输出: 3.141590, float64

	// *** 浮点数 -> 字符串 ***
	// strconv.FormatFloat(f float64, fmt byte, prec, bitSize int) string
	// fmt: 格式字符
	//   'f' (-ddd.dddd)
	//   'b' (-ddddp±ddd, 二进制指数)
	//   'e' (-d.dddde±dd, 十进制指数)
	//   'E' (-d.ddddE±dd, 十进制指数)
	//   'g' (大指数用'e'，否则用'f')
	//   'G' (大指数用'E'，否则用'f')
	// prec: 精度（对于'f','e','E'，表示小数点后的位数；对于'g','G'，表示总位数）
	// bitSize: 32 或 64
	pi := 3.1415926535
	sF1 := strconv.FormatFloat(pi, 'f', 2, 64) // 保留2位小数
	fmt.Println(sF1)                           // 输出: 3.14

	sF2 := strconv.FormatFloat(pi, 'e', -1, 64) // 科学计数法，最小精度
	fmt.Println(sF2)                            // 输出: 3.1415926535e+00
}

func StringsCase01() {
	fmt.Println("Case1: 查找与判断")
	s := "leetcode"

	// 判断包含
	fmt.Printf("是否包含：%t\n", strings.Contains(s, "code"))
	fmt.Printf("是否包含：%t\n", strings.Contains(s, "abcd"))

	// 判断前后缀
	fmt.Printf("是否存在前缀：%t\n", strings.HasPrefix(s, "leet"))
	fmt.Printf("是否存在后缀：%t\n", strings.HasSuffix(s, "code"))

	// 查找索引位置
	s = "hello leetcode! hello go!"
	fmt.Printf("第一次出现的位置：%d\n", strings.Index(s, "hello"))
	fmt.Printf("最后一次出现的位置：%d\n", strings.LastIndex(s, "hello"))
}

func StringsCase02() {
	fmt.Println("Case2: 分割、修剪、替换与连接")
	input := "nameless 无名客 jiligulu zzzzzz"
	strs := strings.Split(input, " ")
	fmt.Println(strs, len(strs))

	// 特殊 —— 空分隔符 - 将字符串拆分为单个字符的切片
	chars := strings.Split("nameless", "")
	fmt.Printf("%v, %T", chars, chars[0])

	input = "    无名客     "
	// 修剪首尾的空格字符
	trim1 := strings.TrimSpace(input)

	input = "!!无!!名!客!!!"
	// 修剪首尾的 !
	trim2 := strings.Trim(input, "!")
	// 只修剪左边的 !
	trim3 := strings.TrimLeft(input, "!")
	fmt.Println(trim1, trim2, trim3)

	// 连接
	words := []string{"I", "love", "U"}
	sentence := strings.Join(words, " ")
	fmt.Println(sentence)

	// 替换与大小写
	s := "go is good, go is great"
	replaced := strings.Replace(s, "go", "Go", -1) // 替换所有
	fmt.Println(replaced)
	replacedN := strings.Replace(s, "go", "Go", 1) // 只替换第一个
	fmt.Println(replacedN)

	// 大小写
	lower := strings.ToLower("HELLO")
	fmt.Println(lower) // 输出: "hello"

	upper := strings.ToUpper("world")
	fmt.Println(upper) // 输出: "WORLD"
}
