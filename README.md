# filecount-api

Stupid simple REST API to list count of files with
specified extension in specified directories.

This does not include any error handling and security stuff,
so please don't actually use this for any public-available endpoints.
I've created this small tool to monitor some directory sizes and 
display them in my [HomeAssistant](https://www.home-assistant.io/) instance.

## How to run
1. Build with Go `go build`
2. Create `filecount-api-config.json` (see example below)
3. Run `filecount-api [PORT]`
4. Consume `http://localhost:[PORT]`

## Example config
```json
{
  "Directories": [
    {
      "Path": "/var/www/foo",
      "Extension": ".xml",
      "FriendlyName": "foo"
    },
    {
      "Path": "/var/www/foo/bar",
      "Extension": ".xml",
      "FriendlyName": "bar"
    }
  ]
}
```

## Example response
```json
{
  "DirectoriesCounts": {
    "foo": {
      "Count": 2
    },
    "bar": {
      "Count": 3
    }
  }
}
```