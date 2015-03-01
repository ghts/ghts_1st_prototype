package common

import (
	"math/big"
	"time"
)

/****************************
*               정수
****************************/
type C정수 struct{ 값 int }
func (c *C정수) G값() int { return c.값 }
func F정수_생성(값 int) *C정수 {
	c := C정수{값}
	return &c
}

type C정수64 struct{ 값 int64 }
func (c *C정수64) G값() int64 { return c.값 }
func F정수64_생성(값 int64) *C정수64 {
	c := C정수64{값}
	return &c
}

type C정수32 struct{ 값 int32 }
func (c *C정수32) G값() int32 { return c.값 }
func F정수32_생성(값 int32) *C정수32 {
	c := C정수32{값}
	return &c
}

type C정수16 struct{ 값 int16 }
func (c *C정수16) G값() int16 { return c.값 }
func F정수16_생성(값 int16) *C정수16 {
	c := C정수16{값}
	return &c
}

type C정수8 struct{ 값 int8 }
func (c *C정수8) G값() int8 { return c.값 }
func F정수8_생성(값 int8) *C정수8 {
	c := C정수8{값}
	return &c
}

/****************************
*            무부호 정수
****************************/
type C부호없는_정수 struct{ 값 uint }
func (c *C부호없는_정수) G값() uint { return c.값 }
func F부호없는_정수_생성(값 uint) *C부호없는_정수 {
	c := C부호없는_정수{값}
	return &c
}

type C부호없는_정수64 struct{ 값 uint64 }
func (c *C부호없는_정수64) G값() uint64 { return c.값 }
func F부호없는_정수64_생성(값 uint64) *C부호없는_정수64 {
	c := C부호없는_정수64{값}
	return &c
}

type C부호없는_정수32 struct{ 값 uint32 }
func (c *C부호없는_정수32) G값() uint32 { return c.값 }
func F부호없는_정수32_생성(값 uint32) *C부호없는_정수32 {
	c := C부호없는_정수32{값}
	return &c
}

type C부호없는_정수16 struct{ 값 uint16 }
func (c *C부호없는_정수16) G값() uint16 { return c.값 }
func F부호없는_정수16_생성(값 uint16) *C부호없는_정수16 {
	c := C부호없는_정수16{값}
	return &c
}

type C부호없는_정수8 struct{ 값 uint8 }
func (c *C부호없는_정수8) G값() uint8 { return c.값 }
func F부호없는_정수8_생성(값 uint8) *C부호없는_정수8 {
	c := C부호없는_정수8{값}
	return &c
}

/****************************
*            실수
****************************/
/* func F실수_생성(값 float64) *C실수64 {
	return F실수64_생성(값)
} */

type C실수64 struct{ 값 float64 }
func (c *C실수64) G값() float64 { return c.값 }
//func (c *C실수64) G금액() float64 { return F반올림_통화(c.값) }
func F실수64_생성(값 float64) *C실수64 {
	c := C실수64{값}
	return &c
}

type C실수32 struct{ 값 float32 }
func (c *C실수32) G값() float32 { return c.값 }
func F실수32_생성(값 float32) *C실수32 {
	c := C실수32{값}
	return &c
}

/****************************
*            문자열
****************************/
type C문자열 struct{ 값 string }
func (c *C문자열) G값() string { return c.값 }
func F문자열_생성(값 string) *C문자열 {
	c := C문자열{값}
	return &c
}

/****************************
*            참거짓
****************************/

var c참 *C참거짓 = &C참거짓{true}
var c거짓 *C참거짓 = &C참거짓{false}

type C참거짓 struct{ 값 bool }
func (c *C참거짓) G값() bool { return c.값 }
func F참거짓_생성(값 bool) *C참거짓 {
	if 값 {
		return c참
	} else {
		return c거짓
	}
}

/****************************
*            time.Time
****************************/

type C시점 struct{ 값 time.Time }
func (c *C시점) G값() time.Time { return c.값 }
func F시점_생성(값 time.Time) *C시점 {
	c := C시점{값}
	return &c
}


/************************
*         C정확한_실수
*************************/

type C정확한_실수 struct { 값 *big.Rat }
func (c *C정확한_실수) G값() float64 {
	실수값, _ := c.값.Float64()
	
	return 실수값
}
func (c *C정확한_실수) G정확한_값() *big.Rat {
	반환값 := new(big.Rat)
	반환값.Set(c.값)
	
	return 반환값
}
func (c *C정확한_실수) G반올림값(소숫점_이하_자릿수 int) float64 {
	문자열 := c.G문자열(소숫점_이하_자릿수)
	
	반올림값 := new(big.Rat)
	반올림값.SetString(문자열)
	실수값, _ := 반올림값.Float64()
	
	return 실수값
}
func (c *C정확한_실수) G문자열(소숫점_이하_자릿수 int) string {
	return c.값.FloatString(소숫점_이하_자릿수)
}
func F정확한_실수_생성(값 float64) *C정확한_실수 {
	c := new(C정확한_실수)
	c.값.SetFloat64(값)
	
	return c
}

/********************************************
*         C고정소숫점 (소수점이 있으니 당연히 실수임.)
********************************************/

type C고정소숫점 struct {
	소숫점_이하_자릿수 int
	값 *big.Rat
}
func (c *C고정소숫점) G소숫점_이하_자릿수() int { return c.소숫점_이하_자릿수 }
func (c *C고정소숫점) G값() *big.Rat {
	// 원본 변경을 방지하기 위해서 복사한 후 넘겨준다.
	복사본 := new(big.Rat)
	복사본.Set(c.값)
	
	return 복사본
}
func (c *C고정소숫점) G실수값() float64 {
	실수값, _ := c.값.Float64()
	
	return 실수값
}
func (c *C고정소숫점) G문자열() string { return c.값.FloatString(c.소숫점_이하_자릿수) }
func (c *C고정소숫점) String() string { return c.G문자열() }
func F고정소숫점_생성(값 float64, 소숫점_이하_자릿수 int) *C고정소숫점 {	
	// 반올림한 결과를 문자열 형태로 얻어서 값을 설정함.
	임시값 := new(big.Rat)
	임시값.SetFloat64(값)
	값_문자열 := 임시값.FloatString(소숫점_이하_자릿수)
	
	고정소숫점 := new(C고정소숫점)
	고정소숫점.소숫점_이하_자릿수 = 소숫점_이하_자릿수
	
	고정소숫점.값 = new(big.Rat)
	고정소숫점.값.SetString(값_문자열)
	
	return 고정소숫점
}