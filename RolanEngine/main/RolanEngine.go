package main
//http://www.360doc.cn/mip/752185515.html tutorial
import (
	"fmt"
	"github.com/g3n/RolanEngine/baseMesh"
	"github.com/g3n/RolanEngine/config"
	m "github.com/g3n/RolanEngine/math32"
	shader "github.com/g3n/RolanEngine/shader/shaderManager"
	"github.com/g3n/threeParty/gl/v4.6-core/gl"
	"github.com/g3n/threeParty/glfw"
	"log"
	"runtime"
	"strings"
)

var meshdata baseMesh.RoBaseMesh
var window *glfw.Window


func main() {
	shader.ShaderLoad("Unlit")
	runtime.LockOSThread() //LockOSThread() 这能确保我们总是在操作系统的同一个线程中运行代码，这对 GLFW 来说很重要
	window = initGlfw() //接下来我们调用 initGlfw 来获得一个窗口的引用

	defer glfw.Terminate() //并且推迟（defer）其终止
	program := initOpenGL()

	// Get a handle for our "MVP" uniform
	MatrixID := gl.GetUniformLocation(program,gl.Str("MVP"+"\x00"))
	println(MatrixID)

	meshdata:=new(baseMesh.RoBaseMesh)
	meshdata.CreateTriangle()
	log.Println("triangle data",meshdata.CreateTriangle())
	//-----------------------------------------

	vao:=makeVAO(meshdata.CreateTriangle())

	//look at data
	cameraPos:= m.NewVector3(4.0,3.0,3.0)
	LookAtOrigin := m.NewVector3(0.0,0.0,0.0)
	HeadisUp := m.NewVector3(0,1,0)
////////////////////////////////////////////////////////
	mat4 := new(m.Matrix4)
	//v3 := new(m.Vector3)
	ModelMatrix := mat4.Identity()
	ViewMatrix    := mat4.LookAt(cameraPos,LookAtOrigin,HeadisUp)
	ProjectMatrix := mat4.MakePerspective(45.0,config.WindowWidth/config.WindowHeight,0.01,1000.0)

	//MV := m.MultiplyMatrices(ModelMatrix,ViewMatrix)
	VP := mat4.MultiplyMatrices(ProjectMatrix,ViewMatrix)
	MVP:= mat4.MultiplyMatrices(VP,ModelMatrix)

	//println(ModelMatrix[0],ModelMatrix[1],ModelMatrix[2])

    println(MVP[0],MVP[1],MVP[2])
	//窗口的引用会被用在一个 for 循环中，只要窗口处于打开的状态，就执行某些事情。
	for !window.ShouldClose() {
		//draw(window, program)
		draw(vao,window, program,MatrixID,&MVP[0])

		update()

	}
}

func update(){
    //fmt.Println("xxx")
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
	win, err := glfw.CreateWindow( config.WindowWidth , config.WindowHeight, "RolanEngine", nil, nil)
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

	vertexShader, err := compileShader(shader.Shaderlink.VertexShader, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}
	fragmentShader, err := compileShader(shader.Shaderlink.FragmentShader, gl.FRAGMENT_SHADER)
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
func draw(vao uint32 , window *glfw.Window, program uint32,MatrixID int32,add *float32) {

	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT) //--> clean last frame
	gl.UseProgram(program)



	// Send our transformation to the currently bound shader,
	// in the "MVP" uniform
	gl.UniformMatrix4fv(MatrixID, 1,false, add);

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

// compileShader 检查编译shader是否通过
func compileShader(shaderSourceCode string, shaderType uint32) (uint32, error) {
	//create shader obj
	shaderObj := gl.CreateShader(shaderType)//****************
    //Strs获取Go字符串列表（带或不带null终止）并返回它们的C对应项。一旦使用完字符串，就必须调用返回的free函数以释放内存。如果没有提供字符串作为参数，此函数将死机
	csources, free := gl.Strs(shaderSourceCode)
	gl.ShaderSource(shaderObj, 1, csources, nil)//****************
	free() //调用返回的free函数以释放内存
    //编译shader
	gl.CompileShader(shaderObj)//****************
	var status int32
	//从shader对象返回参数
	gl.GetShaderiv(shaderObj, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shaderObj, gl.INFO_LOG_LENGTH, &logLength)
		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shaderObj, logLength, nil, gl.Str(log))
		return 0, fmt.Errorf("failed to compile %v: %v", shaderSourceCode, log)
	}
	return shaderObj, nil
}

