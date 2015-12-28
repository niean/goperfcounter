Example Of GoPerfcounter
=====


这里是，goperfcounter提供的一个完整的例子。

按照如下指令，运行这个例子。

```bash
# install
cd $GOPATH/src/github.com/niean
git clone https://github.com/niean/goperfcounter.git
cd $GOPATH/src/github.com/niean/goperfcounter
go get ./...

# run
cd $GOPATH/src/github.com/niean/goperfcounter/example/scripts
./debug build && ./debug start

# proc
./debug proc metrics/json   # list all metrics in json 
./debug proc metrics/falcon # list all metrics in falcon-model

```
