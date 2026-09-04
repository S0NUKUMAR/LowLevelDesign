package service

import (
	"fmt"
	"sync/atomic"
)

var counter int64

func nextID(prefix string) string {
	n:=atomic.AddInt64(&counter,1)
	return fmt.Sprintf("%s_%d",prefix,n)
}
