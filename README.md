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
Not-yet-set keys are treated as empty strings, for the purposes of e.g. `GET` and `APPEND`.
- `SET key value` -> JSON `null`
- `GET key` -> JSON "" if the `key` does not exist, or its value if it does exist
- `APPEND key value` -> the new value of the string after append.

## Hashes (of strings)
- `HSET hashname key value` -> JSON `null`
- `HGET hashname key` -> JSON `null` if the `hashname` does not exist or `keyname` does not exist in that hash, or `"thevalue"` if it does
- `HLEN hashname` -> number of entries in "hashname"

## Sets
Sets.
- `SADD keyname value1 [value2 [value3 [...]]` -> number of new elements that got added. e.g. if "xyz" is already a member of "key", then `SADD key xyz` -> 0. Nonexistent sets are treated as empty.
- `SREM keyname value1 [value2 [value3 [...]]` -> number of values actually removed. e.g. if "xyz" is not a member of "key", then `SREM key xyz` -> 0.
- `SISMEMBER key value` -> JSON `true` if "value" is a member of "key", JSON `false` otherwise
- `SMEMBERS key` -> JSON array of all the members in "key"
- `SCARD key` -> cardinality of "key"
- `SRANDMEMBER key` -> a random member of "key"
- `SPOP key` -> a random member of "key" (and removes that member)


## Counters
Counters are integers that automatically decrement themselves after a specified time increment. This is useful for e.g. knowing the number of visitors to your site in the last 5 minutes.

## Lists (of strings)
Linked lists (bidirectional).
- `LPUSH key value` -> the length of "key" after the value has been added to its left side.
- `RPUSH key value` -> the length of "key" after the value has been added to its right side.
- `LPOP key` -> the leftmost element of "key" (and removes it), or empty string if "key" is empty.
- `RPOP key` -> the rightmost element of "key" (and removes it), or empty string if "key" is empty.
- `LLEN key` -> the length of "key".

