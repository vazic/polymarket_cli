# Polymarket Go CLI

A lightweight, zero-dependency (other than `cobra` for CLI routing) Golang CLI tool built for querying Polymarket's Gamma and CLOB APIs, heavily modularized and optimized to output clean JSON for AI/LLM Agent integration.

## Building

```bash
# Recommended build (disables CGO for better cross-platform compatibility)
CGO_ENABLED=0 go build -o polymarket main.go
```

## Commands

All commands output structured JSON by default, making them extremely easy to pipe into `jq` or read programmatically from LLM environments like Claude Code or Aider.

### 1. Search
Search the Gamma API for markets and events.
```bash
./polymarket search --query "election" --limit 3
```

### 2. Market 
Retrieve the specific pricing, odds, and details of a market given its Condition ID.
```bash
./polymarket market --id "0xYourConditionId..."
```

### 3. Orderbook
Retrieve the L2 depth (bids and asks arrays) for a given CLOB token ID.
```bash
./polymarket orderbook --token-id "123456789..."
```

## Claude Code Integration (Skill Installation)

This repository comes pre-packaged with a custom [Skill instruction set](./install/claude/SKILL.md) that teaches Claude Code exactly how to utilize the `polymarket` binary as an external tool.

**To install the tool for Claude Code:**

1. Copy the `install/claude` folder into your global or local Claude skills directory (usually `.claude/skills/`).
2. Alternatively, simply drop the `SKILL.md` anywhere context-heavy, or point Claude to read it.
3. Claude will automatically read the YAML frontmatter and understand how to route "Check polymarket..." prompts into execution bash commands using this binary!
