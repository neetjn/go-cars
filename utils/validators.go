package utils

import (
  regexp "regexp"
)

const VinPattern string = "([a-zA-Z0-9]{5,10})"

func ValidateVin(vin string) bool {
  re := regexp.MustCompile(VinPattern)
  return re.MatchString(vin)
}
