package utils

import (
  "os"
  "strconv"
)

func GetEnv(key string, fallback string) string {
  result, found := os.LookupEnv(key)
  if found {
    return result
  }
  return fallback
}

func ParseInt(value string) int64 {
  result, err := strconv.ParseInt(value, 10, 64)
  if err != nil {
    return result
  }
  return 0
}
