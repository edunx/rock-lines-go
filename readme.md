# 说明
磐石系统lines文件读取

# 配置
```lua
    local file = rock.lines("data.txt")

    local line = file.line()
    while line do
        print(line)
        line = file.line()
    end
    
    file.close()
```