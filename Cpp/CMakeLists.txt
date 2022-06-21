cmake_minimum_required(VERSION 3.17)
project(Cpp)

add_subdirectory(./googletest/googletest)
add_subdirectory(src)
set(LIBRARIES gtest pthread)

set(CMAKE_CXX_STANDARD 11)

add_executable(Cpp main.cpp src/DFS/codec.cpp)
target_link_libraries(Cpp ${LIBRARIES})