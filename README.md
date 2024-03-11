# GolangCodes
Golang Code snippets Project


# GO111MODULE -> on / off 테스트 on 일 때 GOPATH/pkg 폴더 사용함. 

off - 빌드 중에 $GOPATH에 있는 패키지를 사용합니다.
on - 빌드 중에 $GOPATH 대신 모듈에 있는 패키지를 사용합니다.
auto - 현재 디렉터리가 $GOPATH 외부에 있고 go.mod 파일이 포함된 경우 모듈을 사용하고 그렇지 않으면 $GOPATH의 패키지를 사용합니다.