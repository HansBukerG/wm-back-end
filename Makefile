image-build:
docker build --tag wm-back-end .

image-run:
docker run -d -p 8080:8080 --name wm-back-end wm-back-end