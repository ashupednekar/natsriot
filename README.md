# natsriot ⚡

A simple NATS stream setup and test script. Quickly create streams, publish messages, and see how they replicate across NATS servers.


> ### ⚠️ Known Issue: Exponential Message Growth
> Currently, there's a flaw causing the number of messages to grow exponentially due to duplicate cyclic writes. This issue is known and will be fixed soon. Stay tuned! 🛠️

## Setup 🚀

1. Clone this repo:

   ```
   git clone https://github.com/yourusername/natsriot.git
   cd natsriot
   ```
2. up the containers
   
   ```bash
    (base) ashu@ashu-work:~/oss/natsriot$ docker compose -f one.docker-compose.yaml up -d
    WARN[0000] Found orphan containers ([nats2 natsriot2]) for this project. If you removed or renamed this service in your compose file, you can run this command with the --remove-orphans flag to clean it up.
    [+] Running 2/0
     ✔ Container nats1      Running                                                                                                                                                                                                           0.0s
     ✔ Container natsriot1  Running                                                                                                                                                                                                           0.0s
    (base) ashu@ashu-work:~/oss/natsriot$ docker compose -f two.docker-compose.yaml up -d
    WARN[0000] Found orphan containers ([nats1 natsriot1]) for this project. If you removed or renamed this service in your compose file, you can run this command with the --remove-orphans flag to clean it up.
    [+] Running 2/0
     ✔ Container natsriot2  Running                                                                                                                                                                                                           0.0s
     ✔ Container nats2      Running    
    (base) ashu@ashu-work:~/oss/natsriot$ docker logs nats1
    [1] 2024/11/29 14:30:47.133933 [INF] Starting nats-server
    [1] 2024/11/29 14:30:47.134003 [INF]   Version:  2.9.20
    [1] 2024/11/29 14:30:47.134010 [INF]   Git:      [97dd7cb]
    [1] 2024/11/29 14:30:47.134015 [INF]   Name:     nats1
    [1] 2024/11/29 14:30:47.134023 [INF]   Node:     RztkeQup
    [1] 2024/11/29 14:30:47.134029 [INF]   ID:       NAL2K3SVQLQYCRSYXHXZR27PG67PAQPAIDQ7WSKFB6FW64MLYTPG7RZI
    [1] 2024/11/29 14:30:47.134598 [INF] Starting JetStream
    [1] 2024/11/29 14:30:47.135304 [INF]     _ ___ _____ ___ _____ ___ ___   _   __  __
    [1] 2024/11/29 14:30:47.135314 [INF]  _ | | __|_   _/ __|_   _| _ \ __| /_\ |  \/  |
    [1] 2024/11/29 14:30:47.135320 [INF] | || | _|  | | \__ \ | | |   / _| / _ \| |\/| |
    [1] 2024/11/29 14:30:47.135325 [INF]  \__/|___| |_| |___/ |_| |_|_\___/_/ \_\_|  |_|
    [1] 2024/11/29 14:30:47.135330 [INF]
    [1] 2024/11/29 14:30:47.135336 [INF]          https://docs.nats.io/jetstream
    [1] 2024/11/29 14:30:47.135341 [INF]
    [1] 2024/11/29 14:30:47.135346 [INF] ---------------- JETSTREAM ----------------
    [1] 2024/11/29 14:30:47.135355 [INF]   Max Memory:      10.84 GB
    [1] 2024/11/29 14:30:47.135361 [INF]   Max Storage:     238.27 GB
    [1] 2024/11/29 14:30:47.135367 [INF]   Store Directory: "/data/jetstream"
    [1] 2024/11/29 14:30:47.135372 [INF] -------------------------------------------
    [1] 2024/11/29 14:30:47.136197 [INF] Listening for client connections on 0.0.0.0:30041
    [1] 2024/11/29 14:30:47.141123 [INF] Server is ready
    (base) ashu@ashu-work:~/oss/natsriot$ docker logs nats2
    [1] 2024/11/29 14:30:51.165143 [INF] Starting nats-server
    [1] 2024/11/29 14:30:51.165218 [INF]   Version:  2.9.20
    [1] 2024/11/29 14:30:51.165224 [INF]   Git:      [97dd7cb]
    [1] 2024/11/29 14:30:51.165233 [INF]   Name:     nats2
    [1] 2024/11/29 14:30:51.165243 [INF]   Node:     SRLRpmYS
    [1] 2024/11/29 14:30:51.165267 [INF]   ID:       NCVBJNHHZXW32H3RTG4EOZXDUZJYRN6EZYDNUNMYXAHDMYE2DGM5YOKU
    [1] 2024/11/29 14:30:51.165785 [INF] Starting JetStream
    [1] 2024/11/29 14:30:51.166563 [INF]     _ ___ _____ ___ _____ ___ ___   _   __  __
    [1] 2024/11/29 14:30:51.166575 [INF]  _ | | __|_   _/ __|_   _| _ \ __| /_\ |  \/  |
    [1] 2024/11/29 14:30:51.166584 [INF] | || | _|  | | \__ \ | | |   / _| / _ \| |\/| |
    [1] 2024/11/29 14:30:51.166591 [INF]  \__/|___| |_| |___/ |_| |_|_\___/_/ \_\_|  |_|
    [1] 2024/11/29 14:30:51.166596 [INF]
    [1] 2024/11/29 14:30:51.166602 [INF]          https://docs.nats.io/jetstream
    [1] 2024/11/29 14:30:51.166607 [INF]
    [1] 2024/11/29 14:30:51.166611 [INF] ---------------- JETSTREAM ----------------
    [1] 2024/11/29 14:30:51.166620 [INF]   Max Memory:      10.84 GB
    [1] 2024/11/29 14:30:51.166626 [INF]   Max Storage:     238.27 GB
    [1] 2024/11/29 14:30:51.166632 [INF]   Store Directory: "/data/jetstream"
    [1] 2024/11/29 14:30:51.166637 [INF] -------------------------------------------
    [1] 2024/11/29 14:30:51.167582 [INF] Listening for client connections on 0.0.0.0:30042
    [1] 2024/11/29 14:30:51.173087 [INF] Server is ready
   ```

