package urand

import (
	"fmt"
	"math/rand"
)

// Chance 概率计算
//changes [10,40,50]
//random=20 | index = 1
func Chance(changes []int32) (index int) {
	var total int32
	for _, val := range changes {
		total += val
	}
	random := rand.Int31n(total) + 1
	for _, val := range changes {
		if val >= random {
			return index
		}
		index++
		random -= val
	}
	return 0
}

// RandomByMinMax 通过最大最小值随机中间数,包含max
func RandomByMinMax(min, max int32) int32 {
	// +1 解决 min,max值相同并能随机到max
	diff := max - min + 1
	if diff <= 0 {
		return 0
	}
	res := rand.Int31n(diff) + min
	return res
}

// GetRandomString 生成随机字符串
func GetRandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	var result []byte
	for i := 0; i < length; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}

// GetRandomString 生成随机数字字符串
func GetRandomNumString(length int) string {
	str := "0123456789"
	bytes := []byte(str)
	var result []byte
	for i := 0; i < length; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}

//GetRandomBase32String 生成base32随机 密钥
func GetRandomBase32String(length int) string {
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"
	bytes := []byte(str)
	var result []byte
	for i := 0; i < length; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}

// GetRandomCode 随机六位验证码
func GetRandomCode() string {
	vcode := fmt.Sprintf("%06v", rand.Int31n(1000000))
	return vcode
}

// 字符串内容随机打乱
func GetRandomStr(str string) string {
	perm := rand.Perm(len(str))
	var newStr []byte
	newStr = []byte{}
	strs := str[:]
	for _, i2 := range perm {
		newStr = append(newStr, strs[i2])
	}
	return string(newStr)
}
