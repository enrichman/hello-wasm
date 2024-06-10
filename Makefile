
build:
	GOOS=js GOARCH=wasm go build -o ./hello1/main.wasm ./hello1
	GOOS=js GOARCH=wasm go build -o ./hello2/main.wasm ./hello2
	GOOS=js GOARCH=wasm go build -o ./hello3/main.wasm ./hello3

serve:
	docker run -it --rm -p 8080:80 --name nginx \
    	-v $(shell pwd):/usr/share/nginx/html -d \
    	nginx

stop:
	docker stop nginx
