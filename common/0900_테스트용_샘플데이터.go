package common

import (
	"time"
	//"log"
)

func F일일가격정보_테스트용_테이블_준비() {
	F테스트_모드_시작()

	// 혹시 남아있을 지 모르는 예전 테이블 정리
	F테이블_삭제(F일일가격정보_테이블())
	F테이블_삭제(F종목정보_테이블())

	// 테이블 생성
	F종목정보_테이블_생성()
	F일일가격정보_테이블_생성()

	// 데이터 기록
	F종목정보_맵_DB기록(F종목정보_맵_테스트용())
	F일일가격정보_모음_DB기록(F일일가격정보_모음_맵_테스트용().G종목별_내용(F종목_동화약품()))
	F일일가격정보_모음_DB기록(F일일가격정보_모음_맵_테스트용().G종목별_내용(F종목_삼성전자()))
}

func F일일가격정보_테스트용_테이블_정리() {
	F테이블_삭제(F일일가격정보_테이블())
	F테이블_삭제(F종목정보_테이블())
	F테스트_모드_종료()
}

func F증권사_가상_1() *C증권사 {
	return F증권사_생성(uint64(1), "아무개1 증권사", "12345678-90123456")
}

func F증권사_가상_2() *C증권사 {
	return F증권사_생성(uint64(2), "아무개2 증권사", "98765432-109876543")
}

func F계좌_가상_1() *C계좌 {
	return F계좌_생성(9991,
		"테스트용 가상계좌 1",
		F증권사_가상_1(),
		"111-1111-111111")
}

func F계좌_가상_2() *C계좌 {
	return F계좌_생성(9992,
		"테스트용 가상계좌 2",
		F증권사_가상_2(),
		"222-2222-222222")
}

func F종목_가상_1() *C종목 {
	return F종목_생성(999900,
		"A999900",
		"KR7999900001",
		"테스트용 가상종목 1",
		"Non-existent only for test 1",
		"999900",
		"유가증권시장상장")
}

func F종목_가상_2() *C종목 {
	return F종목_생성(999910,
		"A999910",
		"KR7999910001",
		"테스트용 가상종목 2",
		"Non-existent only for test 2",
		"999910",
		"유가증권시장상장")
}

func F종목_가상_3() *C종목 {
	return F종목_생성(999920,
		"A999920",
		"KR7999920001",
		"테스트용 가상종목 3",
		"Non-existent only for test 3",
		"999920",
		"유가증권시장상장")
}

func F종목_동화약품() *C종목 {
	return F종목_생성(1,
		"A000020",
		"KR7000020008",
		"동화약품보통주",
		"DongwhaPharm",
		"2",
		"유가증권시장상장")
}

func F종목_삼성전자() *C종목 {
	return F종목_생성(2,
		"A005930",
		"KR7005930003",
		"삼성전자보통주",
		"SamsungElectronics",
		"593",
		"유가증권시장상장")
}

func F종목정보_맵_테스트용() map[string]*C종목 {
	종목정보_맵 := make(map[string]*C종목)

	var 종목 *C종목

	종목 = F종목_동화약품()
	종목정보_맵[종목.G종목코드()] = 종목

	종목 = F종목_삼성전자()
	종목정보_맵[종목.G종목코드()] = 종목

	return 종목정보_맵
}

func F일일시세_동화약품_20131101() *C일일가격정보 {
	일자, _ := time.Parse("2006-01-02", "2013-11-01")

	일일시세 := new(S일일가격정보)
	일일시세.S식별코드(1)
	일일시세.S종목(F종목_동화약품())
	일일시세.S일자(일자)
	일일시세.S시가(5590)
	일일시세.S고가(5660)
	일일시세.S저가(5550)
	일일시세.S종가(5590)
	일일시세.S조정종가(5516.52)
	일일시세.S거래량(20100)
	일일시세.M조정가격_재계산()

	return 일일시세.G상수형_구조체()
}

