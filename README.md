# natsriot 🚀🌪️

Welcome to **natsriot**, the ultimate NATS-powered chaos! 🌀 This repo lets you easily set up and experiment with NATS Streams, publish and subscribe to subjects across multiple contexts, and see how messages ripple through your NATS cluster. It's time to take your messaging game to the next level with **speed, simplicity, and a bit of fun!**

## Features ✨

- **Stream Management** 📈: Create and manage streams with ease.
- **Cross-Context Replication** 🔄: Publish a message to one context, and watch it replicate across other servers automatically.
- **Flexible Stream Configuration** ⚙️: Define subjects, retention policies, replication settings, and more with a single command.
- **Real-time Messaging** 📨: Publish and subscribe to messages in real time, and watch the magic happen.

## Installation 🚀

### Prerequisites

Make sure you have [NATS CLI](https://docs.nats.io/nats-tools/nats/) installed on your system. This project uses **NATS Streams** and assumes you have a NATS cluster up and running (or you can easily spin one up locally).

### Install the NATS CLI
1. **Homebrew (MacOS/Linux)**

   ```
   brew install nats-io/nats-tools/nats
   ```

2. **From Source** (for the adventurous)

   ```
   git clone https://github.com/nats-io/natscli.git
   cd natscli
   make install
   ```

### Clone the Repo 🛠️

```
git clone https://github.com/yourusername/natsriot.git
cd natsriot
```

## Usage 📡

### 1. **Create a Stream** 💥
Use the NATS CLI to create a stream on the server, like so:

```
nats --context one s create
```

This will walk you through the interactive process of stream creation. Set up subjects, replication, and retention policies to your liking.

### 2. **Publish Messages** 📬
Publish messages to a subject in your stream like so:

```
nats --context one pub eg.``` "Hello, NATS!"
```

This will publish a message to the `eg.```` subject. You can view the stream and messages on the other context servers.

### 3. **Subscribe to Subjects** 🔔
Listen to messages being published to a subject with the following command:

```
nats --context one sub eg.```
```

As messages are published to `eg.````, they will appear in real-time!

## Fun Commands 🎮

Here are some fun things you can do with `natsriot`:

- **Publish to multiple subjects** 📬

   ```
   nats --context one pub eg.``` "Hello, World!" && nats --context one pub eg.bbb "NATS is awesome!"
   ```

- **View Stream Stats** 📊

   ```
   nats --context one s stats
   ```

- **Replicate Streams Across Servers** 🌍
   Watch how messages get replicated between contexts and streams!

## Troubleshooting ⚠️

If you encounter issues, here are a few tips:
- Make sure your NATS servers are running and accessible.
- Check your network connectivity.
- Use `nats --context <context-name> s subjects` to view the subjects in the stream.

## Contributing 🤝

We'd love to see your contributions! Feel free to fork this repo, create a branch, and submit a PR. Whether it's fixing bugs, improving documentation, or adding cool new features — you're welcome here!

### Steps to contribute:
1. Fork this repo
2. Create a new branch (`git checkout -b feature-name`)
3. Make your changes
4. Commit your changes (`git commit -am 'Add cool new feature'`)
5. Push to your branch (`git push origin feature-name`)
6. Create a pull request!

## License 📜

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

Happy NATS-ing! 🎉
