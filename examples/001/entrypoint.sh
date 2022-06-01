#!/usr/bin/env sh

CMD=$1

case "$CMD" in
  "dev" )
    echo "--- dev ---"
    # 加载环境无效
    # /bin/bash -c "source /app/env.sh"
    exec /app/docker-cases -case=1001
    ;;

  "start" )
    echo "--- start ---"
    # 加载环境无效
    # /bin/bash -c "source /app/env.sh"
    exec /app/docker-cases -case=1001 -block=true
    ;;

   * )
    exec $CMD ${@:2}
    ;;
esac