func F일일시세_동화약품_20140312() *C일일가격정보 {
	일자, _ := time.Parse("2006-01-02", "2014-03-12")

	일일시세 := new(S일일가격정보)
	일일시세.S식별코드(2)
	일일시세.S종목(F종목_동화약품())
	일일시세.S일자(일자)
	일일시세.S시가(5690)
	일일시세.S고가(5710)
	일일시세.S저가(5530)
	일일시세.S종가(5650)
	일일시세.S조정종가(5650)
	일일시세.S거래량(148500)
	일일시세.M조정가격_재계산()

	return 일일시세.G상수형_구조체()
}

func F일일시세_동화약품_20140313() *C일일가격정보 {
	일자, _ := time.Parse("2006-01-02", "2014-03-13")

	일일시세 := new(S일일가격정보)
	일일시세.S식별코드(3)
	일일시세.S종목(F종목_동화약품())
	일일시세.S일자(일자)
	일일시세.S시가(5650)
	일일시세.S고가(5890)
	일일시세.S저가(5580)
	일일시세.S종가(5860)
	일일시세.S조정종가(5860)
	일일시세.S거래량(218100)
	일일시세.M조정가격_재계산()

	return 일일시세.G상수형_구조체()
}

func F일일시세_삼성전자_20140425() *C일일가격정보 {
	일자, _ := time.Parse("2006-01-02", "2014-04-25")

	일일시세 := new(S일일가격정보)
	일일시세.S식별코드(4)
	일일시세.S종목(F종목_삼성전자())
	일일시세.S일자(일자)
	일일시세.S시가(1420000.00)
	일일시세.S고가(1428000.00)
	일일시세.S저가(1398000.00)
	일일시세.S종가(1399000.00)
	일일시세.S조정종가(1399000.00)
	일일시세.S거래량(305300)
	일일시세.M조정가격_재계산()

	return 일일시세.G상수형_구조체()
}

func F일일시세_삼성전자_20140708() *C일일가격정보 {
	일자, _ := time.Parse("2006-01-02", "2014-07-08")

	일일시세 := new(S일일가격정보)
	일일시세.S식별코드(5)
	일일시세.S종목(F종목_삼성전자())
	일일시세.S일자(일자)
	일일시세.S시가(1294000.00)
	일일시세.S고가(1319000.00)
	일일시세.S저가(1287000.00)
	일일시세.S종가(1295000.00)
	일일시세.S조정종가(1295000.00)
	일일시세.S거래량(241400)
	일일시세.M조정가격_재계산()

	return 일일시세.G상수형_구조체()
}

func F일일시세_삼성전자_20140709() *C일일가격정보 {
	일자, _ := time.Parse("2006-01-02", "2014-07-09")

	일일시세 := new(S일일가격정보)
	일일시세.S식별코드(5)
	일일시세.S종목(F종목_삼성전자())
	일일시세.S일자(일자)
	일일시세.S시가(1287000.00)
	일일시세.S고가(1308000.00)
	일일시세.S저가(1283000.00)
	일일시세.S종가(1308000.00)
	일일시세.S조정종가(1308000.00)
	일일시세.S거래량(236500)
	일일시세.M조정가격_재계산()

	return 일일시세.G상수형_구조체()
}

func F종목별_일일가격정보_모음_테스트용() *S종목별_일일가격정보_모음 {
	일일가격정보_모음 := make([]*C일일가격정보, 0)
	일일가격정보_모음 = append(일일가격정보_모음, F일일시세_동화약품_20131101())
	일일가격정보_모음 = append(일일가격정보_모음, F일일시세_동화약품_20140312())
	일일가격정보_모음 = append(일일가격정보_모음, F일일시세_동화약품_20140313())

	종목별_일일가격정보_모음 := new(S종목별_일일가격정보_모음)
	종목별_일일가격정보_모음.S종목(F종목_동화약품())
	종목별_일일가격정보_모음.S추가(일일가격정보_모음)

	return 종목별_일일가격정보_모음
}

