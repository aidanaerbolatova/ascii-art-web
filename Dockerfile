FROM golang:1.17-alpine AS builder
WORKDIR /app 
COPY . .
 
RUN  go build -o main cmd/main.go

FROM alpine:3.6
WORKDIR /app
LABEL  captain="Aidana" port:="8080" project="ASCII-ART-WEB"
COPY --from=builder /app/main .
COPY --from=builder /app/ascii/source  /app/ascii/source
COPY --from=builder /app/templates /app/templates



CMD [ "/app/main" ]
