package jianzhi_offer

import (
	"reflect"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("2 字符串空格替换", func(t *testing.T) {
		s := "We are happy"
		want := "We%20are%20happy"
		got := solution(s)
		if !reflect.DeepEqual(got,  want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	字符串空格替换
	先统计字符串中空格的数量，然后可以直接统计字符串的长度以及需要分配的字符串的长度
	最后，遍历原字符串，遇到空格就替换，其他字符直接复制
*/
func solution(s string) string {
	var res string
	for _, s := range s {
		if s == ' ' {
			res = res + "%20"
		} else {
			res = res + string(s)
		}
	}
	return res
}

///*
// C++ version
// */
//public void solution(char* s, int length) {
//	if (s == NULL || length <= 0) {
//		return;
//	}
//	// 假设 length 是输入字符数组的长度这里已经包含了 '\0'
//	blankCount = 0;
//	for (i = 0; i < length; i++) {
//		if (s[i] == ' ') {
//			blankCount++;
//		}
//	// 新的长度 这里只在原数组基础上修改，所以需要从后面遍历才不会因为复制问题导致原来的数据被覆盖了
//    newLength = blankCount * 2 + length;
//    for (i = length - 1, j = i; i >= 0; i--) {
//    	if s[i] == ' ' {
//    		s[j] = '0'
//    		s[j-1] = '2'
//    		s[j-2] = '%'
//    		j -= 3
//		} else {
//			s[j] = s[i]
//			j--
//		}
//    }
//    }
//}
