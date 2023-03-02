jpath
=====

Find the path of a key / value in a JSON hierarchy *easily*.

The original name of the project was "JSON Path", but
it turned out that there is a specification called like this,
so I renamed it to **jpath**. This project has nothing
to do with the JSON path specification.

(It turned out that IBM has a query language called [JPath](https://www.ibm.com/docs/en/dsm?topic=protocol-jpath). This project has nothing to do with that.)

Motivation
----------

When working with big and nested JSON files, sometimes
it's very difficult to figure out the path of a key. You
open the JSON file in a text editor, find the key / value
pair you need, but then it can be a pain to get the full path of the key.

Installation
------------

    $ go install github.com/jabbalaci/go-jsonpath/cmd/jpath@latest

Example
-------

Consider the following JSON file:

```json
{
    "a": 1,
    "b": {
        "c": 2,
        "friends": [
            {
                "best": "Alice"
            },
            {
                "second": "Bob"
            },
            [5, 6, 7],
            [
                {"one": 1},
                {"two": 2}
            ]
        ]
    }
}
```

jpath will traverse it recursively and print every
path / value pair:

```bash
$ jpath samples/short.json
root["a"] => 1
root["b"]["c"] => 2
root["b"]["friends"][0]["best"] => "Alice"
root["b"]["friends"][1]["second"] => "Bob"
root["b"]["friends"][2][0] => 5
root["b"]["friends"][2][1] => 6
root["b"]["friends"][2][2] => 7
root["b"]["friends"][3][0]["one"] => 1
root["b"]["friends"][3][1]["two"] => 2
```

The idea is to combine its usage with the Unix command
`grep`. For instance, what's the path of the key that
leads to Bob?

```bash
$ jpath samples/short.json | grep -i bob
root["b"]["friends"][1]["second"] => "Bob"
```

Then simply paste it in your application (Python example):

```python
>>> import json
>>>
>>> f = open("samples/short.json")
>>> d = json.load(f)
>>> d
{'a': 1, 'b': {'c': 2, 'friends': [{'best': 'Alice'}, {'second': 'Bob'}, [5, 6, 7], [{'one': 1}, {'two': 2}]]}}
>>> d["b"]["friends"][1]["second"]
'Bob'
>>>
```

Related Works
-------------

* [JSON-path](https://github.com/jabbalaci/JSON-path): my original Python implementation. This project is the same,
re-written in Go.
* [gron](https://github.com/tomnomnom/gron): gron makes
JSON greppable. gron can do a lot more than jpath, but
when I made jpath I didn't know about gron.
