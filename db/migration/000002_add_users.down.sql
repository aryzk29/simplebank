ALTER TABLE IF EXISTS "accounts" DROP CONSTRAINT IF EXIST "owner_currency_key";

ALTER TABLE IF EXISTS "accounts" DROP CONSTRAINT IF EXIST "accounts_owner_fkey";

DROP TABLE IF EXISTS "users"