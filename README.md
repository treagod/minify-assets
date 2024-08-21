# minify-assets-go

`minify-assets-go` is a simple Go script designed to minify CSS and JavaScript files in a given directory. It replaces the original asset files with their minified versions, helping to reduce file sizes and improve load times for web applications.

It's a simple wrapper around the [minify tooling](https://github.com/tdewolff/minify){:target="_blank"}.

## Features

- **Minifies CSS and JavaScript**: Automatically minifies all `.css` and `.js` files in the specified directory.
- **Recursive Directory Traversal**: Processes all asset files within the specified root directory and its subdirectories.

## Installation

To use this script, you need to have Go installed on your system. Clone the repository and initialize the Go module:

```bash
git clone https://github.com/treagod/minify-assets.git
cd minify-assets
go mod tidy
```

## Usage

To run the script, provide the root directory containing your assets as an argument:

```go
go run cmd/minify_assets/minify-assets-go.go /path/to/your/assets
```

This will traverse the specified directory, minify all .css and .js files, and overwrite the original files with their minified versions.

### License

This project is licensed under the MIT License. See the LICENSE file for more details.

#### Contributing

Contributions are welcome! Feel free to open issues or submit pull requests to improve the tool.
