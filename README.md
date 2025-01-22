# 🎬 Subtitle Timing Adjuster

A simple Go program to fix out-of-sync subtitles on .srt files! ⚡

## 🎯 Features

- Adjust subtitle timings forward or backward
- Default offset of -500ms when no value specified
- Creates a new file instead of modifying the original, so you will always have a backup
  
## 🚀 Usage

```bash
# Using default offset (-500ms)
fixsub subtitles.srt

# Using custom offset (e.g., delay by 1000ms)
fixsub subtitles.srt 1000
```

The adjusted subtitles will be saved in a new file with `.adjusted.srt` suffix.

## 🛠️ Installation

```bash
# Clone the repository
git clone https://github.com/the-eduardo/Subtitle-Timing-Adjuster

# Build the program
go build -o fixsub main.go

# Move to a directory in your PATH, on linux:
mv fixsub /usr/local/bin/

# On windows:
export PATH=$PATH:$GOPATH/bin
mv fixsub $GOPATH/bin/

# On macOS
export PATH=$PATH:$HOME/go/bin
mv fixsub ~/go/bin/
```

## 🧩 How It Works

The program processes .srt files using the following logic:

1. **Command Line Parsing** 📥
   - Reads the input file name and optional offset value
   - Defaults to -500ms if no offset is specified

2. **File Processing** 📂
   - Opens the source .srt file for reading
   - Creates a new output file with `.adjusted.srt` suffix
   - Uses buffered I/O for efficient file handling

3. **Timestamp Processing** ⏱️
   ```go
   // Example timestamp format: 00:23:33,370 --> 00:23:35,372
   ```
   - Identifies lines containing timestamp patterns (` --> `)
   - Splits timestamps into hours, minutes, seconds, and milliseconds
   - Converts to Duration for precise calculations
   - Applies the specified offset
   - Converts back to .srt format

4. **Content Preservation** 💾
   - Maintains the original formatting and text unchanged, preserving blank lines and special characters
   - Supports Unicode characters and styling tags (e.g., `{\an8}`)

## 📋 File Format Support

Supports standard .srt files with structure like:
```
355
00:23:33,370 --> 00:23:35,372
{\an8}～♪

366
00:23:35,372 --> 00:23:37,374
Next subtitle line
```

## ⚠️ Important Notes

- **Negative** offset values will shift subtitles **earlier**
- **Positive** offset values will shift subtitles **later**
- Program will always create a new file with `.adjusted.srt` suffix

## 🤝 Contributing

Contributions are always welcome! Feel free to:
- Report bugs
- Suggest features
- Submit pull requests

## 📝 License

This project is licensed under the MIT License - see the LICENSE file for details.
