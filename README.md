
# go-rlottie
golang binding rlottie

[![Build Status](https://travis-ci.org/Samsung/rlottie.svg?branch=master)](https://travis-ci.org/Samsung/rlottie)
[![Build status](https://ci.appveyor.com/api/projects/status/n3xonxk1ooo6s7nr?svg=true&passingText=windows%20-%20passing)](https://ci.appveyor.com/project/smohantty/rlottie-mliua)


rlottie is a platform independent standalone c++ library for rendering vector based animations and art in realtime.

Lottie loads and renders animations and vectors exported in the bodymovin JSON format. Bodymovin JSON can be created and exported from After Effects with [bodymovin](https://github.com/bodymovin/bodymovin), Sketch with [Lottie Sketch Export](https://github.com/buba447/Lottie-Sketch-Export), and from [Haiku](https://www.haiku.ai).

For the first time, designers can create and ship beautiful animations without an engineer painstakingly recreating it by hand. Since the animation is backed by JSON they are extremely small in size but can be large in complexity!

rottie supports [meson](https://mesonbuild.com/) and [cmake](https://cmake.org/) build system. rottie is written in `C++14`. and has a public header dependancy of `C++11`.

[rlottie repository](https://github.com/Samsung/rlottie)