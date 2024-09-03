# Steampipe Plugin CycloneDX

This repository contains a steampipe plugin that parses CycloneDX JSON files and allows for easier searching.

# Requirements

* Golang (tested on 1.22.5)
* Steampipe

# Installation

```bash
task install
```

# Example Query

```sql
SELECT * from cyclonedx_bom_info_file where directory = './examples';
```

Finding packages that don't have a reference to `github.com` and `golang.org`

```sql
SELECT * from cyclonedx_dependencies where ref not like '%github.com%' and ref not like '%golang.org%'
```

# Test

```bash
task test
```

# Uninstall

```bash
task clean
```
