<img src="assets/heart-balloon.svg" width="77">[^1]<br>
[![GitHub Pages](https://img.shields.io/badge/-GitHub_Pages-00A98F.svg?logo=github&style=flat)](https://ddddddo.github.io/gtree/)<br>
[![GitHub release](https://img.shields.io/github/release/ddddddO/gtree.svg?label=Release&color=darkcyan)](https://github.com/ddddddO/gtree/releases) [![Go Reference](https://pkg.go.dev/badge/github.com/ddddddO/gtree)](https://pkg.go.dev/github.com/ddddddO/gtree)<br>
[![License](https://img.shields.io/badge/License-BSD_2--Clause-orange.svg?color=darkcyan)](https://github.com/ddddddO/gtree/blob/master/LICENSE) [![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go#uncategorized)<br>
[![codecov](https://codecov.io/gh/ddddddO/gtree/branch/master/graph/badge.svg?token=JLGSLF33RH)](https://codecov.io/gh/ddddddO/gtree) [![Go Report Card](https://goreportcard.com/badge/github.com/ddddddO/gtree)](https://goreportcard.com/report/github.com/ddddddO/gtree) [![ci](https://github.com/ddddddO/gtree/actions/workflows/ci.yaml/badge.svg)](https://github.com/ddddddO/gtree/actions/workflows/ci.yaml)


<img src="assets/demo.gif"><br>

Using either Markdown or Programmatically to generate directory trees🌳 and directories🗂, and to verify directories🔍.
Provide CLI, Golang library and Web.

# Table of Contents
- **[Acknowledgments](https://github.com/ddddddO/gtree?tab=readme-ov-file#acknowledgments)**
- Features
	- **[Web](https://github.com/ddddddO/gtree?tab=readme-ov-file#web)**
	- **[CLI](https://github.com/ddddddO/gtree?tab=readme-ov-file#cli)**
	- Library - **[Markdown to tree structure](https://github.com/ddddddO/gtree?tab=readme-ov-file#library---markdown-to-tree-structure)**
	- Library - **[Programmable tree structure](https://github.com/ddddddO/gtree?tab=readme-ov-file#library---programmable-tree-structure)**
		- Recommended for developers!👍
- [Documents](https://github.com/ddddddO/gtree?tab=readme-ov-file#documents)
- [Process](https://github.com/ddddddO/gtree?tab=readme-ov-file#process)
- [Performance](https://github.com/ddddddO/gtree?tab=readme-ov-file#performance)
- [Test coverage](https://github.com/ddddddO/gtree?tab=readme-ov-file#test-coverage)

# Acknowledgments

## Thanks for providing very useful CLI for cloud storage tree output🤩🎉

Everyone is encouraged to use them!

### ⭐[_orangekame3/stree_](https://github.com/orangekame3/stree)
CLI for **Amazon S3** tree output.</br>
[_aws s3_](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/s3/index.html) command does not do what `tree` command does, but [**_stree_**](https://github.com/orangekame3/stree) command can display tree!

### ⭐[_owlinux1000/gcstree_](https://github.com/owlinux1000/gcstree)
CLI for **Google Cloud Storage** tree output.</br>
[_gcloud storage_](https://cloud.google.com/sdk/gcloud/reference/storage) command does not do what `tree` command does, but [**_gcstree_**](https://github.com/owlinux1000/gcstree) command can display tree!

## gtree packeage has been utilized for other tools as well🚀
I hope you will use these tools as well!

- [Tools](https://github.com/ddddddO/gtree/network/dependents)

# Web

https://ddddddo.github.io/gtree/

This page is that converts from Markdown to tree!<br>
This page calls a function that outputs tree. This function is a Go package compiled as WebAssembly.<br>
The symbols that can be used in Markdown are `*`, `-`, `+`, and `#`.<br>
Indentation represents hierarchy. The indentation can be whatever you specify, but use the same pattern.<br>
You can change the branches like in the image below.<br>
Also, once loaded, you can enjoy offline!<br>

![](assets/web_example.gif)

You can open it in your browser with
```console
$ gtree web
```

[source code](cmd/gtree-wasm/)


# CLI

## Installation

<pre>
<b>Go</b>
$ go install github.com/ddddddO/gtree/cmd/gtree@latest

<a href="https://github.com/aquaproj/aqua-registry/blob/main/pkgs/ddddddO/gtree/pkg.yaml"><b>aqua</b></a>
$ aqua g -i ddddddO/gtree

<a href="https://github.com/Homebrew/homebrew-core/blob/master/Formula/g/gtree.rb"><b>Homebrew</b></a>
$ brew install gtree

<a href="https://github.com/NixOS/nixpkgs/blob/master/pkgs/by-name/gt/gtree/package.nix"><b>Nix</b></a>
$ nix-env -i gtree
or
$ nix-shell -p gtree

<a href="https://github.com/macports/macports-ports/blob/master/sysutils/gtree/Portfile"><b>MacPorts</b></a>
$ port install gtree

<s><a href="https://aur.archlinux.org/packages/gtree"><b>AUR</b></a></s>
$ wip...

<b>Scoop</b>
$ scoop bucket add ddddddO https://github.com/ddddddO/scoop-bucket.git
$ scoop install ddddddO/gtree

<b>deb</b>
$ export GTREE_VERSION=X.X.X
$ curl -o gtree.deb -L https://github.com/ddddddO/gtree/releases/download/v$GTREE_VERSION/gtree_$GTREE_VERSION-1_amd64.deb
$ dpkg -i gtree.deb

<b>rpm</b>
$ export GTREE_VERSION=X.X.X
$ yum install https://github.com/ddddddO/gtree/releases/download/v$GTREE_VERSION/gtree_$GTREE_VERSION-1_amd64.rpm

<b>apk</b>
$ export GTREE_VERSION=X.X.X
$ curl -o gtree.apk -L https://github.com/ddddddO/gtree/releases/download/v$GTREE_VERSION/gtree_$GTREE_VERSION-1_amd64.apk
$ apk add --allow-untrusted gtree.apk

<a href="https://github.com/ddddddO/gtree/pkgs/container/gtree"><b>Docker</b></a>
$ docker pull ghcr.io/ddddddo/gtree:latest
$ docker run ghcr.io/ddddddo/gtree:latest template | docker run -i ghcr.io/ddddddo/gtree:latest output
gtree
├── cmd
│   └── gtree
│       └── main.go
├── testdata
│   ├── sample1.md
│   └── sample2.md
├── Makefile
└── tree.go
</pre>

### etc

**download binary from [here](https://github.com/ddddddO/gtree/releases).**

## Usage

```console
$ gtree --help
NAME:
   gtree - This CLI uses Markdown to generate directory trees and directories itself, and also verifies directories.
           The symbols that can be used in Markdown are '-', '+', '*', and '#'.
           Within Markdown, indentation represents hierarchy. The indentation can be whatever you specify, but use the same pattern.

USAGE:
   gtree [global options] command [command options] [arguments...]

VERSION:
   1.10.2 / revision 85520a1

COMMANDS:
   output, o, out     Outputs tree from markdown.
                      Let's try 'gtree template | gtree output'.
   mkdir, m           Makes directories and files from markdown. It is possible to dry run.
                      Let's try 'gtree template | gtree mkdir -e .go -e .md -e Makefile'.
   verify, vf         Verifies tree structure represented in markdown by comparing it with existing directories.
                      Let's try 'gtree template | gtree verify'.
   template, t, tmpl  Outputs markdown template. Use it to try out gtree CLI.
   web, w, www        Opens "Tree Maker" in your browser and shows the URL in terminal.
   version, v         Prints the version.
   help, h            Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

### *Output* subcommand
```console
$ gtree output --help
NAME:
   gtree output - Outputs tree from markdown.
                  Let's try 'gtree template | gtree output'.

USAGE:
   gtree output [command options] [arguments...]

OPTIONS:
   --file value, -f value               specify the path to markdown file. (default: stdin)
   --massive, -m                        set this option when there are very many blocks of markdown. (default: false)
   --massive-timeout value, --mt value  set this option if you want to set a timeout. (default: 0s)
   --format value                       set this option when specifying output format. "json", "yaml", "toml"
   --watch, -w                          follow changes in markdown file. (default: false)
   --help, -h                           show help
```

#### Try it!

```console
$ gtree template
- gtree
        - cmd
                - gtree
                        - main.go
        - testdata
                - sample1.md
                - sample2.md
        - Makefile
        - tree.go
$ gtree template | gtree output
gtree
├── cmd
│   └── gtree
│       └── main.go
├── testdata
│   ├── sample1.md
│   └── sample2.md
├── Makefile
└── tree.go
```

Other pattern.

```
├── gtree output -f testdata/sample1.md
├── cat testdata/sample1.md | gtree output -f -
└── cat testdata/sample1.md | gtree output
```


#### Usage other than representing a directory.

```console
$ cat testdata/sample2.md | gtree output
k8s_resources
├── (Tier3)
│   └── (Tier2)
│       └── (Tier1)
│           └── (Tier0)
├── Deployment
│   └── ReplicaSet
│       └── Pod
│           └── container(s)
├── CronJob
│   └── Job
│       └── Pod
│           └── container(s)
└── (empty)
    ├── DaemonSet
    │   └── Pod
    │       └── container(s)
    └── StatefulSet
        └── Pod
            └── container(s)
```


#### Multiple roots

```console
$ cat testdata/sample6.md | gtree output
Artiodactyla
├── Artiofabula
│   ├── Cetruminantia
│   │   ├── Whippomorpha
│   │   │   ├── Hippopotamidae
│   │   │   └── Cetacea
│   │   └── Ruminantia
│   └── Suina
└── Tylopoda
Carnivora
├── Feliformia
└── Caniformia
    ├── Canidae
    └── Arctoidea
        ├── Ursidae
        └── x
            ├── Pinnipedia
            └── Musteloidea
                ├── Ailuridae
                └── x
                    ├── Mephitidae
                    └── x
                        ├── Procyonidae
                        └── Mustelidae
```

#### Output JSON

<details><summary>see</summary>

```console
$ cat testdata/sample5.md | gtree output --format json | jq
{
  "value": "a",
  "children": [
    {
      "value": "i",
      "children": [
        {
          "value": "u",
          "children": [
            {
              "value": "k",
              "children": null
            },
            {
              "value": "kk",
              "children": null
            }
          ]
        },
        {
          "value": "t",
          "children": null
        }
      ]
    },
    {
      "value": "e",
      "children": [
        {
          "value": "o",
          "children": null
        }
      ]
    },
    {
      "value": "g",
      "children": null
    }
  ]
}
```

</details>

#### Output YAML

<details><summary>see</summary>

```console
$ cat testdata/sample5.md | gtree output --format yaml
value: a
children:
- value: i
  children:
  - value: u
    children:
    - value: k
      children: []
    - value: kk
      children: []
  - value: t
    children: []
- value: e
  children:
  - value: o
    children: []
- value: g
  children: []
```

</details>

#### Output TOML

<details><summary>see</summary>

```console
$ cat testdata/sample5.md | gtree output --format toml
value = 'a'
[[children]]
value = 'i'
[[children.children]]
value = 'u'
[[children.children.children]]
value = 'k'
children = []
[[children.children.children]]
value = 'kk'
children = []

[[children.children]]
value = 't'
children = []

[[children]]
value = 'e'
[[children.children]]
value = 'o'
children = []

[[children]]
value = 'g'
children = []

```

</details>


### *Mkdir* subcommand

```console
$ gtree mkdir --help
NAME:
   gtree mkdir - Makes directories and files from markdown. It is possible to dry run.
                 Let's try 'gtree template | gtree mkdir -e .go -e .md -e Makefile'.

USAGE:
   gtree mkdir [command options] [arguments...]

OPTIONS:
   --file value, -f value                                       specify the path to markdown file. (default: stdin)
   --dry-run, -d                                                dry run. detects node that is invalid for directory generation. the order of the output and made directories does not always match. (default: false)
   --extension value, -e value [ --extension value, -e value ]  set this option if you want to create file instead of directory. for example, if you want to generate files with ".go" extension: "-e .go"
   --target-dir value                                           set this option if you want to specify the directory you want to make directory. (default: current directory)
   --help, -h                                                   show help
```

#### Try it!

```console
$ gtree template
- gtree
        - cmd
                - gtree
                        - main.go
        - testdata
                - sample1.md
                - sample2.md
        - Makefile
        - tree.go
$ gtree template | gtree mkdir
$ tree gtree/
gtree/
├── cmd
│   └── gtree
│       └── main.go
├── Makefile
├── testdata
│   ├── sample1.md
│   └── sample2.md
└── tree.go

8 directories, 0 files
```

#### *make directories and files*
```console
$ gtree template
- gtree
        - cmd
                - gtree
                        - main.go
        - testdata
                - sample1.md
                - sample2.md
        - Makefile
        - tree.go
$ gtree template | gtree mkdir -e .go -e .md -e Makefile
$ tree gtree/
gtree/
├── cmd
│   └── gtree
│       └── main.go
├── Makefile
├── testdata
│   ├── sample1.md
│   └── sample2.md
└── tree.go

3 directories, 5 files
```

#### *dry run*
Does not create a file and directory.

```console
$ gtree template | gtree mkdir --dry-run -e .go -e .md -e Makefile
gtree
├── cmd
│   └── gtree
│       └── main.go
├── testdata
│   ├── sample1.md
│   └── sample2.md
├── Makefile
└── tree.go

4 directories, 5 files
```

![](assets/cli_mkdir_dryrun.png)


Any invalid file or directory name will result in an error.

```console
$ gtree mkdir --dry-run <<EOS
- root
  - aa
  - bb
    - b/b
EOS
invalid node name: b/b
```

```console
$ gtree mkdir --dry-run <<EOS
- /root
  - aa
  - bb
    - bb
EOS
invalid node name: /root
```

### *Verify* subcommand
```console
$ gtree verify --help
NAME:
   gtree verify - Verifies tree structure represented in markdown by comparing it with existing directories.
                  Let's try 'gtree template | gtree verify'.

USAGE:
   gtree verify [command options] [arguments...]

OPTIONS:
   --file value, -f value  specify the path to markdown file. (default: stdin)
   --target-dir value      set this option if you want to specify the directory you want to verify. (default: current directory)
   --strict                set this option if you want strict directory match validation. (default: non strict)
   --help, -h              show help
```

#### Try it!

```console
$ tree example
example
├── README.md
├── find_pipe_programmable-gtree
│   ├── README.md
│   ├── go.mod
│   ├── go.sum
│   └── main.go
├── go-list_pipe_programmable-gtree
│   ├── README.md
│   ├── go.mod
│   ├── go.sum
│   └── main.go
├── like_cli
│   ├── adapter
│   │   ├── executor.go
│   │   └── indentation.go
│   └── main.go
├── noexist
│   └── xxx
└── programmable
    └── main.go

6 directories, 14 files
$ cat testdata/sample9.md
- example
        - README.md
        - find_pipe_programmable-gtree
                - README.md
                - go.mod
                - go.sum
                - main.go
        - go-list_pipe_programmable-gtree
                - README.md
                - go.mod
                - go.sum
                - main.go
        - like_cli
                - adapter
                        - executor.go
                        - indentation.go
                - main.go
                - kkk
        - programmable
                - main.go
$ cat testdata/sample9.md | gtree verify --strict
Extra paths exist:
        example/noexist
        example/noexist/xxx
Required paths does not exist:
        example/like_cli/kkk
```

inspired by [mactat/framed](https://github.com/mactat/framed) !

# Library - Markdown to tree structure

## Installation

Go version requires 1.24 or later.

```console
$ go get github.com/ddddddO/gtree
```

## Usage

The symbols that can be used in Markdown are `*`, `-`, `+`, and `#`.

|Function|Description|Available optional functions|
|--|--|--|
|*[OutputFromMarkdown](https://pkg.go.dev/github.com/ddddddO/gtree#OutputFromMarkdown)*|can output trees|WithBranchFormatIntermedialNode<br>WithBranchFormatLastNode<br>WithEncodeJSON<br>WithEncodeTOML<br>WithEncodeYAML<br>WithMassive|
|*[MkdirFromMarkdown](https://pkg.go.dev/github.com/ddddddO/gtree#MkdirFromMarkdown)*|can create directories|WithTargetDir<br>WithFileExtensions<br>WithDryRun<br>WithMassive|
|*[VerifyFromMarkdown](https://pkg.go.dev/github.com/ddddddO/gtree#VerifyFromMarkdown)*|can output the difference between markdown and directories|WithTargetDir<br>WithStrictVerify<br>WithMassive|
|*[WalkFromMarkdown](https://pkg.go.dev/github.com/ddddddO/gtree#WalkFromMarkdown)*|can execute user-defined function while traversing tree structure recursively|WithBranchFormatIntermedialNode<br>WithBranchFormatLastNode<br>WithMassive|

# Library - Programmable tree structure

> [!NOTE]
> There are sample repositories that use gtree library.<br>
> See [here](https://github.com/ddddddO/gtree/blob/master/example/README.md) for details.

## Installation

Go version requires 1.24 or later.

```console
$ go get github.com/ddddddO/gtree
```

## Usage

|Function|Description|Available optional functions|
|--|--|--|
|*[OutputFromRoot](https://pkg.go.dev/github.com/ddddddO/gtree#OutputFromRoot)*|can output tree|WithBranchFormatIntermedialNode<br>WithBranchFormatLastNode<br>WithEncodeJSON<br>WithEncodeTOML<br>WithEncodeYAML|
|*[MkdirFromRoot](https://pkg.go.dev/github.com/ddddddO/gtree#MkdirFromRoot)*|can create directories|WithTargetDir<br>WithFileExtensions<br>WithDryRun|
|*[VerifyFromRoot](https://pkg.go.dev/github.com/ddddddO/gtree#VerifyFromRoot)*|can output the difference between tree you composed and directories|WithTargetDir<br>WithStrictVerify|
|*[WalkFromRoot](https://pkg.go.dev/github.com/ddddddO/gtree#WalkFromRoot)*|can execute user-defined function while traversing tree structure recursively|WithBranchFormatIntermedialNode<br>WithBranchFormatLastNode|
|*[WalkIterFromRoot](https://pkg.go.dev/github.com/ddddddO/gtree#WalkIterFromRoot)*|it returns each node resulting from a recursive traversal of the tree structure, so you can process on each node|WithBranchFormatIntermedialNode<br>WithBranchFormatLastNode|

# Process

> [!NOTE]
> This process is for the Massive Roots mode.

## e.g. [*gtree/pipeline_tree.go*](https://github.com/ddddddO/gtree/blob/master/pipeline_tree.go)

<image src="assets/process.svg" width=100%>


# Performance

> [!WARNING]
> The following benchmarks are for simple implementation before iterator implementation. The simple implementation is now an iterator implementation, and the performance of the simple implementation is better.
> Depends on the environment.

<details><summary>old data</summary>

- Comparison simple implementation and pipeline implementation.
- In the case of few Roots, simple implementation is faster in execution!
	- Use this one by default.
- However, for multiple Roots, pipeline implementation execution speed tends to be faster💪✨
	- In the CLI, it is available by specifying `--massive`.
	- In the Go program, it is available by specifying `WithMassive` func.

<image src="assets/performance.svg" width=100%>

<details><summary>Benchmark log</summary>

## Simple implementation
```console
11:19:22 > go test -benchmem -bench Benchmark -benchtime 100x benchmark_simple_test.go
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-7200U CPU @ 2.50GHz
BenchmarkOutput_singleRoot-4                 100             35375 ns/op           13856 B/op        171 allocs/op
BenchmarkOutput_tenRoots-4                   100            200540 ns/op           72920 B/op       1597 allocs/op
BenchmarkOutput_fiftyRoots-4                 100            730156 ns/op          569851 B/op       7919 allocs/op
BenchmarkOutput_hundredRoots-4               100           1706493 ns/op         1714260 B/op      15820 allocs/op
BenchmarkOutput_fiveHundredsRoots-4          100          16412090 ns/op        32245140 B/op      79022 allocs/op
BenchmarkOutput_thousandRoots-4              100          55142492 ns/op        120929674 B/op    158025 allocs/op
BenchmarkOutput_3000Roots-4                  100         489121246 ns/op        1035617527 B/op   474029 allocs/op
BenchmarkOutput_6000Roots-4                  100        1613641261 ns/op        4087694372 B/op   948033 allocs/op
BenchmarkOutput_10000Roots-4                 100        3913090646 ns/op        11293191221 B/op         1580035 allocs/op
PASS
ok      command-line-arguments  614.944s
```

## Pipeline implementation
```console
11:29:43 > go test -benchmem -bench Benchmark -benchtime 100x benchmark_pipeline_test.go
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-7200U CPU @ 2.50GHz
BenchmarkOutput_pipeline_singleRoot-4                100            188706 ns/op           24236 B/op        300 allocs/op
BenchmarkOutput_pipeline_tenRoots-4                  100            367758 ns/op          115970 B/op       2186 allocs/op
BenchmarkOutput_pipeline_fiftyRoots-4                100            947879 ns/op          542188 B/op      10592 allocs/op
BenchmarkOutput_pipeline_hundredRoots-4              100           1711537 ns/op         1099636 B/op      21094 allocs/op
BenchmarkOutput_pipeline_fiveHundredsRoots-4         100           6892261 ns/op         5524905 B/op     105107 allocs/op
BenchmarkOutput_pipeline_thousandRoots-4             100          13100335 ns/op        11225942 B/op     210115 allocs/op
BenchmarkOutput_pipeline_3000Roots-4                 100          40694497 ns/op        33399766 B/op     630142 allocs/op
BenchmarkOutput_pipeline_6000Roots-4                 100          85807944 ns/op        66974524 B/op    1260171 allocs/op
BenchmarkOutput_pipeline_10000Roots-4                100         151486713 ns/op        113908462 B/op   2100208 allocs/op
PASS
ok      command-line-arguments  30.670s
```

</details>

</details>

# Documents
- [GoDoc](https://pkg.go.dev/github.com/ddddddO/gtree)

## English
- [Want to output a tree in Go?](https://medium.com/@allowing_whip_guineapig_430/want-to-output-a-tree-in-go-1851f9fc9900)
- [Generate directory trees🌳 and the directories itself📁 using Markdown or Programmatically.](https://www.reddit.com/r/commandline/comments/146nk54/generate_directory_trees_and_the_directories/)

## Japanese
- [Goでtreeを表現する](https://zenn.dev/ddddddo/articles/8cd85c68763f2e)
- [Markdown形式の入力からtreeを出力するCLI/Web](https://zenn.dev/ddddddo/articles/ad97623a004496)
- [Markdown形式の入力からファイル/ディレクトリを生成するCLI/Goパッケージ](https://zenn.dev/ddddddo/articles/460d12e8c07763)
- [盆栽 (節目の記事)](https://zenn.dev/openlogi/articles/f6cc91ac413c8f)
- [感想](https://scrapbox.io/ddddddo/useful_tools)

# Test coverage
![treemap](/assets/test_treemap.svg)

...generated by [nikolaydubina/go-cover-treemap](https://github.com/nikolaydubina/go-cover-treemap) !

# Stargazers over time
[![Stargazers over time](https://starchart.cc/ddddddO/gtree.svg?variant=adaptive)](https://starchart.cc/ddddddO/gtree)

[^1]: Gopher retrieved from [egonelbre/gophers](https://github.com/egonelbre/gophers) !