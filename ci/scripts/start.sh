#!/bin/sh -e

/usr/local/bin/envoy -c /var/www/traderstack/scripts/envoy.yaml &
cd /var/www/traderstack/web
yarn start &
/var/www/traderstack/traderstack