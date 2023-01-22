docker build --tag wm-back-end .

docker run -d -p 8000:8000 --name wm-back-end wm-back-end