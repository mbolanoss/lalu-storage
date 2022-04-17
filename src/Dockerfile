FROM golang:latest

RUN mkdir /code
WORKDIR /code

RUN export GO111MODULE=on

#Install dependencies
COPY go.mod .
RUN go mod download

#Build binary
COPY . .
RUN go build

EXPOSE 3000

ENTRYPOINT [ "/code/lalu-storage" ]

# Build and run image commands
# docker build -t lalu-storage .
# docker run -p 4000:3000 -tid lalu-storage