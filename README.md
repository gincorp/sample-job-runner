Custom Jobs
==

The default set of jobs in `github.com/gincorp/gin/taskmanager` are sparse and few. This is for many reasons:

* Domain specific and specialist code aren't appropriate for every use
* Basic tasks may be a little too basic for some uses
* Overloading the basic client with a multitude of tasks can be confusing

Thus; we provide an interface to adding tasks via the package.

Usage
--

```golang
n := node.NewNode(*amqpURI, "", "job")
jobManager := taskmanager.NewJobManager()
jobManager.AddJob("write-to-file", dumpToFile)
n.TaskManager = jobManager
```

In this case `dumpToFile` is a function:

```golang
func dumpToFile(jn taskmanager.JobNotification) (output map[string]interface{}, err error) {
    body := []byte(jn.Context["body"])
    path := jn.Context["path"]

    err = ioutil.WriteFile(path, body, 0644)
    return
}
```

And is called via the task definition:

```json
{
		"name": "Writer",
		"type": "write-to-file",
		"context": {
			"body": "Hello World!",
			"path": "{{.Defaults.path}}"
		}
}
```

Demo Code
--

See `main.go` in this project

Licence
--

MIT License

Copyright (c) 2016 jspc

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
