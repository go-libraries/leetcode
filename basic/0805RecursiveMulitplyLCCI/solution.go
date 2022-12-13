package _805RecursiveMulitplyLCCI

func multiply(A int, B int) int {
	// if(B==1){
	//     return A
	// }else{
	//     if(B&1==1){
	//         return A+multiply(A<<1,B>>1)
	//     }else{
	//         // return 0
	//         return multiply(A<<1,B>>1)
	//     }
	// }
	// B*A = B/2 * A*2
	var C = 0
	for B != 1 {
		//C = A
		if B&1 == 1 {
			C += A
			A = A << 1
		} else {
			A = A << 1
		}
		B = B >> 1
	}
	return A + C
}