func F일일가격정보_모음_맵_테스트용() *S일일가격정보_모음_맵 {
	일일가격정보_모음 := make([]*C일일가격정보, 0)
	일일가격정보_모음 = append(일일가격정보_모음, F일일시세_동화약품_20131101())
	일일가격정보_모음 = append(일일가격정보_모음, F일일시세_동화약품_20140312())
	일일가격정보_모음 = append(일일가격정보_모음, F일일시세_동화약품_20140313())
	일일가격정보_모음 = append(일일가격정보_모음, F일일시세_삼성전자_20140425())
	일일가격정보_모음 = append(일일가격정보_모음, F일일시세_삼성전자_20140708())
	일일가격정보_모음 = append(일일가격정보_모음, F일일시세_삼성전자_20140709())

	일일가격정보_모음_맵 := new(S일일가격정보_모음_맵)
	일일가격정보_모음_맵.S추가(일일가격정보_모음)

	return 일일가격정보_모음_맵
}

func F테스트용_비어있는_전략(식별코드 uint64) I전략 {
	s := new(테스트용_비어있는_전략)
	s.식별코드 = 식별코드
	s.위험관리 = F테스트용_비어있는_위험관리()
	
	return s
}
type 테스트용_비어있는_전략 struct {
	식별코드 uint64
	위험관리 I위험관리
	자본배분액 float64
}
func (s *테스트용_비어있는_전략) G식별코드() uint64 { return s.식별코드 }
func (s *테스트용_비어있는_전략) G포트폴리오() I포트폴리오 { return nil }
func (s *테스트용_비어있는_전략) G위험관리() I위험관리 { return nil }

func F테스트용_비어있는_위험관리() I위험관리 {
	필요한_매개변수_모음 := make([]string, 0)
	필요한_매개변수_모음 = append(필요한_매개변수_모음, "정수형_매개변수")
	필요한_매개변수_모음 = append(필요한_매개변수_모음, "실수형_매개변수")
	필요한_매개변수_모음 = append(필요한_매개변수_모음, "문자열형_매개변수")
	
	s := new(S테스트용_위험관리)
	s.M식별코드 = 100
	s.M이름 = "테스트용_비어있는_위험관리"
	s.M필요한_매개변수_모음 = 필요한_매개변수_모음
	s.M검토_펑션 = func(위험관리_매개변수 *S위험관리_매개변수) (*S위험관리_검토결과, error) {
						return nil, nil
					} 

	return s
}
type S테스트용_위험관리 struct {
	M식별코드 uint64
	M이름 string
	M필요한_매개변수_모음 []string
	M검토_펑션 func(위험관리_매개변수 *S위험관리_매개변수) (*S위험관리_검토결과, error)
}
func (s *S테스트용_위험관리) G식별코드() uint64 { return s.M식별코드 }
func (s *S테스트용_위험관리) G이름() string { return s.M이름 }
func (s *S테스트용_위험관리) G필요한_매개변수_모음() []string { return s.M필요한_매개변수_모음 }
func (s *S테스트용_위험관리) G검토(위험관리_매개변수 *S위험관리_매개변수) (*S위험관리_검토결과, error) {
	if s.M검토_펑션 == nil {
		return nil, nil
	} else {
		return s.M검토_펑션(위험관리_매개변수)
	}
}

