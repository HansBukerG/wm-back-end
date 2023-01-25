docker build --tag wm-back-end .

docker run -d -p 8080:8080 --name wm-back-end wm-back-end