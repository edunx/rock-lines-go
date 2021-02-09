# 说明
磐石系统lines文件读取

# 配置
```lua
    local file = rock.lines("data.txt")
    for line in file.line() do
        print(line)
    end
    file.close()
```