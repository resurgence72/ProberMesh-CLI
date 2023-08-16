### ProberMesh-CLI 

> ProberMesh 分布式拨测工具的命令行 cli 工具
>
> ```sh
> go run main.go -h
> cli for probermesh
> 
> Usage:
>   probermesh-cli [command]
> 
> Available Commands:
>   completion  Generate the autocompletion script for the specified shell
>   help        Help about any command
>   target      show probermesh current target
>   task        send task for probermesh agent
>   upgrade     upgrade probermesh agent
> 
> Flags:
>   -h, --help                     help for probermesh-cli
>       --server.http.url string   probermesh server http url path; exclude http:// or https:// prefix (default "0.0.0.0:6001")
> 
> Use "probermesh-cli [command] --help" for more information about a command.
> ```
>
> #### 1. target 子命令
>
> ```sh
> go run main.go target -h
> show probermesh current target
> 
> Usage:
>   probermesh-cli target [flags]
> 
> Flags:
>   -h, --help                  help for target
>       --prober.type string    filter targets by prober type
>       --region.match string   filter targets by region (default ".+")
> 
> Global Flags:
>       --server.http.url string   probermesh server http url path; exclude http:// or https:// prefix (default "0.0.0.0:6001")
> ```
>
> #### 2. task 子命令
>
> ```sh
> go run main.go task -h
> send task for probermesh agent
> 
> Usage:
>   probermesh-cli task [flags]
> 
> Flags:
>       --cmd string            need agent exec command (default "ls -lh /tmp")
>   -h, --help                  help for task
>       --operator string       support =/!=/=~/!~ operators (default "=~")
>       --region.match string   need exec task region, regular expression support (default ".*")
> 
> Global Flags:
>       --server.http.url string   probermesh server http url path; exclude http:// or https:// prefix (default "0.0.0.0:6001")
> ```
>
> #### 3. upgrade 子命令
>
> ```sh
> go run main.go upgrade -h
> upgrade probermesh agent version
> 
> Usage:
>   probermesh-cli upgrade [flags]
> 
> Flags:
>       --download.url string   probermesh binary download url (default "https://github.com/resurgence72/ProberMesh/releases/download/v0.0.1/probermesh")
>       --force                 whether to force upgrade
>   -h, --help                  help for upgrade
>       --md5sum string         probermesh binary md5sum
>       --version string        need upgrade agent version (default "0.0.1")
> 
> Global Flags:
>       --server.http.url string   probermesh server http url path; exclude http:// or https:// prefix (default "0.0.0.0:6001")
> ```