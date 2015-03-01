package common

import (
	"errors"
	"fmt"
	"log"
	"math/big"
	"reflect"
	"time"
)

type S미정 struct{} // 설계 중에 아직 자료구조를 결정하지 못했을 때 기록용도로 사용.

type S정확한_실수 struct { 값 *big.Rat }
func (s *S정확한_실수) G값() float64 {
	실수값, _ := s.값.Float64()
	
	return 실수값
}
func (s *S정확한_실수) G정확한_값() *big.Rat {
	반환값 := new(big.Rat)
	반환값.Set(s.값)
	
	return 반환값
}
func (s *S정확한_실수) G반올림값(소숫점_이하_자릿수 int) float64 {
	반올림값 := new(S정확한_실수)
	반올림값.S정확한_값(s)
	반올림값.S반올림(소숫점_이하_자릿수)
	
	return 반올림값.G값()
}
func (s *S정확한_실수) G문자열(소숫점_이하_자릿수 int) string {
	// 반올림한 값을 문자열 형태로 구함. 반올림 할 떄 유용함.
	return s.값.FloatString(소숫점_이하_자릿수)
}
func (s *S정확한_실수) S값(값 float64) { s.값.SetFloat64(값) }
func (s *S정확한_실수) S정확한_값(값 *S정확한_실수) { s.값.Set(값.G정확한_값()) }
func (s *S정확한_실수) S반올림(소숫점_이하_자릿수 int) {
	문자열 := s.G문자열(소숫점_이하_자릿수)	// 반올림값을 문자열 형태로 구함.
	s.값.SetString(문자열)
}
func F정확한_실수_구조체_생성(값 float64) *S정확한_실수 {
	s := new(S정확한_실수)
	s.S값(값)
	
	return s
}

type S고정소숫점 struct {
	소숫점_이하_자릿수 int
	값 *big.Rat
}
func (s *S고정소숫점) G소숫점_이하_자릿수() int { return s.소숫점_이하_자릿수 }
func (s *S고정소숫점) G값() *big.Rat {
	// 원본 변경을 방지하기 위해서 복사한 후 넘겨준다.
	복사본 := new(big.Rat)
	복사본.Set(s.값)
	
	return 복사본
}
func (s *S고정소숫점) G실수값() float64 {
	실수값, _ := s.값.Float64()
	
	return 실수값
}
func (s *S고정소숫점) G문자열() string { return s.값.FloatString(s.소숫점_이하_자릿수) }
func (s *S고정소숫점) String() string { return s.G문자열() }

func (s *S고정소숫점) S소숫점_이하_자릿수(소숫점_이하_자릿수 int) {
	s.소숫점_이하_자릿수 = 소숫점_이하_자릿수
	s.S값(s.G값())	// 바뀐 정밀도에 값의 정밀도를 재조정.
}
func (s *S고정소숫점) S값(값 *big.Rat) {
	// 반올림 처리.
	값_문자열 := 값.FloatString(s.소숫점_이하_자릿수)
	
	// 원본과의 독립성 유지를 위해서 새로 생성한 후 값만 같게 한다.
	s.값 = new(big.Rat)
	s.값.SetString(값_문자열)
}
func (s *S고정소숫점) S실수값(실수값 float64) {
	// 반올림한 결과를 문자열 형태로 얻어서 값을 설정함.
	반올림_이전값 := new(big.Rat)
	반올림_이전값.SetFloat64(실수값)
	
	s.S값(반올림_이전값)	// S값()에서 반올림한다.
}

func F고정소숫점_구조체_생성(값 float64, 소숫점_이하_자릿수 int) *S고정소숫점 {	
	// 반올림한 결과를 문자열 형태로 얻어서 값을 설정함.
	임시값 := new(big.Rat)
	임시값.SetFloat64(값)
	최종값_문자열 := 임시값.FloatString(소숫점_이하_자릿수)
	
	s := new(S고정소숫점)
	s.소숫점_이하_자릿수 = 소숫점_이하_자릿수
	s.값 = new(big.Rat)
	s.값.SetString(최종값_문자열)
	
	return s
}

type S시점별실수값 struct {
	시점 time.Time
	값  float64
}

