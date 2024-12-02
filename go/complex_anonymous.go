package gostrusct

import (
	"fmt"
	"strings"
	"sync"
)

type Pool2 struct {
	New func() any

	New2 func(int, string) any

	New3 func() (any, error)

	New4 func(*URI, map[string]interface{}) string

	New5 func(...string) any
}

// Without parameters: func() interface{}
var uriPool1 = &Pool2{
	New: func() interface{} {
		return &URI{}
	},
}

// With parameters: func(param1 int, param2 string) interface{}
var uriPool2 = &Pool2{
	New2: func(param1 int, param2 string) interface{} {
		return fmt.Sprintf("param1: %d, param2: %s", param1, param2)
	},
}

// With return values: func() (interface{}, error)
var uriPool3 = &Pool2{
	New3: func() (interface{}, error) {
		return &URI{}, nil
	},
}

// With method invocation (but no receiver): func() interface{} { someInstance.SomeMethod() }
var uriPool4 = &sync.Pool{
	New: func() interface{} {
		u := &URI{}
		u.Reset() // Using method with receiver inside anonymous function
		return u
	},
}

// With contextual parameters (closure): func() interface{} { return somethingFromContext }
var prefix = "/api/v1"
var uriPool5 = &sync.Pool{
	New: func() interface{} {
		return &URI{path: []byte(prefix + "/resource")}
	},
}

// With complex parameters: func(someStruct *StructType, someMap map[string]interface{}) string
var uriPool6 = &Pool2{
	New4: func(uri *URI, extraData map[string]interface{}) string {
		return fmt.Sprintf("URI: %s, Extra: %v", uri.Path, extraData)
	},
}

// With variadic parameters: func(args ...string) interface{}
var uriPool7 = &Pool2{
	New5: func(paths ...string) interface{} {
		return &URI{path: []byte(strings.Join(paths, "/"))}
	},
}

// With closures returned: func() func() int { return func() int { return someClosure }}
var createCounter = func() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

var urii = &URI{path: []byte("s")}

var createCounter2 = func() func(uri *URI) func() int {
	count := 0
	return func(uri *URI) func() int {
		return func() int {
			count++
			return count
		}
	}
}

// just a func
func AcquireURI2() *URI {
	return uriPool.Get().(*URI)
}

func (u *URI) SetHash(hash string) {
	u.hash = append(u.hash[:0], hash...)
}

func (u *URI) SetHashRetMultWithReceiver(hash string) (interface{}, error) {
	u.hash = append(u.hash[:0], hash...)
	return nil, nil
}

func (u *URI) SetHashRetWithReceiver(hash string) interface{} {
	u.hash = append(u.hash[:0], hash...)
	return nil
}

func (s Stru) sayGoodbye() {}

func SetHashRet(hash string) (interface{}, error) {
	hash = append(hash[:0], hash...)
	return hash, nil
}

type MyComparator struct{}

type Comparator interface {
	Handle(func(int) string)
}

func (m MyComparator) Handle(f func(int) string) {}

type Geometry struct {
	isEqual func(float64, float64) error
}

func (g *Geometry) sayGoodbye() {}

// func (__receiver__) __name__(__args__)
// func (__receiver__) __name__()
// func () __name__()
// func __name__() __return_type__
// func __name__() (__return_type__, __return_type__)
// func (__receiver__) __name__(__args__) (__return_type__, __return_type__)
// func (__receiver__) __name__(__args__) __return_type__
//
// anon:
// func() __return_type__
// func() (__return_type__, __return_type__)
// func(__receiver__) __return_type__
// func() func() __return_type__
