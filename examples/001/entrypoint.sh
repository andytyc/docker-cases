#!/usr/bin/env sh

CMD=$1

case "$CMD" in
  "dev" )
    echo "--- dev ---"
    source /app/env.sh
    exec /app/docker-cases -case=1001
    ;;

  "start" )
    echo "--- start ---"
    source /app/env.sh
    exec /app/docker-cases -case=1001 -block=true
    ;;

   * )
    exec $CMD ${@:2}
    ;;
esac
