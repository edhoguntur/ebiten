// SPDX-License-Identifier: MIT

//go:build !android && !darwin && !js && !windows && !opengles

package gl

// #cgo linux,!nintendosdk freebsd,!nintendosdk openbsd,!nintendosdk pkg-config: gl
//
// #ifndef APIENTRY
//   #define APIENTRY
// #endif
//
// #ifndef APIENTRYP
//   #define APIENTRYP APIENTRY *
// #endif
//
// #ifndef GLAPI
//   #define GLAPI extern
// #endif
//
// #include <stdint.h>
//
// typedef unsigned int GLenum;
// typedef unsigned char GLboolean;
// typedef unsigned int GLbitfield;
// typedef int GLint;
// typedef unsigned int GLuint;
// typedef int GLsizei;
// typedef float GLfloat;
// typedef char GLchar;
// typedef ptrdiff_t GLintptr;
// typedef ptrdiff_t GLsizeiptr;
//
// typedef void (APIENTRY *GLDEBUGPROC)(GLenum source,GLenum type,GLuint id,GLenum severity,GLsizei length,const GLchar *message,const void *userParam);
// typedef void (APIENTRY *GLDEBUGPROCARB)(GLenum source,GLenum type,GLuint id,GLenum severity,GLsizei length,const GLchar *message,const void *userParam);
// typedef void (APIENTRY *GLDEBUGPROCKHR)(GLenum source,GLenum type,GLuint id,GLenum severity,GLsizei length,const GLchar *message,const void *userParam);
// typedef void (APIENTRY *GLDEBUGPROCAMD)(GLuint id,GLenum category,GLenum severity,GLsizei length,const GLchar *message,void *userParam);
// typedef unsigned short GLhalfNV;
// typedef GLintptr GLvdpauSurfaceNV;
//
// typedef void  (APIENTRYP GPACTIVETEXTURE)(GLenum  texture);
// typedef void  (APIENTRYP GPATTACHSHADER)(GLuint  program, GLuint  shader);
// typedef void  (APIENTRYP GPBINDATTRIBLOCATION)(GLuint  program, GLuint  index, const GLchar * name);
// typedef void  (APIENTRYP GPBINDBUFFER)(GLenum  target, GLuint  buffer);
// typedef void  (APIENTRYP GPBINDFRAMEBUFFEREXT)(GLenum  target, GLuint  framebuffer);
// typedef void  (APIENTRYP GPBINDRENDERBUFFEREXT)(GLenum  target, GLuint  renderbuffer);
// typedef void  (APIENTRYP GPBINDTEXTURE)(GLenum  target, GLuint  texture);
// typedef void  (APIENTRYP GPBLENDEQUATIONSEPARATE)(GLenum  modeRGB, GLenum  modeAlpha);
// typedef void  (APIENTRYP GPBLENDFUNCSEPARATE)(GLenum  srcRGB, GLenum  dstRGB, GLenum  srcAlpha, GLenum  dstAlpha);
// typedef void  (APIENTRYP GPBUFFERDATA)(GLenum  target, GLsizeiptr  size, const void * data, GLenum  usage);
// typedef void  (APIENTRYP GPBUFFERSUBDATA)(GLenum  target, GLintptr  offset, GLsizeiptr  size, const void * data);
// typedef GLenum  (APIENTRYP GPCHECKFRAMEBUFFERSTATUSEXT)(GLenum  target);
// typedef void  (APIENTRYP GPCLEAR)(GLbitfield  mask);
// typedef void  (APIENTRYP GPCOLORMASK)(GLboolean  red, GLboolean  green, GLboolean  blue, GLboolean  alpha);
// typedef void  (APIENTRYP GPCOMPILESHADER)(GLuint  shader);
// typedef GLuint  (APIENTRYP GPCREATEPROGRAM)();
// typedef GLuint  (APIENTRYP GPCREATESHADER)(GLenum  type);
// typedef void  (APIENTRYP GPDELETEBUFFERS)(GLsizei  n, const GLuint * buffers);
// typedef void  (APIENTRYP GPDELETEFRAMEBUFFERSEXT)(GLsizei  n, const GLuint * framebuffers);
// typedef void  (APIENTRYP GPDELETEPROGRAM)(GLuint  program);
// typedef void  (APIENTRYP GPDELETERENDERBUFFERSEXT)(GLsizei  n, const GLuint * renderbuffers);
// typedef void  (APIENTRYP GPDELETESHADER)(GLuint  shader);
// typedef void  (APIENTRYP GPDELETETEXTURES)(GLsizei  n, const GLuint * textures);
// typedef void  (APIENTRYP GPDISABLE)(GLenum  cap);
// typedef void  (APIENTRYP GPDISABLEVERTEXATTRIBARRAY)(GLuint  index);
// typedef void  (APIENTRYP GPDRAWELEMENTS)(GLenum  mode, GLsizei  count, GLenum  type, const uintptr_t indices);
// typedef void  (APIENTRYP GPENABLE)(GLenum  cap);
// typedef void  (APIENTRYP GPENABLEVERTEXATTRIBARRAY)(GLuint  index);
// typedef void  (APIENTRYP GPFLUSH)();
// typedef void  (APIENTRYP GPFRAMEBUFFERRENDERBUFFEREXT)(GLenum  target, GLenum  attachment, GLenum  renderbuffertarget, GLuint  renderbuffer);
// typedef void  (APIENTRYP GPFRAMEBUFFERTEXTURE2DEXT)(GLenum  target, GLenum  attachment, GLenum  textarget, GLuint  texture, GLint  level);
// typedef void  (APIENTRYP GPGENBUFFERS)(GLsizei  n, GLuint * buffers);
// typedef void  (APIENTRYP GPGENFRAMEBUFFERSEXT)(GLsizei  n, GLuint * framebuffers);
// typedef void  (APIENTRYP GPGENRENDERBUFFERSEXT)(GLsizei  n, GLuint * renderbuffers);
// typedef void  (APIENTRYP GPGENTEXTURES)(GLsizei  n, GLuint * textures);
// typedef GLenum  (APIENTRYP GPGETERROR)();
// typedef void  (APIENTRYP GPGETINTEGERV)(GLenum  pname, GLint * data);
// typedef void  (APIENTRYP GPGETPROGRAMINFOLOG)(GLuint  program, GLsizei  bufSize, GLsizei * length, GLchar * infoLog);
// typedef void  (APIENTRYP GPGETPROGRAMIV)(GLuint  program, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETSHADERINFOLOG)(GLuint  shader, GLsizei  bufSize, GLsizei * length, GLchar * infoLog);
// typedef void  (APIENTRYP GPGETSHADERIV)(GLuint  shader, GLenum  pname, GLint * params);
// typedef GLint  (APIENTRYP GPGETUNIFORMLOCATION)(GLuint  program, const GLchar * name);
// typedef GLboolean  (APIENTRYP GPISFRAMEBUFFEREXT)(GLuint  framebuffer);
// typedef GLboolean  (APIENTRYP GPISPROGRAM)(GLuint  program);
// typedef GLboolean  (APIENTRYP GPISRENDERBUFFEREXT)(GLuint  renderbuffer);
// typedef GLboolean  (APIENTRYP GPISTEXTURE)(GLuint  texture);
// typedef void  (APIENTRYP GPLINKPROGRAM)(GLuint  program);
// typedef void  (APIENTRYP GPPIXELSTOREI)(GLenum  pname, GLint  param);
// typedef void  (APIENTRYP GPREADPIXELS)(GLint  x, GLint  y, GLsizei  width, GLsizei  height, GLenum  format, GLenum  type, void * pixels);
// typedef void  (APIENTRYP GPRENDERBUFFERSTORAGEEXT)(GLenum  target, GLenum  internalformat, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPSCISSOR)(GLint  x, GLint  y, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPSHADERSOURCE)(GLuint  shader, GLsizei  count, const GLchar *const* string, const GLint * length);
// typedef void  (APIENTRYP GPSTENCILFUNC)(GLenum  func, GLint  ref, GLuint  mask);
// typedef void  (APIENTRYP GPSTENCILOP)(GLenum  fail, GLenum  zfail, GLenum  zpass);
// typedef void  (APIENTRYP GPTEXIMAGE2D)(GLenum  target, GLint  level, GLint  internalformat, GLsizei  width, GLsizei  height, GLint  border, GLenum  format, GLenum  type, const void * pixels);
// typedef void  (APIENTRYP GPTEXPARAMETERI)(GLenum  target, GLenum  pname, GLint  param);
// typedef void  (APIENTRYP GPTEXSUBIMAGE2D)(GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLsizei  width, GLsizei  height, GLenum  format, GLenum  type, const void * pixels);
// typedef void  (APIENTRYP GPUNIFORM1FV)(GLint  location, GLsizei  count, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORM1I)(GLint  location, GLint  v0);
// typedef void  (APIENTRYP GPUNIFORM1IV)(GLint  location, GLsizei  count, const GLint * value);
// typedef void  (APIENTRYP GPUNIFORM2FV)(GLint  location, GLsizei  count, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORM3FV)(GLint  location, GLsizei  count, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORM4FV)(GLint  location, GLsizei  count, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX2FV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX3FV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX4FV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPUSEPROGRAM)(GLuint  program);
// typedef void  (APIENTRYP GPVERTEXATTRIBPOINTER)(GLuint  index, GLint  size, GLenum  type, GLboolean  normalized, GLsizei  stride, const uintptr_t pointer);
// typedef void  (APIENTRYP GPVIEWPORT)(GLint  x, GLint  y, GLsizei  width, GLsizei  height);
//
// static void  glowActiveTexture(GPACTIVETEXTURE fnptr, GLenum  texture) {
//   (*fnptr)(texture);
// }
// static void  glowAttachShader(GPATTACHSHADER fnptr, GLuint  program, GLuint  shader) {
//   (*fnptr)(program, shader);
// }
// static void  glowBindAttribLocation(GPBINDATTRIBLOCATION fnptr, GLuint  program, GLuint  index, const GLchar * name) {
//   (*fnptr)(program, index, name);
// }
// static void  glowBindBuffer(GPBINDBUFFER fnptr, GLenum  target, GLuint  buffer) {
//   (*fnptr)(target, buffer);
// }
// static void  glowBindFramebufferEXT(GPBINDFRAMEBUFFEREXT fnptr, GLenum  target, GLuint  framebuffer) {
//   (*fnptr)(target, framebuffer);
// }
// static void  glowBindRenderbufferEXT(GPBINDRENDERBUFFEREXT fnptr, GLenum  target, GLuint  renderbuffer) {
//   (*fnptr)(target, renderbuffer);
// }
// static void  glowBindTexture(GPBINDTEXTURE fnptr, GLenum  target, GLuint  texture) {
//   (*fnptr)(target, texture);
// }
// static void  glowBlendEquationSeparate(GPBLENDEQUATIONSEPARATE fnptr, GLenum  modeRGB, GLenum  modeAlpha) {
//   (*fnptr)(modeRGB, modeAlpha);
// }
// static void  glowBlendFuncSeparate(GPBLENDFUNCSEPARATE fnptr, GLenum  srcRGB, GLenum  dstRGB, GLenum  srcAlpha, GLenum  dstAlpha) {
//   (*fnptr)(srcRGB, dstRGB, srcAlpha, dstAlpha);
// }
// static void  glowBufferData(GPBUFFERDATA fnptr, GLenum  target, GLsizeiptr  size, const void * data, GLenum  usage) {
//   (*fnptr)(target, size, data, usage);
// }
// static void  glowBufferSubData(GPBUFFERSUBDATA fnptr, GLenum  target, GLintptr  offset, GLsizeiptr  size, const void * data) {
//   (*fnptr)(target, offset, size, data);
// }
// static GLenum  glowCheckFramebufferStatusEXT(GPCHECKFRAMEBUFFERSTATUSEXT fnptr, GLenum  target) {
//   return (*fnptr)(target);
// }
// static void  glowClear(GPCLEAR fnptr, GLbitfield  mask) {
//   (*fnptr)(mask);
// }
// static void  glowColorMask(GPCOLORMASK fnptr, GLboolean  red, GLboolean  green, GLboolean  blue, GLboolean  alpha) {
//   (*fnptr)(red, green, blue, alpha);
// }
// static void  glowCompileShader(GPCOMPILESHADER fnptr, GLuint  shader) {
//   (*fnptr)(shader);
// }
// static GLuint  glowCreateProgram(GPCREATEPROGRAM fnptr) {
//   return (*fnptr)();
// }
// static GLuint  glowCreateShader(GPCREATESHADER fnptr, GLenum  type) {
//   return (*fnptr)(type);
// }
// static void  glowDeleteBuffers(GPDELETEBUFFERS fnptr, GLsizei  n, const GLuint * buffers) {
//   (*fnptr)(n, buffers);
// }
// static void  glowDeleteFramebuffersEXT(GPDELETEFRAMEBUFFERSEXT fnptr, GLsizei  n, const GLuint * framebuffers) {
//   (*fnptr)(n, framebuffers);
// }
// static void  glowDeleteProgram(GPDELETEPROGRAM fnptr, GLuint  program) {
//   (*fnptr)(program);
// }
// static void  glowDeleteRenderbuffersEXT(GPDELETERENDERBUFFERSEXT fnptr, GLsizei  n, const GLuint * renderbuffers) {
//   (*fnptr)(n, renderbuffers);
// }
// static void  glowDeleteShader(GPDELETESHADER fnptr, GLuint  shader) {
//   (*fnptr)(shader);
// }
// static void  glowDeleteTextures(GPDELETETEXTURES fnptr, GLsizei  n, const GLuint * textures) {
//   (*fnptr)(n, textures);
// }
// static void  glowDisable(GPDISABLE fnptr, GLenum  cap) {
//   (*fnptr)(cap);
// }
// static void  glowDisableVertexAttribArray(GPDISABLEVERTEXATTRIBARRAY fnptr, GLuint  index) {
//   (*fnptr)(index);
// }
// static void  glowDrawElements(GPDRAWELEMENTS fnptr, GLenum  mode, GLsizei  count, GLenum  type, const uintptr_t indices) {
//   (*fnptr)(mode, count, type, indices);
// }
// static void  glowEnable(GPENABLE fnptr, GLenum  cap) {
//   (*fnptr)(cap);
// }
// static void  glowEnableVertexAttribArray(GPENABLEVERTEXATTRIBARRAY fnptr, GLuint  index) {
//   (*fnptr)(index);
// }
// static void  glowFlush(GPFLUSH fnptr) {
//   (*fnptr)();
// }
// static void  glowFramebufferRenderbufferEXT(GPFRAMEBUFFERRENDERBUFFEREXT fnptr, GLenum  target, GLenum  attachment, GLenum  renderbuffertarget, GLuint  renderbuffer) {
//   (*fnptr)(target, attachment, renderbuffertarget, renderbuffer);
// }
// static void  glowFramebufferTexture2DEXT(GPFRAMEBUFFERTEXTURE2DEXT fnptr, GLenum  target, GLenum  attachment, GLenum  textarget, GLuint  texture, GLint  level) {
//   (*fnptr)(target, attachment, textarget, texture, level);
// }
// static void  glowGenBuffers(GPGENBUFFERS fnptr, GLsizei  n, GLuint * buffers) {
//   (*fnptr)(n, buffers);
// }
// static void  glowGenFramebuffersEXT(GPGENFRAMEBUFFERSEXT fnptr, GLsizei  n, GLuint * framebuffers) {
//   (*fnptr)(n, framebuffers);
// }
// static void  glowGenRenderbuffersEXT(GPGENRENDERBUFFERSEXT fnptr, GLsizei  n, GLuint * renderbuffers) {
//   (*fnptr)(n, renderbuffers);
// }
// static void  glowGenTextures(GPGENTEXTURES fnptr, GLsizei  n, GLuint * textures) {
//   (*fnptr)(n, textures);
// }
// static GLenum  glowGetError(GPGETERROR fnptr) {
//   return (*fnptr)();
// }
// static void  glowGetIntegerv(GPGETINTEGERV fnptr, GLenum  pname, GLint * data) {
//   (*fnptr)(pname, data);
// }
// static void  glowGetProgramInfoLog(GPGETPROGRAMINFOLOG fnptr, GLuint  program, GLsizei  bufSize, GLsizei * length, GLchar * infoLog) {
//   (*fnptr)(program, bufSize, length, infoLog);
// }
// static void  glowGetProgramiv(GPGETPROGRAMIV fnptr, GLuint  program, GLenum  pname, GLint * params) {
//   (*fnptr)(program, pname, params);
// }
// static void  glowGetShaderInfoLog(GPGETSHADERINFOLOG fnptr, GLuint  shader, GLsizei  bufSize, GLsizei * length, GLchar * infoLog) {
//   (*fnptr)(shader, bufSize, length, infoLog);
// }
// static void  glowGetShaderiv(GPGETSHADERIV fnptr, GLuint  shader, GLenum  pname, GLint * params) {
//   (*fnptr)(shader, pname, params);
// }
// static GLint  glowGetUniformLocation(GPGETUNIFORMLOCATION fnptr, GLuint  program, const GLchar * name) {
//   return (*fnptr)(program, name);
// }
// static GLboolean  glowIsFramebufferEXT(GPISFRAMEBUFFEREXT fnptr, GLuint  framebuffer) {
//   return (*fnptr)(framebuffer);
// }
// static GLboolean  glowIsProgram(GPISPROGRAM fnptr, GLuint  program) {
//   return (*fnptr)(program);
// }
// static GLboolean  glowIsRenderbufferEXT(GPISRENDERBUFFEREXT fnptr, GLuint  renderbuffer) {
//   return (*fnptr)(renderbuffer);
// }
// static GLboolean  glowIsTexture(GPISTEXTURE fnptr, GLuint  texture) {
//   return (*fnptr)(texture);
// }
// static void  glowLinkProgram(GPLINKPROGRAM fnptr, GLuint  program) {
//   (*fnptr)(program);
// }
// static void  glowPixelStorei(GPPIXELSTOREI fnptr, GLenum  pname, GLint  param) {
//   (*fnptr)(pname, param);
// }
// static void  glowReadPixels(GPREADPIXELS fnptr, GLint  x, GLint  y, GLsizei  width, GLsizei  height, GLenum  format, GLenum  type, void * pixels) {
//   (*fnptr)(x, y, width, height, format, type, pixels);
// }
// static void  glowRenderbufferStorageEXT(GPRENDERBUFFERSTORAGEEXT fnptr, GLenum  target, GLenum  internalformat, GLsizei  width, GLsizei  height) {
//   (*fnptr)(target, internalformat, width, height);
// }
// static void  glowScissor(GPSCISSOR fnptr, GLint  x, GLint  y, GLsizei  width, GLsizei  height) {
//   (*fnptr)(x, y, width, height);
// }
// static void  glowShaderSource(GPSHADERSOURCE fnptr, GLuint  shader, GLsizei  count, const GLchar *const* string, const GLint * length) {
//   (*fnptr)(shader, count, string, length);
// }
// static void  glowStencilFunc(GPSTENCILFUNC fnptr, GLenum  func, GLint  ref, GLuint  mask) {
//   (*fnptr)(func, ref, mask);
// }
// static void  glowStencilOp(GPSTENCILOP fnptr, GLenum  fail, GLenum  zfail, GLenum  zpass) {
//   (*fnptr)(fail, zfail, zpass);
// }
// static void  glowTexImage2D(GPTEXIMAGE2D fnptr, GLenum  target, GLint  level, GLint  internalformat, GLsizei  width, GLsizei  height, GLint  border, GLenum  format, GLenum  type, const void * pixels) {
//   (*fnptr)(target, level, internalformat, width, height, border, format, type, pixels);
// }
// static void  glowTexParameteri(GPTEXPARAMETERI fnptr, GLenum  target, GLenum  pname, GLint  param) {
//   (*fnptr)(target, pname, param);
// }
// static void  glowTexSubImage2D(GPTEXSUBIMAGE2D fnptr, GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLsizei  width, GLsizei  height, GLenum  format, GLenum  type, const void * pixels) {
//   (*fnptr)(target, level, xoffset, yoffset, width, height, format, type, pixels);
// }
// static void  glowUniform1fv(GPUNIFORM1FV fnptr, GLint  location, GLsizei  count, const GLfloat * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniform1i(GPUNIFORM1I fnptr, GLint  location, GLint  v0) {
//   (*fnptr)(location, v0);
// }
// static void  glowUniform1iv(GPUNIFORM1IV fnptr, GLint  location, GLsizei  count, const GLint * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniform2fv(GPUNIFORM2FV fnptr, GLint  location, GLsizei  count, const GLfloat * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniform3fv(GPUNIFORM3FV fnptr, GLint  location, GLsizei  count, const GLfloat * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniform4fv(GPUNIFORM4FV fnptr, GLint  location, GLsizei  count, const GLfloat * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniformMatrix2fv(GPUNIFORMMATRIX2FV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static void  glowUniformMatrix3fv(GPUNIFORMMATRIX3FV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static void  glowUniformMatrix4fv(GPUNIFORMMATRIX4FV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static void  glowUseProgram(GPUSEPROGRAM fnptr, GLuint  program) {
//   (*fnptr)(program);
// }
// static void  glowVertexAttribPointer(GPVERTEXATTRIBPOINTER fnptr, GLuint  index, GLint  size, GLenum  type, GLboolean  normalized, GLsizei  stride, const uintptr_t pointer) {
//   (*fnptr)(index, size, type, normalized, stride, pointer);
// }
// static void  glowViewport(GPVIEWPORT fnptr, GLint  x, GLint  y, GLsizei  width, GLsizei  height) {
//   (*fnptr)(x, y, width, height);
// }
import "C"

