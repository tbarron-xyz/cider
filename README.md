# cider
A lightweight in-memory data store. (`redis` spelled backwards.) Connections are handled via websockets.

# Command line flags
- `-port 1234` (default: `6969`) specifies the port to listen on
- `-v` (default: false), if enabled, will print all incoming messages to the console

# Usage
A message can consist of either 
- a single command (e.g. `SET keyname value`), in which case the response will be a status-response JSON object (see below), or
- a JSON array of commands in string format (`["SET keyname value1", "GET keyname"]`, in which case the response will be a status-response object where the inner response is an array of status-response objects.

Arguments after the command name can be enclosed by quotes, in case e.g. they contain spaces. Acceptable quote characters for this purpose are double quote `"`, single quote `'`, or backtick <code>`</code>, inside of which you can freely use the other two quote characters without ending the quote block, or use a backslash to escape a single character. See the comments in [strparse.go](https://github.com/tbarron-xyz/cider/blob/master/strparse.go) for further documentation on this syntax.

## Examples
- `GET keyname` -> `{"status": "success", "response": "the value stored at keyname"}`
- `["SET keyname value1", "GET keyname"]` -> `{"status": "success", "response": [{"status": "success", "response": null}, {"status": "success", "response": "value1"}]}`

# Commands
(documentation in progress)
## Strings
- `SET keyname thevalue` -> JSON `null`
- `GET keyname` -> JSON `null` if the `keyname` does not exist, or `"thevalue"` if it does exist

## Hashes (of strings)
- `HSET hashname keyname thevalue` -> JSON `null`
- `HGET hashname keyname` -> JSON `null` if the `hashname` does not exist or `keyname` does not exist in that hash, or `"thevalue"` if it does

## Sets

## Counters
Counters are integers that automatically decrement themselves after a specified time increment. This is useful for e.g. knowing the number of visitors to your site in the last 5 minutes.

## Lists (of strings)
Linked lists (bidirectional).
