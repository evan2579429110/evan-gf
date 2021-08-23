package service

import (
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/frame/g"
	"reflect"
	"testing"
)

func TestRedisGet(t *testing.T) {
	c := g.Config()
	_ = c.AddPath("../../config")

	res, err := g.Redis().DoVar("GET", "cache-demo")

	fmt.Println(err, res)
}

func TestRedisSet(t *testing.T) {
	c := g.Config()
	_ = c.AddPath("../../config")

	res, err := g.Redis().DoVar("SET", "cache-cs", "dsd")
	fmt.Println(err, res)
}

func TestRedisKeyPattern(t *testing.T) {
	c := g.Config()
	_ = c.AddPath("../../config")

	res, err := g.Redis().DoVar("KEYS", "cache*")
	fmt.Println(json.Marshal(res))
	//var d map[string]interface{}
	//_ = res.Map(&d)
	//if reflect.TypeOf(res).Kind() == reflect.Slice {
	//	for _,val := range res.(g.Slice){
	//		fmt.Println(val,reflect.TypeOf(val).Kind())
	//	}
	//
	//}
	//for _,val := range res {
	//	val.
	//}
	fmt.Println(err, reflect.TypeOf(res).Kind())
}