import (
	"errors"
	"unsafe"
)

var (
	gpActiveTexture              C.GPACTIVETEXTURE
	gpAttachShader               C.GPATTACHSHADER
	gpBindAttribLocation         C.GPBINDATTRIBLOCATION
	gpBindBuffer                 C.GPBINDBUFFER
	gpBindFramebufferEXT         C.GPBINDFRAMEBUFFEREXT
	gpBindRenderbufferEXT        C.GPBINDRENDERBUFFEREXT
	gpBindTexture                C.GPBINDTEXTURE
	gpBlendEquationSeparate      C.GPBLENDEQUATIONSEPARATE
	gpBlendFuncSeparate          C.GPBLENDFUNCSEPARATE
	gpBufferData                 C.GPBUFFERDATA
	gpBufferSubData              C.GPBUFFERSUBDATA
	gpCheckFramebufferStatusEXT  C.GPCHECKFRAMEBUFFERSTATUSEXT
	gpClear                      C.GPCLEAR
	gpColorMask                  C.GPCOLORMASK
	gpCompileShader              C.GPCOMPILESHADER
	gpCreateProgram              C.GPCREATEPROGRAM
	gpCreateShader               C.GPCREATESHADER
	gpDeleteBuffers              C.GPDELETEBUFFERS
	gpDeleteFramebuffersEXT      C.GPDELETEFRAMEBUFFERSEXT
	gpDeleteProgram              C.GPDELETEPROGRAM
	gpDeleteRenderbuffersEXT     C.GPDELETERENDERBUFFERSEXT
	gpDeleteShader               C.GPDELETESHADER
	gpDeleteTextures             C.GPDELETETEXTURES
	gpDisable                    C.GPDISABLE
	gpDisableVertexAttribArray   C.GPDISABLEVERTEXATTRIBARRAY
	gpDrawElements               C.GPDRAWELEMENTS
	gpEnable                     C.GPENABLE
	gpEnableVertexAttribArray    C.GPENABLEVERTEXATTRIBARRAY
	gpFlush                      C.GPFLUSH
	gpFramebufferRenderbufferEXT C.GPFRAMEBUFFERRENDERBUFFEREXT
	gpFramebufferTexture2DEXT    C.GPFRAMEBUFFERTEXTURE2DEXT
	gpGenBuffers                 C.GPGENBUFFERS
	gpGenFramebuffersEXT         C.GPGENFRAMEBUFFERSEXT
	gpGenRenderbuffersEXT        C.GPGENRENDERBUFFERSEXT
	gpGenTextures                C.GPGENTEXTURES
	gpGetError                   C.GPGETERROR
	gpGetIntegerv                C.GPGETINTEGERV
	gpGetProgramInfoLog          C.GPGETPROGRAMINFOLOG
	gpGetProgramiv               C.GPGETPROGRAMIV
	gpGetShaderInfoLog           C.GPGETSHADERINFOLOG
	gpGetShaderiv                C.GPGETSHADERIV
	gpGetUniformLocation         C.GPGETUNIFORMLOCATION
	gpIsFramebufferEXT           C.GPISFRAMEBUFFEREXT
	gpIsProgram                  C.GPISPROGRAM
	gpIsRenderbufferEXT          C.GPISRENDERBUFFEREXT
	gpIsTexture                  C.GPISTEXTURE
	gpLinkProgram                C.GPLINKPROGRAM
	gpPixelStorei                C.GPPIXELSTOREI
	gpReadPixels                 C.GPREADPIXELS
	gpRenderbufferStorageEXT     C.GPRENDERBUFFERSTORAGEEXT
	gpScissor                    C.GPSCISSOR
	gpShaderSource               C.GPSHADERSOURCE
	gpStencilFunc                C.GPSTENCILFUNC
	gpStencilOp                  C.GPSTENCILOP
	gpTexImage2D                 C.GPTEXIMAGE2D
	gpTexParameteri              C.GPTEXPARAMETERI
	gpTexSubImage2D              C.GPTEXSUBIMAGE2D
	gpUniform1fv                 C.GPUNIFORM1FV
	gpUniform1i                  C.GPUNIFORM1I
	gpUniform1iv                 C.GPUNIFORM1IV
	gpUniform2fv                 C.GPUNIFORM2FV
	gpUniform3fv                 C.GPUNIFORM3FV
	gpUniform4fv                 C.GPUNIFORM4FV
	gpUniformMatrix2fv           C.GPUNIFORMMATRIX2FV
	gpUniformMatrix3fv           C.GPUNIFORMMATRIX3FV
	gpUniformMatrix4fv           C.GPUNIFORMMATRIX4FV
	gpUseProgram                 C.GPUSEPROGRAM
	gpVertexAttribPointer        C.GPVERTEXATTRIBPOINTER
	gpViewport                   C.GPVIEWPORT
)

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func ActiveTexture(texture uint32) {
	C.glowActiveTexture(gpActiveTexture, (C.GLenum)(texture))
}

