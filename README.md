## cgncode
> 自动生成curd代码。
> 已经生成的代码，不会重复生成

<a href="https://github.com/jeffcail/cgncode/releases">
    <img src="https://img.shields.io/github/release/cgncode/releases.svg" alt="GitHub release">
</a>
<a href="https://github.com/jeffcail/cgncode/blob/master/LICENSE">
    <img src="https://img.shields.io/github/license/mashape/apistatus.svg" alt="license">
</a>

## 帮助
```shell
make help
```
```shell
Usage:
  cgncode [flags]
  cgncode [command]

Available Commands:
  code        A brief description of your command
  completion  Generate the autocompletion script for the specified shell
  dto         A brief description of your command
  handler     A brief description of your command
  help        Help about any command
  service     A brief description of your command

Flags:
  -h, --help   help for cgncode
```

## 使用
1. 生成 controller、service、dto代码
```shell
make code
```

2. 生成 controller代码
```shell
make h 
```

3. 生成 service代码
```shell
make s 
```

4. 生成 dto代码
```shell
make d
```
