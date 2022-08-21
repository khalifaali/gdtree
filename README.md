# gdtree
## Description
Go module built to enhance readability of `go mod graph` by printing tree-like output.

## How to install
Configure `$GOBIN` directory on the terminal.<br>
`go env -w GOBIN=/path/to/your/bin`

Install go module.<br>
`go install github.com/khalifaali/gdtree@latest`

## Example
### Before
```
[+]go mod graph
github.com/mewkiz/flac github.com/go-audio/audio@v1.0.0
github.com/mewkiz/flac github.com/go-audio/wav@v1.0.0
github.com/mewkiz/flac github.com/icza/bitio@v1.0.0
github.com/mewkiz/flac github.com/mewkiz/pkg@v0.0.0-20190919212034-518ade7978e2
github.com/mewkiz/flac github.com/pkg/errors@v0.8.1
github.com/go-audio/wav@v1.0.0 github.com/go-audio/audio@v1.0.0
github.com/go-audio/wav@v1.0.0 github.com/go-audio/riff@v1.0.0
github.com/icza/bitio@v1.0.0 github.com/icza/mighty@v0.0.0-20180919140131-cfd07d671de6
github.com/mewkiz/pkg@v0.0.0-20190919212034-518ade7978e2 github.com/d4l3k/messagediff@v1.2.2-0.20190829033028-7e0a312ae40b
github.com/mewkiz/pkg@v0.0.0-20190919212034-518ade7978e2 github.com/pkg/errors@v0.8.1
github.com/mewkiz/pkg@v0.0.0-20190919212034-518ade7978e2 golang.org/x/image@v0.0.0-20190220214146-31aff87c08e9
github.com/mewkiz/pkg@v0.0.0-20190919212034-518ade7978e2 golang.org/x/net@v0.0.0-20190213061140-3a22650c66bd
golang.org/x/image@v0.0.0-20190220214146-31aff87c08e9 golang.org/x/text@v0.3.0
```

### After
Call `gdtree` in any directory with a go.mod, and go.sum file present.
```
[+]gdtree
github.com/mewkiz/flac
├── github.com/go-audio/audio@v1.0.0
├── github.com/go-audio/wav@v1.0.0
│    │    ├── github.com/go-audio/audio@v1.0.0
│    │    └── github.com/go-audio/riff@v1.0.0
├── github.com/icza/bitio@v1.0.0
│    │    └── github.com/icza/mighty@v0.0.0-20180919140131-cfd07d671de6
├── github.com/mewkiz/pkg@v0.0.0-20190919212034-518ade7978e2
│    │    ├── github.com/d4l3k/messagediff@v1.2.2-0.20190829033028-7e0a312ae40b
│    │    ├── github.com/pkg/errors@v0.8.1
│    │    ├── golang.org/x/image@v0.0.0-20190220214146-31aff87c08e9
│    │    │    └── golang.org/x/text@v0.3.0
│    │    └── golang.org/x/net@v0.0.0-20190213061140-3a22650c66bd
└── github.com/pkg/errors@v0.8.1
```
