package common

func F종목별_포트폴리오_구조체_생성(식별코드 uint64, 통화_식별코드 P통화, 종목 *C종목) *S종목별_포트폴리오 {
	s := new(S종목별_포트폴리오)
	s.식별코드 = 식별코드
	s.통화_식별코드 = 통화_식별코드
	s.종목 = 종목
	s.단가 = F통화_구조체_생성(통화_식별코드, 0.0)
	s.롱포지션_수량 = 0
	s.숏포지션_수량 = 0
	
	return s
}

type S종목별_포트폴리오 struct {
	식별코드 uint64
	통화_식별코드 P통화
	종목 *C종목	
	단가 I통화_구조체	
	롱포지션_수량 uint64
	숏포지션_수량 uint64
}
func (s *S종목별_포트폴리오) G식별코드() uint64 { return s.식별코드 }
func (s *S종목별_포트폴리오) G통화_식별코드() P통화 { return s.통화_식별코드 }
func (s *S종목별_포트폴리오) G종목() *C종목 { return s.종목 }
func (s *S종목별_포트폴리오) G단가() I통화 { return s.단가.G상수형() }
func (s *S종목별_포트폴리오) G롱포지션_수량() uint64 { return s.롱포지션_수량 }
func (s *S종목별_포트폴리오) G숏포지션_수량() uint64 { return s.숏포지션_수량 }
func (s *S종목별_포트폴리오) G순_수량() uint64 { return s.롱포지션_수량 - s.숏포지션_수량 }
func (s *S종목별_포트폴리오) G총_수량() uint64 { return s.롱포지션_수량 + s.숏포지션_수량 }
func (s *S종목별_포트폴리오) G롱포지션_금액() I통화 { return F통화_생성(s.통화_식별코드, float64(s.롱포지션_수량) * s.G단가().G값()) }
func (s *S종목별_포트폴리오) G숏포지션_금액() I통화 { return F통화_생성(s.통화_식별코드, float64(s.숏포지션_수량) * s.G단가().G값()) }
func (s *S종목별_포트폴리오) G순_금액() I통화 { return F통화_생성(s.통화_식별코드, float64(s.G순_수량()) * s.G단가().G값()) }
func (s *S종목별_포트폴리오) G총_금액() I통화 { return F통화_생성(s.통화_식별코드, float64(s.G총_수량()) * s.G단가().G값()) }
func (s *S종목별_포트폴리오) S단가(단가 float64) { s.단가.S값(단가) }
func (s *S종목별_포트폴리오) S롱포지션_수량(수량 uint64) { s.롱포지션_수량 = 수량 }
func (s *S종목별_포트폴리오) S숏포지션_수량(수량 uint64) { s.숏포지션_수량 = 수량 }