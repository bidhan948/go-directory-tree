# Directory Tree Explorer 🌲

A versatile Go utility that recursively traverses a directory structure and displays it as a tree. It offers flexible configuration options including HTML output and maximum depth filtering.

## Overview 🚀

**Directory Tree Explorer** is a command-line tool written in Go that:
- Recursively traverses directories.
- Displays the directory and file structure in a tree-like format.
- Provides configuration options via command-line flags.
- Optionally outputs the result as an HTML file with preserved formatting.
- Allows setting a maximum traversal depth to limit output.

This tool is useful for quickly visualizing directory structures and can be further extended or integrated into batch operations.

## Features ✨

- **Recursive Traversal:** Walk through directories and subdirectories.
- **Tree-Like Display:** Show directories and files in a structured tree format with indentation.
- **Human-Readable File Sizes:** Display file sizes in bytes, KB, MB, etc.
- **Configurable Output:** 
  - Output directly to the console or save as an HTML file.
  - Limit the traversal depth using a command-line flag.
- **Easy to Extend:** Built with Go, making it simple to add additional features or customizations.

## Installation 🛠️

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/yourusername/go-directory-tree.git
   cd go-directory-tree
   ```

2. **Build the Application:**

   Make sure you have [Go installed](https://golang.org/dl/). Then run:

   ```bash
   go build -o treeexplorer main.go
   ```

   This will generate an executable named `treeexplorer`.

## Usage 💻

The tool accepts a directory path along with optional flags. Below are some examples:

### Basic Console Output

Display the directory tree of `/path/to/directory` in the console:

```bash
./treeexplorer /path/to/directory
```

### HTML Output

Generate an HTML file (`directory_tree.html`) containing the directory tree:

```bash
./treeexplorer -html /path/to/directory
```

### Limit Maximum Depth

Display only up to 2 levels deep:

```bash
./treeexplorer -maxdepth=2 /path/to/directory
```

### Combined Options

Generate an HTML file for `/path/to/directory` and limit the depth to 3:

```bash
./treeexplorer -html -maxdepth=3 /path/to/directory
```

## Command-Line Flags ⚙️

- `-html`  
  Output the directory tree to an HTML file instead of printing to the console.

- `-maxdepth`  
  Limit the directory traversal to a maximum depth. Use `-1` for no limit (default is `-1`).

## Code Structure 📂

- **`main.go`**  
  Contains the main logic for traversing directories, formatting the output, and handling configuration options.
- **Helper Functions:**  
  - **`humanSize`**: Converts file sizes from bytes to a human-readable format (e.g., KB, MB).

## Contributing 🤝

Contributions are welcome! Feel free to open issues or submit pull requests for enhancements, bug fixes, or additional features.

1. Fork the repository.
2. Create your feature branch: `git checkout -b feature/your-feature`.
3. Commit your changes: `git commit -am 'Add some feature'`.
4. Push to the branch: `git push origin feature/your-feature`.
5. Open a pull request.

## License 📄

Distributed under the MIT License. See [LICENSE](LICENSE) for more information.

## Acknowledgements 🙌

- Built with [Go](https://golang.org/).
- Inspired by various file management and tree visualization tools.

---

This README now includes emojis to make it more engaging and visually appealing. Feel free to adjust any sections to better fit your project's specifics!