func AttachShader(program uint32, shader uint32) {
	C.glowAttachShader(gpAttachShader, (C.GLuint)(program), (C.GLuint)(shader))
}

func BindAttribLocation(program uint32, index uint32, name *uint8) {
	C.glowBindAttribLocation(gpBindAttribLocation, (C.GLuint)(program), (C.GLuint)(index), (*C.GLchar)(unsafe.Pointer(name)))
}

func BindBuffer(target uint32, buffer uint32) {
	C.glowBindBuffer(gpBindBuffer, (C.GLenum)(target), (C.GLuint)(buffer))
}

func BindFramebufferEXT(target uint32, framebuffer uint32) {
	C.glowBindFramebufferEXT(gpBindFramebufferEXT, (C.GLenum)(target), (C.GLuint)(framebuffer))
}

func BindRenderbufferEXT(target uint32, renderbuffer uint32) {
	C.glowBindRenderbufferEXT(gpBindRenderbufferEXT, (C.GLenum)(target), (C.GLuint)(renderbuffer))
}

func BindTexture(target uint32, texture uint32) {
	C.glowBindTexture(gpBindTexture, (C.GLenum)(target), (C.GLuint)(texture))
}

func BlendEquationSeparate(modeRGB uint32, modeAlpha uint32) {
	C.glowBlendEquationSeparate(gpBlendEquationSeparate, (C.GLenum)(modeRGB), (C.GLenum)(modeAlpha))
}

