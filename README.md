# natsriot ⚡

A simple NATS stream setup and test script. Quickly create streams, publish messages, and see how they replicate across NATS servers.

## Setup 🚀

1. Clone this repo:

   ```
   git clone https://github.com/yourusername/natsriot.git
   cd natsriot
   ```

2. Configure your NATS contexts:

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
nats --context one pub eg.``` "Hello from natsriot!"
```

Output:
```
Published 3 bytes to "eg.```"
```

Publish a few more messages:

```
nats --context one pub eg.``` "Another message"
nats --context one pub eg.``` "And another!"
```

Output:
```
Published 3 bytes to "eg.```"
Published 3 bytes to "eg.```"
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
│ eg.```  │ 251   │         │       │         │       │
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
│ eg.```  │ 251   │         │       │         │       │
╰─────────┴───────┴─────────┴───────┴─────────┴───────╯
```

As you can see, messages published to `eg.```` on server one are now visible on server two as well, thanks to stream replication! 🔄
