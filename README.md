# deCHAR

## Introduction

A simple Go CLI used to "decode" SQL `CHAR()` commands. These are often used in SQL Injection attacks (SQLi) to obfuscate the commands/queries being injected.

Logs like

```sql
CHAR(45,120,49,45,81,45),CHAR(45,120,50,45,81,45),CHAR(45,120,51,45,81,45),CHAR(45,120,52,45,81,45),CHAR(45,120,53,45,81,45),CHAR(45,120,54,45,81,45),CHAR(45,120,55,45,81,45),...
```

are pretty unhelpful, so run it through `deCHAR` to get the "real" output.

## Usage

To run, simply run `./dechar` and provide the comma-separated `CHAR()` text as argument (support for piping/stdin will hopefully come in future):

```bash
./dechar "CHAR(45,120,49,45,81,45),CHAR(45,120,50,45,81,45),CHAR(45,120,51,45,81,45)"
```

You can also provide `-o simple|lines|table` to change the output format. The default (`simple`) is a comma-separated list, while `lines` is newline-separated and `table` prints both the original and decoded text in a table.