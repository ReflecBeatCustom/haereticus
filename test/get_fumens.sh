#/bin/bash
wget http://175.24.50.211:8080/api --post-file=get_fumens.json -O get_fumens.rsp
cat get_fumens.rsp
echo