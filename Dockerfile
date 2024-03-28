FROM golang:latest


WORKDIR /app

# Get rlm util
RUN apt-get update && apt-get install wget -y
RUN wget https://reprisesoftware.com/wp-content/uploads/2023/v15-2/x64_l1.admin.tar.gz
RUN tar -xvf x64_l1.admin.tar.gz && rm x64_l1.admin.tar.gz
RUN mv x64_l1.admin/rlmutil /usr/local/bin/rlmutil
RUN chmod +x /usr/local/bin/rlmutil

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV RLMUTIL_BIN_ROOT "/usr/local/bin"
ENV RLM_SERVER_HOST ""
ENV RLM_SERVER_PORT ""
ENV RLM_SERVER_USERNAME ""
ENV RLM_SERVER_PASSWORD ""
ENV PORT 8081

CMD ["go", "run", "./main.go"]