func (s *S시점별실수값) G시점() time.Time   { return s.시점 }
func (s *S시점별실수값) G값() float64      { return s.값 }
func (s *S시점별실수값) S시점(시점 time.Time) { s.시점 = 시점 }
func (s *S시점별실수값) S값(실수값 float64)   { s.값 = 실수값 }
func F시점별실수값생성(시점 time.Time, 실수값 float64) *S시점별실수값 {
	s := new(S시점별실수값)
	s.시점 = 시점
	s.값 = 실수값

	return s
}

type S에러내역 struct {
	에러코드 string
	에러설명 string
	발생횟수 int64
}

func (s *S에러내역) G에러코드() string { return s.에러코드 }
func (s *S에러내역) G에러설명() string { return s.에러설명 }
func (s *S에러내역) G발생횟수() int64  { return s.발생횟수 }
func (s *S에러내역) S발생횟수_증가()     { s.발생횟수++ }
func F에러내역_생성(에러코드 string, 에러설명 string) *S에러내역 {
	s := new(S에러내역)
	s.에러코드 = 에러코드
	s.에러설명 = 에러설명
	s.발생횟수 = 1

	return s
}

type S식별코드별_에러내역_맵 struct {
	에러내역_맵 map[uint64][]*S에러내역
}

func (s *S식별코드별_에러내역_맵) 초기화() {
	if s.에러내역_맵 == nil {
		s.에러내역_맵 = make(map[uint64][]*S에러내역)
	}
}
func (s *S식별코드별_에러내역_맵) G에러난_식별코드_수량() int {
	return len(s.에러내역_맵)
}
func (s *S식별코드별_에러내역_맵) G맵() map[uint64][]*S에러내역 {
	if s.에러내역_맵 == nil {
		s.초기화()
	}

	// 외부에서 변경하는 것을 막기 위해서 복사본을 전달한다.
	에러내역_맵 := make(map[uint64][]*S에러내역)

	for 키, 에러내역_슬라이스 := range s.에러내역_맵 {
		에러내역_맵[키] = 에러내역_슬라이스
	}

	return 에러내역_맵
}
func (s *S식별코드별_에러내역_맵) G에러내역_모음(식별코드 uint64) []*S에러내역 {
	if s.에러내역_맵 == nil {
		s.초기화()
	}

	// 외부에서 변경하는 것을 막기 위해서 복사본을 전달한다.
	에러내역_슬라이스 := s.에러내역_맵[식별코드]

	if 에러내역_슬라이스 == nil {
		return nil
	}

	return F슬라이스_복사(에러내역_슬라이스).([]*S에러내역)
}
func (s *S식별코드별_에러내역_맵) S에러내역_추가(식별코드 uint64, 에러내역 *S에러내역) {
	에러내역_모음 := s.G에러내역_모음(식별코드)

	if 에러내역_모음 == nil {
		에러내역_모음 = make([]*S에러내역, 0)
	}

	for _, 기존_에러내역 := range 에러내역_모음 {
		if 기존_에러내역.G에러코드() == 에러내역.G에러코드() &&
			기존_에러내역.G에러설명() == 에러내역.G에러설명() {
			기존_에러내역.S발생횟수_증가()

			return
		}
	}

	에러내역_모음 = append(에러내역_모음, 에러내역)
	s.에러내역_맵[식별코드] = 에러내역_모음
}
func (s *S식별코드별_에러내역_맵) F에러내역_출력() {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("----------------------------------------------")
	fmt.Println("식별코드별 에러내역 출력")
	fmt.Println("----------------------------------------------")
	fmt.Println("")
	for 식별코드, 에러내역_슬라이스 := range s.에러내역_맵 {
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("----------------------------------------------")
		fmt.Printf("식별코드 %v", 식별코드)
		fmt.Println("----------------------------------------------")
		fmt.Println("")
		for 인덱스, 에러내역 := range 에러내역_슬라이스 {
			fmt.Printf("%v %v, %v (%v회 발생)", 인덱스, 에러내역.G에러코드(), 에러내역.G에러설명(), 에러내역.G발생횟수())
		}
		fmt.Println("")
		fmt.Println("")
	}
}

type S종목코드별_에러내역_맵 struct {
	에러내역_맵 map[string][]*S에러내역
}

