package common

type I서버 interface {
	// 서버구동
	M실행(서버설정 C서버설정)

	// 증권사 모듈 리스트
	G증권사모음()

	// 지정 증권사가 없을 경우 주문이 전송될 기본 증권사
	G증권사_기본값()

	// 각 전략이 서버에 등록하는 채널
	G전략등록채널() chan *I전략

	// 서버의 하위모듈이 전략정보를 얻고자 구독신청하는 채널
	G전략정보구독신청채널() chan (chan []I전략)

	// 서버의 증권사 모듈들이 전송한 가격정보를 취합하는 채널.
	// 여기를 거친 후 가격정보는 각 전략으로 배포된다.
	G가격정보채널() chan *C가격정보 // 그냥 각 전략에서 바로 받는 게 간단할 것 같다.

	// 전략이 전송한 주문을 취합하는 채널.
	// 여기를 거친 주문정보는 이후 적절한 증권사 모듈로 전송된다.
	G주문접수채널() chan *C주문

	// 비상탈출 및 종료신호가 취합되는 곳.
	// 여기를 거친 공통 이벤트는 이후 모든 전략과 서버의 모든 하위모듈에 배포된다.
	G공통이벤트채널() chan int8
}

type I서버하위모듈 interface {
	// 모듈을 구동하는 메소드
	M실행(서버 I서버)

	// 비상탈출, 종료등의 이벤트를 전송하는 채널
	G공통이벤트채널() chan int8
}

type I증권사 interface {
	I서버하위모듈

	// 어느 증권사의 모듈인지 확인할 떄 사용.
	G증권사() *C증권사
	G주문접수채널() chan *C주문
	G주문전송_처리권한() chan int8
}
