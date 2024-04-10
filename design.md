### 2024/02/23

- RBAC システムで権限周りは制御する

### 2024/04/05

できること

- 個人事業主（仕入れなし）の仕訳
- csv のインポートをしたらある程度の仕訳を自動で行う
- pdf での出力 etax に入力する方法を出す
- 収入・支出のグラフ推移をみることができる
- 事業モードと個人モードがある
- 事業モードと個人モードをシームレスに切り替えることができる
- ログインができる single sign on ができるとよい
- 大雑把な操作履歴を残す
- 個人の google drive との連携をして書類や履歴を残す

すぐにできなくても問題ないがいずれできること

- 細かい仕訳
- ocr
- 複数人での共有
  - 事業モードのみ
- multi tenancy?
- 料金徴収
- 未来予測を提示すること
- 統計処理
- 所得税などの税率が上がりそうなタイミングなどを教えてくれる
- 節税できそうなところを教える
- LLM の組み込み
- 画像のアップロード
- 細かいセキュリティ対応
- ログを残す

できなくてもよいこと

- 他のサービスとの対抗
- 他のサービスとの API 連携
- 多人数の利用 5 人ぐらいしか使わない想定

---

インフラ想定

- db
  無料の db saas をいったん利用とかで考えている

- api server
  cloud run or gcs or lambda or cloud function or さくらのレンタルサーバーなどの 500yen/1month 程度のサーバー(この場合はほかのサービスと共存させる予定)

- login auth
  firebase or single sign on

- frontend
  gcs or s3 or レンタルサーバー

- desktop
  これは作りたい　配布だけでいいので楽

- mobile
  不明 作るのに対して返ってくる利益が見合わなさそう

## テストについては必ず書いておくこと　後々の資料にもなるため

```mermaid
erDiagram

master_roles {
    int id PK
    varchar name
}
master_roles ||--o{ users : hasMany

master_account_types {
    int id PK
    varchar name "借方,貸方"
}
master_account_types ||--o{ user_account_details : hasMany

master_expense_lists {
    int id PK
    varchar name "現金など"
}
master_expense_lists ||--o{ user_account_details : hasMany

user_payment_methods {
    int id PK
    int user_id FK
    varchar name "現金,iD,Suicaなど"
}
user_payment_methods ||--o{ user_account_details : hasMany

user_payment_locations {
    int id PK
    varchar name "支払い場所"
}
user_payment_locations ||--o{ user_account_details : hasMany

users {
    int id PK
    int master_role_id FK
    varchar email
    varchar name
}
users ||--o{ user_accounts : hasMany
users ||--o{ user_payment_methods : hasMany
users ||--o{ user_payment_locations : hasMany

user_accounts {
    int id PK
    int user_id FK
    int yaer "勘定年"
}
user_accounts ||--o{ user_account_details : hasMany

user_account_details {
    int id PK
    int user_account_id FK
    int master_account_type_id FK
    int master_expense_list_id FK
    int master_payment_method_id FK
    int amount "金額"
    text abstract "摘要"
    text note "備考"
    datetime pay_day "勘定日"
    timestamp created_at
    timestamp updated_at
    int created_by FK
    int updated_by FK
    bool is_lock "データの更新を不可にする"
}
```
