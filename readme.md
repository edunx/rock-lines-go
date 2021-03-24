# rock-lines-go 

## 说明
磐石系统lines文件读取

##函数说明 
- 函数: rock.lines
- 语法: rock.lines( string )
- 用法: rock.lines( "data.txt" )
```lua
    local li = rock.lines("resource/logs/data.txt")

    local line = file.line()
    while line do
        print(line)
        line = file.line()
    end
    
    file.close()
```

## 安装使用
```go
    import (
        lines "github.com/edunx/rock-lines-go"
    )
    
    //注入方法
    lines.LuaInjectApi(L , rock)
    
```