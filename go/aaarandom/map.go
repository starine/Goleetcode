package main

import (
	"fmt"
	"sort"
	"unsafe"
)

type hmap struct {
	count      int            //map中键值对数量
	flags      uint8          //map当前是否处于写入状态等
	B          uint8          //2的B次幂表示当前map中桶的数量
	noverflow  uint16         //map中溢出桶的数量，当溢出桶太多时，map会进行等量扩容
	hash0      uint32         //生成hash的随机数种子
	buckets    unsafe.Pointer //当前map对应的桶的指针
	oldbuckets unsafe.Pointer //map扩容时指向旧桶的指针，当所有旧桶中的数据转移到新桶时，清空
	nevacuate  uintptr        //扩容时，用于标记当前旧桶中小于nevacute的数据都已经转移到了新桶
	//extra *mapextra	//存储map的溢出桶
}
type bmap struct {
	tophash  [8]uint8 //存储Hash值的高8位
	data     []byte   //key value数据：key/key/key.../value/value/value...
	overflow *bmap    //溢出bucket的地址
}

func main() {
	var m1 map[string]int
	m2 := make(map[string]int)
	m3 := make(map[string]int, 10)

	fmt.Println(m1 == nil) //true
	fmt.Println(m2 == nil) //false
	fmt.Println(m3 == nil) //false

	fmt.Println(m1["a"] == 0) //true
	fmt.Println(m2["a"] == 0) //true
	fmt.Println(m3["a"] == 0) //true

	fmt.Println(len(m1)) //0
	fmt.Println(len(m2)) //0
	fmt.Println(len(m3)) //0

	delete(m1, "a") //不会报错

	mapSlice := make([]map[string]int, 8, 8) //只初始化了切片
	fmt.Println(mapSlice[0] == nil)          //true
	fmt.Printf("mapSlice:%v\n", mapSlice)
	//mapSlice:[map[] map[] map[] map[] map[] map[] map[] map[]]

	for i := range mapSlice {
		mapSlice[i] = make(map[string]int, 8) //初始化map
	}

	fmt.Println(mapSlice[0] == nil) //false
	mapSlice[0]["apple"] = 100
	fmt.Printf("mapSlice:%v\n", mapSlice)
	//mapSlice:[map[apple:100] map[] map[] map[] map[] map[] map[] map[]]

	capitals := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo"}
	for key := range capitals {
		fmt.Println("Map item: Capital of", key, "is", capitals[key])
	}

	barVal := map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12,
		"golf": 34, "hotel": 16, "indio": 87,
		"juliet": 65, "kili": 43, "lima": 98}

	fmt.Println("unsorted:")
	for k, v := range barVal {
		fmt.Printf("Key: %v, Value: %v / ", k, v)
	}
	keys := make([]string, len(barVal))
	i := 0
	for k, _ := range barVal {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println()
	fmt.Println("sorted:")
	for _, k := range keys {
		fmt.Printf("Key: %v, Value: %v / ", k, barVal[k])
	}

	invMap := make(map[int]string, len(barVal))
	for k, v := range barVal {
		invMap[v] = k
	}
	fmt.Println("inverted:")
	for k, v := range invMap {
		fmt.Printf("Key: %v, Value: %v / ", k, v)
	}
}
