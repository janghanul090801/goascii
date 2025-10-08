# goascii

이미지를 ASCII 아트로 변환하는 CLI 도구입니다.

`goascii`는 Go로 작성된 CLI 도구로, 이미지(JPEG, PNG)를 ASCII 아트로 변환하여 HTML 파일로 저장할 수 있습니다.

## 설치

실행 파일을 빌드하려면 다음 명령을 실행하십시오.

```bash
go build -o goascii.exe
```

## 사용법

이미지를 변환하려면 `convert` 명령을 사용하십시오.

```bash
.\goascii.exe convert <이미지-경로> [플래그]
```

### 플래그

- `--resolution`, `-r`: ASCII 아트의 해상도를 설정합니다 (기본값: 60).
- `--output`, `-o`: 출력 HTML 파일 이름을 지정합니다 (기본값: "templated.html").

### 예시

```bash
.\goascii.exe convert ana.png -r 120 -o output.html
```

위 명령어는 `ana.png`를 120의 해상도를 가진 ASCII 아트로 변환하여 `output.html`로 저장합니다.