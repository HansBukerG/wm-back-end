docker build --tag wm-back-end .

docker run -d -p 8000:8000 --name wm-back-end wm-back-end

docker push wm-back-end

docker tag wm-back-end hansbukerg/wm-back-end

docker push hansbukerg/wm-back-end

docker tag mongo:3.6.8 hansbukerg/local-mongo-db
docker push hansbukerg/local-mongo-db


docker tag local-image:tagname new-repo:tagname
docker push new-repo:tagname


docker tag mongo:3.6.8 registry.heroku.com/wm-mongo-db/local-db