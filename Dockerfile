# --- Build image ---
FROM golang:1.17.1-alpine AS build

RUN apk --no-cache add ca-certificates git make bash

WORKDIR /src/UpstartGriefBot
COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux
RUN make

# --- Execution image ---
FROM scratch
LABEL org.opencontainers.image.source https://github.com/ThomasK33/UpstartGriefBot

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /src/UpstartGriefBot/bin/UpstartGriefBot /UpstartGriefBot

CMD [ "/UpstartGriefBot" ]
