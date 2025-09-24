# Sample invoke for pushing QR detail onchain

``` sh
curl -X POST 'http://localhost:3001/invoke' \
  -H 'Content-Type: application/x-www-form-urlencoded' \
  --data 'channelid=mychannel' \
  --data 'chaincodeid=horticulture' \
  --data 'function=CreateAsset' \
  --data-urlencode 'args@asset.json'
```

# Sample query for getting QR details

``` sh
curl 'http://localhost:3001/query?channelid=mychannel&chaincodeid=horticulture&function=ReadAsset&args=2' 
```