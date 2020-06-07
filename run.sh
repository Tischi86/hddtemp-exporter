#!/bin/bash
if [ -z "$DATABASE_FILE" ]
then
	hddtemp -d --listen localhost --port 7634 $HDDTEMP_ARGS
else
	hddtemp -d --listen localhost --port 7634 --file=$DATABASE_FILE $HDDTEMP_ARGS
fi

go run hddexporter.go
