#version 410
// Output data
out vec3 color;
in vec3 customColor;
void main(){
    // Output color = red
    color = customColor;//vec3(1,0,0);
}