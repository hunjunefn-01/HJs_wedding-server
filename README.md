# 모바일 청첩장 서버

이 프로젝트는 [모바일 청첩장](https://github.com/juhonamnam/wedding-invitation) 웹 애플리케이션의 백엔드 서버입니다. 모바일 청첩장에 필요한 API 엔드포인트와 데이터베이스 관리 기능을 제공합니다. 모바일 청첩장에 필요한 간단한 기능만 구현하였으며, 트래픽이 많지 않은 환경이기에 SQLite를 사용합니다.

## 사전 요구사항

- Python 3.10+

## 제공 기능

- 방명록 작성 및 조회 API
  - 관리자 비밀번호를 통한 방명록 강제 삭제 기능
- 참석 의사 전달 API
  - 참석자 조회 기능은 현재 미구현 상태

## 시작하기

1. 가상환경 생성 및 활성화:

   ```bash
   python -m venv .venv
   .venv\Scripts\activate  # Windows
   source .venv/bin/activate  # macOS/Linux
   ```

2. 의존성 설치:

   ```bash
   pip install -r requirements.txt
   ```

3. 환경변수 설정 (`.env` 파일 생성):

   - `ALLOW_ORIGIN`
     - 허용할 도메인 (예: `http://localhost:3000`)
   - `ADMIN_PASSWORD`
     - 관리자 전용 비밀번호
     - 방명록 강제 삭제를 원하는 경우 해당 비밀번호로 삭제 가능

4. 서버 실행:
   ```bash
   python app.py
   ```

   서버가 기본적으로 `http://localhost:8080`에서 실행됩니다. 첫 실행 시 `sql.db` (SQLite) 파일과 테이블이 자동으로 생성됩니다.
