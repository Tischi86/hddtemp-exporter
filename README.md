# hddtemp-exporter
```
docker run --privileged=true -p 8080:8080 -e HDDTEMP_ARGS=/dev/sd[cd] docker.pkg.github.com/tischi86/hddtemp-exporter/hddtemp-exporter:latest
```
or with your own hddtemp database file
```
docker run --privileged=true -p 8080:8080 -e HDDTEMP_ARGS=/dev/sd[cd] -e DATABASE_FILE=/mydb.db -v $(pwd)/mydb.db:/mydb.db docker.pkg.github.com/tischi86/hddtemp-exporter/hddtemp-exporter:latest
```
