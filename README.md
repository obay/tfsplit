# Terraform Split (tfsplit)

Terraform split is a tool to split your Terraform files into multiple files. If you have a main.tf file for example that includes multiple resource definitions, you can split them into multiple files using this tool.

## Installation

### On Linux & MacOS using [Homebrew](https://brew.sh)

```bash
brew install obay/tap/tfsplit
```

### On Windows using [Scoop](https://scoop.sh)

```powershell
scoop bucket add obay https://github.com/obay/scoop-bucket.git
scoop install obay/tfsplit
```

## Usage

Simply switch to the directory containing your Terraform files and run `tfsplit`.

```bash
tftr
```
