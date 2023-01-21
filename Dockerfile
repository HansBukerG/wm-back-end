FROM golang:1.19-alpine

# Work directory
RUN Mkdir -p /go/HansBukerG/wm-back-end
WORKDIR /go/HansBukerG/wm-back-end

# Copy everything from my project to my wor directory
COPY . ./go/HansBukerG/wm-back-end

# I run my command from Go to the my prod-version
RUN Go build ./main.go

RUN adduser --disabled-password productUser

USER productUser

RUN chown -R productUser:productUser ./go/HansBukerG/wm-back-end

EXPOSE 8000

RUN ./go/HansBukerG/wm-back-end/main.exe