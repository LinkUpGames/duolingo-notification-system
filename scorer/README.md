# Scorer

The scorer is responsible for updating the scores of a notification in a script that runs after a certain event is called. This is controlled by using a Python script with the actual algorithm implementation in a **C** extension.

## Development Setup

If you want to setup the development environment locally, you can use the following commands:

```bash
# Start a virtual environment
python -m venv env
source env/bin/activate

# Install requirements
pip install -r requirements.txt

# Compile Module
cd extensions/algorithm/
make # Compiles the C module and installs it in the virtual environment
```

To check out the example you can use the `test.py` script to see the script in action.

## Docker Container
If you wish to do the same but in a Docker container instead you can just run the Docker container with the following:
```bash
docker build -t scorer .
docker run -it scorer sh
python test.py # optional: If you wish to test the container immediately
```

