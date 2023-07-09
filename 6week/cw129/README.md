# Cobra CLI 구조
![image](https://github.com/CW129/go_lang_study/assets/104714337/b987cff7-2942-4c70-9db7-9a7abe97d057)

    Cobra-cli init 명령어를 사용하면 cobra 구조로 디렉토리 생성(주로 cmd 디렉토리 하위에 관련 파일을 작성)
    cobra-cli add 명령어로 기본 구조로 작성되어 있는 .go 파일 생성 가능
    go 프로그램의 entrypoint에서 cmd 패키지를 실행시켜 커맨드를 등록시킴
    계층 형태로 cli를 만들어서 Rootcmd 및으로 커맨드를 생성

## Cobra CLI와 Flag의 차이점
    Cobra CLI : 멀티 커맨드 패턴
                여러가지의 작업을 해야할때 사용, 계층적으로 작업이 분리되어있어 유지보수에 용이
                help, complete 기능을 프레임워크단에서 제공
    flag : 싱글 커맨드 패턴
           단순하게 하나의 작업만 커맨드로 만들때 사용, Cobra-cli에 비해 간단하게 구성할수 있음


              
