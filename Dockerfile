FROM alpine

RUN apk add go nodejs npm mariadb-client bash python3 py3-pyaml

WORKDIR /app

COPY . .

RUN npm install
RUN npx @tailwindcss/cli -i ./static/style/tw.tailwind.css -o ./static/style/dist/store.css
RUN ln -s /app/config/config.yaml /app/config.yaml
RUN go build -o main .

ENTRYPOINT [ "/app/entrypoint.sh" ]
