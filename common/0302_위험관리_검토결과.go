package common

func F위험관리_검토결과_생성(통과 bool, 
			수량조정_필요 bool, 수량 int64, 
			금액조정_필요 bool, 금액 I통화,
			매개변수_모음 *S매개변수_모음) *S위험관리_검토결과 {
	s := new(S위험관리_검토결과)
	s.통과 = F참거짓_생성(통과)
	s.수량조정_필요 = F참거짓_생성(수량조정_필요)
	s.수량 = F정수64_생성(수량)
	s.금액조정_필요 = F참거짓_생성(금액조정_필요)
	s.금액 = 금액
	
	// 원본 훼손 가능성을 최소화 하기 위해서 복제해서 넘겨받음.
	// 복제() 펑션이 deep copy가 아니고 shallow copy라서 그 효과가 의문시 되기는 함.
	s.매개변수_모음 = 매개변수_모음.G복제()
	
	return s
}

type S위험관리_검토결과 struct {
	통과 *C참거짓
	수량조정_필요 *C참거짓
	수량 *C정수64
	금액조정_필요 *C참거짓
	금액 I통화
	매개변수_모음 *S매개변수_모음
}
func (s *S위험관리_검토결과) G통과() bool { return s.통과.G값() }
func (s *S위험관리_검토결과) G수량조정_필요() bool { return s.수량조정_필요.G값() }
func (s *S위험관리_검토결과) G수량() int64 { return s.수량.G값() }
func (s *S위험관리_검토결과) G금액조정_필요() bool { return s.금액조정_필요.G값() }
func (s *S위험관리_검토결과) G금액() I통화 { return s.금액 }
func (s *S위험관리_검토결과) G값(이름 string) interface{} { return s.매개변수_모음.G값(이름) }
func (s *S위험관리_검토결과) G가변형(이름 string) *S가변형 { return s.매개변수_모음.G가변형(이름) }
func (s *S위험관리_검토결과) G매개변수_모음() []*S매개변수 { return s.매개변수_모음.G매개변수_모음() }