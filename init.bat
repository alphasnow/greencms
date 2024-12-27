cd server
if not exist .env (
    copy .env.example .env
)
go mod tidy
cd ../
cd admin
if not exist .env (
    copy .env.example .env
)
npm install
cd ../
cd web
if not exist .env (
    copy .env.example .env
)
npm install
cd ../