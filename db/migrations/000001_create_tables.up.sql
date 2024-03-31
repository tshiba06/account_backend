-- 権限マスタ
CREATE TABLE public.master_roles (
  id SERIAL PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  display_name VARCHAR(50) NOT NULL
);
INSERT INTO public.master_roles (name, display_name) VALUES
  ('owner', 'オーナー'),
  ('editor', '編集者'),
  ('reader', '閲覧者')
;

-- 勘定貸借タイプマスタ
CREATE TABLE public.master_account_types (
  id SERIAL PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  display_name VARCHAR(50) NOT NULL
);
INSERT INTO public.master_account_types (name, display_name) VALUES
  ('debit', '借方'),
  ('credit', '貸方')
;

-- 勘定費目マスタ
CREATE TABLE public.master_expense_lists (
  id SERIAL PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  display_name VARCHAR(50) NOT NULL
);
INSERT INTO public.master_expense_lists (name, display_name) VALUES
  ('cash', '現金'),
  ('saving accounts', '普通預金'),
  ('upplies expenses', '消耗品費'),
  ('library expense', '新聞図書費'),
  ('withdrawals by owner', '事業主貸'),
  ('investments by owner', '事業主借'),
  ('rents', '地代家賃'),
  ('utilities expense', '水道光熱費'),
  ('communication expenses', '通信費'),
  ('entertainment expense', '接待交際費'),
  ('sales', '売上'),
  ('trade accounts receivable', '売掛金'),
  ('traveling expense', '旅費交通費'),
  ('depreciation expense', '減価償却費'),
  ('accumulated depreciation', '減価償却累計額'),
  ('start-up costs', '開業費'),
  ('miscellaneous expenses', '雑費')
;

-- ユーザー
CREATE TABLE public.users (
  id SERIAL PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  email VARCHAR(100) NOT NULL UNIQUE,
  master_role_id INTEGER NOT NULL REFERENCES public.master_roles(id),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_users_role_id ON public.users(master_role_id);

-- 更新時間関数
CREATE OR REPLACE FUNCTION update_timestamp() RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ language 'plpgsql';

-- ユーザー更新時間記録トリガー
CREATE TRIGGER update_users_modtime
BEFORE UPDATE ON public.users
FOR EACH ROW
EXECUTE PROCEDURE update_timestamp();

-- ユーザーごとの支払い方法管理
CREATE TABLE public.user_payment_methods (
  id SERIAL PRIMARY KEY,
  user_id INTEGER NOT NULL REFERENCES public.users(id),
  name VARCHAR(50) NOT NULL,
  display_name VARCHAR(50) NOT NULL
);
CREATE INDEX idx_user_payment_methods_user_id ON public.user_payment_methods(user_id);

-- ユーザーごとの支払い場所管理
CREATE TABLE public.user_payment_locations (
  id SERIAL PRIMARY KEY,
  user_id INTEGER NOT NULL REFERENCES public.users(id),
  name VARCHAR(50) NOT NULL,
  display_name VARCHAR(50) NOT NULL
);
CREATE INDEX idx_user_payment_locations_user_id ON public.user_payment_locations(user_id);

-- 1年毎の勘定まとめテーブル
CREATE TABLE public.user_accounts (
  id SERIAL PRIMARY KEY,
  user_id INTEGER NOT NULL REFERENCES public.users(id),
  year INTEGER NOT NULL
);
CREATE INDEX idx_user_accounts_user_id ON public.user_accounts(user_id);

-- 勘定テーブル
CREATE TABLE public.user_account_details (
  id SERIAL PRIMARY KEY,
  user_account_id INTEGER NOT NULL REFERENCES public.user_accounts(id),
  master_account_type_id INTEGER NOT NULL REFERENCES public.master_account_types(id),
  master_expense_list_id INTEGER NOT NULL REFERENCES public.master_expense_lists(id),
  user_payment_method_id INTEGER NOT NULL REFERENCES public.user_payment_methods(id),
  user_payment_location_id INTEGER NOT NULL REFERENCES public.user_payment_locations(id),
  amount INTEGER NOT NULL,
  abstract TEXT,
  note TEXT,
  pay_day DATE,
  is_lock BOOLEAN NOT NULL DEFAULT FALSE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  created_by INTEGER NOT NULL REFERENCES public.users(id),
  updated_by INTEGER NOT NULL REFERENCES public.users(id)
);
CREATE INDEX idx_user_account_details_user_account_id ON public.user_account_details(user_account_id);
CREATE INDEX idx_user_account_details_master_account_type_id ON public.user_account_details(master_account_type_id);
CREATE INDEX idx_user_account_details_master_expense_list_id ON public.user_account_details(master_expense_list_id);
CREATE INDEX idx_user_account_details_user_payment_method_id ON public.user_account_details(user_payment_method_id);
CREATE INDEX idx_user_account_details_user_payment_location_id ON public.user_account_details(user_payment_location_id);
CREATE INDEX idx_user_account_details_created_by ON public.user_account_details(created_by);
CREATE INDEX idx_user_account_details_updated_by ON public.user_account_details(updated_by);

-- 勘定詳細テーブル
CREATE TRIGGER update_user_account_details_modtime
BEFORE UPDATE ON public.user_account_details
FOR EACH ROW
EXECUTE PROCEDURE update_timestamp();
