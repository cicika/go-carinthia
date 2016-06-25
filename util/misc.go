package util

import "crypto/md5"

func MD5(input string)(string) {
  var bytes []byte
  copy(bytes[:], input)
  hash := md5.Sum(bytes)
  return string(hash[:])
}
