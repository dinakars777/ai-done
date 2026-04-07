# 🔔 AI Done

> macOS menu bar app that notifies you when AI coding assistants finish

Never miss when your AI finishes generating code. AI Done sits in your menu bar and gives you instant notifications when Cursor, Windsurf, Kiro, or other AI IDEs complete their responses.

![Demo](demo.gif)

## ✨ Features

- 🎯 **Menu bar indicator** - See AI status at a glance
- 🔊 **Sound notifications** - Hear when AI finishes
- 💬 **Desktop notifications** - Get popup alerts
- 🎨 **Customizable** - Choose your sounds and notification style
- ⚡ **Lightweight** - Minimal resource usage
- 🔌 **Multi-IDE support** - Works with Cursor, Windsurf, Kiro, and more

## 🚀 Installation

### Homebrew (Coming Soon)
```bash
brew install --cask ai-done
```

### Manual Install
1. Download the latest release from [Releases](https://github.com/dinakars777/ai-done/releases)
2. Drag `AI Done.app` to your Applications folder
3. Launch AI Done
4. Grant necessary permissions when prompted

## 🎯 How It Works

AI Done monitors your AI IDE's activity by:
1. Watching for hook events from supported IDEs
2. Detecting when AI starts/stops generating code
3. Showing status in menu bar
4. Playing sound and showing notification when done

## 🔧 Setup

### For Kiro

AI Done automatically detects Kiro. Just make sure Kiro is running.

### For Cursor

1. Open Cursor settings
2. Add AI Done integration (automatic on first launch)
3. Restart Cursor

### For Windsurf

1. Open Windsurf settings
2. Enable AI Done notifications
3. Restart Windsurf

## ⚙️ Configuration

Click the menu bar icon → Preferences to customize:

- **Notification Style**: Sound only, popup only, or both
- **Sound**: Choose from system sounds or custom audio files
- **Quiet Hours**: Disable notifications during specific times
- **IDE Detection**: Enable/disable specific IDEs

## 🎨 Menu Bar States

- 💤 **Gray** - No AI activity
- 🤔 **Yellow** - AI is thinking/generating
- ✅ **Green** - AI just finished (briefly)
- ❌ **Red** - AI encountered an error

## 🛠️ Building from Source

```bash
# Clone the repo
git clone https://github.com/dinakars777/ai-done
cd ai-done

# Install dependencies
npm install

# Run in development
npm run dev

# Build for production
npm run build
```

## 📋 Requirements

- macOS 12.0 or later
- One of: Cursor, Windsurf, Kiro, or other supported AI IDE

## 🤝 Contributing

Contributions welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md).

### Adding IDE Support

Want to add support for another AI IDE? Check out [IDE_INTEGRATION.md](docs/IDE_INTEGRATION.md) for guidelines.

## 📝 License

MIT

## 🌟 Related Projects

- [ai-done-hooks](https://github.com/dinakars777/ai-done-hooks) - Simple hook configs (no app needed)
- [moody](https://github.com/dinakars777/moody) - macOS menu bar mood indicator
- [USB-Clawd](https://x.com/i/grok/share/4489f72c13b6405a9d8024b5f08c2247) - Physical notification device

---

**Star this repo** if it saves you time! ⭐