func BlendFuncSeparate(srcRGB uint32, dstRGB uint32, srcAlpha uint32, dstAlpha uint32) {
	C.glowBlendFuncSeparate(gpBlendFuncSeparate, (C.GLenum)(srcRGB), (C.GLenum)(dstRGB), (C.GLenum)(srcAlpha), (C.GLenum)(dstAlpha))
}

func BufferData(target uint32, size int, data unsafe.Pointer, usage uint32) {
	C.glowBufferData(gpBufferData, (C.GLenum)(target), (C.GLsizeiptr)(size), data, (C.GLenum)(usage))
}

func BufferSubData(target uint32, offset int, size int, data unsafe.Pointer) {
	C.glowBufferSubData(gpBufferSubData, (C.GLenum)(target), (C.GLintptr)(offset), (C.GLsizeiptr)(size), data)
}

func CheckFramebufferStatusEXT(target uint32) uint32 {
	ret := C.glowCheckFramebufferStatusEXT(gpCheckFramebufferStatusEXT, (C.GLenum)(target))
	return (uint32)(ret)
}

func Clear(mask uint32) {
	C.glowClear(gpClear, (C.GLbitfield)(mask))
}

func ColorMask(red bool, green bool, blue bool, alpha bool) {
	C.glowColorMask(gpColorMask, (C.GLboolean)(boolToInt(red)), (C.GLboolean)(boolToInt(green)), (C.GLboolean)(boolToInt(blue)), (C.GLboolean)(boolToInt(alpha)))
}

