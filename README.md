# jsonlines2json
Tool to convert JSON lines to array of JSON objects

# Background

1. What's JSON Lines

JSON line is lines of JSON but not valid JSON. See http://jsonlines.org/

> JSON Lines is a convenient format for storing structured data that may be processed one record at a time. It works well with unix-style text processing tools and shell pipelines. It's a great format for log files. It's also a flexible format for passing messages between cooperating processes.

2. This tool doesn't do anything other than conversion. Since the output is valid JSON (array), there're many tools we could use to do further processing upon the output, for example., ```jq```

3. For JSON log (AKA structured log) tool, see https://github.com/qiangyt/jog

# Features

- Read JSON line file, output converted array of JSON objects to stdout
- Ignore non-JSON prefix in a line but keep converting remaining lines
- Ignore completely-non-JSON line
- Outputs an empty JSON array, even no lines input or parsing failure

## Usage:
  Download the executable binary to $PATH. For ex.

  ```shell
     curl -L https://github.com/qiangyt/jsonlines2json/releases/download/v0.9.0/jl2ja.darwin -o /usr/local/bin/jl2ja
     chmod +x /usr/local/bin/jl2ja
  ```

   * Convert a JSON line file: `jl2ja sample.json.log`

   * Pretty-print a JSON line file: `jl2ja sample.json.log | jq`

## Build

   *  Install GOLANG

   *  In current directory, run `./build.sh`

## License

[MIT](/LICENSE)
