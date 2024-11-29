# natsriot ⚡

A simple NATS stream setup and test script. Quickly create streams, publish messages, and see how they replicate across NATS servers.

## Setup 🚀

1. Clone this repo:

   aaa
   git clone https://github.com/yourusername/natsriot.git
   cd natsriot
   aaa

2. Configure your NATS contexts:

   aaa
   nats context add one -s localhost:30041
   nats context add two -s localhost:30042
   aaa

Now you're ready to create streams and publish messages!

## Usage 📡

### Create a Stream

Run the following to create a stream:

aaa
nats --context one s create
aaa

Output:
aaa
[one] ? Stream Name example
[one] ? Subjects eg.>
[one] ? Storage file
[one] ? Replication 1
[one] ? Retention Policy Work Queue
...
Stream example was created
aaa

### Publish Messages

Publish messages to a subject on the first NATS server:

aaa
nats --context one pub eg.aaa "Hello from natsriot!"
aaa

Output:
aaa
Published 3 bytes to "eg.aaa"
aaa

Publish a few more messages:

aaa
nats --context one pub eg.aaa "Another message"
nats --context one pub eg.aaa "And another!"
aaa

Output:
aaa
Published 3 bytes to "eg.aaa"
Published 3 bytes to "eg.aaa"
aaa

### View Stream Subjects on Server One

Check out the subjects on server one:

aaa
nats --context one s subjects example
aaa

Output:
aaa
╭─────────────────────────────────────────────────────╮
│             1 Subjects in stream example            │
├─────────┬───────┬─────────┬───────┬─────────┬───────┤
│ Subject │ Count │ Subject │ Count │ Subject │ Count │
├─────────┼───────┼─────────┼───────┼─────────┼───────┤
│ eg.aaa  │ 251   │         │       │         │       │
╰─────────┴───────┴─────────┴───────┴─────────┴───────╯
aaa

### View Stream Subjects on Server Two

Check out the subjects on server two, where the messages are replicated:

aaa
nats --context two s subjects example
aaa

Output:
aaa
╭─────────────────────────────────────────────────────╮
│             1 Subjects in stream example            │
├─────────┬───────┬─────────┬───────┬─────────┬───────┤
│ Subject │ Count │ Subject │ Count │ Subject │ Count │
├─────────┼───────┼─────────┼───────┼─────────┼───────┤
│ eg.aaa  │ 251   │         │       │         │       │
╰─────────┴───────┴─────────┴───────┴─────────┴───────╯
aaa

As you can see, messages published to `eg.aaa` on server one are now visible on server two as well, thanks to stream replication! 🔄
