FROM golang

COPY . .

ENV SLEEP_TIME 5
ENV API_KEY test
ENV CHANNEL_ID test

CMD [ "go" "run" "." ]