package lines

import (
	"bufio"
	"github.com/edunx/lua"
	"os"
)

const (
	MT string = "ROCK_LINES_GO_MT"
)

func CheckLinesUserData( L *lua.LState , idx int ) *Lines {

	ud := L.CheckUserData( idx )

	v  , ok := ud.Value.(*Lines)
	if ok {
		return v
	}

	L.TypeError(idx , lua.LTUserData)
	return nil
}

func CreateLinesUserData(L *lua.LState) int {
	filename := L.CheckString(1)

	//_ , err := os.Stat( filename )
	//stat , err := os.Stat( filename )
	//if os.IsNotExist(err) {
	//	L.RaiseError("%s not found" , filename )
	//}
	//size := stat.Size()

	file , err := os.Open( filename )
	if err != nil {
		L.RaiseError("%s open fail" , filename)
		return 0
	}

	lines := &Lines{
		filename: filename,
		fd: file,
		scanner: bufio.NewScanner(file),
	}

	ud := L.NewUserDataByInterface(lines , MT)
	L.Push(ud)
	return 1

}


func LuaInjectApi(L *lua.LState , parent *lua.LTable) {
	mt := L.NewTypeMetatable( MT )

	L.SetField(mt , "__index" , L.NewFunction(Get))
	L.SetField(mt , "__newindex" , L.NewFunction(Set))

	L.SetField(parent , "lines" , L.NewFunction(CreateLinesUserData))
}

func Get(L *lua.LState) int {
	self := CheckLinesUserData(L , 1)
	name := L.CheckString(2)
	switch name {
	case "line":

		L.Push(L.NewFunction( func (L *lua.LState) int {
			if self.scanner.Scan() {
				L.Push(lua.LString(self.scanner.Text()))
				return 1
			}

			L.Push(lua.LNil)
			return 1
		}))

	case "close":

		L.Push(L.NewFunction(func(L *lua.LState) int {
			self.fd.Close()
			return 0
		}))

	default:
		L.Push(lua.LNil)
	}

	return 1
}

func Set(L *lua.LState) int {
	return 0
}

func (self *Lines) ToUserData(L *lua.LState) *lua.LUserData {
	return L.NewUserDataByInterface( self , MT )
}
