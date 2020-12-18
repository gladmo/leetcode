package store

import (
	"github.com/boltdb/bolt"
)

// 存储用户数据
// var private *bolt.DB
// var leetcode *bolt.DB
//
// func init() {
// 	var err error
// 	// private, err = bolt.Open(".private.db", 0600, nil)
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
//
// 	// leetcode, err = bolt.Open(".leetcode.db", 0600, nil)
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// }

// Customer db
func privateDB() (private *bolt.DB, err error) {
	return bolt.Open(".private.db", 0600, nil)
}

// leetcode db
func leetcodeDB() (private *bolt.DB, err error) {
	return bolt.Open(".leetcode.db", 0600, nil)
}