func (s *S종목코드별_에러내역_맵) 초기화() {
	if s.에러내역_맵 == nil {
		s.에러내역_맵 = make(map[string][]*S에러내역)
	}
}
func (s *S종목코드별_에러내역_맵) G에러난_종목코드_수량() int {
	return len(s.에러내역_맵)
}
func (s *S종목코드별_에러내역_맵) G맵() map[string][]*S에러내역 {
	if s.에러내역_맵 == nil {
		s.초기화()
	}

	// 외부에서 변경하는 것을 막기 위해서 복사본을 전달한다.
	에러내역_맵 := make(map[string][]*S에러내역)

	for 키, 에러내역_슬라이스 := range s.에러내역_맵 {
		에러내역_맵[키] = 에러내역_슬라이스
	}

	return 에러내역_맵
}
func (s *S종목코드별_에러내역_맵) G에러내역_모음(종목코드 string) []*S에러내역 {
	if s.에러내역_맵 == nil {
		s.초기화()
	}

	// 외부에서 변경하는 것을 막기 위해서 복사본을 전달한다.
	에러내역_슬라이스 := s.에러내역_맵[종목코드]

	if 에러내역_슬라이스 == nil {
		return nil
	}

	return F슬라이스_복사(에러내역_슬라이스).([]*S에러내역)
}
func (s *S종목코드별_에러내역_맵) S에러내역_추가(종목코드 string, 에러내역 *S에러내역) {
	에러내역_모음 := s.G에러내역_모음(종목코드)

	if 에러내역_모음 == nil {
		에러내역_모음 = make([]*S에러내역, 0)
	}

	for _, 기존_에러내역 := range 에러내역_모음 {
		if 기존_에러내역.G에러코드() == 에러내역.G에러코드() &&
			기존_에러내역.G에러설명() == 에러내역.G에러설명() {
			기존_에러내역.S발생횟수_증가()

			return
		}
	}

	에러내역_모음 = append(에러내역_모음, 에러내역)
	s.에러내역_맵[종목코드] = 에러내역_모음
}
func (s *S종목코드별_에러내역_맵) F에러내역_출력() {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("----------------------------------------------")
	fmt.Println("종목코드별 에러내역 출력")
	fmt.Println("----------------------------------------------")
	fmt.Println("")
	for 종목코드, 에러내역_슬라이스 := range s.에러내역_맵 {
		종목 := F종목_검색(종목코드)
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("----------------------------------------------")
		fmt.Printf("종목코드 %v, 종목명칭 %v", 종목코드, 종목.G종목명칭())
		fmt.Println("----------------------------------------------")
		fmt.Println("")
		for 인덱스, 에러내역 := range 에러내역_슬라이스 {
			fmt.Printf("%v %v, %v (%v회 발생)", 인덱스, 에러내역.G에러코드(), 에러내역.G에러설명(), 에러내역.G발생횟수())
		}
		fmt.Println("")
		fmt.Println("")
	}
}

