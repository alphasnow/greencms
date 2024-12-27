cd server
start cmd /k "go run main.go -env=dev"
cd ../admin
start cmd /k "npm run dev"
cd ../web
start cmd /k "npm run dev"
exit