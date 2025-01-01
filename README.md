# to-pdf

A simple web service that converts Microsoft Word documents (.doc, .docx) to PDF format using LibreOffice.

## Features

- Web-based interface for file upload
- Supports .doc and .docx file formats
- Automatic conversion to PDF
- Direct download of converted files

## Prerequisites

Before running this application, make sure you have the following installed:

- Go (1.16 or later)
- LibreOffice (for document conversion)

## Installation

1. Clone this repository:
```bash
git clone https://github.com/emre-guler/to-pdf.git
cd to-pdf
```

2. Install dependencies:
```bash
go mod download
```

3. Create a directory for file uploads:
```bash
mkdir files
```

## Usage

1. Start the server:
```bash
go run main.go
```

2. Open your web browser and navigate to `http://localhost:9091`

3. Upload a Word document through the web interface

4. The converted PDF will be automatically downloaded

## Configuration

The application uses the following default settings:

- Port: 9091 (can be changed using the PORT environment variable)
- Upload folder: `./files`
- Supported file extensions: .doc, .docx

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request. 