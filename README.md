# jsonlines2json
Tool to convert JSON lines to array of JSON objects

# Why

JSON line (http://jsonlines.org/) is used in logging and data processing, however, JSON line is not valid JSON so cannot be parsed by regular JSON parser. See http://jsonlines.org/:

> JSON Lines is a convenient format for storing structured data that may be processed one record at a time. It works well with unix-style text processing tools and shell pipelines. It's a great format for log files. It's also a flexible format for passing messages between cooperating processes.

This tool does nothing than conversion - its output is a valid JSON array, each of original JSON line is converted to an object of the whole JSON array, so then are able to use other JSON tools to do further processing, for example., ```jq```

3. For JSON log (AKA structured log) tool, see https://github.com/qiangyt/jog

# Features

- Read JSON line file, convert the content and output as an array of JSON objects, to stdout
- Able to ignore non-JSON prefix and suffix in a line but keep converting remaining of that line
- Able to ignore completely-non-JSON line
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
