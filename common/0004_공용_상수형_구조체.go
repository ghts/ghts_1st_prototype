package common

import (
	"time"
)

/*	상수형이 필요한 이유.
	동시처리(concurrent) 모듈에서 공유하는 모든 자료는 경쟁조건(racing condition)이 발생할
	 가능성을 가지고 있다.
	이러한 자료공유 문제를 원천적으로 해결하는 방법은 공유되는 자료를 복사해서 전달하는 방법과
	 자료가 아예 변경되지 않도록 보장하는 방법 2가지가 있다.
	 Go 언어는 공유자료를 복사에 의해서 전달하는 채널(channel)을 자체적으로 제공하지만,
	 변경할 수 없는 자료형식(immutable)을 제공하지는 않으므로 이 부분은 자체적으로 구현해야 한다.

	모든 상수형 구조체는 또 다른 상수형 구조체의 포인터로 이루어져야 한다.
	또한, slice, map은 그 자체로 reference타입이므로 주고 받을 때는 반드시 복사를 거친 후, 복사본을 주고 받도록 하여서 상수형 구조체의 변경불가능성(immutable)을 유지하도록 한다.
*/

type C증권사 struct {
	식별코드    *C부호없는_정수64
	표시명칭    *C문자열
	사업자등록번호 *C문자열
}
func (c *C증권사) G상수형() bool { return true }
func (c *C증권사) G식별코드() uint64    { return c.식별코드.G값() }
func (c *C증권사) G표시명칭() string    { return c.표시명칭.G값() }
func (c *C증권사) G사업자등록번호() string { return c.사업자등록번호.G값() }
func F증권사_생성(식별코드 uint64, 표시명칭 string, 사업자등록번호 string) *C증권사 {
	c := new(C증권사)
	c.식별코드 = F부호없는_정수64_생성(식별코드)
	c.표시명칭 = F문자열_생성(표시명칭)
	c.사업자등록번호 = F문자열_생성(사업자등록번호)

	return c
}

type C계좌 struct {
	식별코드 *C부호없는_정수64
	표시명칭 *C문자열
	증권사  *C증권사 // 상수형의 포인터를 포함해도 상수형임.
	계좌번호 *C문자열
}

func (c *C계좌) G식별코드() uint64 { return c.식별코드.G값() }
func (c *C계좌) G표시명칭() string { return c.표시명칭.G값() }
func (c *C계좌) G증권사() *C증권사   { return c.증권사 }
func (c *C계좌) G계좌번호() string { return c.계좌번호.G값() }
func F계좌_생성(
	식별코드 uint64,
	표시명칭 string,
	증권사 *C증권사,
	계좌번호 string) *C계좌 {
	c := new(C계좌)
	c.식별코드 = F부호없는_정수64_생성(식별코드)
	c.표시명칭 = F문자열_생성(표시명칭)
	c.증권사 = 증권사
	c.계좌번호 = F문자열_생성(계좌번호)

	return c
}

type C가격정보 struct {
	시점  *C시점
	종목  *C종목
	가격  *C실수64
	거래량 *C부호없는_정수64 // 이게 항상 얻을 수 있는 값인가?
}

func (c *C가격정보) G시점() time.Time { return c.시점.G값() }
func (c *C가격정보) G종목() *C종목      { return c.종목 }
func (c *C가격정보) G가격() float64   { return c.가격.G값() }
func (c *C가격정보) G거래량() uint64   { return c.거래량.G값() }
func F가격정보_생성(시점 time.Time, 종목 *C종목, 가격 float64, 거래량 uint64) *C가격정보 {
	c := new(C가격정보)
	c.시점 = F시점_생성(시점)
	c.종목 = 종목
	c.가격 = F실수64_생성(가격)
	c.거래량 = F부호없는_정수64_생성(거래량)

	return c
}