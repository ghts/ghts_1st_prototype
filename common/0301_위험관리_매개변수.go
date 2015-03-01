package common

import (
)

func F위험관리_매개변수_생성(
	전략 I전략, 
	종목 *C종목, 
	수량 int64, 
	단가 I통화, 
	매개변수_모음 *S매개변수_모음) *S위험관리_매개변수 {
	s := new(S위험관리_매개변수)
	s.전략 = 전략
	s.종목 = 종목
	s.수량 = 수량
	s.단가 = 단가
	s.금액 = F통화_생성(단가.G식별코드(), float64(수량) * 단가.G값())
	
	// 원본 훼손 가능성을 최소화 하기 위해서 복제해서 넘겨받음.
	s.매개변수_모음 = 매개변수_모음.G복제()
	
	return s
}

type S위험관리_매개변수 struct {
	전략 I전략
	종목 *C종목
	수량 int64
	단가 I통화
	금액 I통화
	매개변수_모음 *S매개변수_모음
}
func (s *S위험관리_매개변수) G전략() I전략 { return s.전략 }
func (s *S위험관리_매개변수) G종목() *C종목 { return s.종목 }
func (s *S위험관리_매개변수) G수량() int64 { return s.수량 }
func (s *S위험관리_매개변수) G단가() I통화 { return s.단가 }
func (s *S위험관리_매개변수) G금액() I통화 { return s.금액 }
func (s *S위험관리_매개변수) G값(이름 string) interface{} { return s.매개변수_모음.G값(이름) }
func (s *S위험관리_매개변수) G가변형(이름 string) *S가변형 { return s.매개변수_모음.G가변형(이름) }
func (s *S위험관리_매개변수) G매개변수_모음() []*S매개변수 { return s.매개변수_모음.G매개변수_모음() }