# echo-boiler-plate

## 실행 방법 - 테스트 서버
개발 중에는 서버를 실행하는 대신 테스트를 지속적으로 돌려주는 서버를 실행 시켜 테스트를 진행하면서 개발합니다.

### with docker
- [`docker-desktop`](docker.com/products/docker-desktop/)을 설치 합니다.    
- 최초에는 `$ docker-compose up --build` 를 통해서 실행
- 이후에는 `$ docker-compose up`을 통해서 실행합니다.

### without docker
- `go mod download`를 통해서 패키지 설치를 합니다.
- `go install https://github.com/smartystreets/goconvey@latest` 명령어를 통해서 goconvey를 설치
- echo 'export PATH="$PATH:$GOPATH/bin"' >> ~/.zshrc `$GOPATH`는 `$ go env` 를 통해 나온 `path`를 적습니다.

## 테스트 서버 포트 변경
goconvey 포트를 변경하기 위해서는 `.env` 파일의 `PORT` 값을 변경하면 됩니다.

## wire로 dp injection 생성
- wire_gen.sh 를 실행하여서 코드 생성
- 추가적으로 파일 생성시 해당 패키지 명을 뒤에 추가

## code convention
- 파일, 패키지 이름은 `camel case`를 이용합니다.
- 함수, 인터페이스, 구조체는 `pascal case`를 이용합니다.
- 테스트 함수는 `prefix`로 `Test`를 붙입니다.