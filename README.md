# Lua Includer
Lua Includer is a basic pre-processor for lua, with Lua Includer you can use the macro `--include` similar with c/c++

For example:

File 1 (epic_variable.lua)<br>
```lua
local str = "Lua"
```

File 2 (main.lua)<br>
```lua
--include "./epic_variable.lua"

print(str);
```

Output<br>
```lua
local str = "Lua"

print(str);
```

# How To Use
to use Lua Includer (it is recommended to put it as an environment variable but it is not mandatory) you just need to use the command `lual <main_file_path>`, 
using the above example just use `lual main.lua`
