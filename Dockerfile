# Pull base image.
FROM ubuntu

ENV DEBIAN_FRONTEND noninteractiv

# Install hddtemp
RUN apt-get update && apt-get -y install build-essential hddtemp golang-go git

RUN go get github.com/prometheus/client_golang/prometheus
RUN go get github.com/prometheus/client_golang/prometheus/promauto
RUN go get github.com/prometheus/client_golang/prometheus/promhttp

COPY hddexporter.go .

# Define default command.
# CMD hddtemp -d --listen localhost --port 7634 /dev/sd[cd] && go run hddexporter.go
CMD hddtemp -d --listen localhost --port 7634 $HDDTEMP_ARGS && go run hddexporter.go
