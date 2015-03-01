package common

import (
	"log"
	"math"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"
)

func F슬라이스_복사(원본slice interface{}) interface{} {
	원본값 := reflect.ValueOf(원본slice)

	// 원본값 검사.
	if 원본값.IsNil() {
		panic("F슬라이스복사 : 원본이 nil임.")
	}
	if !원본값.IsValid() {
		panic("F슬라이스복사 : 원본값이 유효하지 않은 zero값임.")
	}
	if 원본값.Kind() != reflect.Slice {
		panic("F슬라이스복사 : 원본이 slice가 아님.")
	}
	if 원본값.Len() == 0 {
		panic("F슬라이스복사 : 원본 슬라이스 길이가 0임.")
	}

	구성요소형식 := 원본값.Index(0).Type()
	슬라이스형식 := reflect.SliceOf(구성요소형식)
	복사본 := reflect.MakeSlice(슬라이스형식, 원본값.Len(), 원본값.Cap())
	reflect.Copy(복사본, 원본값)

	return 복사본.Interface()
}

const 이벤트전송_최대동시처리갯수 = 100

var 이벤트전송_처리권한 chan int8 = make(chan int8, 이벤트전송_최대동시처리갯수)

// int8형식 이벤트 전송
// 채널을 세마포어처럼 사용해서 최대 동시처리 갯수를 제어하는 패턴
// 이 패턴은 "Effective Go"에서 추천하는 패턴 중의 하나이다.
func F이벤트전송(채널 chan int8, 이벤트 int8) {
	var 단_한번만_실행 sync.Once

	단_한번만_실행.Do(func() {
		for i := 0; i < 이벤트전송_최대동시처리갯수; i++ {
			이벤트전송_처리권한 <- 1
		}
	})

	<-이벤트전송_처리권한

	// 이 펑션은 단 1개의 go루틴만 생성하며,
	// go루틴에 사용되는 변수는 모두 인수로 전달받으면서,
	// 복사본이 생성되었으므로, go루틴의 독립성이 보장되며,
	// 굳이 새로운 변수를 생성할 필요는 없다.

	go func() {
		채널 <- 이벤트
		이벤트전송_처리권한 <- 1
	}()
}

func F반올림(숫자 float64, 소숫점이하_자리수 int) float64 {
	if math.IsNaN(숫자) || math.IsInf(숫자, 0) {
		return 숫자
	}

	부호 := 1.0
	if 숫자 < 0.0 {
		부호 = -1.0
		숫자 *= -1.0
	}

	var 변환2 float64
	승수 := math.Pow(10, float64(소숫점이하_자리수))
	변환1 := 숫자 * 승수
	_, 나머지 := math.Modf(변환1)

	if 나머지 >= 0.5 {
		변환2 = math.Ceil(변환1)
	} else {
		변환2 = math.Floor(변환1)
	}

	return 변환2 / 승수 * 부호
}

func F반올림_통화(금액 float64) float64 {
	return F반올림(금액, 2)
}

func F실수64to문자열(숫자 float64) string {
	return strconv.FormatFloat(숫자, 'G', -1, 64)
	//return strconv.FormatFloat(숫자, 'f', -1, 64)
}

func F부호없는정수64to문자열(숫자 uint64) string {
	return strconv.FormatUint(숫자, 10)
}

func F정수64to문자열(숫자 int64) string {
	return strconv.FormatInt(숫자, 10)
}

func F일자2문자열(일자 time.Time) string {
	return 일자.Format("2006-01-02")
}

func F문자열2실수64(문자열 string) (float64, error) {
	숫자, 에러 := strconv.ParseFloat(strings.Replace(문자열, ",", "", -1), 64)
	if 에러 != nil {
		return 0.0, 에러
	}

	return 숫자, nil
}

func F문자열2정수64(문자열 string) (int64, error) {
	숫자, 에러 := strconv.ParseInt(문자열, 0, 64)
	if 에러 != nil {
		return 0, 에러
	}

	return 숫자, nil
}

func F문자열2부호없는정수64(문자열 string) (uint64, error) {
	숫자, 에러 := strconv.ParseUint(strings.Replace(문자열, ",", "", -1), 0, 64)
	if 에러 != nil {
		return 0, 에러
	}

	return 숫자, nil
}

