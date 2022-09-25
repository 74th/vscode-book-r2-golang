# Visual Studio Code 実践ガイド 第 2 版 サンプルコード

Go 言語の API 実装。

## パッケージ

- cmd
  - server: 実行プログラム
- domain: ドメイン
  - entity: ドメイン内外で、データの受け渡しに使う構造体。
  - usecase: ビジネスロジックを格納する。
  - repository: ドメイン層から利用するリソースの定義。Database が定義されている。
- memdb: インメモリデータベース。domain/repository の定義を実装する。
- server: Web サーバ。
