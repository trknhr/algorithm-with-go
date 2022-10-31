package main

import "strconv"

func Add(num1, num2 string) string {
	ans := ""
	carry := 0
	lenNum1, lenNum2 := len(num1), len(num2)
	if lenNum1 > lenNum2 {
		for i := 0; i < lenNum1-lenNum2; i++ {
			num2 = "0" + num2
		}
	}
	if lenNum2 > lenNum1 {
		for i := 0; i < lenNum2-lenNum1; i++ {
			num1 = "0" + num1
		}
	}
	for i := len(num1) - 1; i >= 0; i-- {
		num1Num, _ := strconv.Atoi(string(num1[i]))
		num2Num, _ := strconv.Atoi(string(num2[i]))
		sum := num1Num + num2Num + carry
		ans = strconv.Itoa(sum%10) + ans
		carry = int(sum / 10)
	}

	if carry > 0 {
		ans = strconv.Itoa(carry) + ans
	}

	// remove head zero
	for string(ans[0]) == "0" {
		ans = ans[1:]
	}
	return ans
}

func AddWithShift(num1, num2 string, shiftForNum2 int) string {
	for i := 0; i < shiftForNum2; i++ {
		num2 += "0"
	}

	return Add(num1, num2)
}
