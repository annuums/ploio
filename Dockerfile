# Build Go Application
FROM golang:alpine AS builder

LABEL email="dev.whoan@gmail.com"
LABEL name="Eugene Minwhoan Kim"
LABEL version="0.0.1"
LABEL description="Solanum - Web Server Framework Based on Gin"

# 작업 디렉토리 생성
WORKDIR /app

# Go 모듈을 복사하고 종속성 다운로드
COPY go.mod go.sum ./
RUN go mod download

# 애플리케이션 소스 코드 복사
COPY . .

# 애플리케이션 빌드
# 이 예제에서는 main.go가 메인 애플리케이션 파일로 가정합니다.
RUN CGO_ENABLED=0 GOOS=linux go build -o annuums

# 두 번째 스테이지: 최종 컨테이너 이미지 빌드
FROM scratch

# 애플리케이션 실행 파일 복사
COPY --from=builder /app/annuums /annuums

# 웹 서버를 8080 포트로 노출
EXPOSE 8080

# 애플리케이션 실행
CMD ["/annuums"]