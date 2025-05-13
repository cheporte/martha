# 🖤 Martha

> _“A vault for your code. A memory for your soul.”_

**Martha** is a personal CLI tool written in Go.
It lets you create, organize, and manage a vault of your projects—sorted by language, tools, and tags—with Git integration, templates, and a poetic soul.

Named after someone who reignited the spark of artistry in the author’s heart, this tool is both utility and tribute.

---

## 🌿 Features

- Initialize a dedicated **vault folder** for your code
- Create projects with metadata (language, tools, description)
- Git integration: init, status, commit, push
- Search/filter projects by tag/tool/lang
- Templates, tags, and archiving
- More features coming soon...

---

## 📦 What We're Trying to Achive

```bash
$ martha bloom -vault-path ~/vault

$ martha new cherry-v --lang cpp --tools cmake,git --desc "RISC-V assembler"

$ martha show --lang cpp
📦 cherry-v | C++ | Tools: cmake, git | "RISC-V assembler"
```

---

## How to Use Locally
While developing, you can test Martha in multiple ways depending on your needs.

First of all, run the following command to generate `go.sum` file:
```bash
go mod tidy
```

### Option 1: Quick Run with `go run`
Great for testing commands while hacking on the code:

```bash
go run .
```

Use specific commands and flags:

```bash
go run . init
go run . init --vault-path ~/myvault
```

### Option 2: Build a Local Binary
Compile Martha into a single executable:

```bash
go build -o martha
```

To run it:

```bash
./martha
./martha bloom
```

### Option 3: Install Globally (to $PATH)
Want to call `martha` from anywhere in your terminal?

```bash
go build -o martha
sudo mv martha /usr/local/bin/
```

Then use it globally:

```bash
martha bloom
```

---

## 💬 About the Name
**Martha** is not just a tool. She is a memory.
A gesture of gratitude for someone who rekindled the beauty of creation in a world that had gone a little numb.

This vault holds code, but behind the code, it holds moments.
Unseen. Unsaid. Unforgotten.

For her, even if she never sees it.

---

## 🔮 License

Released under the MIT License. For more details read [here](./LICENSE).

This tool is free to use, fork, and shape—just as inspiration is meant to be.

Build with it. Break with it. Feel with it.

> _“As I was ignited, let your feelings ignite you as well.”_
