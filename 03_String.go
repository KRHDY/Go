package main

import "fmt"

func main()  {

	// *********************************
	//********** 문 자 열 *************
	// *********************************

	fmt.Println(len("ハローワルド"));
		// 문자열 길이를 반환 ( 공백도 포함 )
	fmt.Println("ハローワルド"[1]);
		// 문자열중 0 번째가 아닌 1번째의 바이트(ロ)로 표현(정수)
	fmt.Println("ハロー "+"ワルド");
		// 문자열과 문자열의 연결 (문자열을 더하는것은 무의미)

}
