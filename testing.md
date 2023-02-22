## NVME vs. PIOPS



### PIOPS cluster 1 (graviton) Single Shard
- R80, Max IOPS
```
Time: 2022-12-13 23:31:35.377319284 +0000 UTC m=+132.834573757
┌────────────────────────────────────────────────────────────────────────────────────────┐
| Operation       | Per Second | Avg Latency (us) | Errors                               |
| Insert          | 5,752      | 27,814           | map[]                                |
| Reads by _id    | 11,486     | 5,223            | map[]                                |
| Secondary Reads | 33,122     | 4,830            | map[mongo: no documents in result:3] |
| Aggregations    | 10         | 184,024          | map[]                                |
| Updates         | 72         | 27,630           | map[]                                |
└────────────────────────────────────────────────────────────────────────────────────────┘

```
### NVME Single Shard
- R80 NVME

```
Time: 2022-12-13 23:41:36.581069579 +0000 UTC m=+396.544879255
┌──────────────────────────────────────────────────────────┐
| Operation       | Per Second | Avg Latency (us) | Errors |
| Insert          | 4,644      | 34,448           | map[]  |
| Reads by _id    | 8,625      | 6,955            | map[]  |
| Secondary Reads | 21,736     | 7,360            | map[]  |
| Aggregations    | 5          | 356,649          | map[]  |
| Updates         | 57         | 34,651           | map[]  |
└──────────────────────────────────────────────────────────┘
```

### Options

```
--insert-workers
160
--update-workers
1
--id-read-workers
60
--aggregation-works
3
--secondary-id-read-workers
160
```


# M40 Round

## NVME Single Shard N40

```
┌─────────────────────────────────────────────────────────────────────────────────────────┐
| Operation       | Per Second | Avg Latency (us) | Errors                                |
| Insert          | 481        | 332,260          | map[]                                 |
| Reads by _id    | 1,153      | 51,994           | map[]                                 |
| Secondary Reads | 3,208      | 49,866           | map[mongo: no documents in result:92] |
| Aggregations    | 3          | 599,071          | map[]                                 |
| Updates         | 6          | 324,804          | map[]                                 |
└─────────────────────────────────────────────────────────────────────────────────────────┘
```

## R40 3000k IOPS

```
Time: 2022-12-14 21:12:38.058167309 +0000 UTC m=+457.358814111
┌─────────────────────────────────────────────────────────────────────────────────────────┐
| Operation       | Per Second | Avg Latency (us) | Errors                                |
| Insert          | 696        | 229,765          | map[]                                 |
| Reads by _id    | 1,617      | 37,103           | map[]                                 |
| Secondary Reads | 5,059      | 31,626           | map[mongo: no documents in result:85] |
| Aggregations    | 17         | 116,028          | map[]                                 |
| Updates         | 8          | 223,123          | map[]                                 |
└─────────────────────────────────────────────────────────────────────────────────────────┘
```

## R40 3K IOPS Intel

```
Time: 2022-12-15 20:05:07.624525797 +0000 UTC m=+1862.475209380
┌────────────────────────────────────────────────────────────────────────────────────────┐
| Operation       | Per Second | Avg Latency (us) | Errors                               |
| Insert          | 542        | 294,919          | map[]                                |
| Reads by _id    | 1,616      | 37,106           | map[]                                |
| Secondary Reads | 5,053      | 31,661           | map[mongo: no documents in result:1] |
| Aggregations    | 14         | 136,336          | map[]                                |
| Updates         | 6          | 298,716          | map[]                                |
└────────────────────────────────────────────────────────────────────────────────────────┘
```