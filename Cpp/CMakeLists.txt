cmake_minimum_required(VERSION 3.17)
project(Cpp)

add_subdirectory(./googletest/googletest)
set(LIBRARIES gtest pthread)

set(CMAKE_CXX_STANDARD 11)

add_executable(Cpp main.cpp)
target_link_libraries(Cpp ${LIBRARIES})