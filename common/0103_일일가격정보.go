package common

import (
	"bytes"
	"errors"
	"log"
	"time"
)

type C일일가격정보 struct {
	식별코드 *C부호없는_정수64
	종목   *C종목
	일자   *C시점
	시가   *C실수64
	고가   *C실수64
	저가   *C실수64
	종가   *C실수64
	조정종가 *C실수64
	거래량  *C실수64
	조정계수 *C실수64
	조정시가 *C실수64
	조정고가 *C실수64
	조정저가 *C실수64
}

func (c *C일일가격정보) G식별코드() uint64  { return c.식별코드.G값() }
func (c *C일일가격정보) G종목() *C종목      { return c.종목 }
func (c *C일일가격정보) G일자() time.Time { return c.일자.G값() }
func (c *C일일가격정보) G시가() float64   { return c.시가.G값() }
func (c *C일일가격정보) G고가() float64   { return c.고가.G값() }
func (c *C일일가격정보) G저가() float64   { return c.저가.G값() }
func (c *C일일가격정보) G종가() float64   { return c.종가.G값() }
func (c *C일일가격정보) G조정종가() float64 { return c.조정종가.G값() }
func (c *C일일가격정보) G거래량() float64  { return c.거래량.G값() }
func (c *C일일가격정보) G조정계수() float64 { return c.조정계수.G값() }
func (c *C일일가격정보) G조정시가() float64 { return c.조정시가.G값() }
func (c *C일일가격정보) G조정고가() float64 { return c.조정고가.G값() }
func (c *C일일가격정보) G조정저가() float64 { return c.조정저가.G값() }
func (c *C일일가격정보) G같음(일일가격정보 *C일일가격정보) bool {
	c1 := c
	c2 := 일일가격정보

	if c1.G식별코드() == c2.G식별코드() ||
		c1.G종목() == c2.G종목() ||
		c1.G일자() == c2.G일자() ||
		c1.G시가() == c2.G시가() ||
		c1.G고가() == c2.G고가() ||
		c1.G저가() == c2.G저가() ||
		c1.G종가() == c2.G종가() ||
		c1.G조정종가() == c2.G조정종가() ||
		c1.G거래량() == c2.G거래량() ||
		c1.G조정계수() == c2.G조정계수() ||
		c1.G조정시가() == c2.G조정시가() ||
		c1.G조정고가() == c2.G조정고가() ||
		c1.G조정저가() == c2.G조정저가() {
		return true
	}

	return false
}
func (c *C일일가격정보) G구조체() *S일일가격정보 {
	s := new(S일일가격정보)
	s.S식별코드(c.G식별코드())
	s.S종목(c.G종목())
	s.S일자(F일자(c.G일자()))
	s.S시가(c.G시가())
	s.S고가(c.G고가())
	s.S저가(c.G저가())
	s.S종가(c.G종가())
	s.S조정종가(c.G조정종가())
	s.S거래량(c.G거래량())
	s.M조정가격_재계산()

	return s
}
func (c *C일일가격정보) G키() string {
	if c.종목 == nil {
		c.G일자().Format("2006-01-02")
	}

	return c.종목.G종목코드() + "_" + c.G일자().Format("2006-01-02")
}

func F일일가격정보_생성(
	식별코드 uint64,
	종목 *C종목,
	일자 time.Time,
	시가 float64,
	고가 float64,
	저가 float64,
	종가 float64,
	조정종가 float64,
	거래량 float64) *C일일가격정보 {
	c := new(C일일가격정보)
	c.식별코드 = F부호없는_정수64_생성(식별코드)
	c.종목 = 종목
	c.일자 = F시점_생성(F일자(일자))
	c.시가 = F실수64_생성(시가)
	c.고가 = F실수64_생성(고가)
	c.저가 = F실수64_생성(저가)
	c.종가 = F실수64_생성(종가)
	c.조정종가 = F실수64_생성(F반올림_통화(조정종가))
	c.거래량 = F실수64_생성(거래량)

	c.조정계수 = F실수64_생성(조정종가 / 종가)

	c.조정시가 = F실수64_생성(F반올림_통화(시가 * c.조정계수.G값()))
	c.조정고가 = F실수64_생성(F반올림_통화(고가 * c.조정계수.G값()))
	c.조정저가 = F실수64_생성(F반올림_통화(저가 * c.조정계수.G값()))

	c.조정계수 = F실수64_생성(F반올림(c.조정계수.G값(), 4))

	return c
}

