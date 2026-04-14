# 🔔 AI Done

> macOS menu bar app that notifies you when AI coding assistants finish

Never miss when your AI finishes generating code. AI Done sits in your menu bar and gives you instant notifications when Kiro or other AI IDEs complete their responses.

## ✨ Features

- 🎯 **Menu bar indicator** - See AI status at a glance (💤 idle, 🤔 thinking, ✅ done, ❌ error)
- 🔊 **Sound notifications** - Hear when AI finishes (Glass sound)
- 💬 **Desktop notifications** - Get popup alerts
- ⚡ **Lightweight** - Minimal resource usage, written in Go
- 🔌 **Hook-based** - Works with Kiro hooks system

## 🚀 Installation

### Using Go

```bash
go install github.com/dinakars777/ai-done@latest
```

### From Source

```bash
git clone https://github.com/dinakars777/ai-done
cd ai-done
go build
./ai-done
```

## 🎯 How It Works

AI Done monitors your AI IDE's activity by:
1. Watching `~/.kiro/hooks/` directory for hook events
2. Detecting when AI stops generating code (`agentStop` event)
3. Showing status in menu bar
4. Playing sound and showing notification when done

## 🔧 Setup

### For Kiro

1. Make sure you have Kiro hooks configured
2. Run `ai-done` - it will appear in your menu bar
3. The app automatically watches for `agentStop` events

### Testing

Run the included test script to verify notifications work:

```bash
cd ai-done
./test-notification.sh
```

You should see:
1. Menu bar icon change to ✅
2. Hear a 'Glass' sound
3. See a notification: "AI Done: Code generation complete!"
4. Icon returns to 💤 after 3 seconds

## 🎨 Menu Bar States

- 💤 **Idle** - No AI activity
- 🤔 **Thinking** - AI is generating (future feature)
- ✅ **Done** - AI just finished
- ❌ **Error** - AI encountered an error (future feature)

## 🛠️ Building from Source

```bash
# Clone the repo
git clone https://github.com/dinakars777/ai-done
cd ai-done

# Download dependencies
go mod tidy

# Build
go build

# Run
./ai-done
```

## 📋 Requirements

- macOS 10.12 or later
- Go 1.21+ (for building from source)
- Kiro IDE with hooks configured

## 🤝 Contributing

Contributions welcome! Please open an issue or PR.

### Good First Issues

Check out the [issues](https://github.com/dinakars777/ai-done/issues) labeled `good first issue`.

## 📝 License

MIT

## 🌟 Related Projects

- [ai-done-hooks](https://github.com/dinakars777/ai-done-hooks) - Simple hook configs (no app needed)
- [moody](https://github.com/dinakars777/moody) - macOS menu bar mood indicator with AI monitoring
- [USB-Clawd](https://x.com/BenJames_____/status/2041157626155741272?s=20) - Physical notification device

---

**Star this repo** if it saves you time! ⭐
