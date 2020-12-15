package main

import (
	"github.com/g3n/RolanEngine/baseMesh"
	"github.com/g3n/threeParty/gl/v4.6-core/gl"
	"github.com/g3n/threeParty/glfw"
	"log"
	"runtime"

)

const (
	width  = 1024
	height = 768
)

func main() {
	//LockOSThread() 这能确保我们总是在操作系统的同一个线程中运行代码，这对 GLFW 来说很重要
	runtime.LockOSThread()
	//接下来我们调用 initGlfw 来获得一个窗口的引用
	window := initGlfw()
	//并且推迟（defer）其终止
	defer glfw.Terminate()
	program := initOpenGL()
	//窗口的引用会被用在一个 for 循环中，只要窗口处于打开的状态，就执行某些事情。
	for !window.ShouldClose() {
		draw(window, program)
	}
}

func initGlfw() *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4) // OR 2
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	//create window
	win, err := glfw.CreateWindow(width, height, "RolanEngine", nil, nil)
	if err != nil {
		panic(err)
	}
	//告诉它宽度和高度，以及标题，然后调用 window.MakeContextCurrent，将窗口绑定到当前的线程中。最后就是返回窗口的引用了。
	win.MakeContextCurrent()
	return win
}

//initOpenGL 就像之前的 initGlfw 函数一样，初始化 OpenGL 库，创建一个程序program。
//“程序”是一个包含了着色器shader的引用，稍后会用着色器shader绘图。待会儿会讲这一点，现在只用知道
//OpenGL 已经初始化完成了，我们有一个程序的引用。我们还打印了 OpenGL 的版本，可以用于之后的调试。
// initOpenGL 初始化 OpenGL 并且返回一个初始化了的程序。
func initOpenGL() uint32 {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	var version string = gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)
	//-----------------------------------------
	meshdata:=new(baseMesh.RoBaseMesh)
	meshdata.CreateTriangle()
	log.Println("triangle data",meshdata.CreateTriangle())
	//-----------------------------------------

	var prog uint32 = gl.CreateProgram()
	gl.LinkProgram(prog) //--> bind program
	return prog
}

//你应该注意到了现在我们有 program 的引用，在我们的窗口循环中，
//调用新的 draw 函数。最终这个函数会绘制出所有细胞，让游戏状态变得可视化，
//但是现在它做的仅仅是清除窗口，所以我们只能看到一个全黑的屏幕：
func draw(window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT) //--> clean last frame
	//gl.UseProgram(prog)
	glfw.PollEvents() //--> PollEvents check if have mouse&keyboard input
	//GLFW（像其他图形库一样）使用双缓冲，也就是说你绘制的所有东西实际上是绘制到一个不可见的画布上，当你准备好进行展示的时候就把绘制的这些东西放到可见的画布中 —— 这种情况下，就需要调用 SwapBuffers 函数
	window.SwapBuffers()

}

func makeVAO(points []float32) uint32{
	var vbo uint32
	gl.GenBuffers(1,&vbo) //--> 生成vbo
	gl.BindBuffer(gl.ARRAY_BUFFER,vbo)
	gl.BufferData(gl.ARRAY_BUFFER,4*len(points),gl.Ptr(points),gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1,&vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER,vbo)
	gl.VertexAttribPointer(0,3,gl.FLOAT,false,0,nil)

	return vao  //
}