type S일일가격정보 struct {
	식별코드 uint64
	종목   *C종목
	일자   time.Time
	시가   float64
	고가   float64
	저가   float64
	종가   float64
	조정계수 float64
	조정시가 float64
	조정고가 float64
	조정저가 float64
	조정종가 float64
	거래량  float64
}

func (s *S일일가격정보) G식별코드() uint64  { return s.식별코드 }
func (s *S일일가격정보) G종목() *C종목      { return s.종목 }
func (s *S일일가격정보) G일자() time.Time { return s.일자 }
func (s *S일일가격정보) G시가() float64   { return s.시가 }
func (s *S일일가격정보) G고가() float64   { return s.고가 }
func (s *S일일가격정보) G저가() float64   { return s.저가 }
func (s *S일일가격정보) G종가() float64   { return s.종가 }
func (s *S일일가격정보) G조정계수() float64 { return s.조정계수 }
func (s *S일일가격정보) G조정시가() float64 { return s.조정시가 }
func (s *S일일가격정보) G조정고가() float64 { return s.조정고가 }
func (s *S일일가격정보) G조정저가() float64 { return s.조정저가 }
func (s *S일일가격정보) G조정종가() float64 { return s.조정종가 }
func (s *S일일가격정보) G거래량() float64  { return s.거래량 }
func (s *S일일가격정보) G키() string {
	if s.종목 == nil {
		return ""
	}
	//if s.일자 == nil { return "" }

	return s.종목.G종목코드() + "_" + s.G일자().Format("2006-01-02")
}

func (s *S일일가격정보) S식별코드(식별코드 uint64) { s.식별코드 = 식별코드 }
func (s *S일일가격정보) S종목(종목 *C종목)       { s.종목 = 종목 }
func (s *S일일가격정보) S일자(일자 time.Time)  { s.일자 = F일자(일자) }
func (s *S일일가격정보) S시가(시가 float64)    { s.시가 = F반올림_통화(시가) }
func (s *S일일가격정보) S고가(고가 float64)    { s.고가 = F반올림_통화(고가) }
func (s *S일일가격정보) S저가(저가 float64)    { s.저가 = F반올림_통화(저가) }
func (s *S일일가격정보) S종가(종가 float64)    { s.종가 = F반올림_통화(종가) }
func (s *S일일가격정보) S조정종가(조정종가 float64) {
	s.조정종가 = F반올림_통화(조정종가)
}
func (s *S일일가격정보) S거래량(거래량 float64) { s.거래량 = F반올림(거래량, 0) }
func (s *S일일가격정보) M조정가격_재계산() {
	if s.종가 <= 0.0 {
		s.조정계수 = 0.0
		s.조정시가 = 0.0
		s.조정고가 = 0.0
		s.조정저가 = 0.0
	} else {
		s.조정계수 = s.조정종가 / s.종가

		s.조정시가 = F반올림_통화(s.시가 * s.조정계수)
		s.조정고가 = F반올림_통화(s.고가 * s.조정계수)
		s.조정저가 = F반올림_통화(s.저가 * s.조정계수)
		s.조정종가 = F반올림_통화(s.조정종가)

		s.조정계수 = F반올림(s.조정계수, 4)
	}
}
func (s *S일일가격정보) G상수형_구조체() *C일일가격정보 {
	return F일일가격정보_생성(s.식별코드,
		s.종목,
		s.일자,
		s.시가,
		s.고가,
		s.저가,
		s.종가,
		s.조정종가,
		s.거래량)
}

// 일일가격정보를 하나의 종목에 대하여 날짜순으로 정렬하여 보관하는 구조체
type S종목별_일일가격정보_모음 struct {
	종목        *C종목
	일일가격정보_모음 []*C일일가격정보
}

