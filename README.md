# piano

A cross-platform pluggable midi piano controller. This is similar to https://github.com/schollz/PIanoAI, except that it should work on Windows without problems.

For Windows users, first get http://tdm-gcc.tdragon.net/.

Then do

```bash
cd $GOPATH/src/github.com/schollz/piano
go get -d -v github.com/jteeuwen/go-bindata/...
go install -v github.com/jteeuwen/go-bindata/go-bindata
rm -rf assets 
cp -r static assets
cd assets && gzip -9 -r *
cp templates/index.html assets/index.html
go-bindata -nocompress assets 
go build -v 
```

Then you can run it with 

```bash
./piano
```

and open up a browser to localhost:8152 to hear your playing.
