package lines

import (
	"bufio"
	"github.com/edunx/lua"
	"os"
)

func (self *Lines) Line(L *lua.LState , args *lua.Args) lua.LValue {
	if self.scanner.Scan() {
		return lua.LString(self.scanner.Text())
	}
	return lua.LNil
}

func (self *Lines) Close(L *lua.LState , args *lua.Args) lua.LValue {
	self.fd.Close()
	return lua.LNil
}

func (self *Lines) Index(L *lua.LState , key string) lua.LValue {
	if key == "line"  { return lua.NewGFunction( self.Line  ) }
	if key == "close" { return lua.NewGFunction( self.Close ) }
	return lua.LNil
}

func createLinesLightUserData(L *lua.LState  , args *lua.Args ) lua.LValue {
	filename := args.CheckString(L , 1)

	file , err := os.Open( filename )
	if err != nil {
		L.RaiseError("%s open fail" , filename)
		return lua.LNil
	}

	lines := &Lines{
		filename: filename,
		fd: file,
		scanner: bufio.NewScanner(file),
	}

	return lines.ToLightUserData(L)
}

func LuaInjectApi(L *lua.LState , parent *lua.LTable) {
	L.SetField(parent , "lines" , lua.NewGFunction( createLinesLightUserData ))
}