func (s *S종목별_일일가격정보_모음) 초기화() {
	s.일일가격정보_모음 = make([]*C일일가격정보, 0)
}
func (s *S종목별_일일가격정보_모음) G종목() *C종목 { return s.종목 }
func (s *S종목별_일일가격정보_모음) G수량() int  { return len(s.일일가격정보_모음) }
func (s *S종목별_일일가격정보_모음) G일일가격정보(일자 time.Time) *C일일가격정보 {
	if s.일일가격정보_모음 == nil {
		s.초기화()
	}

	var 일일가격정보 *C일일가격정보

	일자 = F일자(일자)

	for _, 일일가격정보 = range s.일일가격정보_모음 {
		if 일자.Equal(일일가격정보.G일자()) {
			return 일일가격정보
		}
	}

	return nil
}
func (s *S종목별_일일가격정보_모음) G슬라이스() []*C일일가격정보 {
	if s.일일가격정보_모음 == nil {
		s.초기화()
	}

	return s.일일가격정보_모음
}
func (s *S종목별_일일가격정보_모음) G맵() map[string]*C일일가격정보 {
	if s.일일가격정보_모음 == nil {
		s.초기화()
	}

	맵 := make(map[string]*C일일가격정보)

	for _, 일일가격정보 := range s.일일가격정보_모음 {
		맵[일일가격정보.G키()] = 일일가격정보
	}

	return 맵
}
func (s *S종목별_일일가격정보_모음) S종목(종목 *C종목) {
	if s.종목 != nil &&
		s.일일가격정보_모음 != nil &&
		len(s.일일가격정보_모음) > 0 {
		log.Println("S종목별_일일가격정보_모음.S종목() : 종목이 이미 설정되어 있습니다. 코딩 에러일 가능성이 높습니다.")
	}

	s.종목 = 종목
}
func (s *S종목별_일일가격정보_모음) S단일내용_추가(일일가격정보 *C일일가격정보) {
	if s.일일가격정보_모음 == nil {
		s.초기화()
	}

	일일가격정보_모음 := make([]*C일일가격정보, 1)
	일일가격정보_모음[0] = 일일가격정보

	s.S추가(일일가격정보_모음)
}
func (s *S종목별_일일가격정보_모음) S추가(일일가격정보_모음 []*C일일가격정보) {
	if s.일일가격정보_모음 == nil {
		s.초기화()
	}

	var 중복 bool = false
	var 중복_인덱스 int

	// 종목별로 종목코드에 따라 별도의 슬라이스에 원소로 추가.
	for 인덱스, 일일가격정보 := range 일일가격정보_모음 {
		if 일일가격정보 == nil {
			log.Println("common.S종목별_일일가격정보_모음.S추가() : nil 원소가 있습니다.", 인덱스)
			continue
		}
		if 일일가격정보.G종목() == nil {
			log.Println("common.S종목별_일일가격정보_모음.S추가() : 종목이 nil인 원소가 있습니다.", 인덱스)
			continue
		}
		if 일일가격정보.G종목().G종목코드() == "" {
			log.Println("common.S종목별_일일가격정보_모음.S추가() : 종목코드가 없는 원소가 있습니다.", 인덱스)
			continue
		}

		if s.종목 != nil &&
			s.종목.G종목코드() != "" &&
			s.종목.G종목코드() != 일일가격정보.G종목().G종목코드() {
			if !F테스트_모드() {
				log.Printf("common.S종목별_일일가격정보_모음.S추가() : "+
					"종목이 일치하지 않습니다. 인덱스 %v, 예상값 %v, 실제 종목코드 %v",
					인덱스, s.종목.G종목코드(), 일일가격정보.G종목().G종목코드())
			}

			continue
		}

		// 무효한 데이터는 추가하지 않음.
		if 일일가격정보.G시가() == 0.0 ||
			일일가격정보.G고가() == 0.0 ||
			일일가격정보.G저가() == 0.0 ||
			일일가격정보.G종가() == 0.0 {
			continue
		}

		중복 = false

		for 내부_인덱스, 기존_일일가격정보 := range s.일일가격정보_모음 {
			if 기존_일일가격정보.G일자().Equal(일일가격정보.G일자()) {
				//log.Println("S종목별_일일가격정보_모음.S추가() : 중복 데이터.", 일일가격정보.G키())
				중복 = true
				중복_인덱스 = 내부_인덱스

				continue
			}
		}

		if 중복 {
			s.일일가격정보_모음[중복_인덱스] = 일일가격정보
		} else {
			s.일일가격정보_모음 = append(s.일일가격정보_모음, 일일가격정보)
		}
	}

	s.정렬()
}
func (s *S종목별_일일가격정보_모음) 정렬() {
	var 일자1, 일자2 time.Time
	슬라이스_길이 := len(s.일일가격정보_모음)

	새로_정렬할_필요없음 := true
	for 인덱스 := 0; 인덱스 < (슬라이스_길이 - 1); 인덱스++ {
		일자1 = s.일일가격정보_모음[인덱스].G일자()
		일자2 = s.일일가격정보_모음[인덱스+1].G일자()

		if 일자1.After(일자2) {
			새로_정렬할_필요없음 = false
			break
		}
	}

	// 이미 정렬되어 있으면 현재 종목을 건너뛴다.
	if 새로_정렬할_필요없음 {
		return
	}

	// 정렬 작업 시작.
	for 인덱스1 := 0; 인덱스1 < 슬라이스_길이; 인덱스1++ {
		for 인덱스2 := 인덱스1; 인덱스2 < 슬라이스_길이; 인덱스2++ {
			일자1 = s.일일가격정보_모음[인덱스1].G일자()
			일자2 = s.일일가격정보_모음[인덱스2].G일자()

			if 일자1.After(일자2) {
				일일가격정보_임시 := s.일일가격정보_모음[인덱스1]
				s.일일가격정보_모음[인덱스1] = s.일일가격정보_모음[인덱스2]
				s.일일가격정보_모음[인덱스2] = 일일가격정보_임시
			}
		}
	}
}

