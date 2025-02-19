# Readme
Babel is a high-performance "dummy data load" generator, which mimics the schema of supported real-world data sources for testing, analysis, and security application development. The current version (v1.0, released 2025-02-17) generates logs in CSV, JSON, or RAW format with support for file size limits and rate limiting (coming soon). 

The current version was built as a utility to support a security related homelab project, where volume, depth and closeness of data are required to adequately benchmark bandwidth, throughput and compute performance. However, due Babel's high performance and low compute requirement, it became of a project of its own. Currently, I'm working to enahnce the capabilities of Babel to offer a more broad range of data srouces, and to be run as a public cloud resource. 

Check out [the project writeup on medium](https://t2-security.com/babel-a-high-performance-dummy-data-load-generator-214b153d2cf8). 

#### Supported Datasources
- Fotrinet FortiOS - 48 Log Message Fields
- Zscaler ZIA DNS Logs - 17 Log Message Fields (in development)
- MISP - 30 Log Message Fields (future development)
#### Features Currently in Development
- Enhanced rate limiting to more closely mimic Events per Second(EPS) rate of continuous data sources.
- Ability to set environmental variables from the CLI for the session; such as URI's, rate limites and paths by data source.



## Example Usage

### Optinas & Examples Menu
This menu menu prints information on flags adn outlines example usage. It's accessed with the following:
```
./bable --help
```
The above command will return the following: 
```
Options:
  -n       <int>    Number of log events to generate (default: 10)
  -fmt  <string> Output format: csv, json, raw (default: csv)
  -f    <string> Output filename for CSV (default: logs.csv)
  -r    <int>    Maximum events per second (0 = unlimited, default: 0)
  -vB      <int>    Maximum volume in bytes before stopping (overrides -n)
  -vT      <float>  Maximum volume in terabytes before stopping (overrides -n)
  --help            Show this help message and exit

Examples:
  go run main.go -n 100 -format json
  go run main.go -vB 1000000 -rate 5 -format csv
  go run main.go -vT 1 -format raw
```

### Generate Logs by Number
The following command uses the -n flag to specify that 20000 events will be created, with the -fmt 
```
./babel -n 20000 -fmt json
```

### Generate Logs by Volume
The follow following uses -vT 500 to to specify that total volume of data to be produced will not exceed 500TB of space. Following is the -fmt flag indicating the desired output is in CSV format.
```
./babel -vT 500 -fmt csv
```

## Installation 

### 1). Clone the repositroy
```
git clone https://github.com/yourusername/babel.git
cd babel

```

### 2). Install dependencies
```
go mod tidy
```

### 3). Build the dependencies
```
go build -o babel main.go
```
Now, you can run ./babel as a standalone executable!
