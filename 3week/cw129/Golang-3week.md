# 18장
## 슬라이스 초기화 방법
    {}로 초기화
      var slice1 = []int{1,2,3}
      var slice2 = []int{1, 5:2, 10:3}
    make()로 초기화
      var slice = make([]int,3) # 첫번째 인수 = 타입 두번째 인수 = 길이 * 세번째 인수로 슬라이스의 크기도 지정이 가능
## 슬라이드 동작 원리
<img width="519" alt="image" src="https://github.com/CW129/GOLANG_study/assets/104714337/1879e501-8fc3-4d4a-b64f-1ea6481e76b0">

    슬라이스에 값을 넣는 append를 사용시
    1. Len, Cap의 값으로 배열의 공간을 확인
    2. 빈 공간이 없을 경우 기존 배열을 두배로 확장
    3. 기존 배열의 인덱스를 새로운 배열로 복사
    
    * 주의)
    슬라이스는 한 배열을 포인터로 값을 참조함으로, 슬라이스를 대입하여 사용할 경우 참조하고 있는 배열이 변경될 동작에 신경써야함, 혹은 슬라이스를 복제하여 변경 작업 수행
    
# 19장
## 메서드 선언
<img width="519" alt="image" src="https://github.com/CW129/GOLANG_study/assets/104714337/5525c594-1901-431e-8109-9c1aa86d2165">
    
    리시버는 해당 패키지 안의 모든 type 키워드 선언 타입(구조체, 사용자 정의 타입).(리시버가 있으면 메소드, 없으면 함수)
    * 리시버 = 어떤 구조체를 위한 메소드인지 표현
    * 타입이 같은 패키지 안에 있으면 어디에도 선언이 가능하지만, 해당 타입과 같은 위치에 위치하는걸 권장
    
## 메서드 사용 이유
    코드의 응집도를 향상시키기 위한 용도(객체 지향을 위해)
    기존의 객체지향과의 가장 큰 차이는 상속을 지원하지 않는 점이 객체지향이 아니라는 의견이 있지만
    상속은 객체지향의 설계를 벗어나버리는 문제점을 가지고 있음
    * 리스코프 치환 원칙
    
    Golang은 의존도를 낮추고 응집도를 높이는 부분에 초점
    객체란?
      객체는 데이터, 메소드를 묶음으로 하나의 기능이 되는것
## Pointer, Value 메소드
    Value 메소드 : 구조체의 데이터를 복사하여 전달, 즉 구조체의 데이터가 변경되어도 호출자의 구조체 데이터는 변경되지 않음
    *값 중심의 메소드
    Pointer 메소드 : 구조체의 포인터를 전달, 즉 구조체의 데이터가 변경되면 호출자의 데이터도 변경됨
    *인스턴스 중심의 메소드
    
# 20장
## 인터페이스 선언
<img width="519" alt="image" src="https://github.com/CW129/GOLANG_study/assets/104714337/3d02a1f6-89e9-43ce-8699-16670276451e">

    인터페이스 : 구현을 포함하지 않는 메서드의 집합(추상화된 객체)
    유의사항
      1. 메서드는 반드시 메서드 명이 있어야함
      2. 이름이 같은 메서드는 허용하지 않음
      3. 인터페이스엔 메서드 구현을 포함하지 않음
    * Golang에선 ~er을 붙여 인터페이스를 만드는것을 권장(메소드명~er)  
## 인터페이스 사용 이유
    호출에 구체화된 객체를 사용할 경우 변경 요청이 생길때 해당 객체를 사용하는 모든 부분을 고쳐야함
    추상화된 객체를 사용하는 경우 변경 요청이 있을 경우, 추상 객체만 변경해도 모든 부분에 적용
    개발자의 역할을 분리하기 좋다
    1. 내부 동작의 개발(메소드, 인터페이스)
    2. 비즈니스 로직의 개발(함수, 클래스, API)
    3. 프론트 앱(브라우저 동작)

