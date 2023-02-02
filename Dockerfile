FROM golang:1.19-alpine

# Work directory
WORKDIR /go/HansBukerG/wm-back-end

# Copy everything from my project to my wor directory
COPY go.mod ./
COPY go.sum ./
COPY .env ./

RUN go mod download

COPY . ./

RUN go build -o /wm-back-end

EXPOSE 8080

CMD [ "/wm-back-end" ]