func F가변형_생성(값 interface{}) *S가변형 {
	var 가변형_변수 *S가변형
	
	switch 값.(type) {
	case S가변형:
		가변형 := 값.(S가변형)
		가변형_변수 = &가변형
	case *S가변형:
		가변형_변수 = 값.(*S가변형)
	default:
		가변형_변수 = new(S가변형)
		가변형_변수.값 = 값
	}
	
	return 가변형_변수
}
type S가변형 struct {
	값 interface{}
}
func (s *S가변형) G값() interface{} { return s.값 }
func (s *S가변형) G종류() reflect.Kind {
	if s.값 == nil {
		return reflect.Invalid
	} else {
		return reflect.TypeOf(s.값).Kind()
	}
}
func (s *S가변형) G형식() string {
	if s.값 == nil {
		return "nil"
	} else {
		return reflect.TypeOf(s.값).String()
	}
}
func (s *S가변형) G상수형임() bool {
	종류 := s.G종류()
	형식 := s.G형식()
	
	if 종류 == reflect.Bool ||
		종류 == reflect.Int ||
		종류 == reflect.Int8 ||
		종류 == reflect.Int16 ||
		종류 == reflect.Int32 ||
		종류 == reflect.Int64 ||
		종류 == reflect.Uint ||
		종류 == reflect.Uint8 ||
		종류 == reflect.Uint16 ||
		종류 == reflect.Uint32 ||
		종류 == reflect.Uint64 ||
		종류 == reflect.Float32 ||
		종류 == reflect.Float64 ||
		종류 == reflect.String ||
		(종류 == reflect.Struct && 형식 == "time.Time") {
		return true
	} else {
		return false
	}
		
	// 차후 추가하는 것을 고려해 볼 항목.	
    //Complex64
    //Complex128
}
func (s *S가변형) G참거짓() (bool, error) {
	if s.G종류() != reflect.Bool {
		에러 := errors.New("common.S가변형.G참거짓() : 에러 발생. 종류 " + s.G종류().String() + ", 형식 " + s.G형식())
		
		return false, 에러
	}
	
	return s.값.(bool), nil
}
func (s *S가변형) G부호없는_정수() (uint, error) {
	if s.G종류() != reflect.Uint {
		에러 := errors.New("common.S가변형.G부호없는_정수() : 에러 발생. 종류 " + s.G종류().String() + ", 형식 " + s.G형식())
		
		return 0, 에러
	}
	
	return s.값.(uint), nil
}
func (s *S가변형) G부호없는_정수8() (uint8, error) {
	if s.G종류() != reflect.Uint8 {
		에러 := errors.New("common.S가변형.G부호없는_정수8() : 에러 발생. 종류 " + s.G종류().String() + ", 형식 " + s.G형식())
		
		return 0, 에러
	}
	
	return s.값.(uint8), nil
}
func (s *S가변형) G부호없는_정수16() (uint16, error) {
	if s.G종류() != reflect.Uint16 {
		에러 := errors.New("common.S가변형.G부호없는_정수16() : 에러 발생. 종류 " + s.G종류().String() + ", 형식 " + s.G형식())
		
		return 0, 에러
	}
	
	return s.값.(uint16), nil
}
func (s *S가변형) G부호없는_정수32() (uint32, error) {
	if s.G종류() != reflect.Uint32 {
		에러 := errors.New("common.S가변형.G부호없는_정수32() : 에러 발생. 종류 " + s.G종류().String() + ", 형식 " + s.G형식())
		
		return 0, 에러
	}
	
	return s.값.(uint32), nil
}
func (s *S가변형) G부호없는_정수64() (uint64, error) {
	if s.G종류() != reflect.Uint64 {
		에러 := errors.New("common.S가변형.G부호없는_정수64() : 에러 발생. 종류 " + s.G종류().String() + ", 형식 " + s.G형식())
		
		return 0, 에러
	}
	
	return s.값.(uint64), nil
}
func (s *S가변형) G정수() (int, error) {
	if s.G종류() != reflect.Int {
		에러 := errors.New("common.S가변형.G정수() : 에러 발생. 종류 " + s.G종류().String() + ", 형식 " + s.G형식())
		
		return 0, 에러
	}
	
	return s.값.(int), nil
}
func (s *S가변형) G정수8() (int8, error) {
	if s.G종류() != reflect.Int8 {
		에러 := errors.New("common.S가변형.G정수8() : 에러 발생. 종류 " + s.G종류().String() + ", 형식 " + s.G형식())
		
		return 0, 에러
	}
	
	return s.값.(int8), nil
}
func (s *S가변형) G정수16() (int16, error) {
	if s.G종류() != reflect.Int16 {
		에러 := errors.New("common.S가변형.G정수16() : 에러 발생. 종류 " + s.G종류().String() + ", 형식 " + s.G형식())
		
		return 0, 에러
	}
	
	return s.값.(int16), nil
}
func (s *S가변형) G정수32() (int32, error) {
	if s.G종류() != reflect.Int32 {
		에러 := errors.New("common.S가변형.G정수32() : 에러 발생. 종류 " + s.G종류().String() + ", 형식 " + s.G형식())
		
		return 0, 에러
	}
	
	return s.값.(int32), nil
}
func (s *S가변형) G정수64() (int64, error) {
	if s.G종류() != reflect.Int64 {
		에러 := errors.New("common.S가변형.G정수64() : 에러 발생. 종류 " + s.G종류().String() + ", 형식 " + s.G형식())
		
		return 0, 에러
	}
	
	return s.값.(int64), nil
}
func (s *S가변형) G실수() (float64, error) {
	if s.G종류() != reflect.Float64 {
		에러 := errors.New("common.S가변형.G실수() : 에러 발생. 종류 " + s.G종류().String() + ", 형식 " + s.G형식())
		
		return 0.0, 에러
	}
	
	return s.값.(float64), nil
}
func (s *S가변형) G실수32() (float32, error) {
	if s.G종류() != reflect.Float32 {
		에러 := errors.New("common.S가변형.G실수32() : 에러 발생. 종류 " + s.G종류().String() + ", 형식 " + s.G형식())
		
		return 0.0, 에러
	}
	
	return s.값.(float32), nil
}
func (s *S가변형) G실수64() (float64, error) {
	if s.G종류() != reflect.Float64 {
		에러 := errors.New("common.S가변형.G실수64() : 에러 발생. 종류 " + s.G종류().String() + ", 형식 " + s.G형식())
		
		return 0.0, 에러
	}
	
	return s.값.(float64), nil
}
func (s *S가변형) G문자열() (string, error) {
if s.G종류() != reflect.String {
		에러 := errors.New("common.S가변형.G문자열() : 에러 발생. 종류 " + s.G종류().String() + ", 형식 " + s.G형식())
		
		return "", 에러
	}
	
	return s.값.(string), nil
}
func (s *S가변형) G시점() (time.Time, error) {
	if s.G종류() != reflect.Struct ||
		s.G형식() != "time.Time" {
		에러 := errors.New("common.S가변형.G시점() : 에러 발생. 종류 " + s.G종류().String() + ", 형식 " + s.G형식())
		
		에러_일자, _ := F문자열2일자("1900-01-01")
		return 에러_일자, 에러
	}
	
	return s.값.(time.Time), nil
}

