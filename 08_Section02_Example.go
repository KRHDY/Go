package main

import "fmt"

func main()  {

	// ***********************************
	// ********** 연 습 문 제 *************
	// ***********************************

	var a = 10 ;
		// 첫번째 방법

	ab := 10;
		// 두번째 방법

	fmt.Println(a+ab);
		// 연습문제 1 : 새로운 변수를 생성하는 두 가지 방법은 무엇인가?


		// 연습문제 2 : 다음 코드를 실행한 후 x의 값은 무엇인가? x := 5; x += 1
		// X의 값은 6일것이다

	var C,F float64;
		// 변수 C,F 를 선언 합니다

	fmt.Scanf("%f",&F);
		// 변수 F값을 받아옵니다

	C = (F - 32) * 5/9;
		// 변수 C값을 설정합니다

	fmt.Println(C);
		// 출력합니다
		// 연습문제 5 :예제 프로그램을 출발점으로 삼아 화씨를 섭씨로 변환하는 프로그램을 작성하라. (C = (F - 32) * 5/9))

	var  M	float64;
		// 변수 M을 선언합니다

	fmt.Scanf("%f",&F);
		// 변수 F에 값을 받습니다

	M = 0.3048 * F;
		// 변수 M값을 설정합니다

	fmt.Println(M);
		// 출력합니다
		// 연습문제 6 : 피트를 미터로 변환하는 프로그램을 하나 더 작성하라. (1 ft = 0.3048 m)
}
