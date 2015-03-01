package common

import (
	"math/big"
)

type P통화 string

const (
	KRW P통화 = "KRW"
	USD    = "USD"
	CNY    = "CNY"
	EUR    = "EUR"
)

func F원화_생성(금액 float64) I통화 { return F통화_생성(KRW, 금액) }
func F달러_생성(금액 float64) I통화 { return F통화_생성(USD, 금액) }
func F위안화_생성(금액 float64) I통화 { return F통화_생성(CNY, 금액) }
func F유로화_생성(금액 float64) I통화 { return F통화_생성(EUR, 금액) }
func F통화_생성(종류 P통화, 금액 float64) I통화 {
	통화 := new(c통화)
	통화.종류 = 종류
	통화.고정소숫점 = F고정소숫점_생성(금액, F통화별_정밀도(종류))
	
	return 통화
}

func F원화_구조체_생성(금액 float64) I통화_구조체 { return F통화_구조체_생성(KRW, 금액) }
func F달러_구조체_생성(금액 float64) I통화_구조체 { return F통화_구조체_생성(USD, 금액) }
func F위안화_구조체_생성(금액 float64) I통화_구조체 { return F통화_구조체_생성(CNY, 금액) }
func F유로화_구조체_생성(금액 float64) I통화_구조체 { return F통화_구조체_생성(EUR, 금액) }
func F통화_구조체_생성(종류 P통화, 금액 float64) I통화_구조체 {
	통화 := new(s통화)
	통화.종류 = 종류
	통화.고정소숫점 = F고정소숫점_구조체_생성(금액, F통화별_정밀도(종류))
	
	return 통화
}

func F통화별_정밀도(통화 P통화) int {
	switch 통화 {
	case KRW:
		return 0
	case USD:
		return 2
	case CNY:
		return 2
	case EUR:
		return 2
	default:
		return 2
	}
}

// 통화 상수형 구조체는 각 통화에 해당되는 펑션을 사용해서만 생성이 가능하도록 한다.
// 이를 위해서, 통화 상수형 구조체를 new()로 생성하지 못하게 한다. (외부에 공개하지 않는다.)
// 구조체를 외부에 숨기는 대신 'I통화'를 공개하고 이것을 사용한다.
type c통화 struct {
	종류 P통화
	고정소숫점 *C고정소숫점
}
func (c *c통화) G식별코드() P통화 { return c.종류 }
func (c *c통화) G소숫점_이하_자릿수() int { return c.고정소숫점.G소숫점_이하_자릿수() }
func (c *c통화) G값() float64 { return c.고정소숫점.G실수값() }
func (c *c통화) G정확한_값() *big.Rat { return c.고정소숫점.G값() }

// 통화 구조체는 각 통화에 해당되는 펑션을 사용해서만 생성이 가능하도록 한다.
// 이를 위해서, 통화 구조체를 new()로 생성하지 못하게 한다. (외부에 공개하지 않는다.)
// 구조체를 외부에 숨기는 대신 'I통화_구조체'를 공개하고 이것을 사용한다.
type s통화 struct {
	종류 P통화
	고정소숫점 *S고정소숫점
}
func (s *s통화) G식별코드() P통화 { return s.종류 }
func (s *s통화) G소숫점_이하_자릿수() int { return s.고정소숫점.G소숫점_이하_자릿수() }
func (s *s통화) G값() float64 { return s.고정소숫점.G실수값() }
func (s *s통화) G정확한_값() *big.Rat { return s.고정소숫점.G값() }
func (s *s통화) G상수형() I통화 { return F통화_생성(s.종류, s.고정소숫점.G실수값()) }	
func (s *s통화) S값(값 float64) { s.고정소숫점.S실수값(값) }
func (s *s통화) S정확한_값(정확한_값 *big.Rat) { s.고정소숫점.S값(정확한_값) }