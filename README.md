# Parrot

Example app used to test docker and kubernetes services.
It's a simple go app that runs on port 4300 and renders parrot image.

## Usage

    ./parrot

## Docker usage

    docker build . -t andrzejtrzaska/parrot
    docker run -p 4300:4300 andrzejtrzaska/parrot
    curl -v http://localhost:4300

## License

MIT
