# goascii

코딩애플 영상보고 만들었습니다.

## 설치

```bash
go install github.com/janghanul090801/goascii
```

## 빌드

```bash
go build -o goascii.exe
```

## 사용법

```bash
goascii convert <이미지-경로> [플래그]
```

### 플래그

- `--resolution`, `-r`: ASCII 아트의 해상도를 설정 (기본값: 60).
- `--output`, `-o`: 출력 HTML 파일 이름을 지정 (기본값: "templated.html").

### 예시

```bash
goascii convert ana.png -r 120 -o output.html
```

위 명령어는 `ana.png`를 120의 해상도를 가진 ASCII 아트로 변환하여 `output.html`로 저장합니다.