type S일일가격정보_모음_맵 struct {
	일일가격정보_모음_맵 map[string]*S종목별_일일가격정보_모음
}

func (s *S일일가격정보_모음_맵) 초기화() {
	s.일일가격정보_모음_맵 = make(map[string]*S종목별_일일가격정보_모음)
}
func (s *S일일가격정보_모음_맵) G맵() map[string]*S종목별_일일가격정보_모음 {
	if s.일일가격정보_모음_맵 == nil {
		s.초기화()
	}

	return s.일일가격정보_모음_맵
}
func (s *S일일가격정보_모음_맵) G종목별_내용(종목 *C종목) *S종목별_일일가격정보_모음 {
	if s.일일가격정보_모음_맵 == nil {
		s.초기화()
	}

	if 종목 == nil {
		log.Println("S일일가격정보_모음_맵.G종목별_내용() : 종목이 nil입니다.")
		return nil
	}

	if 종목.G종목코드() == "" {
		log.Println("S일일가격정보_모음_맵.G종목별_내용() : 종목코드가 비어있습니다.")
		return nil
	}

	종목별_일일가격정보_모음 := s.일일가격정보_모음_맵[종목.G종목코드()]

	if 종목별_일일가격정보_모음 == nil {
		종목별_일일가격정보_모음 = new(S종목별_일일가격정보_모음)
		종목별_일일가격정보_모음.S종목(종목)
		종목별_일일가격정보_모음.S추가(make([]*C일일가격정보, 0))

		s.일일가격정보_모음_맵[종목.G종목코드()] = 종목별_일일가격정보_모음
	}

	return 종목별_일일가격정보_모음
}
func (s *S일일가격정보_모음_맵) G슬라이스() []*C일일가격정보 {
	if s.일일가격정보_모음_맵 == nil {
		s.초기화()
	}

	일일가격정보_모음 := make([]*C일일가격정보, 0)

	for _, 종목별_일일가격정보_모음 := range s.일일가격정보_모음_맵 {
		일일가격정보_모음 = append(일일가격정보_모음, 종목별_일일가격정보_모음.G슬라이스()...)
	}

	return 일일가격정보_모음
}
func (s *S일일가격정보_모음_맵) S단일내용_추가(일일가격정보 *C일일가격정보) {
	if s.일일가격정보_모음_맵 == nil {
		s.초기화()
	}

	일일가격정보_모음 := make([]*C일일가격정보, 1)
	일일가격정보_모음[0] = 일일가격정보

	s.S추가(일일가격정보_모음)
}
func (s *S일일가격정보_모음_맵) S추가(일일가격정보_모음 []*C일일가격정보) {
	if s.일일가격정보_모음_맵 == nil {
		s.초기화()
	}

	일일가격정보_임시저장소 := make(map[string][]*C일일가격정보)

	for 인덱스, 일일가격정보 := range 일일가격정보_모음 {
		if 일일가격정보 == nil {
			log.Println("S일일가격정보_모음_맵.S내용() : 일일가격정보가 nil입니다.", 인덱스)
			continue
		}

		if 일일가격정보.G종목() == nil {
			log.Println("S일일가격정보_모음_맵.S내용() : 일일가격정보 종목이 nil입니다.", 인덱스)
			continue
		}

		if 일일가격정보.G종목().G종목코드() == "" {
			log.Println("S일일가격정보_모음_맵.S내용() : 일일가격정보 종목코드가 비어있습니다.", 인덱스)
			continue
		}

		종목코드 := 일일가격정보.G종목().G종목코드()
		일일가격정보_모음 := 일일가격정보_임시저장소[종목코드]

		if 일일가격정보_모음 == nil {
			일일가격정보_모음 = make([]*C일일가격정보, 0)
		}

		일일가격정보_모음 = append(일일가격정보_모음, 일일가격정보)
		일일가격정보_임시저장소[종목코드] = 일일가격정보_모음
	}

	for 종목코드, 일일가격정보_모음 := range 일일가격정보_임시저장소 {
		종목별_일일가격정보_모음 := s.일일가격정보_모음_맵[종목코드]

		if 종목별_일일가격정보_모음 == nil {
			종목별_일일가격정보_모음 = new(S종목별_일일가격정보_모음)
			종목별_일일가격정보_모음.S종목(F종목_검색(종목코드))
			s.일일가격정보_모음_맵[종목코드] = 종목별_일일가격정보_모음
		}

		종목별_일일가격정보_모음.S추가(일일가격정보_모음)
	}
}

