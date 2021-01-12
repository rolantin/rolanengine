#ifdef VERTEX
#version 410
in vec3 vp;
void main() {
    gl_Position = vec4(vp, 1.0);
}
#endif

#ifdef FRAGMENT
#version 410
out vec4 frag_colour;
void main() {
    frag_colour = vec4(1, 1, 1, 1);
}
#endif