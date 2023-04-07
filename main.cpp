#include <iostream>
#include <boost/program_options.hpp>

#define VERSION "0.0.1"
#define CMD_CAPTION "gazer"

int main(int argc, char **argv) {
    using namespace boost::program_options;
    options_description optionsDescription(CMD_CAPTION);
    optionsDescription.add_options()
            ("help,H", "show help")
            ("version,v", "show version");

    variables_map variablesMap;
    basic_parsed_options<char> basicParsedOptions = parse_command_line(argc, argv, optionsDescription);
    store(basicParsedOptions, variablesMap);
    notify(variablesMap);

    if (variablesMap.count("help")) {
        std::cout << optionsDescription << std::endl;
        return 0;
    }

    if (variablesMap.count("version")) {
        std::cout << VERSION << std::endl;
        return 0;
    }

    return 0;
}
