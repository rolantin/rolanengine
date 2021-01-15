#version 410
uniform mat4 MVP;
in vec3 position;
out vec3 customColor;
void main(){
    // Output position of the vertex, in clip space : MVP * position
    //gl_Position =  MVP * vec4(vertexPosition_modelspace,1);
    gl_Position = MVP * vec4(position, 1.0);
    //customColor = vec3(MVP[0][0],MVP[0][1],MVP[0][2]);
    customColor = vec3(MVP[0].x,MVP[0].y,MVP[0].z);
}