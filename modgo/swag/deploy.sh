#!/bin/bash
uname=$(uname)

checkresult() {
  if [[ $1 -ne 0 ]]; then
    exit 1
  fi
}
dist_dir="$GOPATH/src/deploy/Dist"
if [ ! -d "$dist_dir" ]; then
  mkdir -p $dist_dir
fi

go_modules=$1
os=$2
target_env=$3
gcflags=''
if [ "$3" = "prod" ]; then
  gcflags='-gcflags=-trimpath=$GOPATH'
fi

for m in $go_modules; do
#  cd $GOPATH/src/$m
#  count=$(git rev-list HEAD | wc -l | sed -e 's/ *//g' | xargs -n1 printf %04d)
#  commit=$(git show --abbrev-commit HEAD | grep '^commit' | sed -e 's/commit //')
#  branch=$(git branch | grep "\*" | sed -e 's/\* //' | sed -e 's/release\///')
#  buildno="${branch}-build${count}.${commit}"
  case $uname in
  Darwin)
#  sed -i "" "s/const GIT_VERSION = \".*\"/const GIT_VERSION = \"$buildno\"/g" $GOPATH/src/vendor/laoyuegou.com/version/binary_version.go
    ;;

  Linux)
#   sed -i "s/const GIT_VERSION = \".*\"/const GIT_VERSION = \"$buildno\"/g" $GOPATH/src/vendor/laoyuegou.com/version/binary_version.go
    ;;

  *)
    echo "unsupport system..."
    exit 1
    ;;
  esac
# 编译
  gox -os="$os" -arch="amd64"   $gcflags -output="$dist_dir/"$go_modules"_{{.OS}}_{{.Arch}}"
  checkresult $?

  cd $dist_dir
  module_name=$(basename $m)
  tar cvzf "$module_name"_linux_amd64.tar.gz "$module_name"_linux_amd64

  scp "$module_name"_linux_amd64.tar.gz root@119.28.10.43:~/projects
  ssh root@119.28.10.43 /root/projects/restart.sh $module_name

done
