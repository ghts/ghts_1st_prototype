package common

import (
	"sync"
	"time"
)

// 서버에서 신호발생 및 주문을 전송한 전략으로 결과를 회신하는 데이터
type C주문처리결과 struct {
	식별코드   *C부호없는_정수64
	주문     *C주문
	처리결과종류 *C부호없는_정수8
	체결단가   *C실수64
	체결수량   *C부호없는_정수64
	체결시점   *C시점
}

func (c *C주문처리결과) G식별코드() uint64 { return c.식별코드.G값() }
func (c *C주문처리결과) G주문() *C주문     { return c.주문 }
func (c *C주문처리결과) G처리결과종류() uint8 {
	return c.처리결과종류.G값()
}
func (c *C주문처리결과) G체결단가() float64   { return c.체결단가.G값() }
func (c *C주문처리결과) G체결수량() uint64    { return c.체결수량.G값() }
func (c *C주문처리결과) G체결시점() time.Time { return c.체결시점.G값() }
func F주문처리결과생성(
	식별코드 uint64,
	주문 *C주문,
	처리결과종류 uint8,
	체결수량 uint64,
	체결시점 time.Time) *C주문처리결과 {
	c := new(C주문처리결과)
	c.식별코드 = F부호없는_정수64_생성(식별코드)
	c.주문 = 주문
	c.처리결과종류 = F부호없는_정수8_생성(처리결과종류)
	c.체결수량 = F부호없는_정수64_생성(체결수량)
	c.체결시점 = F시점_생성(체결시점)

	return c
}

type C서버설정 struct {
	전략관리  I서버하위모듈
	가격정보제공  I서버하위모듈
	주문처리    I서버하위모듈
	증권사모음   []I증권사
	증권사_기본값 I증권사

	전략등록채널     chan I전략
	전략정보구독신청채널 chan (chan []I전략)
	가격정보채널       chan *C가격정보
	주문접수채널       chan *C주문
	공통이벤트채널      chan int8

	가격정보제공모듈설정 *C가격정보제공모듈설정
	주문처리모듈설정   *C주문처리모듈설정

	종료대기열 *sync.WaitGroup
}

func (c *C서버설정) G전략관리() I서버하위모듈 { return c.전략관리 }
func (c *C서버설정) G가격정보제공() I서버하위모듈 { return c.가격정보제공 }
func (c *C서버설정) G주문처리() I서버하위모듈   { return c.주문처리 }
func (c *C서버설정) G증권사모음() []I증권사   { return c.증권사모음 }
func (c *C서버설정) G증권사_기본값() I증권사   { return c.증권사_기본값 }
func (c *C서버설정) G전략등록채널() chan I전략 {
	return c.전략등록채널
}
func (c *C서버설정) G전략정보구독신청채널() chan (chan []I전략) {
	return c.전략정보구독신청채널
}
func (c *C서버설정) G가격정보채널() chan *C가격정보 { return c.가격정보채널 }
func (c *C서버설정) G주문접수채널() chan *C주문   { return c.주문접수채널 }
func (c *C서버설정) G공통이벤트채널() chan int8  { return c.공통이벤트채널 }
func (c *C서버설정) G가격정보제공모듈설정() *C가격정보제공모듈설정 {
	return c.가격정보제공모듈설정
}
func (c *C서버설정) G주문처리모듈설정() *C주문처리모듈설정 {
	return c.주문처리모듈설정
}
func (c *C서버설정) G종료대기열() *sync.WaitGroup { return c.종료대기열 }
func F서버설정생성(
	전략관리 I서버하위모듈,
	가격정보제공 I서버하위모듈,
	주문처리 I서버하위모듈,
	증권사모음 []I증권사,
	증권사_기본값 I증권사,
	전략등록채널 chan I전략,
	전략정보구독신청채널 chan (chan []I전략),
	가격정보채널 chan *C가격정보,
	주문접수채널 chan *C주문,
	공통이벤트채널 chan int8,
	가격정보제공모듈설정 *C가격정보제공모듈설정,
	주문처리모듈설정 *C주문처리모듈설정,
	종료대기열 *sync.WaitGroup) *C서버설정 {
	c := new(C서버설정)
	c.전략관리 = 전략관리
	c.가격정보제공 = 가격정보제공
	c.주문처리 = 주문처리
	c.증권사모음 = 증권사모음
	c.증권사_기본값 = 증권사_기본값
	c.전략등록채널 = 전략등록채널
	c.전략정보구독신청채널 = 전략정보구독신청채널
	c.가격정보채널 = 가격정보채널
	c.주문접수채널 = 주문접수채널
	c.공통이벤트채널 = 공통이벤트채널
	c.가격정보제공모듈설정 = 가격정보제공모듈설정
	c.주문처리모듈설정 = 주문처리모듈설정
	c.종료대기열 = 종료대기열

	return c
}

type C가격정보제공모듈설정 struct {
	// 모의서버에서는 동시처리숫자 0개
	가격정보전송_최대동시처리숫자 *C부호없는_정수64
}

func (c *C가격정보제공모듈설정) G가격정보전송_최대동시처리숫자() uint64 {
	return c.가격정보전송_최대동시처리숫자.G값()
}
func F가격정보제공모듈설정생성(가격정보전송_최대동시처리숫자 uint64) *C가격정보제공모듈설정 {
	c := new(C가격정보제공모듈설정)
	c.가격정보전송_최대동시처리숫자 = F부호없는_정수64_생성(가격정보전송_최대동시처리숫자)

	return c
}

type C주문처리모듈설정 struct {
	주문처리_최대동시처리숫자 *C부호없는_정수64
}

func (c *C주문처리모듈설정) G주문처리_최대동시처리숫자() uint64 {
	return c.주문처리_최대동시처리숫자.G값()
}
func F주문처리모듈설정생성(주문처리_최대동시처리숫자 uint64) *C주문처리모듈설정 {
	c := new(C주문처리모듈설정)
	c.주문처리_최대동시처리숫자 = F부호없는_정수64_생성(주문처리_최대동시처리숫자)

	return c
}

type C데이터베이스_연결정보 struct {
	드라이버   *C문자열
	사용자이름  *C문자열
	암호     *C문자열
	DB이름   *C문자열
	추가파라메터 *C문자열
}

func (c *C데이터베이스_연결정보) G드라이버() string { return c.드라이버.G값() }
func (c *C데이터베이스_연결정보) G연결문자열() string {
	return c.사용자이름.G값() + ":" +
		c.암호.G값() + "@/" +
		c.DB이름.G값() + "?" +
		c.추가파라메터.G값()
}

var mySQL연결정보 *C데이터베이스_연결정보 = nil

func F_MySQL연결정보() *C데이터베이스_연결정보 {
	if mySQL연결정보 == nil {
		mySQL연결정보 = new(C데이터베이스_연결정보)
		mySQL연결정보.드라이버 = F문자열_생성("mysql")
		mySQL연결정보.사용자이름 = F문자열_생성("pts")
		mySQL연결정보.암호 = F문자열_생성("FqQrDsEm9f2Vw6pR")
		mySQL연결정보.DB이름 = F문자열_생성("pts")
		mySQL연결정보.추가파라메터 =
			F문자열_생성("charset=utf8mp4,utf8&parseTime=true&strict=true&autocommit=true")
	}

	return mySQL연결정보
}
