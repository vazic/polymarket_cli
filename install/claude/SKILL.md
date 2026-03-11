---
name: Polymarket CLI
description: A CLI tool for querying Polymarket odds, markets, and orderbooks
---

# Polymarket CLI Skill

This skill allows Claude to interact with Polymarket, the world's largest prediction market, directly from the console. 

## Prerequisites
Before using this tool, you must ensure the user has compiled the binary and it is available in the project, or they have built it globally. 

If the `polymarket` binary is not present in the current directory, instruct the user to build it:
```bash
go build -o polymarket main.go
```

## Available Commands

Always use the `--json` flag (which is default, but good fully-qualified practice) to get structured data back.

### 1. Search Markets
Use to find a specific event or market ID based on a keyword.
```bash
./polymarket search --query "<search_term>" --limit 5
```
**Output:** JSON array of simplified market events. Look for the `conditionId` or `clobTokenIds` fields to use in subsequent commands.

### 2. Get Market Details
Use to fetch the exact odds (prices) for an outcome using a `conditionId`.
```bash
./polymarket market --id "<conditionId>"
```

### 3. Get Orderbook
Use to view the depth of the market (bids and asks) for a specific outcome using a `clobTokenId`. Note: Some tokens may return a 404 error if their orderbook is inactive; this is expected behavior.
```bash
./polymarket orderbook --token-id "<clobTokenId>"
```

## How to use this skill
When a user asks about the odds on Polymarket for a specific event (e.g., "What are the odds Trump wins the election?"):
1. Propose running `./polymarket search --query "Trump election" --limit 5`
2. Parse the output to find the exact `conditionId` for the relevant market.
3. Propose running `./polymarket market --id "<conditionId>"` to get the exact "Yes" / "No" pricing.
4. Report the resulting prices back to the user in a natural format (e.g. "The market is currently pricing a Yes at 60 cents!").