func CompileShader(shader uint32) {
	C.glowCompileShader(gpCompileShader, (C.GLuint)(shader))
}

func CreateProgram() uint32 {
	ret := C.glowCreateProgram(gpCreateProgram)
	return (uint32)(ret)
}

func CreateShader(xtype uint32) uint32 {
	ret := C.glowCreateShader(gpCreateShader, (C.GLenum)(xtype))
	return (uint32)(ret)
}

func DeleteBuffers(n int32, buffers *uint32) {
	C.glowDeleteBuffers(gpDeleteBuffers, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(buffers)))
}

func DeleteFramebuffersEXT(n int32, framebuffers *uint32) {
	C.glowDeleteFramebuffersEXT(gpDeleteFramebuffersEXT, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(framebuffers)))
}

func DeleteProgram(program uint32) {
	C.glowDeleteProgram(gpDeleteProgram, (C.GLuint)(program))
}

func DeleteRenderbuffersEXT(n int32, renderbuffers *uint32) {
	C.glowDeleteRenderbuffersEXT(gpDeleteRenderbuffersEXT, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(renderbuffers)))
}

func DeleteShader(shader uint32) {
	C.glowDeleteShader(gpDeleteShader, (C.GLuint)(shader))
}

func DeleteTextures(n int32, textures *uint32) {
	C.glowDeleteTextures(gpDeleteTextures, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(textures)))
}

func Disable(cap uint32) {
	C.glowDisable(gpDisable, (C.GLenum)(cap))
}

func DisableVertexAttribArray(index uint32) {
	C.glowDisableVertexAttribArray(gpDisableVertexAttribArray, (C.GLuint)(index))
}

func DrawElements(mode uint32, count int32, xtype uint32, indices uintptr) {
	C.glowDrawElements(gpDrawElements, (C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(xtype), C.uintptr_t(indices))
}

func Enable(cap uint32) {
	C.glowEnable(gpEnable, (C.GLenum)(cap))
}

func EnableVertexAttribArray(index uint32) {
	C.glowEnableVertexAttribArray(gpEnableVertexAttribArray, (C.GLuint)(index))
}

func Flush() {
	C.glowFlush(gpFlush)
}

func FramebufferRenderbufferEXT(target uint32, attachment uint32, renderbuffertarget uint32, renderbuffer uint32) {
	C.glowFramebufferRenderbufferEXT(gpFramebufferRenderbufferEXT, (C.GLenum)(target), (C.GLenum)(attachment), (C.GLenum)(renderbuffertarget), (C.GLuint)(renderbuffer))
}

func FramebufferTexture2DEXT(target uint32, attachment uint32, textarget uint32, texture uint32, level int32) {
	C.glowFramebufferTexture2DEXT(gpFramebufferTexture2DEXT, (C.GLenum)(target), (C.GLenum)(attachment), (C.GLenum)(textarget), (C.GLuint)(texture), (C.GLint)(level))
}

func GenBuffers(n int32, buffers *uint32) {
	C.glowGenBuffers(gpGenBuffers, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(buffers)))
}

func GenFramebuffersEXT(n int32, framebuffers *uint32) {
	C.glowGenFramebuffersEXT(gpGenFramebuffersEXT, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(framebuffers)))
}

func GenRenderbuffersEXT(n int32, renderbuffers *uint32) {
	C.glowGenRenderbuffersEXT(gpGenRenderbuffersEXT, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(renderbuffers)))
}

func GenTextures(n int32, textures *uint32) {
	C.glowGenTextures(gpGenTextures, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(textures)))
}

func GetError() uint32 {
	ret := C.glowGetError(gpGetError)
	return (uint32)(ret)
}

func GetIntegerv(pname uint32, data *int32) {
	C.glowGetIntegerv(gpGetIntegerv, (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(data)))
}

