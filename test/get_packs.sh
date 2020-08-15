#/bin/bash
wget http://175.24.50.211:8080/api --post-file=get_packs.json -O get_packs.rsp
cat get_packs.rsp
echo