func F문자열2일자(일자_문자열 string) (time.Time, error) {
	일자, 에러 := time.Parse("2006-01-02", 일자_문자열)
	if 에러 != nil {
		if !F테스트_모드() {
			log.Println("common.F문자열2일자() : time.Parse() 에러.")
			log.Println(에러)
		}

		return 일자, 에러
	}

	return 일자, nil
}

func F일자(시점 time.Time) time.Time {
	if 시점.Hour() == 0 &&
		시점.Minute() == 0 &&
		시점.Second() == 0 &&
		시점.Nanosecond() == 0 {
		return 시점
	}

	일자 := time.Date(시점.Year(), 시점.Month(), 시점.Day(), 0, 0, 0, 0, 시점.Location())

	return 일자
}

var 찾고자_하는_파일경로의_끝부분_소문자 string
var 찾은_파일경로 string

func F파일경로로_파일찾기(파일경로 string) (string, error) {
	찾고자_하는_파일경로의_끝부분_소문자 = strings.ToLower(파일경로)
	검색_시작점 := "."
	찾은_파일경로 = ""

	for {
		에러 := filepath.Walk(검색_시작점, 파일찾기_함수)
		if 에러 != nil {
			log.Println("F파일경로로_파일찾기() : filepath.Walk() 에러.", 에러)

			return "", 에러
		}

		if 찾은_파일경로 != "" {
			return 찾은_파일경로, nil
		}

		// 검색결과가 없으면 한 단계 위에서부터 재검색
		if 검색_시작점 == "." {
			검색_시작점 = "../"
		} else {
			검색_시작점 = "../" + 검색_시작점
		}
	}
}

func 파일찾기_함수(파일경로 string, 파일정보 os.FileInfo, 에러 error) error {
	// 디버깅
	//if 파일정보.IsDir() { log.Printf("파일경로 : %v, 파일이름 : %v", 파일경로, 파일정보.Name()) }

	if 파일정보.IsDir() {
		return nil
	}

	파일경로_소문자1 := strings.ToLower(파일경로)
	파일경로_소문자2 := strings.Replace(파일경로_소문자1, "\\", "/", -1)
	if strings.HasSuffix(파일경로_소문자1, 찾고자_하는_파일경로의_끝부분_소문자) ||
		strings.HasSuffix(파일경로_소문자2, 찾고자_하는_파일경로의_끝부분_소문자) {
		찾은_파일경로 = 파일경로
	}

	return nil
}

func F연월일_문자열(일자 time.Time) (연도 string, 월 string, 일 string) {
	날짜_문자열 := 일자.Format("2006-01-02")
	날짜_문자열_슬라이스 := strings.Split(날짜_문자열, "-")

	연도 = 날짜_문자열_슬라이스[0]
	월 = 날짜_문자열_슬라이스[1]
	일 = 날짜_문자열_슬라이스[2]

	return 연도, 월, 일
}

func F연월일_정수(일자 time.Time) (연도 int, 월 int, 일 int) {
	연도_문자열, 월_문자열, 일_문자열 := F연월일_문자열(일자)

	연도, _ = strconv.Atoi(연도_문자열)
	월, _ = strconv.Atoi(월_문자열)
	일, _ = strconv.Atoi(일_문자열)

	return 연도, 월, 일
}

func F중첩된_외부_슬라이스_제거(중첩된_슬라이스 []interface{}) []interface{} {
	for 반복횟수 := 0; 반복횟수 < 100; 반복횟수++ {
		if 중첩된_슬라이스 == nil { return nil }
		
		if len(중첩된_슬라이스) == 0 ||
			len(중첩된_슬라이스) > 1 {
			// 중첩을 모두 제거한 후 원소가 다수인 슬라이스
			return 중첩된_슬라이스 
		}
		
		// 이제 len(중첩된_슬라이스) == 1 인 경우만 남음.
		원소 := 중첩된_슬라이스[0]
		
		if reflect.TypeOf(원소).Kind() != reflect.Slice {
			// 중첩을 모두 제거한 후 원소가 1개인 슬라이스.
			return 중첩된_슬라이스
		}
		
		// 가장 바깥쪽 중첩 슬라이스 제거.
		중첩된_슬라이스 = 원소.([]interface{})
	}
	
	return nil
}