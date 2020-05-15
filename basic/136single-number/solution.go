package _36single_number

func singleNumber(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
		// 异或 任何数字与0异或为他自身
		//异或  任何数字与他本身异或为0
		//所以有  a ^ b ^ c ^ a^ b = c
	}
	return single
}
