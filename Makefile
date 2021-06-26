build:
	echo "Compiling the backend"
	go build -o dist/bin/app main.go
	echo "Compiling the frontend"
	#npm install --prefix ./front-end/
	npm run build --prefix ./front-end/


run:
	go run main.go
