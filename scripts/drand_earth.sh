#!/bin/bash

set -e

KEYS_DIR=/home/ubuntu/drand/earth_keys
SHARED_SECRET="hATDw3MaFnqzl9Go3yskOrg9eTAs6SXJr80tJcZ9p9x/qDfXg5eEpeundKp9"
SECRET_PATH=/home/ubuntu/drand/secret
EARTH_HOST=$1
EARTH_PORT=8654
CONTROL_PORT=9654
OUTPUT_DIR=/home/ubuntu/drand/output
GROUP_FILEPATH=$OUTPUT_DIR/earth_group.toml
NUM_SECONDS=600

# comment this out when testing both the earth and iss scripts on the
# same machine
# killall drand

rm -rf $KEYS_DIR
mkdir -p $KEYS_DIR
mkdir -p $OUTPUT_DIR

echo $SHARED_SECRET > $SECRET_PATH

drand generate-keypair \
  --folder $KEYS_DIR \
  --tls-disable $EARTH_HOST:$EARTH_PORT
  2>&1 | tee $OUTPUT_DIR/earth_generate-keypair.log

drand start \
  --tls-disable \
  --folder $KEYS_DIR \
  --control $CONTROL_PORT \
  > $OUTPUT_DIR/earth_start.log &

sleep 1

# earth node is the leader
drand share \
  --leader \
  --nodes 2 \
  --threshold 2 \
  --secret-file $SECRET_PATH \
  --period 10s \
  --tls-disable \
  --control $CONTROL_PORT \
  2>&1 | tee $OUTPUT_DIR/earth_share.log

drand show group \
  --out $GROUP_FILEPATH \
  --control $CONTROL_PORT \
  2>&1 | tee $OUTPUT_DIR/earth_show_group.log

for i in $(seq 1 $NUM_SECONDS); do
  drand get public $GROUP_FILEPATH >> $OUTPUT_DIR/earth_get.log
  echo $i/$NUM_SECONDS
  sleep 1
done

echo "Done. Saved output files to $OUTPUT_DIR"

kill %