## 덕타이핑
    Go는 파이썬과 같이 인터페이스에 어떤 타입이 포함되어 있는지 여부를 덕타이핑으로 결정
    * 덕타이핑: 타입 선언시 인터페이스 구현 여부를 명시 할 필요 없이 해당 메소드를 포함하는지 여부로 결정
    
## 3가지 인터페이스 기능
    포함된 인터페이스 : 인터페이스에서 다른 인터페이스 포함
    빈 인터페이스 : 메서드를 가지지 않는 빈 인터페이스(어떤 값이든 받을 수 있는 메서드, 함수, 변수값을 위해 사용)
    인터페이스 기본값 : 인터페이스 변수의 기본겂은 nil (초기화하지 않으면 런타임 에러 발생)
    * 포인터, 인터페이스, 함수 타입, 슬라이스, 맵, 채널은 기본값이 nil 

## 인터페이스 변환


# 21장

## ... 키워드
    가변인수 처리 용도(해당 타입의 인수를 여러개 받음)
    func sum(nums ...int) int{
    }
    함수 내부에서 nums 변수는 해당 타입의 슬라이스 타입으로 처리
    
## defer 지연 실행
    함수가 종료되기 전에 실행되야 하는 코드가 있는 경우 사용
    ex) OS 내부자원을 사용하여 자원을 반납해야하는 경우
<img width="519" alt="image" src="https://github.com/CW129/GOLANG_study/assets/104714337/c0f1576d-81ba-4a38-9712-1479ed54e5fb">

## 함수 타입 변수
    함수의 포인터라고 불림
    프로그램 카운터는 절차에 따라 다음 라인을 실행 (count + 1)
    함수 타입 변수는 프로그램 카운터에게 해당 함수의 시작포인트부터 시작하도록 카운트를 변경
    
## 함수 리터럴
<img width="506" alt="image" src="https://github.com/CW129/GOLANG_study/assets/104714337/f4fa8f53-0cb6-4f3b-80bd-a7b3a6bf6136">    
   
    이름없는 함수, 함수명을 적지 않고 함수 타입 변수값으로 대입되는 함수값을 의미(람다)
    
# 22장

## 리스트
    여러 데이터를 보관하는 자료구조(배열과의 차이점은 리스트는 불연속된 메모리에 데이터를 저장)
    리스트는 각 데이터를 갖고 있는 요소를 포인터로 연결한 자료구조
    각 요소들은 다음, 이전 요소에 접근하는 주소를 가지고있음
    배열과 다르게 중간, 앞에 요소를 추가하기가 더 효율적(배열은 O(n))
    
<img width="506" alt="image" src="https://github.com/CW129/GOLANG_study/assets/104714337/bf7709da-9a1e-47a6-a13a-475784433a9f">
<img width="506" alt="image" src="https://github.com/CW129/GOLANG_study/assets/104714337/a59e8a70-41c1-49f1-bc9b-b877915abff1">
<img width="366" alt="image" src="https://github.com/CW129/GOLANG_study/assets/104714337/565820c5-5447-47b9-82e7-f5aeb90be2a3">

## 링    
    리스트를 기반으로 만들어졌으며, 리스트의 형태에 요소의 맨 앞과 맨 뒤가 연결되어있음
    저장할 개수가 지정되어있고 오래된 요소가 없어져도 되는 경우 사용
    ex) revision
    
## 맵
    해시맵 사용한 자료구조로 키와 값의 형태로 데이터 저장
    
# 23장

## Panic
    Golang 내장 지원 함수로 프로그램의 흐름을 중지시키는 함수로 주로 디버깅에 사용
    panic 함수가 실행되면 에러 메시지를 출력하고 함수의 호출 순서를 나타내는 콜 스택을 출력
    
## Revocer
    panic은 호출 순서를 거꾸로 올라가며 main으로 전달됨
    즉 panic이 발생되기 전에 호출된 함수에 recover() 함수를 이용하면 panic의 에러 메시지는 발생되지만 프로그램 종료는 막을 수 있다
    