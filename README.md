
# YAML Utilities

A command-line tool for validating, comparing, and converting YAML files, with support for directories and various output formats. Built with **Go** for speed and extensibility.

---

## Features

- **Validation**: Ensure YAML files are properly formatted and highlight errors.
- **Diffing**: Compare two YAML files and identify differences.
- **Conversion**: Convert between YAML and JSON formats.
- **Directory Support**: Validate multiple YAML files in a directory with options to exclude specific paths.
- **Flexible Output**: Display validation results in text, JSON, or YAML formats.

---

## Installation

### Prerequisites

- Go version `1.20` or later installed.

### Clone & Build

```bash
git clone https://github.com/Maasym/yaml-utilities
cd yaml-utilities
go build -o yaml-utilities
```

### Binary Usage

Move the binary to your `PATH` for global usage:

```bash
mv yaml-check /usr/local/bin/
```

---

## Usage

### 1. **Validate YAML Files**

#### Validate a Single File
```bash
yaml-utilities validate --path <file.yaml>
```

#### Validate a Directory
```bash
yaml-utilities validate --path <directory>
```

#### Exclude Paths
```bash
yaml-utilities validate --path <directory> --exclude <path1,path2>
```

#### Change Output Format
```bash
yaml-utilities validate --path <file.yaml> --output json
```

---

### 2. **Compare YAML Files**

Compare two YAML files and highlight differences:

```bash
yaml-utilities diff --file1 <file1.yaml> --file2 <file2.yaml>
```

---

### 3. **Convert YAML & JSON**

#### YAML to JSON
```bash
yaml-utilities convert --input <file.yaml> --output <file.json> --format json
```

#### JSON to YAML
```bash
yaml-utilities convert --input <file.json> --output <file.yaml> --format yaml
```

---

## Examples

### Example 1: Validate YAML in a Directory
```bash
yaml-utilities validate --path ./configs --exclude ./configs/ignore
```

### Example 2: Diff Two Files
```bash
yaml-utilities diff --file1 old.yaml --file2 new.yaml
```

### Example 3: Convert YAML to JSON
```bash
yaml-utilities convert --input config.yaml --output config.json --format json
```

---

## Output Formats

- **Text (default)**:
  ```plaintext
  Validation Summary:
  Total files checked: 3
  Valid files: 2
  Invalid files: 1

  Errors:
  Invalid YAML in file 'example.yaml': yaml: line 5: mapping values are not allowed in this context
  ```

- **JSON**:
  ```json
  {
    "total": 3,
    "valid": 2,
    "invalid": 1,
    "errors": [
      "Invalid YAML in file 'example.yaml': yaml: line 5: mapping values are not allowed in this context"
    ]
  }
  ```

- **YAML**:
  ```yaml
  total: 3
  valid: 2
  invalid: 1
  errors:
    - "Invalid YAML in file 'example.yaml': yaml: line 5: mapping values are not allowed in this context"
  ```

---

## Roadmap

- [ ] Add custom linting rules.
- [ ] Introduce `fix` functionality for common YAML errors.
- [ ] Enhance `diff` with detailed reporting and optional merging.
