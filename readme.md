# Go API 100本ノック（コメント＋レスポンス例付き）

## 1〜10: 基礎・HTTP/JSON 入門
1. **GET /ping** — 動作確認用の最小API。例: `{"message":"pong"}`
2. **GET /hello?name=Riku** — クエリパラメータの扱い練習。例: `{"message":"Hello, Riku"}`
3. **POST /echo** — JSON を受け取りそのまま返す。例: リクエスト `{"message":"hi"}` → レスポンス `{"message":"hi"}`
4. **GET /time** — サーバー時刻を返す。例: `{"time":"2025-11-24T12:34:56Z"}`
5. **GET /env** — 環境変数の読み取り練習。例: `{"app_env":"dev"}`
6. **GET /headers** — リクエストヘッダを JSON で返す。例: `{"User-Agent":"curl/8.0.0","Accept":"*/*"}`
7. **GET /status/:code** — 任意ステータスコードとメッセージを返す。例: `/status/404` → 404 + `{"status":404,"message":"not found"}`
8. **GET /random/int?min=1&max=100** — ランダム整数生成。例: `{"value":42}`
9. **GET /fib/:n** — n 番目のフィボナッチ数を返す。例: `/fib/10` → `{"n":10,"value":55}`
10. **GET /health** — 基本のヘルスチェック。例: `{"status":"ok"}`

## 11〜25: CRUD（Todo API）
11. **POST /todos** — Todoの作成。例: リクエスト `{"title":"buy milk"}` → レスポンス `{"id":1,"title":"buy milk","completed":false}`
12. **GET /todos** — 一覧取得。例: `[{"id":1,"title":"buy milk","completed":false}]`
13. **GET /todos/:id** — 個別取得。例: `/todos/1` → `{"id":1,"title":"buy milk","completed":false}`
14. **PUT /todos/:id** — 全体更新。例: `{"title":"buy tea","completed":true}` → 更新後 Todo を返す。
15. **PATCH /todos/:id** — 部分更新。例: リクエスト `{"completed":true}` → レスポンス `{"id":1,"title":"buy milk","completed":true}`
16. **DELETE /todos/:id** — 1件削除。例: 204 No Content または `{"deleted":1}`
17. **DELETE /todos?force=true** — 全削除。例: `{"deleted":10}`
18. **GET /todos?completed=true** — 完了ステータスフィルタ。例: `[{"id":2,"title":"done task","completed":true}]`
19. **POST /todos/:id/complete** — 完了処理。例: `{"id":1,"completed":true}`
20. **POST /todos/:id/uncomplete** — 未完了化。例: `{"id":1,"completed":false}`
21. **GET /todos/stats** — 集計情報。例: `{"total":10,"completed":4,"remaining":6}`
22. **POST /todos/bulk** — 複数登録。例: リクエスト `[{"title":"a"},{"title":"b"}]` → レスポンス 登録済み Todo 配列
23. **GET /todos/search?q=milk** — 部分一致検索。例: `[{"id":1,"title":"buy milk","completed":false}]`
24. **GET /todos/recent?limit=5** — 新しい順の取得。例: `[{"id":10,"title":"latest","created_at":"..."}, ...]`
25. **GET /todos/overdue** — 期限切れ取得。例: `[{"id":3,"title":"expired","due_date":"2025-11-01"}]`

## 26〜35: ページング・ソート・リレーション
26. **GET /todos?page=1&per_page=10** — ページング。例: `{"items":[...],"page":1,"per_page":10,"total":53}`
27. **GET /todos?sort=created_at&order=desc** — ソート。例: `{"items":[{"id":10,...},{"id":9,...}]}`
28. **GET /users** — ユーザ一覧。例: `[{"id":1,"name":"Alice"},{"id":2,"name":"Bob"}]`
29. **GET /users/:id/todos** — リレーション1:n。例: `/users/1/todos` → ユーザ1のTodo配列
30. **POST /users/:id/todos** — ユーザに紐づく Todo 作成。例: `{"id":5,"user_id":1,"title":"user todo"}`
31. **GET /users/:id/todos/stats** — ユーザ別集計。例: `{"user_id":1,"total":7,"completed":3}`
32. **GET /todos?user_id=1** — クエリフィルタ。例: `[{"id":5,"user_id":1,"title":"user todo"}]`
33. **GET /users/:id** — ユーザ詳細。例: `{"id":1,"name":"Alice","todo_count":7}`
34. **GET /users/with_more_than?n=10** — Todo が多いユーザ。例: `[{"id":2,"name":"PowerUser","todo_count":15}]`
35. **GET /feed** — タイムライン風。例: `[{"type":"todo","user":"Alice","title":"buy milk","created_at":"..."}, ...]`

