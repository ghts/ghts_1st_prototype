package common

const (
	종료이벤트   int8 = 0
	비상탈출이벤트 int8 = 1
)

const (
	초도매수   uint8 = 1
	추가매수   uint8 = 2
	전량매도   uint8 = 3
	분할초도매도 uint8 = 4
	분할추가매도 uint8 = 5
)

const (
	현재가매수  uint8 = 1
	가격지정매수 uint8 = 2
	현재가매도  uint8 = 21
	가격지정매도 uint8 = 22
)

const (
	체결완료 uint8 = 1
	체결실패 uint8 = 2
	일부체결 uint8 = 3
)
