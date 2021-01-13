package shaderManager

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// ShaderSource vs ps 文本 与 vs ps 路径
type ShaderSource struct {
	VertexShader string
	FragmentShader string
	VsPath string
	PsPath string
}

var Shaderlink ShaderSource

// ReadShaderLab 分别读取vs 和 ps 两份文件进行加载shader
func ReadShaderLab(VSpath string,PSpath string){
	vsfile,err := os.Open(VSpath)
	psfile,err := os.Open(PSpath)
	if err != nil {
		log.Fatal(err)
	}
	defer vsfile.Close()
	defer psfile.Close()
	//// 使用ioutil读取文件所有内容
	vs, err := ioutil.ReadAll(vsfile)
	ps, err := ioutil.ReadAll(psfile)
	//指针使用
	 //p:= &Shaderlink.VsPath
	 //*p = "x"
	Shaderlink.VertexShader = string(vs) + "\x00"
	Shaderlink.FragmentShader = string(ps) + "\x00"
}

// ClipShaderLab 进行分割一份glsl 进行分别加载 vs 部分 和 ps 部分 //todo
func ClipShaderLab(path string){
	//path := "RolanEngine/shader/Unlit.shader"
	file,err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	//// 使用ioutil读取文件所有内容
	b, err := ioutil.ReadAll(file)
	//fmt.Print(string(b))
	tracer := string(b)
	//comma的意思是从字符串tracer查找第一个 #ifdef FRAGMENT，然后返回他的位置，这里的每个中文是占3个字符，从0开始计算，那么#ifdef FRAGMENT的位置就是12
	comma := strings.Index(tracer, "#ifdef FRAGMENT")
	Shaderlink.FragmentShader = tracer[comma:]
}

// ShaderPath 用于定义shader加载路径
func ShaderPath(shaderName string){
	basePath:= "RolanEngine/shader/shaderType/"
	shaderVSName:= shaderName + ".vs.glsl"
	shaderPSName:= shaderName + ".ps.glsl"

	vsP:= &Shaderlink.VsPath
	*vsP =  basePath + shaderVSName

	Shaderlink.VsPath = basePath + shaderVSName
	Shaderlink.PsPath = basePath + shaderPSName
	ReadShaderLab(Shaderlink.VsPath,Shaderlink.PsPath)
}

// ShaderLoad 加载shader相关的所有项
func ShaderLoad(shaderName string){
	ShaderPath(shaderName)
}


//sample base shader
/*
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
*/