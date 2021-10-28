package dp

/*
信封嵌套问题：给定一些标记了宽度和高度的信封，宽度和高度以整数对的形式(w,h)出现。当另一个信封
的宽度和高度都比这个信封大的时候，这个信封就可以放入另一个信封里，就像俄罗斯套娃一样。

请计算最多能有多少个信封能组成一组“俄罗斯套娃”信封。例如：
输入：envelopes = [[5,4],[6,4],[6,7],[2,3]]
输出：3
解释：最多嵌套信封个数为3，组合为：[2,3] -> [5,4] -> [6,7]
*/

type Envelope struct {
	Width  int
	Height int
}

func (a Envelope) equal(b Envelope) bool {
	if a.Width == b.Width && a.Height == b.Height {
		return true
	}
	return false
}

func envelopeSliceEqual(a, b []Envelope) bool {
	if len(a) != len(b) {
		return false
	}
	//遍历判断每一个元素是否相等
	for i := 0; i < len(a); i++ {
		if a[i].equal(b[i]) == false {
			return false
		}
	}

	return true
}

/*
sortEnvelopes
利用快速排序算法按照信封的宽度对envelopes进行升序排序。当两个信封的宽度相同时，按照
信封高度的降序进行排序。快速排序算法的原理如下：
1. 选择数组的头或尾巴作为基准
2. 根据这个基准将数组分成两部分。比基准小的数在基准数左边，比基准大的数在基准数右边
3. 递归对基准数左右两部分的数组进行快速排序
*/
func sortEnvelopes(envelopes []Envelope, i, j int) {
	if i < j {
		index := partition(envelopes, i, j)
		sortEnvelopes(envelopes, i, index-1)
		sortEnvelopes(envelopes, index+1, j)
	}
}

//partition 将数组envelopes[left,right]部分的数据按照基准分成两部分，并返回基准的索引
func partition(envelopes []Envelope, i, j int) int {
	pivot := envelopes[i]
	for i < j {
		//从尾部开始，找到第一个小于或等于pivot的数的位置并将其值赋值为i
		for i < j {
			//有两种情况会发生赋值：1. j的width小于pivot的width；2.j的width等于pivot的width，但是j的height大于pivot的height
			if envelopes[j].Width < pivot.Width ||
				(envelopes[j].Width == pivot.Width && envelopes[j].Height > pivot.Height) {
				envelopes[i] = envelopes[j]
				i++
				break
			} else {
				j--
			}
		}
		//从头部开始，找到第一个width大于或等于pivot的数的位置
		for i < j {
			//有两种情况会发生赋值：1. i的width大于pivot的width；2.i的width等于j的width，但是i的height小于pivot的height
			if envelopes[i].Width > pivot.Width || (envelopes[i].Width == pivot.Width && envelopes[i].Height < pivot.Height) {
				envelopes[j] = envelopes[i]
				j--
				break
			} else {
				i++
			}
		}
	}
	envelopes[i] = pivot
	return i
}

/*
MaxEnvelopes
像这种求最值的问题，一般都是考虑使用动态规划求解。题目要求只有当一个信封的宽度和高度均小于另一个
信封时，才能将其放入另一个信封当中。因此，此题是最长递增子序列的一个变种，解法也比较巧妙。首先，
按照宽度w对信封进行排序，然后以高度为基准，找出最长递增子序列即可。但需要注意的是，当两个信封的宽
度相同时，这两个信封的顺序需要按照高度的降序排列，因为宽度相等的信封不能嵌套。
					[1,8]
					[2,3]
					[5,4]  这里必须按照高度的降序，否则按照2 4的顺序，会被归入到递增子序列中。
					[5,2]  然而，宽度相同的信封是不允许嵌套的。下面的[6,7]和[6,4]也是如此。
					[6,7]
					[6,4]
*/
func MaxEnvelopes(envelopes []Envelope) int {
	//先对所有信封进行排序
	temp := envelopes
	sortEnvelopes(temp, 0, len(temp)-1)
	//然后对排序后的信封进行按照height的最长递增子序列搜索
	var lis []int
	for _, enl := range temp {
		lis = append(lis, enl.Height)
	}
	return Lis(lis)
}