type S테스트용_종목별_포트폴리오 struct {
	M식별코드 uint64
	M종목 *C종목	
	M단가 I통화
	M롱포지션_수량 int64
	M숏포지션_수량 int64
}
func (s *S테스트용_종목별_포트폴리오) G식별코드() uint64 { return s.M식별코드 }
func (s *S테스트용_종목별_포트폴리오) G종목() *C종목 { return s.M종목 }
func (s *S테스트용_종목별_포트폴리오) G단가() float64 { return s.M단가.G값() }
func (s *S테스트용_종목별_포트폴리오) G롱포지션_수량() int64 { return s.M롱포지션_수량 }
func (s *S테스트용_종목별_포트폴리오) G숏포지션_수량() int64 { return s.M숏포지션_수량 }
func (s *S테스트용_종목별_포트폴리오) G순_수량() int64 { return s.M롱포지션_수량 - s.M숏포지션_수량 }
func (s *S테스트용_종목별_포트폴리오) G총_수량() int64 { return s.M롱포지션_수량 + s.M숏포지션_수량 }
func (s *S테스트용_종목별_포트폴리오) G롱포지션_금액() float64 { return F반올림_통화(float64(s.M롱포지션_수량) * s.M단가.G값()) }
func (s *S테스트용_종목별_포트폴리오) G숏포지션_금액() float64 { return F반올림_통화(float64(s.M숏포지션_수량) * s.M단가.G값()) }
func (s *S테스트용_종목별_포트폴리오) G순_금액() float64 { return F반올림_통화(float64(s.M롱포지션_수량 - s.M숏포지션_수량) * s.M단가.G값()) }
func (s *S테스트용_종목별_포트폴리오) G총_금액() float64 { return F반올림_통화(float64(s.M롱포지션_수량 + s.M숏포지션_수량) * s.M단가.G값()) }

type S테스트용_포트폴리오 struct {
	M식별코드 uint64
	M통화_식별코드 P통화
	M포트폴리오_내용 []I종목별_포트폴리오
}
func (s *S테스트용_포트폴리오) G식별코드() uint64 { return s.M식별코드 }
func (s *S테스트용_포트폴리오) G통화_식별코드() P통화 { return s.M통화_식별코드 }
func (s *S테스트용_포트폴리오) G보유_종목_모음() []*C종목 {
	종목_모음 := make([]*C종목, len(s.M포트폴리오_내용))
	
	for 인덱스, 종목별_포트폴리오 := range s.M포트폴리오_내용 {
		종목_모음[인덱스] = 종목별_포트폴리오.G종목()
	}
	
	return 종목_모음
}
func (s *S테스트용_포트폴리오) G종목별_포트폴리오(종목코드 string) I종목별_포트폴리오 {
	for _, 종목별_포트폴리오 := range s.M포트폴리오_내용 {
		if 종목별_포트폴리오.G종목().G종목코드() == 종목코드 {
			return 종목별_포트폴리오
		}
	}
	
	return nil
}
func (s *S테스트용_포트폴리오) G전종목_포트폴리오() []I종목별_포트폴리오 {
	return F슬라이스_복사(s.M포트폴리오_내용).([]I종목별_포트폴리오)
}
func (s *S테스트용_포트폴리오) G롱포지션_금액() I통화 {
	합계 := float64(0.0)
	
	for _, 종목별_포트폴리오 := range s.M포트폴리오_내용 {
		합계 = 합계 + 종목별_포트폴리오.G롱포지션_금액().G값()
	}
	
	return F통화_생성(s.M통화_식별코드, 합계)
}
func (s *S테스트용_포트폴리오) G숏포지션_금액() I통화 {
	합계 := float64(0.0)
	
	for _, 종목별_포트폴리오 := range s.M포트폴리오_내용 {
		합계 = 합계 + 종목별_포트폴리오.G숏포지션_금액().G값()
	}
	
	return F통화_생성(s.M통화_식별코드, 합계)
}
func (s *S테스트용_포트폴리오) G순_금액() I통화 {
	합계 := float64(0.0)
	
	for _, 종목별_포트폴리오 := range s.M포트폴리오_내용 {
		합계 = 합계 + 종목별_포트폴리오.G순_금액().G값()
	}
	
	return F통화_생성(s.M통화_식별코드, 합계)
}
func (s *S테스트용_포트폴리오) G총_금액() I통화 {
	합계 := float64(0.0)
	
	for _, 종목별_포트폴리오 := range s.M포트폴리오_내용 {
		합계 = 합계 + 종목별_포트폴리오.G총_금액().G값()
	}
	
	return F통화_생성(s.M통화_식별코드, 합계)
}