func F종목별_일일가격정보_모음(종목 *C종목) (*S종목별_일일가격정보_모음, error) {
	데이터베이스, 에러 := F데이터베이스_연결()
	if 에러 != nil {
		log.Print("common.F일일가격정보_모음() : DB열기 에러.", 에러)

		return nil, 에러
	}
	defer 데이터베이스.Close()

	sql := new(bytes.Buffer)
	sql.WriteString("SELECT")
	sql.WriteString("	id,")
	sql.WriteString("	stock_info_id,")
	sql.WriteString("	priced_on,")
	sql.WriteString("	open,")
	sql.WriteString("	high,")
	sql.WriteString("	low,")
	sql.WriteString("	close,")
	sql.WriteString("	adj_coeff,")
	sql.WriteString("	adj_open,")
	sql.WriteString("	adj_high,")
	sql.WriteString("	adj_low,")
	sql.WriteString("	adj_close,")
	sql.WriteString("	volumn ")
	sql.WriteString("FROM " + F일일가격정보_테이블() + " ")
	sql.WriteString("WHERE stock_info_id = ?")
	sql.WriteString("	AND open > 0")
	sql.WriteString("	AND high > 0")
	sql.WriteString("	AND low > 0")
	sql.WriteString("	AND close > 0 ")
	sql.WriteString("ORDER BY")
	sql.WriteString("	priced_on DESC")

	// 테스트할 때 종목 식별코드가 제대로 입력되지 않아서 에러가 발생하므로 DB기준으로 재검색
	종목 = F종목_검색(종목.G종목코드())

	행_모음, 에러 := 데이터베이스.Query(sql.String(), 종목.G식별코드())
	if 에러 != nil {
		return nil, 에러
	}
	defer 행_모음.Close()

	일일가격정보_모음 := make([]*C일일가격정보, 0)

	var 식별코드, 종목_식별코드 uint64
	var 일자 time.Time
	var 시가, 고가, 저가, 종가 float64
	var 조정계수, 조정시가, 조정고가, 조정저가, 조정종가, 거래량 float64

	인덱스 := 0
	for 행_모음.Next() {
		에러 := 행_모음.Scan(
			&식별코드,
			&종목_식별코드,
			&일자,
			&시가,
			&고가,
			&저가,
			&종가,
			&조정계수,
			&조정시가,
			&조정고가,
			&조정저가,
			&조정종가,
			&거래량)

		if 에러 != nil {
			log.Printf("common.F종목별_일일가격정보_모음() 행_모음.Scan() 에러 발생. 인덱스 %v", 인덱스)
			log.Println(에러)

			return nil, 에러
		}

		일일가격정보 := F일일가격정보_생성(
			식별코드,
			종목,
			일자,
			시가,
			고가,
			저가,
			종가,
			조정종가,
			거래량)

		일일가격정보_모음 = append(일일가격정보_모음, 일일가격정보)

		인덱스++
	}

	일일가격정보_모음_반환값 := new(S종목별_일일가격정보_모음)
	일일가격정보_모음_반환값.S종목(종목)
	일일가격정보_모음_반환값.S추가(일일가격정보_모음)

	return 일일가격정보_모음_반환값, nil
}