func GetProgramInfoLog(program uint32, bufSize int32, length *int32, infoLog *uint8) {
	C.glowGetProgramInfoLog(gpGetProgramInfoLog, (C.GLuint)(program), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(infoLog)))
}

func GetProgramiv(program uint32, pname uint32, params *int32) {
	C.glowGetProgramiv(gpGetProgramiv, (C.GLuint)(program), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}

func GetShaderInfoLog(shader uint32, bufSize int32, length *int32, infoLog *uint8) {
	C.glowGetShaderInfoLog(gpGetShaderInfoLog, (C.GLuint)(shader), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(infoLog)))
}

func GetShaderiv(shader uint32, pname uint32, params *int32) {
	C.glowGetShaderiv(gpGetShaderiv, (C.GLuint)(shader), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}

func GetUniformLocation(program uint32, name *uint8) int32 {
	ret := C.glowGetUniformLocation(gpGetUniformLocation, (C.GLuint)(program), (*C.GLchar)(unsafe.Pointer(name)))
	return (int32)(ret)
}

func IsFramebufferEXT(framebuffer uint32) bool {
	ret := C.glowIsFramebufferEXT(gpIsFramebufferEXT, (C.GLuint)(framebuffer))
	return ret == TRUE
}

func IsProgram(program uint32) bool {
	ret := C.glowIsProgram(gpIsProgram, (C.GLuint)(program))
	return ret == TRUE
}

func IsRenderbufferEXT(renderbuffer uint32) bool {
	ret := C.glowIsRenderbufferEXT(gpIsRenderbufferEXT, (C.GLuint)(renderbuffer))
	return ret == TRUE
}

func IsTexture(texture uint32) bool {
	ret := C.glowIsTexture(gpIsTexture, (C.GLuint)(texture))
	return ret == TRUE
}

func LinkProgram(program uint32) {
	C.glowLinkProgram(gpLinkProgram, (C.GLuint)(program))
}

func PixelStorei(pname uint32, param int32) {
	C.glowPixelStorei(gpPixelStorei, (C.GLenum)(pname), (C.GLint)(param))
}

func ReadPixels(x int32, y int32, width int32, height int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	C.glowReadPixels(gpReadPixels, (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLenum)(xtype), pixels)
}

func RenderbufferStorageEXT(target uint32, internalformat uint32, width int32, height int32) {
	C.glowRenderbufferStorageEXT(gpRenderbufferStorageEXT, (C.GLenum)(target), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height))
}

func Scissor(x int32, y int32, width int32, height int32) {
	C.glowScissor(gpScissor, (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}

func ShaderSource(shader uint32, count int32, xstring **uint8, length *int32) {
	C.glowShaderSource(gpShaderSource, (C.GLuint)(shader), (C.GLsizei)(count), (**C.GLchar)(unsafe.Pointer(xstring)), (*C.GLint)(unsafe.Pointer(length)))
}

func StencilFunc(xfunc uint32, ref int32, mask uint32) {
	C.glowStencilFunc(gpStencilFunc, (C.GLenum)(xfunc), (C.GLint)(ref), (C.GLuint)(mask))
}

func StencilOp(fail uint32, zfail uint32, zpass uint32) {
	C.glowStencilOp(gpStencilOp, (C.GLenum)(fail), (C.GLenum)(zfail), (C.GLenum)(zpass))
}

func TexImage2D(target uint32, level int32, internalformat int32, width int32, height int32, border int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	C.glowTexImage2D(gpTexImage2D, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLint)(border), (C.GLenum)(format), (C.GLenum)(xtype), pixels)
}

func TexParameteri(target uint32, pname uint32, param int32) {
	C.glowTexParameteri(gpTexParameteri, (C.GLenum)(target), (C.GLenum)(pname), (C.GLint)(param))
}

func TexSubImage2D(target uint32, level int32, xoffset int32, yoffset int32, width int32, height int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	C.glowTexSubImage2D(gpTexSubImage2D, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLenum)(xtype), pixels)
}

func Uniform1fv(location int32, count int32, value *float32) {
	C.glowUniform1fv(gpUniform1fv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(unsafe.Pointer(value)))
}

func Uniform1i(location int32, v0 int32) {
	C.glowUniform1i(gpUniform1i, (C.GLint)(location), (C.GLint)(v0))
}

func Uniform1iv(location int32, count int32, value *int32) {
	C.glowUniform1iv(gpUniform1iv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(unsafe.Pointer(value)))
}

func Uniform2fv(location int32, count int32, value *float32) {
	C.glowUniform2fv(gpUniform2fv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(unsafe.Pointer(value)))
}

func Uniform3fv(location int32, count int32, value *float32) {
	C.glowUniform3fv(gpUniform3fv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(unsafe.Pointer(value)))
}

func Uniform4fv(location int32, count int32, value *float32) {
	C.glowUniform4fv(gpUniform4fv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(unsafe.Pointer(value)))
}