## 36〜45: 認証・認可
36. **POST /auth/signup** — 新規アカウント登録。例: `{"id":1,"email":"a@example.com"}`
37. **POST /auth/login** — ログインしトークン発行。例: `{"access_token":"xxx","token_type":"bearer"}`
38. **GET /me** — 認証済ユーザ情報。例: `{"id":1,"email":"a@example.com","name":"Alice"}`
39. **GET /me/todos** — 自分のTodo取得。例: `[{"id":1,"user_id":1,"title":"my task"}]`
40. **POST /me/todos** — 自分のTodo作成。例: `{"id":2,"user_id":1,"title":"my new todo"}`
41. **PUT /me/profile** — プロフィール更新。例: リクエスト `{"name":"New Name"}` → `{"id":1,"name":"New Name"}`
42. **GET /admin/users** — 管理者専用API。例: `[{"id":1,"role":"admin"},{"id":2,"role":"user"}]`
43. **POST /auth/logout** — トークン無効化。例: `{"message":"logged out"}`
44. **POST /auth/refresh** — トークン更新。例: `{"access_token":"new-token","token_type":"bearer"}`
45. **GET /auth/check** — トークン有効性確認。例: `{"valid":true}`

## 46〜55: 外部サービス連携（モック可）
46. **GET /external/time** — 外部API呼び出し風。例: `{"source":"worldtimeapi","datetime":"..."}` 
47. **POST /notifications/email** — メール送信モック。例: `{"to":"user@example.com","status":"queued"}`
48. **POST /notifications/slack** — Slack通知風API。例: `{"channel":"#general","status":"sent"}`
49. **GET /weather?city=Tokyo** — 気象API風レスポンス。例: `{"city":"Tokyo","temp":24.5,"condition":"sunny"}`
50. **POST /payments/charge** — 決済処理モック。例: `{"payment_id":"pay_123","status":"succeeded"}`
51. **GET /payments/:id** — 決済ステータス。例: `{"id":"pay_123","amount":1200,"status":"succeeded"}`
52. **POST /webhooks/payment** — 決済Webhook受信。例: `{"received":true}`
53. **POST /webhooks/github** — GitHub webhook練習。例: `{"event":"push","received":true}`
54. **GET /integrations/status** — 外部連携の疎通チェック。例: `{"payment_api":"ok","slack":"ok"}`
55. **GET /metrics/external-calls** — 外部API呼び出し回数。例: `{"payment_api_calls":32,"weather_api_calls":10}`

## 56〜65: ファイル・画像・ダウンロード
56. **POST /files** — ファイルアップロード。例: レスポンス `{"id":"file_abc","filename":"a.txt","size":1234}`
57. **GET /files/:id** — ファイル取得。例: バイナリ or `Content-Disposition`付きレスポンス。
58. **DELETE /files/:id** — 削除。例: `{"deleted":"file_abc"}`
59. **GET /files** — 全ファイル一覧。例: `[{"id":"file_abc","filename":"a.txt"}]`
60. **POST /avatars** — 画像アップロード。例: `{"user_id":1,"avatar_id":"avt_1"}`
61. **GET /avatars/:user_id** — アバター取得。例: 画像バイナリ or デフォルト画像。
62. **GET /reports/todos.csv** — CSVダウンロード。例: `id,title,completed\n1,"buy milk",false`
63. **GET /reports/todos.zip** — ZIPダウンロード。例: `Content-Type: application/zip` でバイナリ。
64. **POST /files/text** — テキスト保存。例: リクエスト `"hello"` → `{"id":"txt_1"}`
65. **GET /files/text/:id** — テキストファイル取得。例: レスポンスボディに `"hello"`。