3. Configure your NATS contexts:

   ```
   nats context add one -s localhost:30041
   nats context add two -s localhost:30042
   ```

Now you're ready to create streams and publish messages!

## Usage 📡

### Create a Stream

Run the following to create a stream:

```
nats --context one s create
```

Output:
```
[one] ? Stream Name example
[one] ? Subjects eg.>
[one] ? Storage file
[one] ? Replication 1
[one] ? Retention Policy Work Queue
...
Stream example was created
```

### Publish Messages

Publish messages to a subject on the first NATS server:

```
nats --context one pub eg.aaa "Hello from natsriot!"
```

Output:
```
Published 3 bytes to "eg.aaa"
```

Publish a few more messages:

```
nats --context one pub eg.aaa "Another message"
nats --context one pub eg.aaa "And another!"
```

Output:
```
Published 3 bytes to "eg.aaa"
Published 3 bytes to "eg.aaa"
```

### View Stream Subjects on Server One

Check out the subjects on server one:

```
nats --context one s subjects example
```

Output:
```
╭─────────────────────────────────────────────────────╮
│             1 Subjects in stream example            │
├─────────┬───────┬─────────┬───────┬─────────┬───────┤
│ Subject │ Count │ Subject │ Count │ Subject │ Count │
├─────────┼───────┼─────────┼───────┼─────────┼───────┤
│ eg.aaa  │ 251   │         │       │         │       │
╰─────────┴───────┴─────────┴───────┴─────────┴───────╯
```

### View Stream Subjects on Server Two

Check out the subjects on server two, where the messages are replicated:

```
nats --context two s subjects example
```

Output:
```
╭─────────────────────────────────────────────────────╮
│             1 Subjects in stream example            │
├─────────┬───────┬─────────┬───────┬─────────┬───────┤
│ Subject │ Count │ Subject │ Count │ Subject │ Count │
├─────────┼───────┼─────────┼───────┼─────────┼───────┤
│ eg.aaa  │ 251   │         │       │         │       │
╰─────────┴───────┴─────────┴───────┴─────────┴───────╯
```

As you can see, messages published to `eg.```` on server one are now visible on server two as well, thanks to stream replication! 🔄 Works both ways 😄

> Ignore the count for now, that's a known flaw, will be fixed soon 🤞 
