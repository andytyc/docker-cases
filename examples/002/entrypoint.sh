#!/usr/bin/env sh

CMD=$1

case "$CMD" in
  "dev" )
    echo "--- dev ---"
    exec /app/docker-cases -case=1001
    ;;

  "pro" )
    echo "--- pro ---"
    exec /app/docker-cases -case=1001 -block=true
    ;;

   * )
    exec $CMD ${@:2}
    ;;
esac
