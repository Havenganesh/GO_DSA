// package main

// import (
//   "fmt"
//   "strconv"
//   "strings"
//   "sort"
// )

// func main() {
//   ips := []string{
//     "192.168.1.5",
//     "69.52.220.44",
//     "10.152.16.23",
//     "192.168.3.10",
//     "192.168.1.4",
//     "192.168.1.41",
//     }

// ipsWithInt := make(map[string]int64)

// for _, ip := range ips {
//     ipStr := strings.Split(ip, ".")

//     oct0, _ := strconv.ParseInt(ipStr[0], 10, 64)
//     ipInt0 := oct0 * 255 * 255 * 255

//     oct1, _ := strconv.ParseInt(ipStr[1], 10, 64)
//     ipInt1 := oct1 * 255 * 255

//     oct2, _ := strconv.ParseInt(ipStr[2], 10, 64)
//     ipInt2 := oct2 * 255

//     oct3, _ := strconv.ParseInt(ipStr[3], 10, 64)
//     ipInt3 := oct3

//     ipInt := ipInt0 + ipInt1 + ipInt2 + ipInt3

//     ipsWithInt[ip] = ipInt
//   }

//   type kv struct {
//     Key   string
//     Value int64
//   }

//   var ss []kv
//   for k, v := range ipsWithInt {
//     ss = append(ss, kv{k, v})
//   }

//   sort.Slice(ss, func(i, j int) bool {
//     return ss[i].Value < ss[j].Value
//   })

//	  for _, kv := range ss {
//	    fmt.Printf("%s\n", kv.Key)
//	  }
//	}
package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	ips := []string{
		"192.168.1.5",
		"69.52.220.44",
		"10.152.16.23",
		"192.168.3.10",
		"192.168.1.4",
		"192.168.1.41",
	}

	ipsWithInt := make(map[string]int64)

	for _, ip := range ips {
		ipStr := strings.Split(ip, ".")

		oct0, _ := strconv.ParseInt(ipStr[0], 10, 64)
		ipInt0 := oct0 * 255 * 255 * 255

		oct1, _ := strconv.ParseInt(ipStr[1], 10, 64)
		ipInt1 := oct1 * 255 * 255

		oct2, _ := strconv.ParseInt(ipStr[2], 10, 64)
		ipInt2 := oct2 * 255

		oct3, _ := strconv.ParseInt(ipStr[3], 10, 64)
		ipInt3 := oct3

		ipInt := ipInt0 + ipInt1 + ipInt2 + ipInt3

		ipsWithInt[ip] = ipInt
	}

	type kv struct {
		Key   string
		Value int64
	}

	var ss []kv
	for k, v := range ipsWithInt {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value < ss[j].Value
	})

	for _, kv := range ss {
		fmt.Printf("%s\n", kv.Key)
	}
}
