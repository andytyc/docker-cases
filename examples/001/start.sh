echo "-- Start --"

source /app/env.sh

echo "----"
echo $ALG_DIR
echo $PATH
echo $DARWIN_ETCD_ENDPOINTS
echo $DARWIN_ETCD_AUTH
echo "----"

/app/docker-cases
