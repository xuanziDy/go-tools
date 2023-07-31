package slice

import (
	"errors"
)

// 删除指定下标元素
func Delete[T any](s []T, index int) ([]T, error) {
	len := len(s)
	if index > len || index < 0 {
		return nil, errors.New("下标越界")
	}

	//第一种
	/*

		new := make([]T, len(s)-1) //不指定容量就与长度一致
		for k, v := range s {
			if k < index {
				new[k] = v
			}
			if k > index {
				new[k-1] = v
			}
		}
	*/

	//第二种
	/*
		new := s[:index]
		for i := index + 1; i < len; i++ {
			new = append(new, s[i])
		}
	*/

	//第三种
	j := 0
	for i, v := range s {
		if i != index {
			s[j] = v
			j++
		}
	}
	new := s[:j]
	return new, nil
}
