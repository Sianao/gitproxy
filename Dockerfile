# Install dependencies only when needed
FROM docker.io/node:alpine AS deps
# Check https://github.com/nodejs/docker-node/tree/b4117f9333da4138b03a546ec926ef50a31506c3#nodealpine to understand why libc6-compat might be needed.
RUN apk add --no-cache libc6-compat
WORKDIR /app
COPY . .
RUN yarn install --frozen-lockfile

# Rebuild the source code only when needed
FROM docker.io/node:alpine AS builder
WORKDIR /app
COPY . .
COPY --from=deps /app/node_modules ./node_modules
RUN npm run build 

FROM golang:alpine AS Gobuilder

LABEL stage=gobuilder

ENV CGO_ENABLED 0

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
RUN go build -ldflags="-s -w" -o /app/gitproxy main.go


FROM alpine

ENV TZ Asia/Shanghai
RUN  apk add  ca-certificates
WORKDIR /app
COPY --from=Gobuilder /app/gitproxy /app/gitproxy
COPY --from=builder  /app/out /app/out/
CMD ["./gitproxy"]