type S매개변수 struct {
	이름 string
	가변형 *S가변형
}
func (s *S매개변수) G이름() string { return s.이름 }
func (s *S매개변수) G값() interface{} { return s.가변형.G값() }
func (s *S매개변수) G가변형() *S가변형 { return s.가변형 }
func F매개변수_생성(이름 string, 값 interface{}) *S매개변수 {	
	매개변수 := new(S매개변수)
	매개변수.이름 = 이름
	매개변수.가변형 = F가변형_생성(값)
	
	return 매개변수
}

type S매개변수_모음 struct {
	매개변수_모음 map[string]*S가변형
}
func (s *S매개변수_모음) 초기화() {
	if s.매개변수_모음 == nil {
		s.매개변수_모음 = make(map[string]*S가변형)
	}
}
func (s *S매개변수_모음) G수량() int { return len(s.매개변수_모음) }
func (s *S매개변수_모음) G이미_있음(이름 string) bool {
	if s.매개변수_모음 == nil { s.초기화() }
	
	_, 이미_있음 := s.매개변수_모음[이름]
	
	return 이미_있음
}
func (s *S매개변수_모음) G값(이름 string) interface{} {
	if s.매개변수_모음 == nil { s.초기화() }
	
	가변형_변수, _ := s.매개변수_모음[이름]
	
	if 가변형_변수 == nil {
		return nil
	} else {
		return 가변형_변수.G값()
	}
}
func (s *S매개변수_모음) G가변형(이름 string) *S가변형 {
	if s.매개변수_모음 == nil { s.초기화() }
	
	가변형, _ := s.매개변수_모음[이름]
	
	return 가변형
}
func (s *S매개변수_모음) G매개변수_모음() []*S매개변수 {
	if s.매개변수_모음 == nil { s.초기화() }
	
	매개변수_모음 := make([]*S매개변수, 0)
	
	for 이름, 가변형_변수 := range s.매개변수_모음 {
		매개변수 := F매개변수_생성(이름, 가변형_변수)
		매개변수_모음 = append(매개변수_모음, 매개변수)
	}
	
	return 매개변수_모음
}
func (s *S매개변수_모음) G복제() *S매개변수_모음 {
	if s.매개변수_모음 == nil { s.초기화() }
	
	복제 := new(S매개변수_모음)
	
	for 이름, 값 := range s.매개변수_모음 {
		복제.S추가(이름, 값)
	}
	
	return 복제
}
func (s *S매개변수_모음) S추가(이름 string, 매개변수 interface{}) {
	if s.매개변수_모음 == nil { s.초기화() }
	
	if s.G이미_있음(이름) {
		// 같은 이름의 매개변수가 이미 존재함.
		// 새로운 입력값으로 대체하면, 전달된 매개변수에 의도하지 않은 변경이 가해질 가능성이 높아짐.
		// 꼭 변경해야 하는 경우가 발생하면, 별도의 S대체()를 구현하는 것도 고려해 볼 것.
		log.Printf("common.S매개변수_모음 : S추가(). 같은 이름의 매개변수가 이미 존재함. 이름 %v", 이름)
		
		return
	}
	
	s.매개변수_모음[이름] = F가변형_생성(매개변수)
}