func F일일가격정보_모음_DB기록(종목별_일일가격정보_모음 *S종목별_일일가격정보_모음) error {
	종목 := 종목별_일일가격정보_모음.G종목()

	if 종목 == nil {
		log.Println("common.F일일가격정보_모음_DB기록() : 종목이 nil입니다.")

		return errors.New("common.F일일가격정보_모음_DB기록() : 종목이 nil입니다.")
	}

	if 종목.G종목코드() == "" {
		log.Println("common.F일일가격정보_모음_DB기록() : 종목코드가 없습니다.")

		return errors.New("common.F일일가격정보_모음_DB기록() : 종목코드가 없습니다.")
	}

	sql := new(bytes.Buffer)
	sql.WriteString("INSERT INTO ")
	sql.WriteString(F일일가격정보_테이블() + " ")
	sql.WriteString("(")
	sql.WriteString("stock_info_id,")
	sql.WriteString("priced_on,")
	sql.WriteString("open,")
	sql.WriteString("high,")
	sql.WriteString("low,")
	sql.WriteString("close,")
	sql.WriteString("adj_coeff,")
	sql.WriteString("adj_open,")
	sql.WriteString("adj_high,")
	sql.WriteString("adj_low,")
	sql.WriteString("adj_close,")
	sql.WriteString("volumn")
	sql.WriteString(") VALUES (")
	sql.WriteString("?,?,?,?,?,?,?,?,?,?,?,?")
	sql.WriteString(")")
	sql추가 := sql.String()

	sql = new(bytes.Buffer)
	sql.WriteString("UPDATE ")
	sql.WriteString(F일일가격정보_테이블() + " ")
	sql.WriteString("SET")
	sql.WriteString("	adj_coeff = ?,")
	sql.WriteString("	adj_open = ?,")
	sql.WriteString("	adj_high = ?,")
	sql.WriteString("	adj_low = ?,")
	sql.WriteString("	adj_close = ? ")
	sql.WriteString("WHERE")
	sql.WriteString("	id = ?")
	sql갱신 := sql.String()

	데이터베이스, 에러 := F데이터베이스_연결()
	if 에러 != nil {
		log.Println("F일일가격정보_모음_DB기록() : F데이터베이스_연결() 에러.", 에러)
		return 에러
	}
	defer 데이터베이스.Close()

	기존_일일가격정보_모음, 에러 := F종목별_일일가격정보_모음(종목)
	if 에러 != nil {
		log.Println("F일일가격정보_모음_DB기록() : F종목별_일일가격정보_모음() 에러.", 에러)
		return 에러
	}

	기존_일일가격정보_맵 := 기존_일일가격정보_모음.G맵()

	트랜잭션, 에러 := 데이터베이스.Begin()
	if 에러 != nil {
		log.Println("F일일가격정보_모음_DB기록() : 트랜잭션 초기화 에러.", 에러)
		return 에러
	}

	stmt추가, 에러 := 트랜잭션.Prepare(sql추가)
	if 에러 != nil {
		log.Println("F일일가격정보_모음_DB기록() stmt추가 준비에러 : ", 에러)
		return 에러
	}
	defer stmt추가.Close()

	stmt갱신, 에러 := 트랜잭션.Prepare(sql갱신)
	if 에러 != nil {
		log.Println("F일일가격정보_모음_DB기록() : stmt갱신 초기화 에러.", 에러)
		return 에러
	}
	defer stmt갱신.Close()

	F종목정보_맵_초기화()

	for _, 일일가격정보 := range 종목별_일일가격정보_모음.G슬라이스() {
		기존_일일가격정보, 이미_존재함 := 기존_일일가격정보_맵[일일가격정보.G키()]

		if 이미_존재함 {
			if 일일가격정보.G종목().G종목코드() != 기존_일일가격정보.G종목().G종목코드() ||
				일일가격정보.G일자() != 기존_일일가격정보.G일자() {
				log.Println("F일일가격정보_모음_DB기록() : " +
					"이미 존재하는 데이터와 종목 혹은 날짜가 일치하지 않음. " +
					일일가격정보.G키() + ", " +
					기존_일일가격정보.G키())
				return errors.New("F일일가격정보_모음_DB기록() : " +
					"이미 존재하는 데이터와 종목 혹은 날짜가 일치하지 않음. " +
					일일가격정보.G키() + ", " +
					기존_일일가격정보.G키())
			}

			// 조정종가가 바뀐 게 없으면 굳이 저장할 필요없음.
			if 일일가격정보.G종가() == 기존_일일가격정보.G종가() &&
				일일가격정보.G조정종가() == 기존_일일가격정보.G조정종가() {
				continue
			}

			// 조정종가가 변한 것은 반영하되 기존의 검증을 거친 데이터는 최대한 살린다.
			새로운_조정계수 := 일일가격정보.G조정종가() / 일일가격정보.G종가()
			새로운_조정종가 := 기존_일일가격정보.G종가() * 새로운_조정계수

			if 새로운_조정종가 == 일일가격정보.G조정종가()-1 ||
				새로운_조정종가 == 일일가격정보.G조정종가()+1 {
				// 올림, 내림, 반올림등의 처리로 인한 1원 차이가 날 경우 처리
				새로운_조정종가 = 일일가격정보.G조정종가()
			}

			s := 기존_일일가격정보.G구조체()
			s.S조정종가(새로운_조정종가)
			s.M조정가격_재계산()

			// 기존 일일가격정보 갱신
			_, 에러 = stmt갱신.Exec(
				s.G조정계수(),
				s.G조정시가(),
				s.G조정고가(),
				s.G조정저가(),
				s.G조정종가(),
				기존_일일가격정보.G식별코드())

			if 에러 != nil {
				log.Println("common.F일일가격정보_모음_DB기록() : stmt갱신 실행에러 : "+일일가격정보.G키(), 에러)

				롤백_에러 := 트랜잭션.Rollback()
				if 롤백_에러 != nil {
					log.Println("common.F일일가격정보_모음_DB기록() : 롤백 에러가 발생하였습니다.", 롤백_에러)
				}

				return 에러
			}
		} else {
			// 혹여 종목의 식별코드(id)가 제대로 설정되지 않았을 경우에 대비하여 종목 재검색
			종목 := F종목_검색(일일가격정보.G종목().G종목코드())

			_, 에러 = stmt추가.Exec(
				종목.G식별코드(),
				F일자(일일가격정보.G일자()),
				일일가격정보.G시가(),
				일일가격정보.G고가(),
				일일가격정보.G저가(),
				일일가격정보.G종가(),
				일일가격정보.G조정계수(),
				일일가격정보.G조정시가(),
				일일가격정보.G조정고가(),
				일일가격정보.G조정저가(),
				일일가격정보.G조정종가(),
				일일가격정보.G거래량())

			if 에러 != nil {
				log.Println("common.F일일가격정보_모음_DB기록() : stmt추가 실행 에러. "+일일가격정보.G키(), 에러)

				롤백_에러 := 트랜잭션.Rollback()
				if 롤백_에러 != nil {
					log.Println("common.F일일가격정보_모음_DB기록() : 롤백 에러가 발생하였습니다.", 롤백_에러)
				}

				return 에러
			}
		}
	}

	에러 = 트랜잭션.Commit()
	if 에러 != nil {
		log.Println("common.F일일가격정보_모음_DB기록() : 트랜잭션.Commit() 에러가 발생하였습니다.", 에러)

		롤백_에러 := 트랜잭션.Rollback()
		if 롤백_에러 != nil {
			log.Println("common.F일일가격정보_모음_DB기록() : 롤백 에러가 발생하였습니다.", 롤백_에러)
		}

		return 에러
	}

	return nil
}