func UniformMatrix2fv(location int32, count int32, transpose bool, value *float32) {
	C.glowUniformMatrix2fv(gpUniformMatrix2fv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}

func UniformMatrix3fv(location int32, count int32, transpose bool, value *float32) {
	C.glowUniformMatrix3fv(gpUniformMatrix3fv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}

func UniformMatrix4fv(location int32, count int32, transpose bool, value *float32) {
	C.glowUniformMatrix4fv(gpUniformMatrix4fv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}

func UseProgram(program uint32) {
	C.glowUseProgram(gpUseProgram, (C.GLuint)(program))
}

func VertexAttribPointer(index uint32, size int32, xtype uint32, normalized bool, stride int32, pointer uintptr) {
	C.glowVertexAttribPointer(gpVertexAttribPointer, (C.GLuint)(index), (C.GLint)(size), (C.GLenum)(xtype), (C.GLboolean)(boolToInt(normalized)), (C.GLsizei)(stride), C.uintptr_t(pointer))
}

func Viewport(x int32, y int32, width int32, height int32) {
	C.glowViewport(gpViewport, (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}

// InitWithProcAddrFunc intializes the package using the specified OpenGL
// function pointer loading function.
//
// For more cases Init should be used.
func InitWithProcAddrFunc(getProcAddr func(name string) unsafe.Pointer) error {
	gpActiveTexture = (C.GPACTIVETEXTURE)(getProcAddr("glActiveTexture"))
	if gpActiveTexture == nil {
		return errors.New("gl: glActiveTexture is missing")
	}
	gpAttachShader = (C.GPATTACHSHADER)(getProcAddr("glAttachShader"))
	if gpAttachShader == nil {
		return errors.New("gl: glAttachShader is missing")
	}
	gpBindAttribLocation = (C.GPBINDATTRIBLOCATION)(getProcAddr("glBindAttribLocation"))
	if gpBindAttribLocation == nil {
		return errors.New("gl: glBindAttribLocation is missing")
	}
	gpBindBuffer = (C.GPBINDBUFFER)(getProcAddr("glBindBuffer"))
	if gpBindBuffer == nil {
		return errors.New("gl: glBindBuffer is missing")
	}
	gpBindFramebufferEXT = (C.GPBINDFRAMEBUFFEREXT)(getProcAddr("glBindFramebufferEXT"))
	gpBindRenderbufferEXT = (C.GPBINDRENDERBUFFEREXT)(getProcAddr("glBindRenderbufferEXT"))
	gpBindTexture = (C.GPBINDTEXTURE)(getProcAddr("glBindTexture"))
	if gpBindTexture == nil {
		return errors.New("gl: glBindTexture is missing")
	}
	gpBlendEquationSeparate = (C.GPBLENDEQUATIONSEPARATE)(getProcAddr("glBlendEquationSeparate"))
	if gpBlendEquationSeparate == nil {
		return errors.New("gl: glBlendEquationSeparate is missing")
	}
	gpBlendFuncSeparate = (C.GPBLENDFUNCSEPARATE)(getProcAddr("glBlendFuncSeparate"))
	if gpBlendFuncSeparate == nil {
		return errors.New("gl: glBlendFuncSeparate is missing")
	}
	gpBufferData = (C.GPBUFFERDATA)(getProcAddr("glBufferData"))
	if gpBufferData == nil {
		return errors.New("gl: glBufferData is missing")
	}
	gpBufferSubData = (C.GPBUFFERSUBDATA)(getProcAddr("glBufferSubData"))
	if gpBufferSubData == nil {
		return errors.New("gl: glBufferSubData is missing")
	}
	gpCheckFramebufferStatusEXT = (C.GPCHECKFRAMEBUFFERSTATUSEXT)(getProcAddr("glCheckFramebufferStatusEXT"))
	gpClear = (C.GPCLEAR)(getProcAddr("glClear"))
	if gpClear == nil {
		return errors.New("gl: glClear is missing")
	}
	gpColorMask = (C.GPCOLORMASK)(getProcAddr("glColorMask"))
	if gpColorMask == nil {
		return errors.New("gl: glColorMask is missing")
	}
	gpCompileShader = (C.GPCOMPILESHADER)(getProcAddr("glCompileShader"))
	if gpCompileShader == nil {
		return errors.New("gl: glCompileShader is missing")
	}
	gpCreateProgram = (C.GPCREATEPROGRAM)(getProcAddr("glCreateProgram"))
	if gpCreateProgram == nil {
		return errors.New("gl: glCreateProgram is missing")
	}
	gpCreateShader = (C.GPCREATESHADER)(getProcAddr("glCreateShader"))
	if gpCreateShader == nil {
		return errors.New("gl: glCreateShader is missing")
	}
	gpDeleteBuffers = (C.GPDELETEBUFFERS)(getProcAddr("glDeleteBuffers"))
	if gpDeleteBuffers == nil {
		return errors.New("gl: glDeleteBuffers is missing")
	}
	gpDeleteFramebuffersEXT = (C.GPDELETEFRAMEBUFFERSEXT)(getProcAddr("glDeleteFramebuffersEXT"))
	gpDeleteProgram = (C.GPDELETEPROGRAM)(getProcAddr("glDeleteProgram"))
	if gpDeleteProgram == nil {
		return errors.New("gl: glDeleteProgram is missing")
	}
	gpDeleteRenderbuffersEXT = (C.GPDELETERENDERBUFFERSEXT)(getProcAddr("glDeleteRenderbuffersEXT"))
	gpDeleteShader = (C.GPDELETESHADER)(getProcAddr("glDeleteShader"))
	if gpDeleteShader == nil {
		return errors.New("gl: glDeleteShader is missing")
	}
	gpDeleteTextures = (C.GPDELETETEXTURES)(getProcAddr("glDeleteTextures"))
	if gpDeleteTextures == nil {
		return errors.New("gl: glDeleteTextures is missing")
	}
	gpDisable = (C.GPDISABLE)(getProcAddr("glDisable"))
	if gpDisable == nil {
		return errors.New("gl: glDisable is missing")
	}
	gpDisableVertexAttribArray = (C.GPDISABLEVERTEXATTRIBARRAY)(getProcAddr("glDisableVertexAttribArray"))
	if gpDisableVertexAttribArray == nil {
		return errors.New("gl: glDisableVertexAttribArray is missing")
	}
	gpDrawElements = (C.GPDRAWELEMENTS)(getProcAddr("glDrawElements"))
	if gpDrawElements == nil {
		return errors.New("gl: glDrawElements is missing")
	}
	gpEnable = (C.GPENABLE)(getProcAddr("glEnable"))
	if gpEnable == nil {
		return errors.New("gl: glEnable is missing")
	}
	gpEnableVertexAttribArray = (C.GPENABLEVERTEXATTRIBARRAY)(getProcAddr("glEnableVertexAttribArray"))
	if gpEnableVertexAttribArray == nil {
		return errors.New("gl: glEnableVertexAttribArray is missing")
	}
	gpFlush = (C.GPFLUSH)(getProcAddr("glFlush"))
	if gpFlush == nil {
		return errors.New("gl: glFlush is missing")
	}
	gpFramebufferRenderbufferEXT = (C.GPFRAMEBUFFERRENDERBUFFEREXT)(getProcAddr("glFramebufferRenderbufferEXT"))
	gpFramebufferTexture2DEXT = (C.GPFRAMEBUFFERTEXTURE2DEXT)(getProcAddr("glFramebufferTexture2DEXT"))
	gpGenBuffers = (C.GPGENBUFFERS)(getProcAddr("glGenBuffers"))
	if gpGenBuffers == nil {
		return errors.New("gl: glGenBuffers is missing")
	}
	gpGenFramebuffersEXT = (C.GPGENFRAMEBUFFERSEXT)(getProcAddr("glGenFramebuffersEXT"))
	gpGenRenderbuffersEXT = (C.GPGENRENDERBUFFERSEXT)(getProcAddr("glGenRenderbuffersEXT"))
	gpGenTextures = (C.GPGENTEXTURES)(getProcAddr("glGenTextures"))
	if gpGenTextures == nil {
		return errors.New("gl: glGenTextures is missing")
	}
	gpGetError = (C.GPGETERROR)(getProcAddr("glGetError"))
	if gpGetError == nil {
		return errors.New("gl: glGetError is missing")
	}
	gpGetIntegerv = (C.GPGETINTEGERV)(getProcAddr("glGetIntegerv"))
	if gpGetIntegerv == nil {
		return errors.New("gl: glGetIntegerv is missing")
	}
	gpGetProgramInfoLog = (C.GPGETPROGRAMINFOLOG)(getProcAddr("glGetProgramInfoLog"))
	if gpGetProgramInfoLog == nil {
		return errors.New("gl: glGetProgramInfoLog is missing")
	}
	gpGetProgramiv = (C.GPGETPROGRAMIV)(getProcAddr("glGetProgramiv"))
	if gpGetProgramiv == nil {
		return errors.New("gl: glGetProgramiv is missing")
	}
	gpGetShaderInfoLog = (C.GPGETSHADERINFOLOG)(getProcAddr("glGetShaderInfoLog"))
	if gpGetShaderInfoLog == nil {
		return errors.New("gl: glGetShaderInfoLog is missing")
	}
	gpGetShaderiv = (C.GPGETSHADERIV)(getProcAddr("glGetShaderiv"))
	if gpGetShaderiv == nil {
		return errors.New("gl: glGetShaderiv is missing")
	}
	gpGetUniformLocation = (C.GPGETUNIFORMLOCATION)(getProcAddr("glGetUniformLocation"))
	if gpGetUniformLocation == nil {
		return errors.New("gl: glGetUniformLocation is missing")
	}
	gpIsFramebufferEXT = (C.GPISFRAMEBUFFEREXT)(getProcAddr("glIsFramebufferEXT"))
	gpIsProgram = (C.GPISPROGRAM)(getProcAddr("glIsProgram"))
	if gpIsProgram == nil {
		return errors.New("gl: glIsProgram is missing")
	}
	gpIsRenderbufferEXT = (C.GPISRENDERBUFFEREXT)(getProcAddr("glIsRenderbufferEXT"))
	gpIsTexture = (C.GPISTEXTURE)(getProcAddr("glIsTexture"))
	if gpIsTexture == nil {
		return errors.New("gl: glIsTexture is missing")
	}
	gpLinkProgram = (C.GPLINKPROGRAM)(getProcAddr("glLinkProgram"))
	if gpLinkProgram == nil {
		return errors.New("gl: glLinkProgram is missing")
	}
	gpPixelStorei = (C.GPPIXELSTOREI)(getProcAddr("glPixelStorei"))
	if gpPixelStorei == nil {
		return errors.New("gl: glPixelStorei is missing")
	}
	gpReadPixels = (C.GPREADPIXELS)(getProcAddr("glReadPixels"))
	if gpReadPixels == nil {
		return errors.New("gl: glReadPixels is missing")
	}
	gpRenderbufferStorageEXT = (C.GPRENDERBUFFERSTORAGEEXT)(getProcAddr("glRenderbufferStorageEXT"))
	gpScissor = (C.GPSCISSOR)(getProcAddr("glScissor"))
	if gpScissor == nil {
		return errors.New("gl: glScissor is missing")
	}
	gpShaderSource = (C.GPSHADERSOURCE)(getProcAddr("glShaderSource"))
	if gpShaderSource == nil {
		return errors.New("gl: glShaderSource is missing")
	}
	gpStencilFunc = (C.GPSTENCILFUNC)(getProcAddr("glStencilFunc"))
	if gpStencilFunc == nil {
		return errors.New("gl: glStencilFunc is missing")
	}
	gpStencilOp = (C.GPSTENCILOP)(getProcAddr("glStencilOp"))
	if gpStencilOp == nil {
		return errors.New("gl: glStencilOp is missing")
	}
	gpTexImage2D = (C.GPTEXIMAGE2D)(getProcAddr("glTexImage2D"))
	if gpTexImage2D == nil {
		return errors.New("gl: glTexImage2D is missing")
	}
	gpTexParameteri = (C.GPTEXPARAMETERI)(getProcAddr("glTexParameteri"))
	if gpTexParameteri == nil {
		return errors.New("gl: glTexParameteri is missing")
	}
	gpTexSubImage2D = (C.GPTEXSUBIMAGE2D)(getProcAddr("glTexSubImage2D"))
	if gpTexSubImage2D == nil {
		return errors.New("gl: glTexSubImage2D is missing")
	}
	gpUniform1fv = (C.GPUNIFORM1FV)(getProcAddr("glUniform1fv"))
	if gpUniform1fv == nil {
		return errors.New("gl: glUniform1fv is missing")
	}
	gpUniform1i = (C.GPUNIFORM1I)(getProcAddr("glUniform1i"))
	if gpUniform1i == nil {
		return errors.New("gl: glUniform1i is missing")
	}
	gpUniform1iv = (C.GPUNIFORM1IV)(getProcAddr("glUniform1iv"))
	if gpUniform1iv == nil {
		return errors.New("gl: glUniform1iv is missing")
	}
	gpUniform2fv = (C.GPUNIFORM2FV)(getProcAddr("glUniform2fv"))
	if gpUniform2fv == nil {
		return errors.New("gl: glUniform2fv is missing")
	}
	gpUniform3fv = (C.GPUNIFORM3FV)(getProcAddr("glUniform3fv"))
	if gpUniform3fv == nil {
		return errors.New("gl: glUniform3fv is missing")
	}
	gpUniform4fv = (C.GPUNIFORM4FV)(getProcAddr("glUniform4fv"))
	if gpUniform4fv == nil {
		return errors.New("gl: glUniform4fv is missing")
	}
	gpUniformMatrix2fv = (C.GPUNIFORMMATRIX2FV)(getProcAddr("glUniformMatrix2fv"))
	if gpUniformMatrix2fv == nil {
		return errors.New("gl: glUniformMatrix2fv is missing")
	}
	gpUniformMatrix3fv = (C.GPUNIFORMMATRIX3FV)(getProcAddr("glUniformMatrix3fv"))
	if gpUniformMatrix3fv == nil {
		return errors.New("gl: glUniformMatrix3fv is missing")
	}
	gpUniformMatrix4fv = (C.GPUNIFORMMATRIX4FV)(getProcAddr("glUniformMatrix4fv"))
	if gpUniformMatrix4fv == nil {
		return errors.New("gl: glUniformMatrix4fv is missing")
	}
	gpUseProgram = (C.GPUSEPROGRAM)(getProcAddr("glUseProgram"))
	if gpUseProgram == nil {
		return errors.New("gl: glUseProgram is missing")
	}
	gpVertexAttribPointer = (C.GPVERTEXATTRIBPOINTER)(getProcAddr("glVertexAttribPointer"))
	if gpVertexAttribPointer == nil {
		return errors.New("gl: glVertexAttribPointer is missing")
	}
	gpViewport = (C.GPVIEWPORT)(getProcAddr("glViewport"))
	if gpViewport == nil {
		return errors.New("gl: glViewport is missing")
	}
	return nil
}
