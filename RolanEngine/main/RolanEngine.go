package main
//http://www.360doc.cn/mip/752185515.html tutorial
import (
	"fmt"
	"github.com/g3n/RolanEngine/baseMesh"
	"github.com/g3n/threeParty/gl/v4.6-core/gl"
	"github.com/g3n/threeParty/glfw"
	"log"
	"runtime"
	"strings"
)

const (
	width  = 1024
	height = 768

	vertexShaderSource = `
        #version 410
        in vec3 vp;
        void main() {
            gl_Position = vec4(vp, 1.0);
        }
    ` + "\x00"
	fragmentShaderSource = `
        #version 410
        out vec4 frag_colour;
        void main() {
            frag_colour = vec4(1, 1, 1, 1);
        }
    ` + "\x00"
)

var meshdata baseMesh.RoBaseMesh

func main() {
	//LockOSThread() 这能确保我们总是在操作系统的同一个线程中运行代码，这对 GLFW 来说很重要
	runtime.LockOSThread()
	//接下来我们调用 initGlfw 来获得一个窗口的引用
	window := initGlfw()
	//并且推迟（defer）其终止
	defer glfw.Terminate()
	program := initOpenGL()

	meshdata:=new(baseMesh.RoBaseMesh)
	meshdata.CreateTriangle()
	log.Println("triangle data",meshdata.CreateTriangle())
	//-----------------------------------------

	vao:=makeVAO(meshdata.CreateTriangle())


	//窗口的引用会被用在一个 for 循环中，只要窗口处于打开的状态，就执行某些事情。
	for !window.ShouldClose() {
		//draw(window, program)
		draw(vao,window, program)
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
	if err := gl.Init(); err != nil{
		panic(err)
	}
	var version string = gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)
	//-----------------------------------------

	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}
	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}


	var prog uint32 = gl.CreateProgram()

	gl.AttachShader(prog, vertexShader)
	gl.AttachShader(prog, fragmentShader)

	gl.LinkProgram(prog) //--> bind program
	return prog
}

//你应该注意到了现在我们有 program 的引用，在我们的窗口循环中，
//调用新的 draw 函数。最终这个函数会绘制出所有细胞，让游戏状态变得可视化，
//但是现在它做的仅仅是清除窗口，所以我们只能看到一个全黑的屏幕：
func draw(vao uint32 , window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT) //--> clean last frame
	gl.UseProgram(program)

	gl.BindVertexArray(vao)
	//然后我们把 OpenGL 绑定到 vao 上，这样当我们告诉 OpenGL 三角形切片的顶点数（除以 3，是因为每一个点有 X、Y、Z 坐标），让它去 DrawArrays ，它就知道要画多少个顶点了。
	gl.DrawArrays(gl.TRIANGLES,0,int32(len(meshdata.CreateTriangle())/3))


	glfw.PollEvents() //--> PollEvents check if have mouse&keyboard input
	//GLFW（像其他图形库一样）使用双缓冲，也就是说你绘制的所有东西实际上是绘制到一个不可见的画布上，当你准备好进行展示的时候就把绘制的这些东西放到可见的画布中 —— 这种情况下，就需要调用 SwapBuffers 函数
	window.SwapBuffers()

}

func makeVAO(points []float32) uint32{

	//首先我们创造了顶点缓冲区对象 或者说 vbo 绑定到我们的 vao 上，vbo 是通过所占空间（也就是 4 倍 len(points) 大小的空间）和一个指向顶点的指针（gl.Ptr(points)）来创建的。你也许会好奇为什么它是 4 倍 —— 而不是 6 或者 3 或者 1078 呢？原因在于我们用的是 float32 切片，32 个位的浮点型变量是 4 个字节，因此我们说这个缓冲区以字节为单位的大小是点个数的 4 倍。
	//现在我们有缓冲区了，可以创建 vao 并用 gl.BindBuffer 把它绑定到缓冲区上，最后返回 vao。这个 vao 将会被用于绘制三角形！
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


func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)
	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)
	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)
		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))
		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}
	return shader, nil
}

