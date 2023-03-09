cmake_minimum_required(VERSION 3.8)

project("gazer" VERSION 0.0.1)

set(CMAKE_CXX_STANDARD 11)
set(CMAKE_CXX_COMPILER "/usr/bin/g++" CACHE STRING "gnu c++ compiler" FORCE)
set(CMAKE_C_COMPILER "/usr/bin/gcc" CACHE STRING "gnu cc compiler" FORCE)
set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -Wall -O3")
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -Wall -O3 -std=c++11")

find_package(Boost 1.74 COMPONENTS program_options REQUIRED)

file(GLOB GAZER_SRC "*.h" "*.cpp")

add_executable(gazer ${GAZER_SRC})

target_include_directories(gazer PRIVATE ${Boost_INCLUDE_DIR})
target_link_libraries(gazer ${Boost_LIBRARIES})

