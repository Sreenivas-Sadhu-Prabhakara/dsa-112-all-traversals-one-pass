package main

import (
	"reflect"
	"testing"
)

func TestAllTraversals(t *testing.T) {
	root := tn(1, tn(2, tn(4, nil, nil), tn(5, nil, nil)), tn(3, nil, nil))
	pre, in, post := AllTraversals(root)
	if !reflect.DeepEqual(pre, []int{1, 2, 4, 5, 3}) {
		t.Errorf("pre = %v", pre)
	}
	if !reflect.DeepEqual(in, []int{4, 2, 5, 1, 3}) {
		t.Errorf("in = %v", in)
	}
	if !reflect.DeepEqual(post, []int{4, 5, 2, 3, 1}) {
		t.Errorf("post = %v", post)
	}

	pre, in, post = AllTraversals(nil)
	if len(pre)+len(in)+len(post) != 0 {
		t.Errorf("nil tree should give three empty slices")
	}
}