## 66〜75: 非同期・ジョブ
66. **POST /jobs/long-task** — 長時間タスク開始。例: `{"job_id":"job_1","status":"queued"}`
67. **GET /jobs/:id** — ジョブ状態確認。例: `{"job_id":"job_1","status":"running"}`
68. **GET /jobs/:id/result** — 結果取得。例: `{"job_id":"job_1","status":"done","result":{"count":123}}`
69. **POST /jobs/bulk-email** — 大量メール送信ジョブ。例: `{"job_id":"bulk_1","status":"queued"}`
70. **GET /jobs** — ジョブ一覧。例: `[{"job_id":"job_1","status":"done"}]`
71. **POST /reindex/todos** — 再インデックス。例: `{"job_id":"reindex_1","status":"queued"}`
72. **POST /cleanup/old-todos** — 古いTodo削除。例: `{"deleted":5,"job_id":"cleanup_1"}`
73. **POST /simulate/timeout** — タイムアウト実験。例: 一定秒数待ってから `{"status":"done"}` を返す。
74. **POST /simulate/error** — 強制エラー。例: 常に 500 + `{"error":"simulated error"}`
75. **GET /simulate/flaky** — 不安定API。例: たまに 200 `{"status":"ok"}`, たまに 500。

## 76〜85: バリデーション・エラー設計
76. **POST /validation/users** — 形式チェック。例: エラー時 422 + `{"errors":{"email":"invalid format"}}`
77. **POST /validation/todos** — 項目バリデーション。例: タイトル未入力で `{"errors":{"title":"required"}}`
78. **GET /validation/query?age=abc** — 不正クエリ検証。例: 400 + `{"error":"age must be integer"}`
79. **GET /validation/date-range?from=...&to=...** — 日付範囲チェック。例: 不正時 400 + `{"error":"from must be before to"}`
80. **POST /validation/nested** — ネスト構造の検証。例: `{"errors":{"profile.name":"required"}}`
81. **GET /errors/example** — エラーレスポンス例。例: 500 + `{"error":"internal server error","request_id":"..."}`
82. **GET /errors/not-found-sample** — 404 の練習。例: 404 + `{"error":"resource not found"}`
83. **GET /errors/unauthorized-sample** — 401 の練習。例: 401 + `{"error":"unauthorized"}`
84. **GET /errors/forbidden-sample** — 403 の練習。例: 403 + `{"error":"forbidden"}`
85. **GET /errors/rate-limit-sample** — 429 の練習。例: 429 + `{"error":"too many requests","retry_after":30}`

## 86〜92: 運用・メトリクス
86. **GET /metrics** — 基本メトリクス。例: `{"requests_total":1024,"uptime_seconds":3600}`
87. **GET /metrics/todos** — Todo系メトリクス。例: `{"todos_total":50,"completed_ratio":0.4}`
88. **GET /logs/recent** — 最近のログ取得。例: `{"logs":["2025-11-24T... INFO started","..."]}`
89. **GET /debug/config** — 設定値デバッグ。例: `{"app_env":"dev","debug":true}`
90. **GET /debug/routes** — ルート一覧。例: `[{"method":"GET","path":"/ping"},{"method":"POST","path":"/todos"}]`
91. **GET /health/deep** — 深いヘルスチェック。例: `{"status":"ok","db":"ok","redis":"ok","external_api":"degraded"}`
92. **GET /version** — バージョン情報。例: `{"version":"1.2.3","commit":"abc123","build_date":"2025-11-24"}`

## 93〜100: ミニプロダクト系
93. **POST /short-urls** — URL短縮。例: リクエスト `{"url":"https://example.com"}` → `{"code":"abc123","short_url":"https://sho.rt/abc123"}`
94. **GET /s/:code** — 302 リダイレクト。例: `/s/abc123` → 302 + Location: `https://example.com`
95. **GET /short-urls/:code/stats** — アクセス数。例: `{"code":"abc123","clicks":42}`
96. **POST /chat/messages** — チャット投稿。例: `{"id":"msg_1","room_id":"room1","text":"hello"}`
97. **GET /chat/messages?room_id=room1** — チャットログ取得。例: `[{"id":"msg_1","room_id":"room1","text":"hello"}]`
98. **GET /ranking/todos/users** — ユーザTodo数のランキング。例: `[{"user_id":1,"name":"Alice","todo_count":20}]`
99. **GET /recommend/todos** — おすすめTodo生成。例: `[{"id":7,"title":"review PR","reason":"high priority"}]`
100. **GET /export/all** — 全データのエクスポート。例: `{"users":[...],"todos":[...],"files":[...]}`